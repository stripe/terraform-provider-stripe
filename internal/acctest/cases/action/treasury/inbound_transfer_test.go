// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import "testing"

func TestAccActionTreasuryInboundTransferBasic(t *testing.T) {
	runRemainingActionCase(
		t,
		"treasury_inbound_transfer_basic",
		"stripe_treasury_inbound_transfer",
		"treasury",
		nil,
		"action/treasury/treasury_inbound_transfer_basic_action.tf",
		nil,
		prepareTreasuryInboundTransferConfig,
		verifyTreasuryInboundTransferAction,
	)
}
