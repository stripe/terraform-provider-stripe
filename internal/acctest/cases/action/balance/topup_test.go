// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import "testing"

func TestAccActionTopupBasic(t *testing.T) {
	runRemainingActionCase(
		t,
		"topup_basic",
		"stripe_topup",
		"connect",
		nil,
		"action/balance/topup_basic_action.tf",
		renderPlatformActionEnv,
		nil,
		verifyTopupAction,
	)
}
