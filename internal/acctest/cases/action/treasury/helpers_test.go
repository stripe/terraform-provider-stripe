// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import (
	"context"
	"fmt"
	"testing"
	"time"

	stripe "github.com/stripe/stripe-go/v86"
	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

type treasuryActionFixtureIDs struct {
	SourceFinancialAccount      string
	DestinationFinancialAccount string
	OriginPaymentMethod         string
	ReceivedCredit              string
	ReceivedDebit               string
}

var treasuryActionFixtures treasuryActionFixtureIDs

func runRemainingActionCase(
	t *testing.T,
	name string,
	surface string,
	group string,
	requiredEnv []string,
	configTemplate string,
	renderEnv runner.ActionRenderEnv,
	configPreparer runner.ActionConfigPreparer,
	verify runner.StateVerifier,
) {
	t.Helper()

	runner.RunActionCase(t, runner.ActionCase{
		Definition: runner.CaseDefinition{
			Name:        name,
			Surface:     surface,
			Group:       group,
			Kind:        "action",
			RequiredEnv: requiredEnv,
		},
		ConfigTemplate: configTemplate,
		ConfigPreparer: configPreparer,
		RenderEnv:      renderEnv,
		VerifyInvoke:   verify,
	})
}

func expectMetadataSubset(
	name string,
	actual map[string]string,
	expected map[string]string,
) error {
	for key, expectedValue := range expected {
		actualValue, ok := actual[key]
		if !ok {
			return fmt.Errorf("%s missing metadata key %q", name, key)
		}
		if actualValue != expectedValue {
			return fmt.Errorf(
				"%s metadata[%q] mismatch: expected %q, got %q",
				name,
				key,
				expectedValue,
				actualValue,
			)
		}
	}

	return nil
}

func prepareTreasuryInboundTransferConfig(
	t *testing.T,
	env runner.TestEnv,
	client *stripe.Client,
) map[string]string {
	t.Helper()

	fixtures := ensureTreasuryActionFixtures(t, env, client, treasuryActionFixtureRequirements{
		SourceFinancialAccount: true,
		OriginPaymentMethod:    true,
	})
	return treasuryActionReplacements(fixtures)
}

func prepareTreasuryOutboundPaymentConfig(
	t *testing.T,
	env runner.TestEnv,
	client *stripe.Client,
) map[string]string {
	t.Helper()

	fixtures := ensureTreasuryActionFixtures(t, env, client, treasuryActionFixtureRequirements{
		SourceFinancialAccount: true,
		FundSourceAccount:      true,
	})
	return treasuryActionReplacements(fixtures)
}

func prepareTreasuryOutboundTransferConfig(
	t *testing.T,
	env runner.TestEnv,
	client *stripe.Client,
) map[string]string {
	t.Helper()

	fixtures := ensureTreasuryActionFixtures(t, env, client, treasuryActionFixtureRequirements{
		SourceFinancialAccount:      true,
		DestinationFinancialAccount: true,
		FundSourceAccount:           true,
	})
	return treasuryActionReplacements(fixtures)
}

func prepareTreasuryCreditReversalConfig(
	t *testing.T,
	env runner.TestEnv,
	client *stripe.Client,
) map[string]string {
	t.Helper()

	fixtures := ensureTreasuryActionFixtures(t, env, client, treasuryActionFixtureRequirements{
		SourceFinancialAccount: true,
		ReceivedCredit:         true,
	})
	return treasuryActionReplacements(fixtures)
}

func prepareTreasuryDebitReversalConfig(
	t *testing.T,
	env runner.TestEnv,
	client *stripe.Client,
) map[string]string {
	t.Helper()

	fixtures := ensureTreasuryActionFixtures(t, env, client, treasuryActionFixtureRequirements{
		SourceFinancialAccount: true,
		FundSourceAccount:      true,
		ReceivedDebit:          true,
	})
	return treasuryActionReplacements(fixtures)
}

type treasuryActionFixtureRequirements struct {
	SourceFinancialAccount      bool
	DestinationFinancialAccount bool
	OriginPaymentMethod         bool
	FundSourceAccount           bool
	ReceivedCredit              bool
	ReceivedDebit               bool
}

func ensureTreasuryActionFixtures(
	t *testing.T,
	env runner.TestEnv,
	client *stripe.Client,
	requirements treasuryActionFixtureRequirements,
) treasuryActionFixtureIDs {
	t.Helper()

	fixtures := treasuryActionFixtureIDs{
		SourceFinancialAccount:      env.TreasurySourceFinancialAccount,
		DestinationFinancialAccount: env.TreasuryDestinationFinancialAccount,
		OriginPaymentMethod:         env.TreasuryOriginPaymentMethod,
		ReceivedCredit:              env.TreasuryReceivedCredit,
		ReceivedDebit:               env.TreasuryReceivedDebit,
	}

	if requirements.SourceFinancialAccount && fixtures.SourceFinancialAccount == "" {
		fixtures.SourceFinancialAccount = ensureTreasuryFinancialAccount(t, client, "source", "")
	}
	if requirements.DestinationFinancialAccount && fixtures.DestinationFinancialAccount == "" {
		fixtures.DestinationFinancialAccount = ensureTreasuryFinancialAccount(t, client, "destination", fixtures.SourceFinancialAccount)
	}
	if requirements.OriginPaymentMethod && fixtures.OriginPaymentMethod == "" {
		fixtures.OriginPaymentMethod = createTreasuryOriginPaymentMethod(t, client)
	}
	if requirements.FundSourceAccount {
		createTreasuryReceivedCredit(t, client, fixtures.SourceFinancialAccount, 5000)
	}
	if requirements.ReceivedCredit && fixtures.ReceivedCredit == "" {
		fixtures.ReceivedCredit = createTreasuryReceivedCredit(t, client, fixtures.SourceFinancialAccount, 320)
	}
	if requirements.ReceivedDebit && fixtures.ReceivedDebit == "" {
		fixtures.ReceivedDebit = createTreasuryReceivedDebit(t, client, fixtures.SourceFinancialAccount, 210)
	}

	treasuryActionFixtures = fixtures
	return fixtures
}

func treasuryActionReplacements(fixtures treasuryActionFixtureIDs) map[string]string {
	return map[string]string{
		"{{TREASURY_SOURCE_FINANCIAL_ACCOUNT}}":      fixtures.SourceFinancialAccount,
		"{{TREASURY_DESTINATION_FINANCIAL_ACCOUNT}}": fixtures.DestinationFinancialAccount,
		"{{TREASURY_ORIGIN_PAYMENT_METHOD}}":         fixtures.OriginPaymentMethod,
		"{{TREASURY_RECEIVED_CREDIT}}":               fixtures.ReceivedCredit,
		"{{TREASURY_RECEIVED_DEBIT}}":                fixtures.ReceivedDebit,
	}
}

func ensureTreasuryFinancialAccount(
	t *testing.T,
	client *stripe.Client,
	role string,
	excludeID string,
) string {
	t.Helper()

	openFinancialAccounts := listOpenTreasuryFinancialAccounts(t, client)
	for _, financialAccount := range openFinancialAccounts {
		if financialAccount.ID == excludeID {
			continue
		}
		if financialAccount.Metadata["suite"] == "sdk-codegen" && financialAccount.Metadata["role"] == role {
			requestTreasuryFinancialAccountFeatures(t, client, financialAccount.ID)
			waitTreasuryFinancialAccountFeaturesActive(t, client, financialAccount.ID)
			return financialAccount.ID
		}
	}
	for _, financialAccount := range openFinancialAccounts {
		if financialAccount.ID == excludeID {
			continue
		}
		requestTreasuryFinancialAccountFeatures(t, client, financialAccount.ID)
		waitTreasuryFinancialAccountFeaturesActive(t, client, financialAccount.ID)
		return financialAccount.ID
	}

	financialAccount, err := client.V1TreasuryFinancialAccounts.Create(
		context.Background(),
		&stripe.TreasuryFinancialAccountCreateParams{
			Features: treasuryFinancialAccountFeatures(),
			Metadata: map[string]string{
				"suite": "sdk-codegen",
				"role":  role,
			},
			Nickname:            stripe.String(fmt.Sprintf("sdk-codegen-%s-%d", role, time.Now().UnixNano())),
			SupportedCurrencies: []*string{stripe.String(string(stripe.CurrencyUSD))},
		},
	)
	if err != nil {
		t.Fatalf("create treasury %s financial account: %v", role, err)
	}
	waitTreasuryFinancialAccountFeaturesActive(t, client, financialAccount.ID)

	return financialAccount.ID
}

func listOpenTreasuryFinancialAccounts(
	t *testing.T,
	client *stripe.Client,
) []*stripe.TreasuryFinancialAccount {
	t.Helper()

	params := &stripe.TreasuryFinancialAccountListParams{
		Status: stripe.String(string(stripe.TreasuryFinancialAccountStatusOpen)),
	}
	params.Limit = stripe.Int64(25)
	financialAccounts := []*stripe.TreasuryFinancialAccount{}
	for financialAccount, err := range client.V1TreasuryFinancialAccounts.List(context.Background(), params).All(context.Background()) {
		if err != nil {
			t.Fatalf("list treasury financial accounts: %v", err)
		}
		if financialAccount == nil {
			continue
		}
		financialAccounts = append(financialAccounts, financialAccount)
	}

	return financialAccounts
}

func requestTreasuryFinancialAccountFeatures(
	t *testing.T,
	client *stripe.Client,
	financialAccountID string,
) {
	t.Helper()

	_, err := client.V1TreasuryFinancialAccounts.UpdateFeatures(
		context.Background(),
		financialAccountID,
		treasuryFinancialAccountUpdateFeatures(),
	)
	if err != nil {
		t.Fatalf("request treasury financial account features for %s: %v", financialAccountID, err)
	}
}

func waitTreasuryFinancialAccountFeaturesActive(
	t *testing.T,
	client *stripe.Client,
	financialAccountID string,
) {
	t.Helper()

	deadline := time.Now().Add(90 * time.Second)
	var lastStatus string
	for {
		features, err := client.V1TreasuryFinancialAccounts.RetrieveFeatures(
			context.Background(),
			financialAccountID,
			nil,
		)
		if err != nil {
			t.Fatalf("retrieve treasury financial account features for %s: %v", financialAccountID, err)
		}
		if treasuryFinancialAccountFeaturesActive(features) {
			return
		}
		lastStatus = treasuryFinancialAccountFeatureStatusSummary(features)
		if time.Now().After(deadline) {
			t.Fatalf("timed out waiting for treasury financial account %s features to become active; last status: %s", financialAccountID, lastStatus)
		}
		time.Sleep(2 * time.Second)
	}
}

func treasuryFinancialAccountFeaturesActive(features *stripe.TreasuryFinancialAccountFeatures) bool {
	if features == nil ||
		features.FinancialAddresses == nil ||
		features.FinancialAddresses.ABA == nil ||
		features.InboundTransfers == nil ||
		features.InboundTransfers.ACH == nil ||
		features.IntraStripeFlows == nil ||
		features.OutboundPayments == nil ||
		features.OutboundPayments.ACH == nil ||
		features.OutboundTransfers == nil ||
		features.OutboundTransfers.ACH == nil {
		return false
	}

	return features.FinancialAddresses.ABA.Status == stripe.TreasuryFinancialAccountFeaturesFinancialAddressesABAStatusActive &&
		features.InboundTransfers.ACH.Status == stripe.TreasuryFinancialAccountFeaturesInboundTransfersACHStatusActive &&
		features.IntraStripeFlows.Status == stripe.TreasuryFinancialAccountFeaturesIntraStripeFlowsStatusActive &&
		features.OutboundPayments.ACH.Status == stripe.TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatusActive &&
		features.OutboundTransfers.ACH.Status == stripe.TreasuryFinancialAccountFeaturesOutboundTransfersACHStatusActive
}

func treasuryFinancialAccountFeatureStatusSummary(features *stripe.TreasuryFinancialAccountFeatures) string {
	if features == nil {
		return "missing features"
	}

	statuses := map[string]string{}
	if features.FinancialAddresses != nil && features.FinancialAddresses.ABA != nil {
		statuses["financial_addresses.aba"] = string(features.FinancialAddresses.ABA.Status)
	}
	if features.InboundTransfers != nil && features.InboundTransfers.ACH != nil {
		statuses["inbound_transfers.ach"] = string(features.InboundTransfers.ACH.Status)
	}
	if features.IntraStripeFlows != nil {
		statuses["intra_stripe_flows"] = string(features.IntraStripeFlows.Status)
	}
	if features.OutboundPayments != nil && features.OutboundPayments.ACH != nil {
		statuses["outbound_payments.ach"] = string(features.OutboundPayments.ACH.Status)
	}
	if features.OutboundTransfers != nil && features.OutboundTransfers.ACH != nil {
		statuses["outbound_transfers.ach"] = string(features.OutboundTransfers.ACH.Status)
	}

	return fmt.Sprintf("%v", statuses)
}

func treasuryFinancialAccountUpdateFeatures() *stripe.TreasuryFinancialAccountUpdateFeaturesParams {
	return &stripe.TreasuryFinancialAccountUpdateFeaturesParams{
		FinancialAddresses: &stripe.TreasuryFinancialAccountUpdateFeaturesFinancialAddressesParams{
			ABA: &stripe.TreasuryFinancialAccountUpdateFeaturesFinancialAddressesABAParams{
				Requested: stripe.Bool(true),
			},
		},
		InboundTransfers: &stripe.TreasuryFinancialAccountUpdateFeaturesInboundTransfersParams{
			ACH: &stripe.TreasuryFinancialAccountUpdateFeaturesInboundTransfersACHParams{
				Requested: stripe.Bool(true),
			},
		},
		IntraStripeFlows: &stripe.TreasuryFinancialAccountUpdateFeaturesIntraStripeFlowsParams{
			Requested: stripe.Bool(true),
		},
		OutboundPayments: &stripe.TreasuryFinancialAccountUpdateFeaturesOutboundPaymentsParams{
			ACH: &stripe.TreasuryFinancialAccountUpdateFeaturesOutboundPaymentsACHParams{
				Requested: stripe.Bool(true),
			},
		},
		OutboundTransfers: &stripe.TreasuryFinancialAccountUpdateFeaturesOutboundTransfersParams{
			ACH: &stripe.TreasuryFinancialAccountUpdateFeaturesOutboundTransfersACHParams{
				Requested: stripe.Bool(true),
			},
		},
	}
}

func treasuryFinancialAccountFeatures() *stripe.TreasuryFinancialAccountCreateFeaturesParams {
	return &stripe.TreasuryFinancialAccountCreateFeaturesParams{
		FinancialAddresses: &stripe.TreasuryFinancialAccountCreateFeaturesFinancialAddressesParams{
			ABA: &stripe.TreasuryFinancialAccountCreateFeaturesFinancialAddressesABAParams{
				Requested: stripe.Bool(true),
			},
		},
		InboundTransfers: &stripe.TreasuryFinancialAccountCreateFeaturesInboundTransfersParams{
			ACH: &stripe.TreasuryFinancialAccountCreateFeaturesInboundTransfersACHParams{
				Requested: stripe.Bool(true),
			},
		},
		IntraStripeFlows: &stripe.TreasuryFinancialAccountCreateFeaturesIntraStripeFlowsParams{
			Requested: stripe.Bool(true),
		},
		OutboundPayments: &stripe.TreasuryFinancialAccountCreateFeaturesOutboundPaymentsParams{
			ACH: &stripe.TreasuryFinancialAccountCreateFeaturesOutboundPaymentsACHParams{
				Requested: stripe.Bool(true),
			},
		},
		OutboundTransfers: &stripe.TreasuryFinancialAccountCreateFeaturesOutboundTransfersParams{
			ACH: &stripe.TreasuryFinancialAccountCreateFeaturesOutboundTransfersACHParams{
				Requested: stripe.Bool(true),
			},
		},
	}
}

func createTreasuryOriginPaymentMethod(t *testing.T, client *stripe.Client) string {
	t.Helper()

	setupIntent, err := client.V1SetupIntents.Create(
		context.Background(),
		&stripe.SetupIntentCreateParams{
			AttachToSelf: stripe.Bool(true),
			Confirm:      stripe.Bool(true),
			FlowDirections: []*string{
				stripe.String(string(stripe.SetupIntentFlowDirectionInbound)),
			},
			MandateData: &stripe.SetupIntentCreateMandateDataParams{
				CustomerAcceptance: &stripe.SetupIntentCreateMandateDataCustomerAcceptanceParams{
					Type: stripe.String("offline"),
				},
			},
			PaymentMethodData: &stripe.SetupIntentCreatePaymentMethodDataParams{
				BillingDetails: &stripe.SetupIntentCreatePaymentMethodDataBillingDetailsParams{
					Name: stripe.String("SDK Codegen Treasury Origin"),
				},
				Type: stripe.String(string(stripe.PaymentMethodTypeUSBankAccount)),
				USBankAccount: &stripe.SetupIntentCreatePaymentMethodDataUSBankAccountParams{
					AccountHolderType: stripe.String("individual"),
					AccountNumber:     stripe.String("000123456789"),
					RoutingNumber:     stripe.String("110000000"),
				},
			},
			PaymentMethodOptions: &stripe.SetupIntentCreatePaymentMethodOptionsParams{
				USBankAccount: &stripe.SetupIntentCreatePaymentMethodOptionsUSBankAccountParams{
					VerificationMethod: stripe.String(string(stripe.SetupIntentPaymentMethodOptionsUSBankAccountVerificationMethodMicrodeposits)),
				},
			},
			PaymentMethodTypes: []*string{stripe.String(string(stripe.PaymentMethodTypeUSBankAccount))},
		},
	)
	if err != nil {
		t.Fatalf("create treasury origin setup intent: %v", err)
	}
	setupIntent, err = client.V1SetupIntents.VerifyMicrodeposits(
		context.Background(),
		setupIntent.ID,
		&stripe.SetupIntentVerifyMicrodepositsParams{
			Amounts: []*int64{stripe.Int64(32), stripe.Int64(45)},
		},
	)
	if err != nil {
		t.Fatalf("verify treasury origin setup intent microdeposits: %v", err)
	}
	if setupIntent.PaymentMethod == nil {
		t.Fatalf("verified treasury origin setup intent %s did not return a payment method", setupIntent.ID)
	}

	return setupIntent.PaymentMethod.ID
}

func createTreasuryReceivedCredit(
	t *testing.T,
	client *stripe.Client,
	financialAccount string,
	amount int64,
) string {
	t.Helper()

	receivedCredit, err := client.V1TestHelpersTreasuryReceivedCredits.Create(
		context.Background(),
		&stripe.TestHelpersTreasuryReceivedCreditCreateParams{
			Amount:           stripe.Int64(amount),
			Currency:         stripe.String(string(stripe.CurrencyUSD)),
			Description:      stripe.String("sdkgen"),
			FinancialAccount: stripe.String(financialAccount),
			Network:          stripe.String(string(stripe.TreasuryReceivedCreditNetworkACH)),
			InitiatingPaymentMethodDetails: &stripe.TestHelpersTreasuryReceivedCreditCreateInitiatingPaymentMethodDetailsParams{
				Type: stripe.String("us_bank_account"),
				USBankAccount: &stripe.TestHelpersTreasuryReceivedCreditCreateInitiatingPaymentMethodDetailsUSBankAccountParams{
					AccountHolderName: stripe.String("SDK Codegen"),
					AccountNumber:     stripe.String("000123456789"),
					RoutingNumber:     stripe.String("110000000"),
				},
			},
		},
	)
	if err != nil {
		t.Fatalf("create treasury received credit fixture: %v", err)
	}

	return receivedCredit.ID
}

func createTreasuryReceivedDebit(
	t *testing.T,
	client *stripe.Client,
	financialAccount string,
	amount int64,
) string {
	t.Helper()

	receivedDebit, err := client.V1TestHelpersTreasuryReceivedDebits.Create(
		context.Background(),
		&stripe.TestHelpersTreasuryReceivedDebitCreateParams{
			Amount:           stripe.Int64(amount),
			Currency:         stripe.String(string(stripe.CurrencyUSD)),
			Description:      stripe.String("sdkgen"),
			FinancialAccount: stripe.String(financialAccount),
			Network:          stripe.String("ach"),
			InitiatingPaymentMethodDetails: &stripe.TestHelpersTreasuryReceivedDebitCreateInitiatingPaymentMethodDetailsParams{
				Type: stripe.String("us_bank_account"),
				USBankAccount: &stripe.TestHelpersTreasuryReceivedDebitCreateInitiatingPaymentMethodDetailsUSBankAccountParams{
					AccountHolderName: stripe.String("SDK Codegen"),
					AccountNumber:     stripe.String("000123456789"),
					RoutingNumber:     stripe.String("110000000"),
				},
			},
		},
	)
	if err != nil {
		t.Fatalf("create treasury received debit fixture: %v", err)
	}

	return receivedDebit.ID
}

func expectedTreasurySourceFinancialAccount(env runner.TestEnv) string {
	if env.TreasurySourceFinancialAccount != "" {
		return env.TreasurySourceFinancialAccount
	}
	return treasuryActionFixtures.SourceFinancialAccount
}

func expectedTreasuryDestinationFinancialAccount(env runner.TestEnv) string {
	if env.TreasuryDestinationFinancialAccount != "" {
		return env.TreasuryDestinationFinancialAccount
	}
	return treasuryActionFixtures.DestinationFinancialAccount
}

func expectedTreasuryOriginPaymentMethod(env runner.TestEnv) string {
	if env.TreasuryOriginPaymentMethod != "" {
		return env.TreasuryOriginPaymentMethod
	}
	return treasuryActionFixtures.OriginPaymentMethod
}

func expectedTreasuryReceivedCredit(env runner.TestEnv) string {
	if env.TreasuryReceivedCredit != "" {
		return env.TreasuryReceivedCredit
	}
	return treasuryActionFixtures.ReceivedCredit
}

func expectedTreasuryReceivedDebit(env runner.TestEnv) string {
	if env.TreasuryReceivedDebit != "" {
		return env.TreasuryReceivedDebit
	}
	return treasuryActionFixtures.ReceivedDebit
}
