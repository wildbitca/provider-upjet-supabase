package function

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds the function group resource configurators. Edge functions and
// their secrets reference the owning project (kind Project, group project) via
// project_ref.
func Configure(p *config.Provider) {
	for _, name := range []string{
		"supabase_edge_function",
		"supabase_edge_function_secrets",
	} {
		resName := name
		p.AddResourceConfigurator(resName, func(r *config.Resource) {
			r.ShortGroup = "function"
			r.References["project_ref"] = config.Reference{
				TerraformName: "supabase_project",
			}
		})
	}
}
