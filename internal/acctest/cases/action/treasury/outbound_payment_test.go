// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import "testing"

func TestAccActionTreasuryOutboundPaymentBasic(t *testing.T) {
	runRemainingActionCase(
		t,
		"treasury_outbound_payment_basic",
		"stripe_treasury_outbound_payment",
		"treasury",
		nil,
		"action/treasury/treasury_outbound_payment_basic_action.tf",
		nil,
		prepareTreasuryOutboundPaymentConfig,
		verifyTreasuryOutboundPaymentAction,
	)
}
