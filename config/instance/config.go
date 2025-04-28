package instance

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_instance", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "linode"
		r.ShortGroup = "instance"
		r.UseAsync = true

		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"disk", "config", "ipv4"},
		}

		r.References["stackscript_id"] = config.Reference{
			Type: "github.com/linode/provider-linode/apis/stackscript/v1alpha1.Stackscript",
		}

	})
}
