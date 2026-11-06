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

func verifyTreasuryOutboundPaymentAction(
	env runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	if err := runner.ExpectResourceAbsent(state, "action.stripe_treasury_outbound_payment.test"); err != nil {
		return err
	}
	payment, err := findLatestTreasuryOutboundPaymentByCase(client, "treasury_outbound_payment_basic")
	if err != nil {
		return err
	}
	if payment.Amount != 55 {
		return fmt.Errorf("treasury outbound payment amount mismatch: expected %d, got %d", 55, payment.Amount)
	}
	expectedFinancialAccount := expectedTreasurySourceFinancialAccount(env)
	if payment.FinancialAccount != expectedFinancialAccount {
		return fmt.Errorf("treasury outbound payment financial_account mismatch: expected %q, got %q", expectedFinancialAccount, payment.FinancialAccount)
	}
	if payment.Description != "sdk-codegen treasury outbound payment basic" {
		return fmt.Errorf("treasury outbound payment description mismatch: expected %q, got %q", "sdk-codegen treasury outbound payment basic", payment.Description)
	}
	if payment.DestinationPaymentMethodDetails == nil || string(payment.DestinationPaymentMethodDetails.Type) != "us_bank_account" {
		actual := ""
		if payment.DestinationPaymentMethodDetails != nil {
			actual = string(payment.DestinationPaymentMethodDetails.Type)
		}
		return fmt.Errorf("treasury outbound payment destination_payment_method_details.type mismatch: expected %q, got %q", "us_bank_account", actual)
	}
	if payment.DestinationPaymentMethodDetails == nil || payment.DestinationPaymentMethodDetails.USBankAccount == nil || payment.DestinationPaymentMethodDetails.USBankAccount.Last4 != "6789" {
		actual := ""
		if payment.DestinationPaymentMethodDetails != nil && payment.DestinationPaymentMethodDetails.USBankAccount != nil {
			actual = payment.DestinationPaymentMethodDetails.USBankAccount.Last4
		}
		return fmt.Errorf("treasury outbound payment destination_payment_method_details.us_bank_account.last4 mismatch: expected %q, got %q", "6789", actual)
	}
	return expectMetadataSubset("treasury_outbound_payment_basic.metadata", payment.Metadata, map[string]string{
		"suite": "sdk-codegen",
		"case":  "treasury_outbound_payment_basic",
	})
}

func findLatestTreasuryOutboundPaymentByCase(client *stripe.Client, caseName string) (*stripe.TreasuryOutboundPayment, error) {
	params := &stripe.TreasuryOutboundPaymentListParams{}
	params.Limit = stripe.Int64(25)
	for payment, err := range client.V1TreasuryOutboundPayments.List(context.Background(), params).All(context.Background()) {
		if err != nil {
			return nil, fmt.Errorf("list treasury outbound payments for %s: %w", caseName, err)
		}
		if payment != nil && payment.Metadata["suite"] == "sdk-codegen" && payment.Metadata["case"] == caseName {
			return payment, nil
		}
	}
	return nil, fmt.Errorf("treasury outbound payment for case %q not found", caseName)
}
