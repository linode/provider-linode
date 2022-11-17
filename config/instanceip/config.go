package instanceip

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_instance_ip", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider

	})
}
