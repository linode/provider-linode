package databaseaccesscontrols

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_database_access_controls", func(r *config.Resource) {
		// this field references the Mysql, Mongodb and Postgresql objects,
		// so we can't set this reference to a single model right now
		// https://github.com/upbound/upjet/issues/95
		//  r.References["database_id"] = config.Reference{
		//	  Type: "Mysql",
		//  }

	})
}
