package nodebalancerconfig

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_nodebalancer_config", func(r *config.Resource) {
		r.References["nodebalancer_id"] = config.Reference{
			Type: "Nodebalancer",
		}
	})
}
