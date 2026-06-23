// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package runner

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
)

func RunEphemeralCase(t *testing.T, tc EphemeralCase) {
	t.Helper()

	if !AcceptanceEnabled() {
		t.Skip("set TF_ACC=1 to run Terraform acceptance tests")
	}
	if !ShouldRunCase(tc.Definition) {
		t.Skip("case filtered out by selectors")
	}

	env := ResolveTestEnv(t, tc.Definition.Group, tc.Definition.RequiredEnv)
	ApplyResolvedAccountEnv(t, env)
	client := NewStripeClient(env)
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: NewProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: RenderEphemeralConfig(env, tc.ConfigTemplate, nil),
				Check:  stateVerifierCheck(env, client, tc.VerifyPersistedState),
				ConfigStateChecks: append(
					[]statecheck.StateCheck{},
					tc.ConfigStateChecks...,
				),
			},
		},
	})
}
