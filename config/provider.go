/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	"context"
	_ "embed"

	"github.com/crossplane/upjet/pkg/config"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/linode/terraform-provider-linode/v2/linode"
	"github.com/linode/terraform-provider-linode/v2/version"
	"github.com/pkg/errors"

	"github.com/linode/provider-linode/config/databaseaccesscontrols"
	"github.com/linode/provider-linode/config/databasemysql"
	"github.com/linode/provider-linode/config/databasepostgresql"
	"github.com/linode/provider-linode/config/domain"
	"github.com/linode/provider-linode/config/domainrecord"
	"github.com/linode/provider-linode/config/firewall"
	"github.com/linode/provider-linode/config/firewalldevice"
	"github.com/linode/provider-linode/config/image"
	"github.com/linode/provider-linode/config/instance"
	"github.com/linode/provider-linode/config/instanceconfig"
	"github.com/linode/provider-linode/config/instancedisk"
	"github.com/linode/provider-linode/config/instanceip"
	"github.com/linode/provider-linode/config/instancesharedips"
	"github.com/linode/provider-linode/config/ipv6range"
	"github.com/linode/provider-linode/config/lkecluster"
	"github.com/linode/provider-linode/config/nodebalancer"
	"github.com/linode/provider-linode/config/nodebalancerconfig"
	"github.com/linode/provider-linode/config/nodebalancernode"
	"github.com/linode/provider-linode/config/objectstoragebucket"
	"github.com/linode/provider-linode/config/objectstoragekey"
	"github.com/linode/provider-linode/config/objectstorageobject"
	"github.com/linode/provider-linode/config/rdns"
	"github.com/linode/provider-linode/config/sshkey"
	"github.com/linode/provider-linode/config/stackscript"
	"github.com/linode/provider-linode/config/token"
	"github.com/linode/provider-linode/config/user"
	"github.com/linode/provider-linode/config/volume"
	"github.com/linode/provider-linode/config/vpc"
	"github.com/linode/provider-linode/config/vpcsubnet"
)

const (
	resourcePrefix = "linode"
	modulePath     = "github.com/linode/provider-linode"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

func externalNameConfig() config.ResourceOption {
	return func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	}
}

func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// GetProvider returns provider configuration
func GetProvider(_ context.Context, generationProvider bool) (*config.Provider, error) {
	var p *schema.Provider
	var err error

	if generationProvider {
		p, err = getProviderSchema(providerSchema)
	} else {
		p = linode.Provider()
	}
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get the Terraform provider schema with generation mode set to %t", generationProvider)
	}

	fwProvider := linode.CreateFrameworkProvider(version.ProviderVersion)

	pc := config.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		config.WithDefaultResourceOptions(
			externalNameConfig(),
			resourceConfigurator(),
		),
		config.WithIncludeList(resourceList(cliReconciledExternalNameConfigs)),
		config.WithTerraformPluginSDKIncludeList(resourceList(terraformPluginSDKIncludeList)),
		config.WithTerraformPluginFrameworkIncludeList(resourceList(terraformPluginFrameworkExternalNameConfigs)),
		config.WithTerraformProvider(p),
		config.WithTerraformPluginFrameworkProvider(fwProvider),
	)

	for _, configure := range []func(provider *config.Provider){
		// add custom config functions
		databaseaccesscontrols.Configure,
		databasemysql.Configure,
		databasepostgresql.Configure,
		domain.Configure,
		domainrecord.Configure,
		firewall.Configure,
		firewalldevice.Configure,
		image.Configure,
		instance.Configure,
		instanceconfig.Configure,
		instancedisk.Configure,
		instanceip.Configure,
		instancesharedips.Configure,
		ipv6range.Configure,
		lkecluster.Configure,
		nodebalancer.Configure,
		nodebalancerconfig.Configure,
		nodebalancernode.Configure,
		objectstoragebucket.Configure,
		objectstoragekey.Configure,
		objectstorageobject.Configure,
		rdns.Configure,
		sshkey.Configure,
		stackscript.Configure,
		token.Configure,
		user.Configure,
		volume.Configure,
		vpc.Configure,
		vpcsubnet.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}

// resourceList returns the list of resources that have external
// name configured in the specified table.
func resourceList(t map[string]config.ExternalName) []string {
	l := make([]string, len(t))
	i := 0
	for n := range t {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = n + "$"
		i++
	}
	return l
}
