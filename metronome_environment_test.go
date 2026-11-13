//
// File generated from our OpenAPI spec
//

package main

import (
	"os"
	"strings"
)

func terraformTestEnvironment(extra ...string) []string {
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
