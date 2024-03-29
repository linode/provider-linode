package rdns

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_rdns", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "linode"
		r.ShortGroup = "rdns"
		r.Kind = "RDNS"
		r.Path = "rdns"
	})
}
