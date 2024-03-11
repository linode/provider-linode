package ipv6range

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_ipv6_range", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "linode"
		r.Kind = "IPv6Range"
		r.References["linode_id"] = config.Reference{
			Type: "github.com/linode/provider-linode/apis/instance/v1alpha1.Instance",
		}
	})
}
