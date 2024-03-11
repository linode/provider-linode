/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
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

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
