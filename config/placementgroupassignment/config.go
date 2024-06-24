package placementgroupassignment

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("linode_placement_group_assignment", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "linode"
		r.ShortGroup = "placementgroupassignment"

		// Overriding "GroupAssignment" to "PlacementGroupAssignment" to make it less generic
		r.Kind = "PlacementGroupAssignment"
	})
}
