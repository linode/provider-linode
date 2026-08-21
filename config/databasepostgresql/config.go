package databasepostgresql

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_database_postgresql_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.Kind = "PostgreSQLv2"
	})
}
