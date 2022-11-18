package instance

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_instance", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "linode"
		r.ShortGroup = "instance"
		r.UseAsync = true

		r.References["image"] = config.Reference{
			Type: "github.com/linode/provider-linode/apis/image/v1alpha1.Image",
		}
	})
}
