/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	_domain "github.com/linode/provider-linode/internal/controller/domain/domain"
	_record "github.com/linode/provider-linode/internal/controller/domain_record/record"
	_firewall "github.com/linode/provider-linode/internal/controller/firewall/firewall"
	_device "github.com/linode/provider-linode/internal/controller/firewall_device/device"
	_image "github.com/linode/provider-linode/internal/controller/image/image"
	_instance "github.com/linode/provider-linode/internal/controller/instance/instance"
	_config "github.com/linode/provider-linode/internal/controller/instance_config/config"
	_disk "github.com/linode/provider-linode/internal/controller/instance_disk/disk"
	_ip "github.com/linode/provider-linode/internal/controller/instance_ip/ip"
	_sharedips "github.com/linode/provider-linode/internal/controller/instance_shared_ips/sharedips"
	_cluster "github.com/linode/provider-linode/internal/controller/lke_cluster/cluster"
	_nodebalancer "github.com/linode/provider-linode/internal/controller/nodebalancer/nodebalancer"
	_confignodebalancer_config "github.com/linode/provider-linode/internal/controller/nodebalancer_config/config"
	_node "github.com/linode/provider-linode/internal/controller/nodebalancer_node/node"
	_providerconfig "github.com/linode/provider-linode/internal/controller/providerconfig"
	_stackscript "github.com/linode/provider-linode/internal/controller/stackscript/stackscript"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		_domain.Setup,
		_record.Setup,
		_firewall.Setup,
		_device.Setup,
		_image.Setup,
		_instance.Setup,
		_config.Setup,
		_disk.Setup,
		_ip.Setup,
		_sharedips.Setup,
		_cluster.Setup,
		_nodebalancer.Setup,
		_confignodebalancer_config.Setup,
		_node.Setup,
		_providerconfig.Setup,
		_stackscript.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
