// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import "testing"

func TestAccActionTreasuryCreditReversalBasic(t *testing.T) {
	runRemainingActionCase(
		t,
		"treasury_credit_reversal_action",
		"stripe_treasury_credit_reversal",
		"treasury",
		nil,
		"action/treasury/treasury_credit_reversal_action.tf",
		nil,
		prepareTreasuryCreditReversalConfig,
		verifyTreasuryCreditReversalAction,
	)
}
