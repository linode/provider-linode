package firewalldevice

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_firewall_device", func(r *config.Resource) {
		r.References["entity_id"] = config.Reference{
			Type: "github.com/linode/provider-linode/apis/instance/v1alpha1.Instance",
		}

		r.References["firewall_id"] = config.Reference{
			Type: "github.com/linode/provider-linode/apis/firewall/v1alpha1.Device",
		}
	})
}
