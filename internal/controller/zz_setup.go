/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	_accesscontrols "github.com/linode/provider-linode/internal/controller/database_access_controls/accesscontrols"
	_mongodb "github.com/linode/provider-linode/internal/controller/database_mongodb/mongodb"
	_mysql "github.com/linode/provider-linode/internal/controller/database_mysql/mysql"
	_postgresql "github.com/linode/provider-linode/internal/controller/database_postgresql/postgresql"
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
	_storagebucket "github.com/linode/provider-linode/internal/controller/object_storage_bucket/storagebucket"
	_storagekey "github.com/linode/provider-linode/internal/controller/object_storage_key/storagekey"
	_storageobject "github.com/linode/provider-linode/internal/controller/object_storage_object/storageobject"
	_providerconfig "github.com/linode/provider-linode/internal/controller/providerconfig"
	_rdns "github.com/linode/provider-linode/internal/controller/rdns/rdns"
	_sshkey "github.com/linode/provider-linode/internal/controller/sshkey/sshkey"
	_stackscript "github.com/linode/provider-linode/internal/controller/stackscript/stackscript"
	_token "github.com/linode/provider-linode/internal/controller/token/token"
	_user "github.com/linode/provider-linode/internal/controller/user/user"
	_volume "github.com/linode/provider-linode/internal/controller/volume/volume"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		_accesscontrols.Setup,
		_mongodb.Setup,
		_mysql.Setup,
		_postgresql.Setup,
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
		_storagebucket.Setup,
		_storagekey.Setup,
		_storageobject.Setup,
		_providerconfig.Setup,
		_rdns.Setup,
		_sshkey.Setup,
		_stackscript.Setup,
		_token.Setup,
		_user.Setup,
		_volume.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
