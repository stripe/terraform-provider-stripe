// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import "testing"

func TestAccActionRefundBasic(t *testing.T) {
	runRemainingActionCaseWithPrerequisite(
		t,
		"refund_basic",
		"stripe_refund",
		"base",
		nil,
		"action/refunds/refund_basic_prerequisite.tf",
		"action/refunds/refund_basic_action.tf",
		nil,
		nil,
		verifyRefundAction,
	)
}

func TestAccActionFeeRefundRegression(t *testing.T) {
	runRemainingActionCase(
		t,
		"fee_refund_action_regression",
		"stripe_fee_refund",
		"base",
		nil,
		"action/refunds/fee_refund_regression_action.tf",
		nil,
		prepareFeeRefundActionConfig,
		verifyFeeRefundAction,
	)
}
