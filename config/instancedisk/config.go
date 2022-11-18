package instancedisk

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_instance_disk", func(r *config.Resource) {

		r.References["image"] = config.Reference{
			Type: "github.com/linode/provider-linode/apis/image/v1alpha1.Image",
		}

	})
}
