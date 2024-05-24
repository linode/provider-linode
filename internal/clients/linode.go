/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

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

		// set provider configuration
		ps.Configuration = map[string]any{}
		if v, ok := creds[keyToken]; ok {
			ps.Configuration[keyToken] = v
		}
		if v, ok := creds[keyURL]; ok {
			ps.Configuration[keyURL] = v
		}
		if v, ok := creds[keyVersion]; ok {
			ps.Configuration[keyVersion] = v
		}

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
	p.Meta().(*helper.ProviderMeta).Client.OnBeforeRequest(apiCallCounterMiddleware)
	fwProvider := linode.CreateFrameworkProviderWithMeta(version.ProviderVersion, p.Meta().(*helper.ProviderMeta))

	ps.FrameworkProvider = fwProvider
	return nil
}

func apiCallCounterMiddleware(r *resty.Request) error {
	metrics.IncLinodeAPICall(r.URL, r.Method)
	return nil
}
