// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import (
	"testing"

	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

func TestAccEphemeralTokenCard(t *testing.T) {
	runner.RunEphemeralCase(t, runner.EphemeralCase{
		Definition: runner.CaseDefinition{
			Name:    "token_card_ephemeral",
			Surface: "stripe_token",
			Group:   "base",
			Kind:    "ephemeral",
		},
		ConfigTemplate:       "ephemeral/core/token_card_ephemeral.tf",
		VerifyPersistedState: verifyTokenEphemeralState,
	})
}

func TestAccEphemeralTokenBankAccount(t *testing.T) {
	runner.RunEphemeralCase(t, runner.EphemeralCase{
		Definition: runner.CaseDefinition{
			Name:    "token_bank_account_ephemeral",
			Surface: "stripe_token",
			Group:   "base",
			Kind:    "ephemeral",
		},
		ConfigTemplate:       "ephemeral/core/token_bank_account_ephemeral.tf",
		VerifyPersistedState: verifyTokenEphemeralState,
	})
}

func TestAccEphemeralTokenPII(t *testing.T) {
	runner.RunEphemeralCase(t, runner.EphemeralCase{
		Definition: runner.CaseDefinition{
			Name:    "token_pii_ephemeral",
			Surface: "stripe_token",
			Group:   "base",
			Kind:    "ephemeral",
		},
		ConfigTemplate:       "ephemeral/core/token_pii_ephemeral.tf",
		VerifyPersistedState: verifyTokenEphemeralState,
	})
}
