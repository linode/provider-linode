/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	accountsettings "github.com/linode/provider-linode/internal/controller/accountsettings/accountsettings"
	accesscontrols "github.com/linode/provider-linode/internal/controller/database/accesscontrols"
	mysql "github.com/linode/provider-linode/internal/controller/database/mysql"
	mysqlv2 "github.com/linode/provider-linode/internal/controller/database/mysqlv2"
	postgresql "github.com/linode/provider-linode/internal/controller/database/postgresql"
	postgresqlv2 "github.com/linode/provider-linode/internal/controller/database/postgresqlv2"
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
	ipv6range "github.com/linode/provider-linode/internal/controller/ipv6/ipv6range"
	cluster "github.com/linode/provider-linode/internal/controller/lke/cluster"
	nodepool "github.com/linode/provider-linode/internal/controller/lke/nodepool"
	ipnetworking "github.com/linode/provider-linode/internal/controller/networking/ip"
	confignodebalancer "github.com/linode/provider-linode/internal/controller/nodebalancer/config"
	node "github.com/linode/provider-linode/internal/controller/nodebalancer/node"
	nodebalancer "github.com/linode/provider-linode/internal/controller/nodebalancer/nodebalancer"
	bucket "github.com/linode/provider-linode/internal/controller/objectstorage/bucket"
	key "github.com/linode/provider-linode/internal/controller/objectstorage/key"
	object "github.com/linode/provider-linode/internal/controller/objectstorage/object"
	placementgroup "github.com/linode/provider-linode/internal/controller/placementgroup/placementgroup"
	placementgroupassignment "github.com/linode/provider-linode/internal/controller/placementgroupassignment/placementgroupassignment"
	providerconfig "github.com/linode/provider-linode/internal/controller/providerconfig"
	rdns "github.com/linode/provider-linode/internal/controller/rdns/rdns"
	ipassignment "github.com/linode/provider-linode/internal/controller/reserved/ipassignment"
	sshkey "github.com/linode/provider-linode/internal/controller/sshkey/sshkey"
	stackscript "github.com/linode/provider-linode/internal/controller/stackscript/stackscript"
	token "github.com/linode/provider-linode/internal/controller/token/token"
	user "github.com/linode/provider-linode/internal/controller/user/user"
	volume "github.com/linode/provider-linode/internal/controller/volume/volume"
	subnet "github.com/linode/provider-linode/internal/controller/vpc/subnet"
	vpc "github.com/linode/provider-linode/internal/controller/vpc/vpc"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accountsettings.Setup,
		accesscontrols.Setup,
		mysql.Setup,
		mysqlv2.Setup,
		postgresql.Setup,
		postgresqlv2.Setup,
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
		ipv6range.Setup,
		cluster.Setup,
		nodepool.Setup,
		ipnetworking.Setup,
		confignodebalancer.Setup,
		node.Setup,
		nodebalancer.Setup,
		bucket.Setup,
		key.Setup,
		object.Setup,
		placementgroup.Setup,
		placementgroupassignment.Setup,
		providerconfig.Setup,
		rdns.Setup,
		ipassignment.Setup,
		sshkey.Setup,
		stackscript.Setup,
		token.Setup,
		user.Setup,
		volume.Setup,
		subnet.Setup,
		vpc.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
