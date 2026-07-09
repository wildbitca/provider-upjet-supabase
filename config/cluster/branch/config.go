package branch

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds the branch group resource configurators. Branches reference
// their parent project (kind Project, group project) via parent_project_ref.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("supabase_branch", func(r *config.Resource) {
		r.ShortGroup = "branch"
		r.References["parent_project_ref"] = config.Reference{
			TerraformName: "supabase_project",
		}
	})
}
