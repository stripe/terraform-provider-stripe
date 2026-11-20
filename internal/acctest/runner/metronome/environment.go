//
// File generated from our OpenAPI spec
//

package runner

import (
	"encoding/json"
	"os"
	"os/exec"
	"strconv"
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
	parts := strings.Split(version.Version, ".")
	major, minor, patch := 0, 0, 0
	if len(parts) == 3 {
		major, _ = strconv.Atoi(parts[0])
		minor, _ = strconv.Atoi(parts[1])
		patch, _ = strconv.Atoi(parts[2])
	}
	if major != 1 || minor < 15 || (minor == 15 && patch < 4) {
		t.Fatalf("Terraform integration tests require Terraform 1.15.4 or later, got %s", version.Version)
	}
}
