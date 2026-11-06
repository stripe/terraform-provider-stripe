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

func verifyPayoutAction(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	if err := runner.ExpectResourceAbsent(state, "action.stripe_payout.test"); err != nil {
		return err
	}
	payout, err := findLatestPayoutByCase(client, "payout_basic")
	if err != nil {
		return err
	}
	if payout.Amount != 101 {
		return fmt.Errorf("payout amount mismatch: expected %d, got %d", 101, payout.Amount)
	}
	if string(payout.Currency) != "usd" {
		return fmt.Errorf("payout currency mismatch: expected %q, got %q", "usd", string(payout.Currency))
	}
	if payout.Description != "sdk-codegen payout basic" {
		return fmt.Errorf("payout description mismatch: expected %q, got %q", "sdk-codegen payout basic", payout.Description)
	}
	if string(payout.Method) != "standard" {
		return fmt.Errorf("payout method mismatch: expected %q, got %q", "standard", string(payout.Method))
	}
	if string(payout.SourceType) != "card" {
		return fmt.Errorf("payout source_type mismatch: expected %q, got %q", "card", string(payout.SourceType))
	}
	if string(payout.Type) != "bank_account" {
		return fmt.Errorf("payout type mismatch: expected %q, got %q", "bank_account", string(payout.Type))
	}
	return expectMetadataSubset("payout_basic.metadata", payout.Metadata, map[string]string{
		"suite": "sdk-codegen",
		"case":  "payout_basic",
	})
}

func findLatestPayoutByCase(client *stripe.Client, caseName string) (*stripe.Payout, error) {
	params := &stripe.PayoutListParams{}
	params.Limit = stripe.Int64(25)
	for payout, err := range client.V1Payouts.List(context.Background(), params).All(context.Background()) {
		if err != nil {
			return nil, fmt.Errorf("list payouts for %s: %w", caseName, err)
		}
		if payout != nil && payout.Metadata["suite"] == "sdk-codegen" && payout.Metadata["case"] == caseName {
			return payout, nil
		}
	}
	return nil, fmt.Errorf("payout for case %q not found", caseName)
}
