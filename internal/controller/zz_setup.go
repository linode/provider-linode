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
	image "github.com/linode/provider-linode/internal/controller/image/image"
	config "github.com/linode/provider-linode/internal/controller/instance/config"
	disk "github.com/linode/provider-linode/internal/controller/instance/disk"
	instance "github.com/linode/provider-linode/internal/controller/instance/instance"
	ip "github.com/linode/provider-linode/internal/controller/instance/ip"
	sharedips "github.com/linode/provider-linode/internal/controller/instance/sharedips"
	cluster "github.com/linode/provider-linode/internal/controller/lke/cluster"
	confignodebalancer "github.com/linode/provider-linode/internal/controller/nodebalancer/config"
	node "github.com/linode/provider-linode/internal/controller/nodebalancer/node"
	nodebalancer "github.com/linode/provider-linode/internal/controller/nodebalancer/nodebalancer"
	bucket "github.com/linode/provider-linode/internal/controller/objectstorage/bucket"
	key "github.com/linode/provider-linode/internal/controller/objectstorage/key"
	object "github.com/linode/provider-linode/internal/controller/objectstorage/object"
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
		image.Setup,
		config.Setup,
		disk.Setup,
		instance.Setup,
		ip.Setup,
		sharedips.Setup,
		cluster.Setup,
		confignodebalancer.Setup,
		node.Setup,
		nodebalancer.Setup,
		bucket.Setup,
		key.Setup,
		object.Setup,
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
