package nodebalancernode

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_nodebalancer_node", func(r *config.Resource) {
		r.References["nodebalancer_id"] = config.Reference{
			Type: "Nodebalancer",
		}

		r.References["config_id"] = config.Reference{
			Type: "Config",
		}
	})
}
