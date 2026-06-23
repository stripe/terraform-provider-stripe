// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	stripe "github.com/stripe/stripe-go/v86"

	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

func prepareTransferActionConfig(
	t *testing.T,
	env runner.TestEnv,
	client *stripe.Client,
) map[string]string {
	t.Helper()

	return map[string]string{
		"{{STRIPE_ACCOUNT}}": ensureTransferDestination(t, env, client),
	}
}

func prepareTransferReversalActionConfig(
	t *testing.T,
	env runner.TestEnv,
	client *stripe.Client,
) map[string]string {
	t.Helper()

	chargeParams := &stripe.ChargeCreateParams{
		Amount:      stripe.Int64(220),
		Currency:    stripe.String("usd"),
		Description: stripe.String("sdk-codegen transfer reversal funding"),
		Metadata: map[string]string{
			"suite": "sdk-codegen",
			"case":  "transfer_reversal_basic_setup",
		},
	}
	if err := chargeParams.SetSource("tok_visa"); err != nil {
		t.Fatal(err)
	}
	charge, err := client.V1Charges.Create(context.Background(), chargeParams)
	if err != nil {
		t.Fatal(err)
	}

	transfer, err := client.V1Transfers.Create(context.Background(), &stripe.TransferCreateParams{
		Amount:            stripe.Int64(110),
		Currency:          stripe.String("usd"),
		Destination:       stripe.String(ensureTransferDestination(t, env, client)),
		SourceTransaction: stripe.String(charge.ID),
		Description:       stripe.String("sdk-codegen transfer reversal parent"),
		TransferGroup:     stripe.String("sdk-codegen-transfer-reversal"),
		Metadata: map[string]string{
			"suite": "sdk-codegen",
			"case":  "transfer_reversal_basic_parent",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	return map[string]string{
		"{{TRANSFER_ID}}": transfer.ID,
	}
}

func ensureTransferDestination(
	t *testing.T,
	env runner.TestEnv,
	client *stripe.Client,
) string {
	t.Helper()

	if env.StripeAccount != "" {
		account, err := client.V1Accounts.GetByID(context.Background(), env.StripeAccount, nil)
		if err == nil && account.Capabilities != nil && account.Capabilities.Transfers == stripe.AccountCapabilityStatusActive {
			return env.StripeAccount
		}
	}

	account, err := client.V1Accounts.Create(context.Background(), &stripe.AccountCreateParams{
		Type:         stripe.String(stripe.AccountTypeCustom),
		Country:      stripe.String("US"),
		Email:        stripe.String("sdk-codegen-transfer-test@example.com"),
		BusinessType: stripe.String("individual"),
		BusinessProfile: &stripe.AccountCreateBusinessProfileParams{
			ProductDescription: stripe.String("sdk-codegen acceptance tests"),
		},
		Capabilities: &stripe.AccountCreateCapabilitiesParams{
			Transfers: &stripe.AccountCreateCapabilitiesTransfersParams{
				Requested: stripe.Bool(true),
			},
		},
		ExternalAccount: &stripe.AccountExternalAccountParams{
			Token: stripe.String("btok_us"),
		},
		Individual: &stripe.PersonParams{
			FirstName: stripe.String("Jenny"),
			LastName:  stripe.String("Rosen"),
			DOB: &stripe.PersonDOBParams{
				Day:   stripe.Int64(1),
				Month: stripe.Int64(1),
				Year:  stripe.Int64(1990),
			},
			SSNLast4: stripe.String("0000"),
		},
		TOSAcceptance: &stripe.AccountCreateTOSAcceptanceParams{
			Date: stripe.Int64(time.Now().Unix()),
			IP:   stripe.String("127.0.0.1"),
		},
		Metadata: map[string]string{
			"suite": "sdk-codegen",
			"case":  "transfer_destination",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if account.Capabilities == nil || account.Capabilities.Transfers != stripe.AccountCapabilityStatusActive {
		t.Fatalf("created transfer destination %s without active transfers capability", account.ID)
	}
	return account.ID
}

func verifyTransferAction(
	env runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	if err := runner.ExpectResourceAbsent(state, "action.stripe_transfer.test"); err != nil {
		return err
	}
	transfer, err := findLatestTransferByCase(client, "transfer_basic")
	if err != nil {
		return err
	}
	expectedChargeID, err := runner.ResourcePrimaryID(state, "stripe_charge.charge")
	if err != nil {
		return err
	}
	if transfer.Amount != 110 {
		return fmt.Errorf("transfer amount mismatch: expected %d, got %d", 110, transfer.Amount)
	}
	if string(transfer.Currency) != "usd" {
		return fmt.Errorf("transfer currency mismatch: expected %q, got %q", "usd", string(transfer.Currency))
	}
	if transfer.Description != "sdk-codegen transfer basic" {
		return fmt.Errorf("transfer description mismatch: expected %q, got %q", "sdk-codegen transfer basic", transfer.Description)
	}
	if transfer.Destination == nil || transfer.Destination.ID == "" {
		return fmt.Errorf("transfer destination missing")
	}
	if transfer.SourceTransaction == nil || transfer.SourceTransaction.ID != expectedChargeID {
		actual := ""
		if transfer.SourceTransaction != nil {
			actual = transfer.SourceTransaction.ID
		}
		return fmt.Errorf("transfer source_transaction mismatch: expected %q, got %q", expectedChargeID, actual)
	}
	if transfer.TransferGroup != "sdk-codegen-transfer" {
		return fmt.Errorf("transfer transfer_group mismatch: expected %q, got %q", "sdk-codegen-transfer", transfer.TransferGroup)
	}
	return expectMetadataSubset("transfer_basic.metadata", transfer.Metadata, map[string]string{
		"suite": "sdk-codegen",
		"case":  "transfer_basic",
	})
}

func verifyTransferReversalAction(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	if err := runner.ExpectResourceAbsent(state, "action.stripe_transfer_reversal.test"); err != nil {
		return err
	}
	transfer, err := findLatestTransferByCase(client, "transfer_reversal_basic_parent")
	if err != nil {
		return err
	}
	reversal, err := findLatestTransferReversalByCase(client, transfer.ID, "transfer_reversal_basic")
	if err != nil {
		return err
	}
	if reversal.Amount != 40 {
		return fmt.Errorf("transfer reversal amount mismatch: expected %d, got %d", 40, reversal.Amount)
	}
	if reversal.Transfer == nil || reversal.Transfer.ID != transfer.ID {
		actual := ""
		if reversal.Transfer != nil {
			actual = reversal.Transfer.ID
		}
		return fmt.Errorf("transfer reversal transfer mismatch: expected %q, got %q", transfer.ID, actual)
	}
	return expectMetadataSubset("transfer_reversal_basic.metadata", reversal.Metadata, map[string]string{
		"suite": "sdk-codegen",
		"case":  "transfer_reversal_basic",
	})
}

func findLatestTransferByCase(client *stripe.Client, caseName string) (*stripe.Transfer, error) {
	params := &stripe.TransferListParams{}
	params.Limit = stripe.Int64(25)
	for transfer, err := range client.V1Transfers.List(context.Background(), params).All(context.Background()) {
		if err != nil {
			return nil, fmt.Errorf("list transfers for %s: %w", caseName, err)
		}
		if transfer != nil && transfer.Metadata["suite"] == "sdk-codegen" && transfer.Metadata["case"] == caseName {
			return transfer, nil
		}
	}
	return nil, fmt.Errorf("transfer for case %q not found", caseName)
}

func findLatestTransferReversalByCase(client *stripe.Client, transferID string, caseName string) (*stripe.TransferReversal, error) {
	params := &stripe.TransferReversalListParams{ID: stripe.String(transferID)}
	params.Limit = stripe.Int64(25)
	for reversal, err := range client.V1TransferReversals.List(context.Background(), params).All(context.Background()) {
		if err != nil {
			return nil, fmt.Errorf("list transfer reversals for %s: %w", caseName, err)
		}
		if reversal != nil && reversal.Metadata["suite"] == "sdk-codegen" && reversal.Metadata["case"] == caseName {
			return reversal, nil
		}
	}
	return nil, fmt.Errorf("transfer reversal for case %q not found", caseName)
}
