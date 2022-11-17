package objectstoragebucket

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_object_storage_bucket", func(r *config.Resource) {
	})
}
