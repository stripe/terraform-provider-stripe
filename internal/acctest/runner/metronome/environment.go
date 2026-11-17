//
// File generated from our OpenAPI spec
//

package runner

import (
	"encoding/json"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func Environment(extra ...string) []string {
	var result []string
	for _, value := range os.Environ() {
		key := strings.SplitN(value, "=", 2)[0]
		if strings.HasPrefix(key, "TF_CLI_ARGS") || strings.HasPrefix(key, "STRIPE_") || strings.HasPrefix(key, "METRONOME_") || key == "TF_PLUGIN_CACHE_DIR" || key == "TF_CLI_CONFIG_FILE" || key == "TF_DATA_DIR" || key == "TF_REATTACH_PROVIDERS" || key == "TF_PLUGIN_CACHE_MAY_BREAK_DEPENDENCY_LOCK_FILE" || key == "TF_WORKSPACE" {
			continue
		}
		result = append(result, value)
	}
	return append(result, extra...)
}

func RequireTerraform(t *testing.T) {
	t.Helper()
	out, err := exec.Command("terraform", "version", "-json").Output()
	if err != nil {
		t.Fatalf("Terraform integration test executable is unavailable: %v", err)
	}
	var version struct {
		Version string `json:"terraform_version"`
	}
	if err = json.Unmarshal(out, &version); err != nil {
		t.Fatal(err)
	}
	if version.Version != "1.16.5" {
		t.Fatalf("Terraform integration tests require version 1.16.5, got %s", version.Version)
	}
}
