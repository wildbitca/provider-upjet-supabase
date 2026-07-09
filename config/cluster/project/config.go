package project

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds the project group resource configurators. The project group
// holds the root project resource and its singleton settings. Project must
// keep this group ("project") because production claims reference
// Project.project.upjet-supabase.[m.]upbound.io/v1alpha1.
func Configure(p *config.Provider) {
	// supabase_project — the root project resource.
	p.AddResourceConfigurator("supabase_project", func(r *config.Resource) {
		r.ShortGroup = "project"
	})

	// supabase_settings — singleton keyed by project_ref.
	p.AddResourceConfigurator("supabase_settings", func(r *config.Resource) {
		r.ShortGroup = "project"
		r.References["project_ref"] = config.Reference{
			TerraformName: "supabase_project",
		}
	})
}
