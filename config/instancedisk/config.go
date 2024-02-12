package instancedisk

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_instance_disk", func(r *config.Resource) {

		r.References["linode_id"] = config.Reference{
			Type: "Instance",
		}

		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"disk", "config"},
		}

		r.References["stackscript_id"] = config.Reference{
			Type: "github.com/linode/provider-linode/apis/stackscript/v1alpha1.Stackscript",
		}

	})
}
