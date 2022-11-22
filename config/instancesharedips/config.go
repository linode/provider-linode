package instancesharedips

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_instance_shared_ips", func(r *config.Resource) {
		r.Kind = "SharedIPs"
	})
}
