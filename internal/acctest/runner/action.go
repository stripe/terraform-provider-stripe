// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package runner

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
)

func RunActionCase(t *testing.T, tc ActionCase) {
	t.Helper()

	if !AcceptanceEnabled() {
		t.Skip("set TF_ACC=1 to run Terraform acceptance tests")
	}
	if !ShouldRunCase(tc.Definition) {
		t.Skip("case filtered out by selectors")
	}

	env := ResolveTestEnv(t, tc.Definition.Group, tc.Definition.RequiredEnv)
	renderEnv := env
	if tc.RenderEnv != nil {
		renderEnv = tc.RenderEnv(env)
	}
	ApplyResolvedAccountEnv(t, renderEnv)
	client := NewStripeClient(renderEnv)
	var extras map[string]string
	if tc.ConfigPreparer != nil {
		extras = tc.ConfigPreparer(t, env, client)
	}
	if os.Getenv("TF_ACC_TERRAFORM_PATH") == "" && os.Getenv("TF_ACC_TERRAFORM_VERSION") == "" {
		terraformPath, err := exec.LookPath("terraform")
		if err == nil {
			t.Setenv("TF_ACC_TERRAFORM_PATH", terraformPath)
		} else {
			t.Setenv("TF_ACC_TERRAFORM_VERSION", "1.15.4")
		}
	}
	configuredActionAddress := fmt.Sprintf("action.%s.test", tc.Definition.Surface)
	originalPlanArgs, originalPlanArgsSet := os.LookupEnv("TF_CLI_ARGS_plan")
	originalApplyArgs, originalApplyArgsSet := os.LookupEnv("TF_CLI_ARGS_apply")
	t.Cleanup(func() {
		restoreEnv("TF_CLI_ARGS_plan", originalPlanArgs, originalPlanArgsSet)
		restoreEnv("TF_CLI_ARGS_apply", originalApplyArgs, originalApplyArgsSet)
	})

	clearInvokeArgs := func() {
		restoreEnv("TF_CLI_ARGS_plan", originalPlanArgs, originalPlanArgsSet)
		restoreEnv("TF_CLI_ARGS_apply", originalApplyArgs, originalApplyArgsSet)
	}
	setInvokeArgs := func() {
		setEnv("TF_CLI_ARGS_plan", appendCliArg(originalPlanArgs, fmt.Sprintf("-invoke=%s", configuredActionAddress)))
		setEnv("TF_CLI_ARGS_apply", appendCliArg(originalApplyArgs, fmt.Sprintf("-invoke=%s", configuredActionAddress)))
	}

	steps := make([]resource.TestStep, 0, 2)
	if tc.PrerequisiteConfigTemplate != "" {
		steps = append(steps, resource.TestStep{
			PreConfig: clearInvokeArgs,
			Config:    RenderActionConfig(renderEnv, tc.PrerequisiteConfigTemplate, extras),
		})
	}
	steps = append(steps, resource.TestStep{
		PreConfig: setInvokeArgs,
		Config:    RenderActionConfig(renderEnv, tc.ConfigTemplate, extras),
		Check:     stateVerifierCheck(env, client, tc.VerifyInvoke),
		ConfigStateChecks: append(
			[]statecheck.StateCheck{},
			tc.ConfigStateChecks...,
		),
	})

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: NewProtoV6ProviderFactories(),
		Steps:                    steps,
	})
}

func appendCliArg(existing string, arg string) string {
	if existing == "" {
		return arg
	}
	return fmt.Sprintf("%s %s", existing, arg)
}

func restoreEnv(key string, value string, set bool) {
	if set {
		setEnv(key, value)
		return
	}
	os.Unsetenv(key)
}

func setEnv(key string, value string) {
	if value == "" {
		os.Unsetenv(key)
		return
	}
	os.Setenv(key, value)
}
