package vpcsubnet

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_vpc_subnet", func(r *config.Resource) {
		r.References["vpc_id"] = config.Reference{
			Type: "github.com/linode/provider-linode/apis/vpc/v1alpha1.VPC",
		}
	})
}
