// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	apikey "github.com/wildbitca/provider-upjet-supabase/internal/controller/cluster/project/apikey"
	branch "github.com/wildbitca/provider-upjet-supabase/internal/controller/cluster/project/branch"
	function "github.com/wildbitca/provider-upjet-supabase/internal/controller/cluster/project/function"
	functionsecrets "github.com/wildbitca/provider-upjet-supabase/internal/controller/cluster/project/functionsecrets"
	project "github.com/wildbitca/provider-upjet-supabase/internal/controller/cluster/project/project"
	settings "github.com/wildbitca/provider-upjet-supabase/internal/controller/cluster/project/settings"
)

// Setup_project creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_project(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		apikey.Setup,
		branch.Setup,
		function.Setup,
		functionsecrets.Setup,
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
		apikey.SetupGated,
		branch.SetupGated,
		function.SetupGated,
		functionsecrets.SetupGated,
		project.SetupGated,
		settings.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
