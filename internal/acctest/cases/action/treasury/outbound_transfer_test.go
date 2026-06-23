// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import "testing"

func TestAccActionTreasuryOutboundTransferBasic(t *testing.T) {
	runRemainingActionCase(
		t,
		"treasury_outbound_transfer_basic",
		"stripe_treasury_outbound_transfer",
		"treasury",
		nil,
		"action/treasury/treasury_outbound_transfer_basic_action.tf",
		nil,
		prepareTreasuryOutboundTransferConfig,
		verifyTreasuryOutboundTransferAction,
	)
}
