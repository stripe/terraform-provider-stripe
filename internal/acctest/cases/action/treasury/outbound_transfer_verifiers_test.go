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

func verifyTreasuryOutboundTransferAction(
	env runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	if err := runner.ExpectResourceAbsent(state, "action.stripe_treasury_outbound_transfer.test"); err != nil {
		return err
	}
	transfer, err := findLatestTreasuryOutboundTransferByCase(client, "treasury_outbound_transfer_basic")
	if err != nil {
		return err
	}
	if transfer.Amount != 65 {
		return fmt.Errorf("treasury outbound transfer amount mismatch: expected %d, got %d", 65, transfer.Amount)
	}
	expectedFinancialAccount := expectedTreasurySourceFinancialAccount(env)
	if transfer.FinancialAccount != expectedFinancialAccount {
		return fmt.Errorf("treasury outbound transfer financial_account mismatch: expected %q, got %q", expectedFinancialAccount, transfer.FinancialAccount)
	}
	if transfer.Description != "sdk-codegen treasury outbound transfer basic" {
		return fmt.Errorf("treasury outbound transfer description mismatch: expected %q, got %q", "sdk-codegen treasury outbound transfer basic", transfer.Description)
	}
	expectedDestinationFinancialAccount := expectedTreasuryDestinationFinancialAccount(env)
	if transfer.DestinationPaymentMethodDetails == nil || transfer.DestinationPaymentMethodDetails.FinancialAccount == nil || transfer.DestinationPaymentMethodDetails.FinancialAccount.ID != expectedDestinationFinancialAccount {
		actual := ""
		if transfer.DestinationPaymentMethodDetails != nil && transfer.DestinationPaymentMethodDetails.FinancialAccount != nil {
			actual = transfer.DestinationPaymentMethodDetails.FinancialAccount.ID
		}
		return fmt.Errorf("treasury outbound transfer destination financial_account mismatch: expected %q, got %q", expectedDestinationFinancialAccount, actual)
	}
	return expectMetadataSubset("treasury_outbound_transfer_basic.metadata", transfer.Metadata, map[string]string{
		"suite": "sdk-codegen",
		"case":  "treasury_outbound_transfer_basic",
	})
}

func findLatestTreasuryOutboundTransferByCase(client *stripe.Client, caseName string) (*stripe.TreasuryOutboundTransfer, error) {
	params := &stripe.TreasuryOutboundTransferListParams{}
	params.Limit = stripe.Int64(25)
	for transfer, err := range client.V1TreasuryOutboundTransfers.List(context.Background(), params).All(context.Background()) {
		if err != nil {
			return nil, fmt.Errorf("list treasury outbound transfers for %s: %w", caseName, err)
		}
		if transfer != nil && transfer.Metadata["suite"] == "sdk-codegen" && transfer.Metadata["case"] == caseName {
			return transfer, nil
		}
	}
	return nil, fmt.Errorf("treasury outbound transfer for case %q not found", caseName)
}
