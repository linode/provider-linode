package domainrecord

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_domain_record", func(r *config.Resource) {
		r.References["domain_id"] = config.Reference{
			Type: "Domain",
		}
	})
}
