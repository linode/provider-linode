/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	cluster "github.com/linode/provider-linode/internal/controller/lke_cluster/cluster"
	providerconfig "github.com/linode/provider-linode/internal/controller/providerconfig"
	stackscript "github.com/linode/provider-linode/internal/controller/stackscript/stackscript"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		providerconfig.Setup,
		stackscript.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
