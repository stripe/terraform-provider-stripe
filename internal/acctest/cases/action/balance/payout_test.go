// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import "testing"

func TestAccActionPayoutBasic(t *testing.T) {
	runRemainingActionCase(
		t,
		"payout_basic",
		"stripe_payout",
		"connect",
		nil,
		"action/balance/payout_basic_action.tf",
		renderPlatformActionEnv,
		nil,
		verifyPayoutAction,
	)
}
