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
		r.TerraformResource.Schema["updates"] = &schema.Schema{
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"day_of_week": {
						Type:     schema.TypeFloat,
						Optional: true,
						Computed: true,
					},
					"duration": {
						Type:     schema.TypeFloat,
						Optional: true,
						Computed: true,
					},
					"frequency": {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
					"hour_of_day": {
						Type:     schema.TypeFloat,
						Optional: true,
						Computed: true,
					},
				},
			},
		}
	})
}
