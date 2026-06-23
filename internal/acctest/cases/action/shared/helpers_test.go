// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import (
	"testing"

	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

func renderPlatformActionEnv(env runner.TestEnv) runner.TestEnv {
	env.StripeAccount = ""
	return env
}

func runRemainingActionCase(
	t *testing.T,
	name string,
	surface string,
	group string,
	requiredEnv []string,
	configTemplate string,
	renderEnv runner.ActionRenderEnv,
	configPreparer runner.ActionConfigPreparer,
	verify runner.StateVerifier,
) {
	t.Helper()

	runRemainingActionCaseWithPrerequisite(
		t,
		name,
		surface,
		group,
		requiredEnv,
		"",
		configTemplate,
		renderEnv,
		configPreparer,
		verify,
	)
}

func runRemainingActionCaseWithPrerequisite(
	t *testing.T,
	name string,
	surface string,
	group string,
	requiredEnv []string,
	prerequisiteConfigTemplate string,
	configTemplate string,
	renderEnv runner.ActionRenderEnv,
	configPreparer runner.ActionConfigPreparer,
	verify runner.StateVerifier,
) {
	t.Helper()

	runner.RunActionCase(t, runner.ActionCase{
		Definition: runner.CaseDefinition{
			Name:        name,
			Surface:     surface,
			Group:       group,
			Kind:        "action",
			RequiredEnv: requiredEnv,
		},
		PrerequisiteConfigTemplate: prerequisiteConfigTemplate,
		ConfigTemplate:             configTemplate,
		ConfigPreparer:             configPreparer,
		RenderEnv:                  renderEnv,
		VerifyInvoke:               verify,
	})
}
