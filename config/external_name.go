package config

import "github.com/crossplane/upjet/v2/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider. supabase/supabase is a thin Terraform Plugin Framework provider.
// Every resource exposes a provider-computed "id" attribute, so all managed
// resources use IdentifierFromProvider (the provider assigns the identifier on
// create and reports it back in state).
//
//   - supabase_project                -> id is the project ref (provider-assigned)
//   - supabase_settings               -> singleton keyed by project_ref; id is provider-assigned
//   - supabase_branch                 -> id is provider-assigned
//   - supabase_edge_function          -> id is provider-assigned
//   - supabase_edge_function_secrets  -> keyed by project_ref; id is provider-assigned
//   - supabase_apikey                 -> id is provider-assigned
//
// See https://registry.terraform.io/providers/supabase/supabase/latest/docs
var ExternalNameConfigs = map[string]config.ExternalName{
	"supabase_project":               config.IdentifierFromProvider,
	"supabase_settings":              config.IdentifierFromProvider,
	"supabase_branch":                config.IdentifierFromProvider,
	"supabase_edge_function":         config.IdentifierFromProvider,
	"supabase_edge_function_secrets": config.IdentifierFromProvider,
	"supabase_apikey":                config.IdentifierFromProvider,
}

// ExternalNameConfigurations returns the list of all ExternalName configured resources.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns a list of all resources with external name configured.
func ExternalNameConfigured() []string {
	l := make([]string, 0, len(ExternalNameConfigs))
	for name := range ExternalNameConfigs {
		l = append(l, name)
	}
	return l
}
