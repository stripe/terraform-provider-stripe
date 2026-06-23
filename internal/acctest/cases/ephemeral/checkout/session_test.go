// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import (
	"testing"

	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

func TestAccEphemeralCheckoutSession(t *testing.T) {
	runner.RunEphemeralCase(t, runner.EphemeralCase{
		Definition: runner.CaseDefinition{
			Name:    "checkout_session_ephemeral",
			Surface: "stripe_checkout_session",
			Group:   "base",
			Kind:    "ephemeral",
		},
		ConfigTemplate:       "ephemeral/checkout/checkout_session_ephemeral.tf",
		VerifyPersistedState: verifyCheckoutSessionEphemeral,
	})
}
