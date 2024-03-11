package objectstoragebucket

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_object_storage_bucket", func(r *config.Resource) {

		r.ShortGroup = "objectstorage"
		r.Kind = "Bucket"
		r.References["access_key"] = config.Reference{
			Type:              "Key",
			RefFieldName:      "AccessKeyRef",
			SelectorFieldName: "AccessKeySelector",
		}
		r.References["secret_key"] = config.Reference{
			Type:              "Key",
			RefFieldName:      "SecretKeyRef",
			SelectorFieldName: "SecretKeySelector",
		}
	})
}
