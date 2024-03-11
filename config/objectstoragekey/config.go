package objectstoragekey

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_object_storage_key", func(r *config.Resource) {

		r.ShortGroup = "objectstorage"
		r.Kind = "Key"
		r.TerraformResource.Schema["access_key"].Sensitive = true
	})
}
