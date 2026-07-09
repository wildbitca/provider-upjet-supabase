package apikey

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds the apikey group resource configurators. API keys reference
// the owning project (kind Project, group project) via project_ref.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("supabase_apikey", func(r *config.Resource) {
		r.ShortGroup = "apikey"
		r.References["project_ref"] = config.Reference{
			TerraformName: "supabase_project",
		}
	})
}
