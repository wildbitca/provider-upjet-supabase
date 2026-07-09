// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	function "github.com/wildbitca/provider-upjet-supabase/internal/controller/namespaced/function/function"
	functionsecrets "github.com/wildbitca/provider-upjet-supabase/internal/controller/namespaced/function/functionsecrets"
)

// Setup_function creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_function(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		function.Setup,
		functionsecrets.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_function creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_function(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		function.SetupGated,
		functionsecrets.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
