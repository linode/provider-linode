/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	accesscontrols "github.com/linode/provider-linode/internal/controller/database_access_controls/accesscontrols"
	mongodb "github.com/linode/provider-linode/internal/controller/database_mongodb/mongodb"
	mysql "github.com/linode/provider-linode/internal/controller/database_mysql/mysql"
	postgresql "github.com/linode/provider-linode/internal/controller/database_postgresql/postgresql"
	domain "github.com/linode/provider-linode/internal/controller/domain/domain"
	record "github.com/linode/provider-linode/internal/controller/domain_record/record"
	firewall "github.com/linode/provider-linode/internal/controller/firewall/firewall"
	device "github.com/linode/provider-linode/internal/controller/firewall_device/device"
	image "github.com/linode/provider-linode/internal/controller/image/image"
	instance "github.com/linode/provider-linode/internal/controller/instance/instance"
	config "github.com/linode/provider-linode/internal/controller/instance_config/config"
	disk "github.com/linode/provider-linode/internal/controller/instance_disk/disk"
	ip "github.com/linode/provider-linode/internal/controller/instance_ip/ip"
	sharedips "github.com/linode/provider-linode/internal/controller/instance_shared_ips/sharedips"
	cluster "github.com/linode/provider-linode/internal/controller/lke_cluster/cluster"
	nodebalancer "github.com/linode/provider-linode/internal/controller/nodebalancer/nodebalancer"
	confignodebalancer_config "github.com/linode/provider-linode/internal/controller/nodebalancer_config/config"
	node "github.com/linode/provider-linode/internal/controller/nodebalancer_node/node"
	storagebucket "github.com/linode/provider-linode/internal/controller/object_storage_bucket/storagebucket"
	storagekey "github.com/linode/provider-linode/internal/controller/object_storage_key/storagekey"
	storageobject "github.com/linode/provider-linode/internal/controller/object_storage_object/storageobject"
	providerconfig "github.com/linode/provider-linode/internal/controller/providerconfig"
	rdns "github.com/linode/provider-linode/internal/controller/rdns/rdns"
	sshkey "github.com/linode/provider-linode/internal/controller/sshkey/sshkey"
	stackscript "github.com/linode/provider-linode/internal/controller/stackscript/stackscript"
	token "github.com/linode/provider-linode/internal/controller/token/token"
	user "github.com/linode/provider-linode/internal/controller/user/user"
	volume "github.com/linode/provider-linode/internal/controller/volume/volume"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesscontrols.Setup,
		mongodb.Setup,
		mysql.Setup,
		postgresql.Setup,
		domain.Setup,
		record.Setup,
		firewall.Setup,
		device.Setup,
		image.Setup,
		instance.Setup,
		config.Setup,
		disk.Setup,
		ip.Setup,
		sharedips.Setup,
		cluster.Setup,
		nodebalancer.Setup,
		confignodebalancer_config.Setup,
		node.Setup,
		storagebucket.Setup,
		storagekey.Setup,
		storageobject.Setup,
		providerconfig.Setup,
		rdns.Setup,
		sshkey.Setup,
		stackscript.Setup,
		token.Setup,
		user.Setup,
		volume.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
