package lkecluster

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_lke_cluster", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider

	})
}
