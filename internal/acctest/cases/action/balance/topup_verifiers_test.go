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

func verifyTopupAction(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	if err := runner.ExpectResourceAbsent(state, "action.stripe_topup.test"); err != nil {
		return err
	}
	topup, err := findLatestTopupByCase(client, "topup_basic")
	if err != nil {
		return err
	}
	if topup.Amount != 5100 {
		return fmt.Errorf("topup amount mismatch: expected %d, got %d", 5100, topup.Amount)
	}
	if string(topup.Currency) != "usd" {
		return fmt.Errorf("topup currency mismatch: expected %q, got %q", "usd", string(topup.Currency))
	}
	if topup.Description != "sdk-codegen topup basic" {
		return fmt.Errorf("topup description mismatch: expected %q, got %q", "sdk-codegen topup basic", topup.Description)
	}
	if topup.TransferGroup != "sdk-codegen-topup" {
		return fmt.Errorf("topup transfer_group mismatch: expected %q, got %q", "sdk-codegen-topup", topup.TransferGroup)
	}
	if string(topup.Status) == "" {
		return fmt.Errorf("topup status missing")
	}
	return expectMetadataSubset("topup_basic.metadata", topup.Metadata, map[string]string{
		"suite": "sdk-codegen",
		"case":  "topup_basic",
	})
}

func findLatestTopupByCase(client *stripe.Client, caseName string) (*stripe.Topup, error) {
	params := &stripe.TopupListParams{}
	params.Limit = stripe.Int64(25)
	for topup, err := range client.V1Topups.List(context.Background(), params).All(context.Background()) {
		if err != nil {
			return nil, fmt.Errorf("list topups for %s: %w", caseName, err)
		}
		if topup != nil && topup.Metadata["suite"] == "sdk-codegen" && topup.Metadata["case"] == caseName {
			return topup, nil
		}
	}
	return nil, fmt.Errorf("topup for case %q not found", caseName)
}
