package databasemysql

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_database_mysql", func(r *config.Resource) {
		r.UseAsync = true
	})
}
