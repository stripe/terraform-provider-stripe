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

func verifyTreasuryCreditReversalAction(
	env runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	if err := runner.ExpectResourceAbsent(state, "action.stripe_treasury_credit_reversal.test"); err != nil {
		return err
	}
	expectedFinancialAccount := expectedTreasurySourceFinancialAccount(env)
	reversal, err := findLatestTreasuryCreditReversalByCase(client, expectedFinancialAccount, "treasury_credit_reversal_action")
	if err != nil {
		return err
	}
	expectedReceivedCredit := expectedTreasuryReceivedCredit(env)
	if reversal.ReceivedCredit != expectedReceivedCredit {
		return fmt.Errorf("treasury credit reversal received_credit mismatch: expected %q, got %q", expectedReceivedCredit, reversal.ReceivedCredit)
	}
	if reversal.FinancialAccount != expectedFinancialAccount {
		return fmt.Errorf("treasury credit reversal financial_account mismatch: expected %q, got %q", expectedFinancialAccount, reversal.FinancialAccount)
	}
	return expectMetadataSubset("treasury_credit_reversal_action.metadata", reversal.Metadata, map[string]string{
		"suite": "sdk-codegen",
		"case":  "treasury_credit_reversal_action",
	})
}

func findLatestTreasuryCreditReversalByCase(
	client *stripe.Client,
	financialAccount string,
	caseName string,
) (*stripe.TreasuryCreditReversal, error) {
	params := &stripe.TreasuryCreditReversalListParams{
		FinancialAccount: stripe.String(financialAccount),
	}
	params.Limit = stripe.Int64(25)
	for reversal, err := range client.V1TreasuryCreditReversals.List(context.Background(), params).All(context.Background()) {
		if err != nil {
			return nil, fmt.Errorf("list treasury credit reversals for %s: %w", caseName, err)
		}
		if reversal != nil && reversal.Metadata["suite"] == "sdk-codegen" && reversal.Metadata["case"] == caseName {
			return reversal, nil
		}
	}
	return nil, fmt.Errorf("treasury credit reversal for case %q not found", caseName)
}
