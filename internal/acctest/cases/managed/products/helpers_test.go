// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import (
	"fmt"
	"testing"

	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

func runManagedCaseForGroup(
	t *testing.T,
	group string,
	requiredEnv []string,
	name string,
	surface string,
	resourceAddress string,
	createTemplate string,
	updateTemplate string,
	importStable bool,
	verifyCreate runner.StateVerifier,
	verifyUpdate runner.StateVerifier,
	verifyDestroy runner.StateVerifier,
) {
	t.Helper()

	runner.RunManagedCase(t, runner.ManagedCase{
		Definition: runner.CaseDefinition{
			Name:        name,
			Surface:     surface,
			Group:       group,
			Kind:        "resource",
			RequiredEnv: requiredEnv,
		},
		ResourceAddress: resourceAddress,
		CreateTemplate:  createTemplate,
		UpdateTemplate:  updateTemplate,
		ImportStable:    importStable,
		VerifyCreate:    verifyCreate,
		VerifyUpdate:    verifyUpdate,
		VerifyDestroy:   verifyDestroy,
	})
}

func expectMetadataSubset(name string, actual map[string]string, expected map[string]string) error {
	for key, expectedValue := range expected {
		actualValue, ok := actual[key]
		if !ok {
			return fmt.Errorf("%s missing metadata key %q", name, key)
		}
		if actualValue != expectedValue {
			return fmt.Errorf("%s metadata[%q] mismatch: expected %q, got %q", name, key, expectedValue, actualValue)
		}
	}

	return nil
}
