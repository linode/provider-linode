/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/go-resty/resty/v2"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"
	terraformsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/linode/provider-linode/apis/v1beta1"
	"github.com/linode/provider-linode/internal/metrics"

	"github.com/linode/terraform-provider-linode/v2/linode"
	"github.com/linode/terraform-provider-linode/v2/linode/helper"
	"github.com/linode/terraform-provider-linode/v2/version"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal linode credentials as JSON"
)

const (
	keyToken   = "token"
	keyURL     = "api_url"
	keyVersion = "api_version"
)

type SetupConfig struct {
	NativeProviderPath    *string
	NativeProviderSource  *string
	NativeProviderVersion *string
	TerraformVersion      *string
	DefaultScheduler      terraform.ProviderScheduler
	TerraformProvider     *schema.Provider
}

func prepareTerraformProviderConfiguration(creds map[string]string, pc v1beta1.ProviderConfiguration) map[string]any {
	config := make(map[string]any)

	if v, ok := creds[keyToken]; ok {
		config[keyToken] = v
	}
	if v, ok := creds[keyURL]; ok {
		config[keyURL] = v
	}
	if v, ok := creds[keyVersion]; ok {
		config[keyVersion] = v
	}

	// Set the booleans as is.
	config["obj_bucket_force_delete"] = pc.ObjForceDelete
	config["obj_use_temp_keys"] = pc.ObjUseTempKeys
	config["skip_instance_ready_poll"] = pc.SkipInstanceReadyPoll
	config["skip_instance_delete_poll"] = pc.SkipInstanceDeletePoll
	config["skip_implicit_reboots"] = pc.SkipImplicitReboots
	config["disable_internal_cache"] = pc.DisableInternalCache

	// do not want to override terraform defaults
	if len(pc.UserAgentPrefix) > 0 {
		config["ua_prefix"] = pc.UserAgentPrefix
	}

	if pc.MinRetryDelayms > 0 {
		config["min_retry_delay_ms"] = pc.MinRetryDelayms
	}

	if pc.MaxRetryDelayms > 0 {
		config["max_retry_delay_ms"] = pc.MaxRetryDelayms
	}

	if pc.EventPollms > 0 {
		config["event_poll_ms"] = pc.MaxRetryDelayms
	}

	if pc.LKEEventPollms > 0 {
		config["lke_event_poll_ms"] = pc.LKEEventPollms
	}

	if pc.LKENodeReadyPollms > 0 {
		config["lke_node_ready_poll_ms"] = pc.LKENodeReadyPollms
	}

	if len(pc.ObjAccessKey) > 0 {
		config["obj_access_key"] = pc.ObjAccessKey
	}

	if len(pc.ObjSecretKey) > 0 {
		config["obj_secret_key"] = pc.ObjSecretKey
	}
	return config
}

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(tfProvider *schema.Provider) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		if pc.Spec.Configuration.ObjForceDelete && !pc.Spec.Configuration.ObjUseTempKeys && (pc.Spec.Configuration.ObjAccessKey == "" || pc.Spec.Configuration.ObjSecretKey == "") {
			return ps, errors.Wrap(err, "if obj_bucket_force_delete is set, then set obj_use_temp_keys or obj_access_key and obj_secret_key")
		}

		// set provider configuration
		ps.Configuration = prepareTerraformProviderConfiguration(creds, pc.Spec.Configuration)

		return ps, errors.Wrap(configureNoForkLinodeclient(ctx, &ps, *tfProvider), "failed to configure the no-fork linode client")
	}
}

func configureNoForkLinodeclient(ctx context.Context, ps *terraform.Setup, p schema.Provider) error {
	// Please be aware that this implementation relies on the schema.Provider
	// parameter `p` being a non-pointer. This is because normally
	// the Terraform plugin SDK normally configures the provider
	// only once and using a pointer argument here will cause
	// race conditions between resources referring to different
	// ProviderConfigs.
	diag := p.Configure(context.WithoutCancel(ctx), &terraformsdk.ResourceConfig{
		Config: ps.Configuration,
	})
	if diag != nil && diag.HasError() {
		return errors.Errorf("failed to configure the provider: %v", diag)
	}

	ps.Meta = p.Meta()
	p.Meta().(*helper.ProviderMeta).Client.SetUserAgent("crossplane-provider-linode")
	p.Meta().(*helper.ProviderMeta).Client.OnAfterResponse(apiResponseCounterMiddleware)
	fwProvider := linode.CreateFrameworkProviderWithMeta(version.ProviderVersion, p.Meta().(*helper.ProviderMeta))
	ps.FrameworkProvider = fwProvider
	return nil
}

func apiResponseCounterMiddleware(r *resty.Response) error {
	url, err := url.ParseRequestURI(r.Request.URL)
	if err != nil {
		return err
	}
	urlParts := strings.Split(strings.TrimLeft(url.Path, "/"), "/")
	service := ""
	if _, err := strconv.Atoi(urlParts[len(urlParts)-1]); err == nil {
		// if the last part is a integer, we want the piece in front of it.
		service = urlParts[len(urlParts)-2]
	} else {
		service = urlParts[len(urlParts)-1]
	}
	apiToken := strings.TrimPrefix(r.Request.Header.Get("Authorization"), "Bearer ")
	return metrics.IncLinodeAPIResp(service, r.Request.Method, fmt.Sprintf("%d", r.StatusCode()), apiToken[:5])
}
