package databasemysql

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_database_mysql", func(r *config.Resource) {
		r.UseAsync = true
	})
	p.AddResourceConfigurator("linode_database_mysql_v2", func(r *config.Resource) {
		r.UseAsync = true
		r.Kind = "MySQLv2"
		r.TerraformResource.Schema["updates"].Type = schema.TypeMap
	})
}
