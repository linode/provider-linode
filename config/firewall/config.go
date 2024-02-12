package firewall

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_firewall", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "linode"
		r.ShortGroup = "firewall"
		r.References["linodes"] = config.Reference{
			Type: "github.com/linode/provider-linode/apis/instance/v1alpha1.Instance",
		}
	})
}
