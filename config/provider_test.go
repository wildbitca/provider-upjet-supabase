package config

import "testing"

// TestGetProviderResources verifies that both provider scopes build and expose
// the same set of wrapped Terraform resources after configuration.
func TestGetProviderResources(t *testing.T) {
	cluster := GetProvider()
	namespaced := GetProviderNamespaced()
	if len(cluster.Resources) == 0 {
		t.Fatal("cluster-scoped provider has no resources")
	}
	if len(cluster.Resources) != len(namespaced.Resources) {
		t.Fatalf("resource count mismatch: cluster=%d namespaced=%d", len(cluster.Resources), len(namespaced.Resources))
	}
	for name := range cluster.Resources {
		if _, ok := namespaced.Resources[name]; !ok {
			t.Errorf("resource %q missing from namespaced provider", name)
		}
	}
}
