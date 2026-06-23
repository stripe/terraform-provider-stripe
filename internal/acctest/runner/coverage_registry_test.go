// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package runner

import "testing"

func TestCoverageRegistryEntriesAreWellFormed(t *testing.T) {
	seen := map[string]struct{}{}
	for _, entry := range CoverageRegistry {
		key := entry.Surface + ":" + entry.Kind
		if _, ok := seen[key]; ok {
			t.Fatalf("duplicate coverage registry entry %s", key)
		}
		seen[key] = struct{}{}

		if entry.Surface == "" || entry.Kind == "" || entry.Group == "" || entry.Status == "" {
			t.Fatalf("coverage registry entry %s is incomplete", key)
		}
		if entry.Status == "skipped" && entry.Reason == "" {
			t.Fatalf("coverage registry entry %s must declare a reason when skipped", key)
		}
	}
}
