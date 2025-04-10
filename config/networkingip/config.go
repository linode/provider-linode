package networkingip

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_networking_ip", func(r *config.Resource) {

		r.References["linode_id"] = config.Reference{
			Type: "github.com/linode/provider-linode/apis/instance/v1alpha1.Instance",
		}

		r.References["rdns"] = config.Reference{
			Type: "github.com/linode/provider-linode/apis/rdns/v1alpha1.RDNS",
		}

		r.TerraformResource.Schema["vpc_nat_1_1"].Type = schema.TypeMap
	})
}
