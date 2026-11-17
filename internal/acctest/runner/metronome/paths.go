//
// File generated from our OpenAPI spec
//

package runner

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

// ProviderRoot locates the provider module regardless of the test package depth.
func ProviderRoot(t *testing.T) string {
	t.Helper()
	directory, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(directory, "go.mod")); err == nil {
			return directory
		}
		parent := filepath.Dir(directory)
		if parent == directory {
			t.Fatal("cannot locate the provider module")
		}
		directory = parent
	}
}

// RenderConfig substitutes quoted values in a Terraform test configuration.
func RenderConfig(t *testing.T, name string, replacements map[string]string) string {
	t.Helper()
	contents, err := os.ReadFile(filepath.Join(ProviderRoot(t), "internal", "acctest", "testdata", "metronome", name))
	if err != nil {
		t.Fatal(err)
	}
	config := string(contents)
	for name, value := range replacements {
		config = strings.ReplaceAll(config, strconv.Quote("{{"+name+"}}"), strconv.Quote(value))
	}
	return config
}
