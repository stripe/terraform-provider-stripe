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

func prepareFeeRefundActionConfig(
	t *testing.T,
	env runner.TestEnv,
	client *stripe.Client,
) map[string]string {
	t.Helper()

	destination := ensureTransferDestination(t, env, client)
	chargeParams := &stripe.ChargeCreateParams{
		Amount:               stripe.Int64(600),
		ApplicationFeeAmount: stripe.Int64(150),
		Currency:             stripe.String("usd"),
		Description:          stripe.String("sdk-codegen fee refund funding"),
		TransferData: &stripe.ChargeCreateTransferDataParams{
			Destination: stripe.String(destination),
		},
		Metadata: map[string]string{
			"suite": "sdk-codegen",
			"case":  "fee_refund_action_regression_setup",
		},
	}
	if err := chargeParams.SetSource("tok_visa"); err != nil {
		t.Fatal(err)
	}
	chargeParams.AddExpand("application_fee")
	charge, err := client.V1Charges.Create(context.Background(), chargeParams)
	if err != nil {
		t.Fatal(err)
	}

	if charge.ApplicationFee == nil || charge.ApplicationFee.ID == "" {
		t.Fatalf("application fee for charge %s missing", charge.ID)
	}

	return map[string]string{
		"{{APPLICATION_FEE_ID}}": charge.ApplicationFee.ID,
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
		Email:        stripe.String("sdk-codegen-fee-refund-test@example.com"),
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
			"case":  "fee_refund_destination",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if account.Capabilities == nil || account.Capabilities.Transfers != stripe.AccountCapabilityStatusActive {
		t.Fatalf("created fee refund destination %s without active transfers capability", account.ID)
	}
	return account.ID
}

func verifyRefundAction(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	if err := runner.ExpectResourceAbsent(state, "action.stripe_refund.test"); err != nil {
		return err
	}
	refund, err := findLatestRefundByCase(client, "refund_basic")
	if err != nil {
		return err
	}
	expectedChargeID, err := runner.ResourcePrimaryID(state, "stripe_charge.charge")
	if err != nil {
		return err
	}
	if refund.Charge == nil || refund.Charge.ID != expectedChargeID {
		actual := ""
		if refund.Charge != nil {
			actual = refund.Charge.ID
		}
		return fmt.Errorf("refund charge mismatch: expected %q, got %q", expectedChargeID, actual)
	}
	if refund.Amount != 220 {
		return fmt.Errorf("refund amount mismatch: expected %d, got %d", 220, refund.Amount)
	}
	if string(refund.Reason) != "requested_by_customer" {
		return fmt.Errorf("refund reason mismatch: expected %q, got %q", "requested_by_customer", string(refund.Reason))
	}
	if err := expectMetadataSubset("refund_basic.metadata", refund.Metadata, map[string]string{
		"suite": "sdk-codegen",
		"case":  "refund_basic",
	}); err != nil {
		return err
	}
	if string(refund.Status) == "" {
		return fmt.Errorf("refund status missing")
	}
	return nil
}

func verifyFeeRefundAction(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	if err := runner.ExpectResourceAbsent(state, "action.stripe_fee_refund.test"); err != nil {
		return err
	}
	refund, err := findLatestFeeRefundByCase(client, "fee_refund_action_regression")
	if err != nil {
		return err
	}
	if refund.Amount != 55 {
		return fmt.Errorf("fee refund amount mismatch: expected %d, got %d", 55, refund.Amount)
	}
	if refund.Fee == nil || refund.Fee.ID == "" {
		return fmt.Errorf("fee refund fee missing")
	}
	return expectMetadataSubset("fee_refund_action_regression.metadata", refund.Metadata, map[string]string{
		"suite": "sdk-codegen",
		"case":  "fee_refund_action_regression",
	})
}

func findLatestRefundByCase(client *stripe.Client, caseName string) (*stripe.Refund, error) {
	params := &stripe.RefundListParams{}
	params.Limit = stripe.Int64(25)
	for refund, err := range client.V1Refunds.List(context.Background(), params).All(context.Background()) {
		if err != nil {
			return nil, fmt.Errorf("list refunds for %s: %w", caseName, err)
		}
		if refund != nil && refund.Metadata["suite"] == "sdk-codegen" && refund.Metadata["case"] == caseName {
			return refund, nil
		}
	}
	return nil, fmt.Errorf("refund for case %q not found", caseName)
}

func findLatestFeeRefundByCase(client *stripe.Client, caseName string) (*stripe.FeeRefund, error) {
	feeParams := &stripe.ApplicationFeeListParams{}
	feeParams.Limit = stripe.Int64(25)
	for fee, err := range client.V1ApplicationFees.List(context.Background(), feeParams).All(context.Background()) {
		if err != nil {
			return nil, fmt.Errorf("list application fees for %s: %w", caseName, err)
		}
		if fee != nil {
			refund, err := findLatestFeeRefundByCaseForFee(client, fee.ID, caseName)
			if err == nil {
				return refund, nil
			}
		}
	}
	return nil, fmt.Errorf("fee refund for case %q not found", caseName)
}

func findLatestFeeRefundByCaseForFee(client *stripe.Client, feeID string, caseName string) (*stripe.FeeRefund, error) {
	params := &stripe.FeeRefundListParams{ID: stripe.String(feeID)}
	params.Limit = stripe.Int64(25)
	for refund, err := range client.V1FeeRefunds.List(context.Background(), params).All(context.Background()) {
		if err != nil {
			return nil, fmt.Errorf("list fee refunds for %s: %w", caseName, err)
		}
		if refund != nil && refund.Metadata["suite"] == "sdk-codegen" && refund.Metadata["case"] == caseName {
			return refund, nil
		}
	}
	return nil, fmt.Errorf("fee refund for case %q not found", caseName)
}
