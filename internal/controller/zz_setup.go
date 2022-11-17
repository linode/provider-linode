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
	domain "github.com/linode/provider-linode/internal/controller/domain/domain"
	record "github.com/linode/provider-linode/internal/controller/domain/record"
	device "github.com/linode/provider-linode/internal/controller/firewall/device"
	firewall "github.com/linode/provider-linode/internal/controller/firewall/firewall"
	config "github.com/linode/provider-linode/internal/controller/instance/config"
	disk "github.com/linode/provider-linode/internal/controller/instance/disk"
	instance "github.com/linode/provider-linode/internal/controller/instance/instance"
	ip "github.com/linode/provider-linode/internal/controller/instance/ip"
	sharedips "github.com/linode/provider-linode/internal/controller/instance/sharedips"
	cluster "github.com/linode/provider-linode/internal/controller/lke/cluster"
	confignodebalancer "github.com/linode/provider-linode/internal/controller/nodebalancer/config"
	image "github.com/linode/provider-linode/internal/controller/nodebalancer/image"
	node "github.com/linode/provider-linode/internal/controller/nodebalancer/node"
	nodebalancer "github.com/linode/provider-linode/internal/controller/nodebalancer/nodebalancer"
	storagebucket "github.com/linode/provider-linode/internal/controller/object/storagebucket"
	storagekey "github.com/linode/provider-linode/internal/controller/object/storagekey"
	storageobject "github.com/linode/provider-linode/internal/controller/object/storageobject"
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
		device.Setup,
		firewall.Setup,
		config.Setup,
		disk.Setup,
		instance.Setup,
		ip.Setup,
		sharedips.Setup,
		cluster.Setup,
		confignodebalancer.Setup,
		image.Setup,
		node.Setup,
		nodebalancer.Setup,
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
