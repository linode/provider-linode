/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	"context"
	_ "embed"

	"github.com/crossplane/upjet/pkg/config"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/linode/terraform-provider-linode/v4/linode"
	"github.com/linode/terraform-provider-linode/v4/version"
	"github.com/pkg/errors"

	"github.com/linode/provider-linode/config/accountsettings"
	"github.com/linode/provider-linode/config/databaseaccesscontrols"
	"github.com/linode/provider-linode/config/databasemysql"
	"github.com/linode/provider-linode/config/databasepostgresql"
	"github.com/linode/provider-linode/config/domain"
	"github.com/linode/provider-linode/config/domainrecord"
	"github.com/linode/provider-linode/config/firewall"
	"github.com/linode/provider-linode/config/firewalldevice"
	"github.com/linode/provider-linode/config/image"
	"github.com/linode/provider-linode/config/instance"
	"github.com/linode/provider-linode/config/instanceconfig"
	"github.com/linode/provider-linode/config/instancedisk"
	"github.com/linode/provider-linode/config/instanceip"
	"github.com/linode/provider-linode/config/instancereservedipassignment"
	"github.com/linode/provider-linode/config/instancesharedips"
	"github.com/linode/provider-linode/config/ipv6range"
	"github.com/linode/provider-linode/config/lkecluster"
	"github.com/linode/provider-linode/config/lkenodepool"
	"github.com/linode/provider-linode/config/networkingip"
	"github.com/linode/provider-linode/config/nodebalancer"
	"github.com/linode/provider-linode/config/nodebalancerconfig"
	"github.com/linode/provider-linode/config/nodebalancernode"
	"github.com/linode/provider-linode/config/objectstoragebucket"
	"github.com/linode/provider-linode/config/objectstoragekey"
	"github.com/linode/provider-linode/config/objectstorageobject"
	"github.com/linode/provider-linode/config/placementgroup"
	"github.com/linode/provider-linode/config/placementgroupassignment"
	"github.com/linode/provider-linode/config/rdns"
	"github.com/linode/provider-linode/config/sshkey"
	"github.com/linode/provider-linode/config/stackscript"
	"github.com/linode/provider-linode/config/token"
	"github.com/linode/provider-linode/config/user"
	"github.com/linode/provider-linode/config/vpc"
	"github.com/linode/provider-linode/config/vpcsubnet"

	"github.com/linode/provider-linode/config/volume"
)

const (
	resourcePrefix = "linode"
	modulePath     = "github.com/linode/provider-linode"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

func externalNameConfig() config.ResourceOption {
	return func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	}
}

func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}

// resourceBlocks returns the raw terraform-json schema block of every resource,
// keyed by Terraform resource name.
func resourceBlocks(s string) map[string]*tfjson.SchemaBlock {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	blocks := map[string]*tfjson.SchemaBlock{}
	for _, v := range ps.Schemas {
		for name, rs := range v.ResourceSchemas {
			if rs != nil {
				blocks[name] = rs.Block
			}
		}
		break
	}
	return blocks
}

// untypedAttributeConfig repairs the attributes that upjet's tfjson to SDKv2
// converter leaves without a type, which the code generator later rejects with
// "invalid schema type TypeInvalid".
//
// The converter only understands attributes carrying a primitive or collection
// cty type, plus nested *blocks*. Terraform Plugin Framework resources have
// two more shapes it does not cover:
//
//   - nested attributes ("nested_type" in the JSON schema), used by among
//     others linode_vpc.ipv6, linode_nodebalancer.vpcs and
//     linode_database_*_v2.updates;
//   - attributes whose cty type is a bare object, used by
//     linode_networking_ip.assigned_entity, its vpc_nat_1_1 and
//     linode_reserved_ip_assignment.assigned_entity.
//
// This runs as a default resource option so that it covers resources
// reconciled through both the Plugin SDK and the Plugin Framework: upjet
// derives the former from config.Provider.TerraformProvider but the latter
// from its own internal conversion, and only Resource.TerraformResource is
// common to both.
func untypedAttributeConfig(blocks map[string]*tfjson.SchemaBlock) config.ResourceOption {
	return func(r *config.Resource) {
		if r.TerraformResource == nil {
			return
		}
		resolveUntypedAttributes(blocks[r.Name], r.TerraformResource)
	}
}

// resolveUntypedAttributes fills in SDKv2 schemas for the attributes of b that
// the converter could not type. Entries that already carry a valid type are
// left alone, so that a hand-written Go schema always wins over the
// reconstructed one.
//
// We deliberately do not route nested attributes through the converter's
// nested block handling: that path infers Optional/Computed from
// MinItems/MaxItems, because blocks carry no such flags, and therefore cannot
// represent a computed-only field like linode_nodebalancer.lke_cluster. Nested
// attributes do carry the flags, so we honour them directly.
func resolveUntypedAttributes(b *tfjson.SchemaBlock, r *schema.Resource) {
	if b == nil || r == nil || r.Schema == nil {
		return
	}
	for name, attr := range b.Attributes {
		if existing, ok := r.Schema[name]; ok && existing != nil && existing.Type != schema.TypeInvalid {
			continue
		}
		if s := attributeToV2Schema(attr); s != nil {
			r.Schema[name] = s
		}
	}
	// Untyped attributes may also appear underneath a nested block.
	for name, nb := range b.NestedBlocks {
		if nb == nil {
			continue
		}
		s, ok := r.Schema[name]
		if !ok || s == nil {
			continue
		}
		if elem, ok := s.Elem.(*schema.Resource); ok {
			resolveUntypedAttributes(nb.Block, elem)
		}
	}
}

// attributeToV2Schema converts a nested or object-typed attribute into an
// SDKv2 schema. It returns nil for any other attribute, which the converter
// already handles.
func attributeToV2Schema(attr *tfjson.SchemaAttribute) *schema.Schema {
	if attr == nil {
		return nil
	}
	nt := attr.AttributeNestedType
	if nt == nil && !attr.AttributeType.IsObjectType() {
		return nil
	}

	s := &schema.Schema{
		Description: attr.Description,
		Required:    attr.Required,
		Optional:    attr.Optional,
		Computed:    attr.Computed,
		Sensitive:   attr.Sensitive,
		ConfigMode:  schema.SchemaConfigModeAttr,
	}
	if attr.Deprecated {
		s.Deprecated = "deprecated"
	}

	if nt == nil {
		// A bare object is a single element, so it is modelled the same way
		// the converter models a "single" nested block. Its fields are
		// presented as plain attributes inheriting the object's own flags,
		// which the SDKv2 schema has no way to express per field.
		s.Type = schema.TypeList
		s.MaxItems = 1
		fields := map[string]*tfjson.SchemaAttribute{}
		for name, t := range attr.AttributeType.AttributeTypes() {
			fields[name] = &tfjson.SchemaAttribute{
				AttributeType: t,
				Optional:      attr.Optional,
				Computed:      attr.Computed,
			}
		}
		s.Elem = elemResource(fields)
		return s
	}

	// Nesting modes follow the same conventions the upjet converter uses for
	// nested blocks, so that "single" becomes a one-element list.
	switch nt.NestingMode {
	case tfjson.SchemaNestingModeSet:
		s.Type = schema.TypeSet
	case tfjson.SchemaNestingModeMap:
		s.Type = schema.TypeMap
	case tfjson.SchemaNestingModeList:
		s.Type = schema.TypeList
	case tfjson.SchemaNestingModeSingle, tfjson.SchemaNestingModeGroup:
		// These hold exactly one element.
		s.Type = schema.TypeList
		s.MaxItems = 1
	default:
		s.Type = schema.TypeList
		s.MaxItems = 1
	}
	if nt.MinItems > 0 {
		s.MinItems = int(nt.MinItems) //nolint:gosec
	}
	if nt.MaxItems > 0 {
		s.MaxItems = int(nt.MaxItems) //nolint:gosec
	}
	s.Elem = elemResource(nt.Attributes)

	return s
}

// elemResource builds the element schema of a nested or object-typed
// attribute. It reuses upjet's converter for the attributes it can type by
// presenting them as a standalone resource schema, then recurses for the rest.
func elemResource(attrs map[string]*tfjson.SchemaAttribute) *schema.Resource {
	elem := &schema.Resource{Schema: map[string]*schema.Schema{}}
	if converted := conversiontfjson.GetV2ResourceMap(map[string]*tfjson.Schema{
		"": {Block: &tfjson.SchemaBlock{Attributes: attrs}},
	})[""]; converted != nil && converted.Schema != nil {
		elem.Schema = converted.Schema
	}
	for name, attr := range attrs {
		if existing, ok := elem.Schema[name]; ok && existing != nil && existing.Type != schema.TypeInvalid {
			continue
		}
		if s := attributeToV2Schema(attr); s != nil {
			elem.Schema[name] = s
		}
	}
	return elem
}

// GetProvider returns provider configuration
func GetProvider(_ context.Context, generationProvider bool) (*config.Provider, error) {
	var p *schema.Provider
	var err error

	if generationProvider {
		p, err = getProviderSchema(providerSchema)
	} else {
		p = linode.Provider()
	}
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get the Terraform provider schema with generation mode set to %t", generationProvider)
	}

	fwProvider := linode.CreateFrameworkProvider(version.ProviderVersion)

	pc := config.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		config.WithDefaultResourceOptions(
			externalNameConfig(),
			resourceConfigurator(),
			untypedAttributeConfig(resourceBlocks(providerSchema)),
		),
		config.WithFeaturesPackage("internal/features"),
		config.WithIncludeList(resourceList(cliReconciledExternalNameConfigs)),
		config.WithTerraformPluginSDKIncludeList(resourceList(terraformSDKIncludeList)),
		config.WithTerraformPluginFrameworkIncludeList(resourceList(terraformPluginFrameworkExternalNameConfigs)),
		config.WithTerraformProvider(p),
		config.WithTerraformPluginFrameworkProvider(fwProvider),
	)

	for _, configure := range []func(provider *config.Provider){
		// add custom config functions
		accountsettings.Configure,
		databaseaccesscontrols.Configure,
		databasemysql.Configure,
		databasepostgresql.Configure,
		domain.Configure,
		domainrecord.Configure,
		firewall.Configure,
		firewalldevice.Configure,
		image.Configure,
		instance.Configure,
		instanceconfig.Configure,
		instancedisk.Configure,
		instanceip.Configure,
		instancesharedips.Configure,
		instancereservedipassignment.Configure,
		ipv6range.Configure,
		lkecluster.Configure,
		lkenodepool.Configure,
		networkingip.Configure,
		nodebalancer.Configure,
		nodebalancerconfig.Configure,
		nodebalancernode.Configure,
		objectstoragebucket.Configure,
		objectstoragekey.Configure,
		objectstorageobject.Configure,
		placementgroup.Configure,
		placementgroupassignment.Configure,
		rdns.Configure,
		sshkey.Configure,
		stackscript.Configure,
		token.Configure,
		user.Configure,
		volume.Configure,
		vpc.Configure,
		vpcsubnet.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}

// resourceList returns the list of resources that have external
// name configured in the specified table.
func resourceList(t map[string]config.ExternalName) []string {
	l := make([]string, len(t))
	i := 0
	for n := range t {
		// Expected format is regex and we'd like to have exact matches.
		l[i] = n + "$"
		i++
	}
	return l
}
