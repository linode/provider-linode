/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/linode/provider-linode/config/domain"
	"github.com/linode/provider-linode/config/domain_record"
	"github.com/linode/provider-linode/config/firewall"
	"github.com/linode/provider-linode/config/firewall_device"
	"github.com/linode/provider-linode/config/image"
	"github.com/linode/provider-linode/config/instance"
	"github.com/linode/provider-linode/config/lke_cluster"
	"github.com/linode/provider-linode/config/stackscript"
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
		domain.Configure,
		domain_record.Configure,
		firewall.Configure,
		firewall_device.Configure,
		image.Configure,
		instance.Configure,
		stackscript.Configure,
		lke_cluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
