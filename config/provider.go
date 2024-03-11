/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/linode/provider-linode/config/ipv6range"

	ujconfig "github.com/crossplane/upjet/pkg/config"

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
	"github.com/linode/provider-linode/config/instancesharedips"
	"github.com/linode/provider-linode/config/lkecluster"
	"github.com/linode/provider-linode/config/nodebalancer"
	"github.com/linode/provider-linode/config/nodebalancerconfig"
	"github.com/linode/provider-linode/config/nodebalancernode"
	"github.com/linode/provider-linode/config/objectstoragebucket"
	"github.com/linode/provider-linode/config/objectstoragekey"
	"github.com/linode/provider-linode/config/objectstorageobject"
	"github.com/linode/provider-linode/config/rdns"
	"github.com/linode/provider-linode/config/sshkey"
	"github.com/linode/provider-linode/config/stackscript"
	"github.com/linode/provider-linode/config/token"
	"github.com/linode/provider-linode/config/user"
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

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
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
		ipv6range.Configure,
		lkecluster.Configure,
		nodebalancer.Configure,
		nodebalancerconfig.Configure,
		nodebalancernode.Configure,
		objectstoragebucket.Configure,
		objectstoragekey.Configure,
		objectstorageobject.Configure,
		rdns.Configure,
		sshkey.Configure,
		stackscript.Configure,
		token.Configure,
		user.Configure,
		volume.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
