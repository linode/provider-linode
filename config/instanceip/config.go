package instanceip

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_instance_ip", func(r *config.Resource) {

		r.References["linode_id"] = config.Reference{
			Type: "Instance",
		}

		r.References["rdns"] = config.Reference{
			Type: "github.com/linode/provider-linode/apis/rdns/v1alpha1.RDNS",
		}
	})
}
