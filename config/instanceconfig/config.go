package instanceconfig

import (
	"fmt"

	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_instance_config", func(r *config.Resource) {
		r.References["linode_id"] = config.Reference{
			Type: "Instance",
		}
		for letter := 'a'; letter < 'i'; letter++ {
			r.References[fmt.Sprintf("devices.sd%c.disk_id", letter)] = config.Reference{
				Type: "Disk",
			}

			// This doesn't currently work because it introduces a circular dependency of
			// instance config -> volume -> instance -> instance config
			// r.References[fmt.Sprintf("devices.sd%c.volume_id", letter)] = config.Reference{
			// 	Type: "github.com/linode/provider-linode/apis/volume/v1alpha1.Volume",
			// }
		}

	})
}
