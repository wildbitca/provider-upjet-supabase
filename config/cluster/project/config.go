package project

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds the project group resource configurators. All Supabase
// resources are grouped under the single "project" API group.
func Configure(p *config.Provider) {
	// supabase_project — the root project resource.
	p.AddResourceConfigurator("supabase_project", func(r *config.Resource) {
		r.ShortGroup = "project"
	})

	// Project-scoped resources referencing a project via project_ref.
	for _, name := range []string{
		"supabase_settings",
		"supabase_edge_function",
		"supabase_edge_function_secrets",
		"supabase_apikey",
	} {
		resName := name
		p.AddResourceConfigurator(resName, func(r *config.Resource) {
			r.ShortGroup = "project"
			r.References["project_ref"] = config.Reference{
				TerraformName: "supabase_project",
			}
		})
	}

	// supabase_branch references its parent project via parent_project_ref.
	p.AddResourceConfigurator("supabase_branch", func(r *config.Resource) {
		r.ShortGroup = "project"
		r.References["parent_project_ref"] = config.Reference{
			TerraformName: "supabase_project",
		}
	})
}
