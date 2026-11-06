// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	stripe "github.com/stripe/stripe-go/v86"

	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

func verifyTreasuryInboundTransferAction(
	env runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	if err := runner.ExpectResourceAbsent(state, "action.stripe_treasury_inbound_transfer.test"); err != nil {
		return err
	}
	transfer, err := findLatestTreasuryInboundTransferByCase(client, "treasury_inbound_transfer_basic")
	if err != nil {
		return err
	}
	if transfer.Amount != 210 {
		return fmt.Errorf("treasury inbound transfer amount mismatch: expected %d, got %d", 210, transfer.Amount)
	}
	expectedFinancialAccount := expectedTreasurySourceFinancialAccount(env)
	if transfer.FinancialAccount != expectedFinancialAccount {
		return fmt.Errorf("treasury inbound transfer financial_account mismatch: expected %q, got %q", expectedFinancialAccount, transfer.FinancialAccount)
	}
	expectedPaymentMethod := expectedTreasuryOriginPaymentMethod(env)
	if transfer.OriginPaymentMethod != expectedPaymentMethod {
		return fmt.Errorf("treasury inbound transfer origin_payment_method mismatch: expected %q, got %q", expectedPaymentMethod, transfer.OriginPaymentMethod)
	}
	if transfer.OriginPaymentMethodDetails == nil || string(transfer.OriginPaymentMethodDetails.Type) != "us_bank_account" {
		actual := ""
		if transfer.OriginPaymentMethodDetails != nil {
			actual = string(transfer.OriginPaymentMethodDetails.Type)
		}
		return fmt.Errorf("treasury inbound transfer origin_payment_method_details.type mismatch: expected %q, got %q", "us_bank_account", actual)
	}
	return expectMetadataSubset("treasury_inbound_transfer_basic.metadata", transfer.Metadata, map[string]string{
		"suite": "sdk-codegen",
		"case":  "treasury_inbound_transfer_basic",
	})
}

func findLatestTreasuryInboundTransferByCase(client *stripe.Client, caseName string) (*stripe.TreasuryInboundTransfer, error) {
	params := &stripe.TreasuryInboundTransferListParams{}
	params.Limit = stripe.Int64(25)
	for transfer, err := range client.V1TreasuryInboundTransfers.List(context.Background(), params).All(context.Background()) {
		if err != nil {
			return nil, fmt.Errorf("list treasury inbound transfers for %s: %w", caseName, err)
		}
		if transfer != nil && transfer.Metadata["suite"] == "sdk-codegen" && transfer.Metadata["case"] == caseName {
			return transfer, nil
		}
	}
	return nil, fmt.Errorf("treasury inbound transfer for case %q not found", caseName)
}
