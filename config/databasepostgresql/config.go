package databasepostgresql

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_database_postgresql", func(r *config.Resource) {
		r.UseAsync = true
		r.Kind = "PostgreSQL"
	})
	p.AddResourceConfigurator("linode_database_postgresql_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.Kind = "PostgreSQLv2"
		r.TerraformResource.Schema["updates"].Type = schema.TypeMap
	})
}
