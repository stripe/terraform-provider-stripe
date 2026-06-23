// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import "testing"

func TestAccActionTreasuryDebitReversalBasic(t *testing.T) {
	runRemainingActionCase(
		t,
		"treasury_debit_reversal_basic",
		"stripe_treasury_debit_reversal",
		"treasury",
		nil,
		"action/treasury/treasury_debit_reversal_basic_action.tf",
		nil,
		prepareTreasuryDebitReversalConfig,
		verifyTreasuryDebitReversalAction,
	)
}
