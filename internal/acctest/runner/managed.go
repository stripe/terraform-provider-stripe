// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package runner

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	stripe "github.com/stripe/stripe-go/v86"

	"github.com/stripe/terraform-provider-stripe/internal/acctest/checks"
)

const legacyStripeProviderSource = "stripe/stripe"

func RunManagedCase(t *testing.T, tc ManagedCase) {
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
	caseReplacements := map[string]string{
		"{{RAND}}": fmt.Sprintf("%d", time.Now().UnixNano()),
	}
	createConfig := RenderManagedConfig(env, tc.CreateTemplate, caseReplacements)

	steps := []resource.TestStep{
		{
			Config: createConfig,
			Check:  stateVerifierCheck(env, client, tc.VerifyCreate),
			ConfigStateChecks: []statecheck.StateCheck{
				checks.ResourceIDIsSet(tc.ResourceAddress),
			},
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PostApplyPostRefresh: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tc.ResourceAddress, plancheck.ResourceActionNoop),
				},
			},
		},
	}

	if tc.UpdateTemplate != "" {
		steps = append(steps, resource.TestStep{
			Config: RenderManagedConfig(env, tc.UpdateTemplate, caseReplacements),
			Check:  stateVerifierCheck(env, client, tc.VerifyUpdate),
			ConfigStateChecks: []statecheck.StateCheck{
				checks.ResourceIDIsSet(tc.ResourceAddress),
			},
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PostApplyPostRefresh: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tc.ResourceAddress, plancheck.ResourceActionNoop),
				},
			},
		})
	}

	if tc.ImportStable {
		steps = append(steps, resource.TestStep{
			ResourceName:            tc.ResourceAddress,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: tc.ImportStateVerifyIgnore,
		})
	}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: NewProtoV6ProviderFactories(),
		CheckDestroy:             stateVerifierCheck(env, client, tc.VerifyDestroy),
		Steps:                    steps,
	})
}

func RunManagedLegacyUpgradeCase(t *testing.T, tc ManagedCase, legacyProviderVersion string) {
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
	caseReplacements := map[string]string{
		"{{RAND}}": fmt.Sprintf("%d", time.Now().UnixNano()),
	}
	createConfig := RenderManagedConfig(env, tc.CreateTemplate, caseReplacements)

	resource.Test(t, resource.TestCase{
		CheckDestroy: stateVerifierCheck(env, client, tc.VerifyDestroy),
		Steps: []resource.TestStep{
			{
				Config:            createConfig,
				ExternalProviders: legacyStripeExternalProviders(legacyProviderVersion),
				ConfigStateChecks: []statecheck.StateCheck{
					checks.ResourceIDIsSet(tc.ResourceAddress),
				},
			},
			{
				Config:                   createConfig,
				ProtoV6ProviderFactories: NewProtoV6ProviderFactories(),
				Check:                    stateVerifierCheck(env, client, tc.VerifyCreate),
				ConfigStateChecks: []statecheck.StateCheck{
					checks.ResourceIDIsSet(tc.ResourceAddress),
				},
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(tc.ResourceAddress, plancheck.ResourceActionNoop),
					},
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(tc.ResourceAddress, plancheck.ResourceActionNoop),
					},
				},
			},
		},
	})
}

func legacyStripeExternalProviders(version string) map[string]resource.ExternalProvider {
	return map[string]resource.ExternalProvider{
		"stripe": {
			Source:            legacyStripeProviderSource,
			VersionConstraint: fmt.Sprintf("= %s", version),
		},
	}
}

func stateVerifierCheck(
	env TestEnv,
	client *stripe.Client,
	verify StateVerifier,
) resource.TestCheckFunc {
	if verify == nil {
		return nil
	}

	return func(state *terraform.State) error {
		return verify(env, client, state)
	}
}
