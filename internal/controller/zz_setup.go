/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	accesscontrols "github.com/linode/provider-linode/internal/controller/database/accesscontrols"
	mongodb "github.com/linode/provider-linode/internal/controller/database/mongodb"
	mysql "github.com/linode/provider-linode/internal/controller/database/mysql"
	postgresql "github.com/linode/provider-linode/internal/controller/database/postgresql"
	record "github.com/linode/provider-linode/internal/controller/domain/record"
	device "github.com/linode/provider-linode/internal/controller/firewall/device"
	config "github.com/linode/provider-linode/internal/controller/instance/config"
	disk "github.com/linode/provider-linode/internal/controller/instance/disk"
	ip "github.com/linode/provider-linode/internal/controller/instance/ip"
	sharedips "github.com/linode/provider-linode/internal/controller/instance/sharedips"
	domain "github.com/linode/provider-linode/internal/controller/linode/domain"
	firewall "github.com/linode/provider-linode/internal/controller/linode/firewall"
	image "github.com/linode/provider-linode/internal/controller/linode/image"
	instance "github.com/linode/provider-linode/internal/controller/linode/instance"
	nodebalancer "github.com/linode/provider-linode/internal/controller/linode/nodebalancer"
	rdns "github.com/linode/provider-linode/internal/controller/linode/rdns"
	sshkey "github.com/linode/provider-linode/internal/controller/linode/sshkey"
	stackscript "github.com/linode/provider-linode/internal/controller/linode/stackscript"
	token "github.com/linode/provider-linode/internal/controller/linode/token"
	user "github.com/linode/provider-linode/internal/controller/linode/user"
	volume "github.com/linode/provider-linode/internal/controller/linode/volume"
	cluster "github.com/linode/provider-linode/internal/controller/lke/cluster"
	confignodebalancer "github.com/linode/provider-linode/internal/controller/nodebalancer/config"
	node "github.com/linode/provider-linode/internal/controller/nodebalancer/node"
	storagebucket "github.com/linode/provider-linode/internal/controller/object/storagebucket"
	storagekey "github.com/linode/provider-linode/internal/controller/object/storagekey"
	storageobject "github.com/linode/provider-linode/internal/controller/object/storageobject"
	providerconfig "github.com/linode/provider-linode/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesscontrols.Setup,
		mongodb.Setup,
		mysql.Setup,
		postgresql.Setup,
		record.Setup,
		device.Setup,
		config.Setup,
		disk.Setup,
		ip.Setup,
		sharedips.Setup,
		domain.Setup,
		firewall.Setup,
		image.Setup,
		instance.Setup,
		nodebalancer.Setup,
		rdns.Setup,
		sshkey.Setup,
		stackscript.Setup,
		token.Setup,
		user.Setup,
		volume.Setup,
		cluster.Setup,
		confignodebalancer.Setup,
		node.Setup,
		storagebucket.Setup,
		storagekey.Setup,
		storageobject.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
