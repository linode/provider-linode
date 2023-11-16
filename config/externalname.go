/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var noForkExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: 1234567
	"linode_database_access_controls": config.IdentifierFromProvider,
	"linode_database_mongodb":         config.IdentifierFromProvider,
	"linode_database_mysql":           config.IdentifierFromProvider,
	"linode_database_postgresql":      config.IdentifierFromProvider,
	"linode_domain":                   config.IdentifierFromProvider,
	"linode_domain_record":            config.IdentifierFromProvider,
	"linode_firewall":                 config.IdentifierFromProvider,
	"linode_firewall_device":          config.IdentifierFromProvider,
	"linode_image":                    config.IdentifierFromProvider,
	"linode_instance":                 config.IdentifierFromProvider,
	"linode_instance_config":          config.IdentifierFromProvider,
	"linode_instance_disk":            config.IdentifierFromProvider,
	"linode_instance_ip":              config.IdentifierFromProvider,
	"linode_instance_shared_ips":      config.IdentifierFromProvider,
	"linode_ipv6_range":               config.IdentifierFromProvider,
	"linode_lke_cluster":              config.IdentifierFromProvider,
	"linode_nodebalancer":             config.IdentifierFromProvider,
	"linode_nodebalancer_config":      config.IdentifierFromProvider,
	"linode_nodebalancer_node":        config.IdentifierFromProvider,
	"linode_object_storage_bucket":    config.IdentifierFromProvider,
	"linode_object_storage_key":       config.IdentifierFromProvider,
	"linode_object_storage_object":    config.IdentifierFromProvider,
	"linode_rdns":                     config.IdentifierFromProvider,
	"linode_sshkey":                   config.IdentifierFromProvider,
	"linode_stackscript":              config.IdentifierFromProvider,
	"linode_token":                    config.IdentifierFromProvider,
	"linode_user":                     config.IdentifierFromProvider,
	"linode_volume":                   config.IdentifierFromProvider,
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

// resourceConfigurator applies all external name configs
// listed in the table NoForkExternalNameConfigs and
// cliReconciledExternalNameConfigs and sets the version
// of those resources to v1alpha1. For those resource in
// noForkExternalNameConfigs, it also sets
// config.Resource.UseNoForkClient to `true`.
func resourceConfigurator() config.ResourceOption {
	return func(r *config.Resource) {
		// if configured both for the no-fork and CLI based architectures,
		// no-fork configuration prevails
		e, configured := noForkExternalNameConfigs[r.Name]
		if !configured {
			e, configured = cliReconciledExternalNameConfigs[r.Name]
		}
		if !configured {
			return
		}
		r.Version = "v1alpha1"
		r.ExternalName = e
	}
}
