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

func verifyTreasuryDebitReversalAction(
	env runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	if err := runner.ExpectResourceAbsent(state, "action.stripe_treasury_debit_reversal.test"); err != nil {
		return err
	}
	reversal, err := findLatestTreasuryDebitReversalByCase(client, "treasury_debit_reversal_basic")
	if err != nil {
		return err
	}
	expectedReceivedDebit := expectedTreasuryReceivedDebit(env)
	if reversal.ReceivedDebit != expectedReceivedDebit {
		return fmt.Errorf("treasury debit reversal received_debit mismatch: expected %q, got %q", expectedReceivedDebit, reversal.ReceivedDebit)
	}
	expectedFinancialAccount := expectedTreasurySourceFinancialAccount(env)
	if reversal.FinancialAccount != expectedFinancialAccount {
		return fmt.Errorf("treasury debit reversal financial_account mismatch: expected %q, got %q", expectedFinancialAccount, reversal.FinancialAccount)
	}
	if reversal.Amount != 210 {
		return fmt.Errorf("treasury debit reversal amount mismatch: expected %d, got %d", 210, reversal.Amount)
	}
	return expectMetadataSubset("treasury_debit_reversal_basic.metadata", reversal.Metadata, map[string]string{
		"suite": "sdk-codegen",
		"case":  "treasury_debit_reversal_basic",
	})
}

func findLatestTreasuryDebitReversalByCase(client *stripe.Client, caseName string) (*stripe.TreasuryDebitReversal, error) {
	params := &stripe.TreasuryDebitReversalListParams{}
	params.Limit = stripe.Int64(25)
	for reversal, err := range client.V1TreasuryDebitReversals.List(context.Background(), params).All(context.Background()) {
		if err != nil {
			return nil, fmt.Errorf("list treasury debit reversals for %s: %w", caseName, err)
		}
		if reversal != nil && reversal.Metadata["suite"] == "sdk-codegen" && reversal.Metadata["case"] == caseName {
			return reversal, nil
		}
	}
	return nil, fmt.Errorf("treasury debit reversal for case %q not found", caseName)
}
