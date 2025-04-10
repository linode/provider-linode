package instancereservedipassignment

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_reserved_ip_assignment", func(r *config.Resource) {

		r.References["linode_id"] = config.Reference{
			Type: "github.com/linode/provider-linode/apis/instance/v1alpha1.Instance",
		}

		r.References["rdns"] = config.Reference{
			Type: "github.com/linode/provider-linode/apis/rdns/v1alpha1.RDNS",
		}
	})
}
