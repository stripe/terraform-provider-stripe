// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import (
	"testing"

	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

func runBaseEphemeralCase(
	t *testing.T,
	name string,
	surface string,
	configTemplate string,
	configStateChecks []runner.StateCheck,
	verify runner.StateVerifier,
) {
	t.Helper()

	convertedChecks := make([]runner.StateCheck, 0, len(configStateChecks))
	convertedChecks = append(convertedChecks, configStateChecks...)

	runner.RunEphemeralCase(t, runner.EphemeralCase{
		Definition: runner.CaseDefinition{
			Name:    name,
			Surface: surface,
			Group:   "base",
			Kind:    "ephemeral",
		},
		ConfigTemplate:       configTemplate,
		ConfigStateChecks:    convertedChecks,
		VerifyPersistedState: verify,
	})
}
