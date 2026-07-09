// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	project "github.com/wildbitca/provider-upjet-supabase/internal/controller/cluster/project/project"
	settings "github.com/wildbitca/provider-upjet-supabase/internal/controller/cluster/project/settings"
)

// Setup_project creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_project(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		project.Setup,
		settings.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_project creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_project(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		project.SetupGated,
		settings.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
