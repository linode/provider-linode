/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"github.com/crossplane/upjet/pkg/config"
)

// terraformPluginFrameworkExternalNameConfigs contains all external
// name configurations belonging to Terraform Plugin Framework
// resources to be reconciled under the no-fork architecture for this
// provider.
var terraformPluginFrameworkExternalNameConfigs = map[string]config.ExternalName{
	"linode_token":                      config.IdentifierFromProvider,
	"linode_stackscript":                config.IdentifierFromProvider,
	"linode_rdns":                       config.IdentifierFromProvider,
	"linode_object_storage_key":         config.IdentifierFromProvider,
	"linode_sshkey":                     config.IdentifierFromProvider,
	"linode_ipv6_range":                 config.IdentifierFromProvider,
	"linode_nodebalancer":               config.IdentifierFromProvider,
	"linode_account_settings":           config.IdentifierFromProvider,
	"linode_vpc_subnet":                 config.IdentifierFromProvider,
	"linode_vpc":                        config.IdentifierFromProvider,
	"linode_instance_ip":                config.IdentifierFromProvider,
	"linode_firewall_device":            config.IdentifierFromProvider,
	"linode_volume":                     config.IdentifierFromProvider,
	"linode_instance_shared_ips":        config.IdentifierFromProvider,
	"linode_instance_disk":              config.IdentifierFromProvider,
	"linode_lke_node_pool":              config.IdentifierFromProvider,
	"linode_image":                      config.IdentifierFromProvider,
	"linode_nodebalancer_config":        config.IdentifierFromProvider,
	"linode_nodebalancer_node":          config.IdentifierFromProvider,
	"linode_firewall":                   config.IdentifierFromProvider,
	"linode_object_storage_object":      config.IdentifierFromProvider,
	"linode_placement_group":            config.IdentifierFromProvider,
	"linode_placement_group_assignment": config.IdentifierFromProvider,
}

// terraformSDKIncludeList contains all external name configurations
// belonging to Terraform Plugin SDKv2 resources to be reconciled
// under the no-fork architecture for this provider.
var terraformSDKIncludeList = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: 1234567
	"linode_database_access_controls": config.IdentifierFromProvider,
	"linode_database_mysql":           config.IdentifierFromProvider,
	"linode_database_postgresql":      config.IdentifierFromProvider,
	"linode_domain":                   config.IdentifierFromProvider,
	"linode_domain_record":            config.IdentifierFromProvider,
	"linode_instance":                 config.IdentifierFromProvider,
	"linode_instance_config":          config.IdentifierFromProvider,
	"linode_lke_cluster":              config.IdentifierFromProvider,
	"linode_object_storage_bucket":    config.IdentifierFromProvider,
	"linode_user":                     config.IdentifierFromProvider,
}

// cliReconciledExternalNameConfigs contains all external name configurations
// belonging to Terraform resources to be reconciled under the CLI-based
// architecture for this provider.
var cliReconciledExternalNameConfigs = map[string]config.ExternalName{}

// TemplatedStringAsIdentifierWithNoName uses TemplatedStringAsIdentifier but
// without the name initializer. This allows it to be used in cases where the ID
// is constructed with parameters and a provider-defined value, meaning no
// user-defined input. Since the external name is not user-defined, the name
// initializer has to be disabled.
func TemplatedStringAsIdentifierWithNoName(tmpl string) config.ExternalName {
	e := config.TemplatedStringAsIdentifier("", tmpl)
	e.DisableNameInitializer = true
	return e
}

// ResourceConfigurator applies all external name configs listed in
// the table NoForkExternalNameConfigs,
// CLIReconciledExternalNameConfigs, and
// TerraformPluginFrameworkExternalNameConfigs and sets the version of
// those resources to v1alpha1.
func resourceConfigurator() config.ResourceOption {
	return func(r *config.Resource) {
		// If an external name is configured for multiple architectures,
		// Terraform Plugin Framework takes precedence over Terraform
		// Plugin SDKv2, which takes precedence over CLI architecture.
		e, configured := terraformPluginFrameworkExternalNameConfigs[r.Name]
		if !configured {
			e, configured = terraformSDKIncludeList[r.Name]
			if !configured {
				e, configured = cliReconciledExternalNameConfigs[r.Name]
			}
		}
		if !configured {
			return
		}
		r.Version = "v1alpha1"
		r.ExternalName = e
	}
}
