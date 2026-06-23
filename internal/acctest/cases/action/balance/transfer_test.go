// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import "testing"

func TestAccActionTransferBasic(t *testing.T) {
	runRemainingActionCaseWithPrerequisite(
		t,
		"transfer_basic",
		"stripe_transfer",
		"base",
		nil,
		"action/balance/transfer_basic_prerequisite.tf",
		"action/balance/transfer_basic_action.tf",
		renderPlatformActionEnv,
		prepareTransferActionConfig,
		verifyTransferAction,
	)
}

func TestAccActionTransferReversalBasic(t *testing.T) {
	runRemainingActionCase(
		t,
		"transfer_reversal_basic",
		"stripe_transfer_reversal",
		"base",
		nil,
		"action/balance/transfer_reversal_basic_action.tf",
		renderPlatformActionEnv,
		prepareTransferReversalActionConfig,
		verifyTransferReversalAction,
	)
}
