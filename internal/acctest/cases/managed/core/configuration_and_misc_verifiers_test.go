// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	stripe "github.com/stripe/stripe-go/v86"

	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

type creditNoteExpectations struct {
	Address                   string
	ExpectedAmount            int64
	ExpectedEffectiveAt       int64
	ExpectedEffectiveAtOffset int64
	ExpectedMemo              string
	ExpectedReason            string
	ExpectedStatus            string
	ExpectedMetadata          map[string]string
}

type customerBalanceTransactionExpectations struct {
	Address                 string
	ExpectedCustomerAddress string
	ExpectedAmount          int64
	ExpectedCurrency        string
	ExpectedDescription     string
	ExpectedType            string
	ExpectedMetadata        map[string]string
}

type fileLinkExpectations struct {
	Address             string
	ExpectedFileAddress string
	ExpectedExpiresAt   int64
	ExpectedMetadata    map[string]string
}

type taxIDExpectations struct {
	Address                 string
	ExpectedCustomerAddress string
	ExpectedType            string
	ExpectedValue           string
	ExpectedCountry         string
}

type billingPortalConfigurationExpectations struct {
	Address                                     string
	ExpectedActive                              bool
	ExpectedName                                string
	ExpectedDefaultReturnURL                    string
	ExpectedHeadline                            string
	CheckLoginPage                              bool
	ExpectedLoginPageEnabled                    bool
	ExpectedPrivacyPolicyURL                    string
	ExpectedTermsOfServiceURL                   string
	CheckSubscriptionCancel                     bool
	ExpectedSubscriptionCancelEnabled           bool
	ExpectedSubscriptionCancelMode              string
	ExpectedSubscriptionCancelProrationBehavior string
	ExpectedCancellationReasonEnabled           bool
	ExpectedCancellationReasonOptions           []string
	ExpectedMetadata                            map[string]string
	ExpectedCustomerUpdateEnabled               bool
	ExpectedAllowedUpdates                      []string
	ExpectedInvoiceHistoryEnabled               bool
	ExpectedPaymentMethodUpdateEnabled          bool
}

type paymentMethodConfigurationExpectations struct {
	Address                    string
	ExpectedActive             bool
	ExpectedName               string
	ExpectedCardPreference     string
	ExpectedCardValue          string
	ExpectedApplePayPreference string
	ExpectedApplePayValue      string
	ExpectedLinkPreference     string
	ExpectedLinkValue          string
}

type paymentMethodDomainExpectations struct {
	Address         string
	ExpectedEnabled bool
}

type productFeatureExpectations struct {
	Address                           string
	ExpectedProductAddress            string
	ExpectedEntitlementFeatureAddress string
}

type testHelpersTestClockExpectations struct {
	Address                 string
	ExpectedCustomerAddress string
	ExpectedName            string
	ExpectedFrozenTime      int64
	ExpectedStatus          string
}

type applePayDomainExpectations struct {
	Address string
}

func verifyCreditNote(expect creditNoteExpectations) runner.StateVerifier {
	return func(
		env runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		creditNote, err := retrieveCreditNote(client, state, expect.Address)
		if err != nil {
			return err
		}
		if env.CreditNoteInvoice != "" {
			if creditNote.Invoice == nil || creditNote.Invoice.ID != env.CreditNoteInvoice {
				actualInvoiceID := ""
				if creditNote.Invoice != nil {
					actualInvoiceID = creditNote.Invoice.ID
				}
				return fmt.Errorf(
					"remote %s.invoice mismatch: expected %q, got %q",
					expect.Address,
					env.CreditNoteInvoice,
					actualInvoiceID,
				)
			}
		}
		if creditNote.Memo != expect.ExpectedMemo {
			return fmt.Errorf(
				"remote %s.memo mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedMemo,
				creditNote.Memo,
			)
		}
		if expect.ExpectedAmount != 0 && creditNote.Amount != expect.ExpectedAmount {
			return fmt.Errorf(
				"remote %s.amount mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedAmount,
				creditNote.Amount,
			)
		}
		expectedEffectiveAt := expect.ExpectedEffectiveAt
		if expect.ExpectedEffectiveAtOffset != 0 {
			if env.CreditNoteInvoiceFinalizedAt == "" {
				return fmt.Errorf(
					"missing STRIPE_CREDIT_NOTE_INVOICE_FINALIZED_AT for %s effective_at verification",
					expect.Address,
				)
			}
			finalizedAt, err := strconv.ParseInt(env.CreditNoteInvoiceFinalizedAt, 10, 64)
			if err != nil {
				return fmt.Errorf(
					"invalid STRIPE_CREDIT_NOTE_INVOICE_FINALIZED_AT %q: %w",
					env.CreditNoteInvoiceFinalizedAt,
					err,
				)
			}
			expectedEffectiveAt = finalizedAt + expect.ExpectedEffectiveAtOffset
		}
		if expectedEffectiveAt != 0 && creditNote.EffectiveAt != expectedEffectiveAt {
			return fmt.Errorf(
				"remote %s.effective_at mismatch: expected %d, got %d",
				expect.Address,
				expectedEffectiveAt,
				creditNote.EffectiveAt,
			)
		}
		if expect.ExpectedReason != "" && string(creditNote.Reason) != expect.ExpectedReason {
			return fmt.Errorf(
				"remote %s.reason mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedReason,
				string(creditNote.Reason),
			)
		}
		if string(creditNote.Status) != expect.ExpectedStatus {
			return fmt.Errorf(
				"remote %s.status mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedStatus,
				string(creditNote.Status),
			)
		}
		if creditNote.PDF == "" {
			return fmt.Errorf("remote %s.pdf missing", expect.Address)
		}
		if err := expectMetadataSubset(
			expect.Address+".metadata",
			creditNote.Metadata,
			expect.ExpectedMetadata,
		); err != nil {
			return err
		}

		return nil
	}
}

func verifyCreditNoteDestroyStillExists(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	_, err := retrieveCreditNote(client, state, "stripe_credit_note.test")
	return err
}

func verifyCustomerBalanceTransaction(
	expect customerBalanceTransactionExpectations,
) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		transaction, err := retrieveCustomerBalanceTransaction(client, state, expect.Address)
		if err != nil {
			return err
		}
		expectedCustomerID, err := runner.ResourcePrimaryID(state, expect.ExpectedCustomerAddress)
		if err != nil {
			return err
		}
		if transaction.Customer == nil || transaction.Customer.ID != expectedCustomerID {
			actualCustomerID := ""
			if transaction.Customer != nil {
				actualCustomerID = transaction.Customer.ID
			}
			return fmt.Errorf(
				"remote %s.customer mismatch: expected %q, got %q",
				expect.Address,
				expectedCustomerID,
				actualCustomerID,
			)
		}
		if transaction.Amount != expect.ExpectedAmount {
			return fmt.Errorf(
				"remote %s.amount mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedAmount,
				transaction.Amount,
			)
		}
		if string(transaction.Currency) != expect.ExpectedCurrency {
			return fmt.Errorf(
				"remote %s.currency mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedCurrency,
				string(transaction.Currency),
			)
		}
		if transaction.Description != expect.ExpectedDescription {
			return fmt.Errorf(
				"remote %s.description mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedDescription,
				transaction.Description,
			)
		}
		if string(transaction.Type) != expect.ExpectedType {
			return fmt.Errorf(
				"remote %s.type mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedType,
				string(transaction.Type),
			)
		}
		if err := expectMetadataSubset(
			expect.Address+".metadata",
			transaction.Metadata,
			expect.ExpectedMetadata,
		); err != nil {
			return err
		}

		return nil
	}
}

func verifyCustomerBalanceTransactionDestroyStillExists(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	_, err := retrieveCustomerBalanceTransaction(
		client,
		state,
		"stripe_customer_balance_transaction.test",
	)
	if err != nil && strings.Contains(err.Error(), "No such customer") {
		return nil
	}
	return err
}

func verifyFileLink(expect fileLinkExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		fileLink, err := retrieveFileLink(client, state, expect.Address)
		if err != nil {
			return err
		}
		expectedFileID, err := runner.ResourcePrimaryID(state, expect.ExpectedFileAddress)
		if err != nil {
			return err
		}
		if fileLink.File == nil || fileLink.File.ID != expectedFileID {
			actualFileID := ""
			if fileLink.File != nil {
				actualFileID = fileLink.File.ID
			}
			return fmt.Errorf(
				"remote %s.file mismatch: expected %q, got %q",
				expect.Address,
				expectedFileID,
				actualFileID,
			)
		}
		if fileLink.ExpiresAt != expect.ExpectedExpiresAt {
			return fmt.Errorf(
				"remote %s.expires_at mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedExpiresAt,
				fileLink.ExpiresAt,
			)
		}
		if fileLink.URL == "" {
			return fmt.Errorf("remote %s.url missing", expect.Address)
		}
		if fileLink.Expired {
			return fmt.Errorf("remote %s unexpectedly expired", expect.Address)
		}
		if err := expectMetadataSubset(
			expect.Address+".metadata",
			fileLink.Metadata,
			expect.ExpectedMetadata,
		); err != nil {
			return err
		}

		return nil
	}
}

func verifyFileLinkDestroyStillExists(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	_, err := retrieveFileLink(client, state, "stripe_file_link.test")
	return err
}

func verifyTaxID(expect taxIDExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		taxID, err := retrieveTaxID(client, state, expect.Address)
		if err != nil {
			return err
		}
		expectedCustomerID, err := runner.ResourcePrimaryID(state, expect.ExpectedCustomerAddress)
		if err != nil {
			return err
		}
		if taxID.Customer == nil || taxID.Customer.ID != expectedCustomerID {
			actualCustomerID := ""
			if taxID.Customer != nil {
				actualCustomerID = taxID.Customer.ID
			}
			return fmt.Errorf(
				"remote %s.customer mismatch: expected %q, got %q",
				expect.Address,
				expectedCustomerID,
				actualCustomerID,
			)
		}
		if string(taxID.Type) != expect.ExpectedType {
			return fmt.Errorf(
				"remote %s.type mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedType,
				string(taxID.Type),
			)
		}
		if taxID.Value != expect.ExpectedValue {
			return fmt.Errorf(
				"remote %s.value mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedValue,
				taxID.Value,
			)
		}
		if taxID.Country != expect.ExpectedCountry {
			return fmt.Errorf(
				"remote %s.country mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedCountry,
				taxID.Country,
			)
		}

		return nil
	}
}

func verifyTaxIDDestroyMissing(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	id, err := runner.ResourcePrimaryID(state, "stripe_tax_id.test")
	if err != nil {
		return err
	}
	customerID, err := runner.ResourceAttribute(state, "stripe_tax_id.test", "customer")
	if err != nil {
		return err
	}

	params := &stripe.TaxIDRetrieveParams{
		Customer: stripe.String(customerID),
	}
	taxID, err := client.V1TaxIDs.Retrieve(context.Background(), id, params)
	if err == nil && taxID != nil && taxID.Deleted {
		return nil
	}
	return expectRemoteMissing("stripe_tax_id.test", id, err)
}

func verifyBillingPortalConfiguration(
	expect billingPortalConfigurationExpectations,
) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		configuration, err := retrieveBillingPortalConfiguration(client, state, expect.Address)
		if err != nil {
			return err
		}
		if configuration.Active != expect.ExpectedActive {
			return fmt.Errorf(
				"remote %s.active mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedActive,
				configuration.Active,
			)
		}
		if configuration.Name != expect.ExpectedName {
			return fmt.Errorf(
				"remote %s.name mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedName,
				configuration.Name,
			)
		}
		if configuration.DefaultReturnURL != expect.ExpectedDefaultReturnURL {
			return fmt.Errorf(
				"remote %s.default_return_url mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedDefaultReturnURL,
				configuration.DefaultReturnURL,
			)
		}
		if configuration.BusinessProfile == nil {
			return fmt.Errorf("remote %s.business_profile missing", expect.Address)
		}
		if configuration.BusinessProfile.Headline != expect.ExpectedHeadline {
			return fmt.Errorf(
				"remote %s.business_profile.headline mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedHeadline,
				configuration.BusinessProfile.Headline,
			)
		}
		if expect.ExpectedPrivacyPolicyURL != "" &&
			configuration.BusinessProfile.PrivacyPolicyURL != expect.ExpectedPrivacyPolicyURL {
			return fmt.Errorf(
				"remote %s.business_profile.privacy_policy_url mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedPrivacyPolicyURL,
				configuration.BusinessProfile.PrivacyPolicyURL,
			)
		}
		if expect.ExpectedTermsOfServiceURL != "" &&
			configuration.BusinessProfile.TermsOfServiceURL != expect.ExpectedTermsOfServiceURL {
			return fmt.Errorf(
				"remote %s.business_profile.terms_of_service_url mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedTermsOfServiceURL,
				configuration.BusinessProfile.TermsOfServiceURL,
			)
		}
		if configuration.Features == nil {
			return fmt.Errorf("remote %s.features missing", expect.Address)
		}
		if configuration.Features.CustomerUpdate == nil {
			return fmt.Errorf("remote %s.features.customer_update missing", expect.Address)
		}
		if configuration.Features.CustomerUpdate.Enabled != expect.ExpectedCustomerUpdateEnabled {
			return fmt.Errorf(
				"remote %s.features.customer_update.enabled mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedCustomerUpdateEnabled,
				configuration.Features.CustomerUpdate.Enabled,
			)
		}
		actualAllowedUpdates := make([]string, 0, len(configuration.Features.CustomerUpdate.AllowedUpdates))
		for _, allowedUpdate := range configuration.Features.CustomerUpdate.AllowedUpdates {
			actualAllowedUpdates = append(actualAllowedUpdates, string(allowedUpdate))
		}
		if err := expectRemoteStringList(
			expect.Address+".features.customer_update.allowed_updates",
			actualAllowedUpdates,
			expect.ExpectedAllowedUpdates,
		); err != nil {
			return err
		}
		if configuration.Features.InvoiceHistory == nil {
			return fmt.Errorf("remote %s.features.invoice_history missing", expect.Address)
		}
		if configuration.Features.InvoiceHistory.Enabled != expect.ExpectedInvoiceHistoryEnabled {
			return fmt.Errorf(
				"remote %s.features.invoice_history.enabled mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedInvoiceHistoryEnabled,
				configuration.Features.InvoiceHistory.Enabled,
			)
		}
		if configuration.Features.PaymentMethodUpdate == nil {
			return fmt.Errorf("remote %s.features.payment_method_update missing", expect.Address)
		}
		if configuration.Features.PaymentMethodUpdate.Enabled != expect.ExpectedPaymentMethodUpdateEnabled {
			return fmt.Errorf(
				"remote %s.features.payment_method_update.enabled mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedPaymentMethodUpdateEnabled,
				configuration.Features.PaymentMethodUpdate.Enabled,
			)
		}
		if expect.CheckSubscriptionCancel {
			if configuration.Features.SubscriptionCancel == nil {
				return fmt.Errorf("remote %s.features.subscription_cancel missing", expect.Address)
			}
			if configuration.Features.SubscriptionCancel.Enabled != expect.ExpectedSubscriptionCancelEnabled {
				return fmt.Errorf(
					"remote %s.features.subscription_cancel.enabled mismatch: expected %t, got %t",
					expect.Address,
					expect.ExpectedSubscriptionCancelEnabled,
					configuration.Features.SubscriptionCancel.Enabled,
				)
			}
			if string(configuration.Features.SubscriptionCancel.Mode) != expect.ExpectedSubscriptionCancelMode {
				return fmt.Errorf(
					"remote %s.features.subscription_cancel.mode mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedSubscriptionCancelMode,
					string(configuration.Features.SubscriptionCancel.Mode),
				)
			}
			if string(configuration.Features.SubscriptionCancel.ProrationBehavior) != expect.ExpectedSubscriptionCancelProrationBehavior {
				return fmt.Errorf(
					"remote %s.features.subscription_cancel.proration_behavior mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedSubscriptionCancelProrationBehavior,
					string(configuration.Features.SubscriptionCancel.ProrationBehavior),
				)
			}
			if configuration.Features.SubscriptionCancel.CancellationReason == nil {
				return fmt.Errorf("remote %s.features.subscription_cancel.cancellation_reason missing", expect.Address)
			}
			if configuration.Features.SubscriptionCancel.CancellationReason.Enabled != expect.ExpectedCancellationReasonEnabled {
				return fmt.Errorf(
					"remote %s.features.subscription_cancel.cancellation_reason.enabled mismatch: expected %t, got %t",
					expect.Address,
					expect.ExpectedCancellationReasonEnabled,
					configuration.Features.SubscriptionCancel.CancellationReason.Enabled,
				)
			}
			actualOptions := make([]string, 0, len(configuration.Features.SubscriptionCancel.CancellationReason.Options))
			for _, option := range configuration.Features.SubscriptionCancel.CancellationReason.Options {
				actualOptions = append(actualOptions, string(option))
			}
			if err := expectRemoteStringList(
				expect.Address+".features.subscription_cancel.cancellation_reason.options",
				actualOptions,
				expect.ExpectedCancellationReasonOptions,
			); err != nil {
				return err
			}
		}
		if expect.CheckLoginPage {
			if configuration.LoginPage == nil {
				return fmt.Errorf("remote %s.login_page missing", expect.Address)
			}
			if configuration.LoginPage.Enabled != expect.ExpectedLoginPageEnabled {
				return fmt.Errorf(
					"remote %s.login_page.enabled mismatch: expected %t, got %t",
					expect.Address,
					expect.ExpectedLoginPageEnabled,
					configuration.LoginPage.Enabled,
				)
			}
		}
		if err := expectMetadataSubset(
			expect.Address+".metadata",
			configuration.Metadata,
			expect.ExpectedMetadata,
		); err != nil {
			return err
		}

		return nil
	}
}

func verifyBillingPortalConfigurationDestroyInactive(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	configuration, err := retrieveBillingPortalConfiguration(
		client,
		state,
		"stripe_billing_portal_configuration.test",
	)
	if err != nil {
		return err
	}
	if configuration.Active {
		return fmt.Errorf(
			"expected stripe_billing_portal_configuration.test to be inactive after destroy",
		)
	}

	return nil
}

func verifyPaymentMethodConfiguration(
	expect paymentMethodConfigurationExpectations,
) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		configuration, err := retrievePaymentMethodConfiguration(client, state, expect.Address)
		if err != nil {
			return err
		}
		if configuration.Active != expect.ExpectedActive {
			return fmt.Errorf(
				"remote %s.active mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedActive,
				configuration.Active,
			)
		}
		if configuration.Name != expect.ExpectedName {
			return fmt.Errorf(
				"remote %s.name mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedName,
				configuration.Name,
			)
		}
		if configuration.Card == nil || configuration.Card.DisplayPreference == nil {
			return fmt.Errorf("remote %s.card.display_preference missing", expect.Address)
		}
		if string(configuration.Card.DisplayPreference.Preference) != expect.ExpectedCardPreference {
			return fmt.Errorf(
				"remote %s.card.display_preference.preference mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedCardPreference,
				string(configuration.Card.DisplayPreference.Preference),
			)
		}
		if string(configuration.Card.DisplayPreference.Value) != expect.ExpectedCardValue {
			return fmt.Errorf(
				"remote %s.card.display_preference.value mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedCardValue,
				string(configuration.Card.DisplayPreference.Value),
			)
		}
		if expect.ExpectedApplePayPreference != "" {
			if configuration.ApplePay == nil || configuration.ApplePay.DisplayPreference == nil {
				return fmt.Errorf("remote %s.apple_pay.display_preference missing", expect.Address)
			}
			if string(configuration.ApplePay.DisplayPreference.Preference) != expect.ExpectedApplePayPreference {
				return fmt.Errorf(
					"remote %s.apple_pay.display_preference.preference mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedApplePayPreference,
					string(configuration.ApplePay.DisplayPreference.Preference),
				)
			}
			if string(configuration.ApplePay.DisplayPreference.Value) != expect.ExpectedApplePayValue {
				return fmt.Errorf(
					"remote %s.apple_pay.display_preference.value mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedApplePayValue,
					string(configuration.ApplePay.DisplayPreference.Value),
				)
			}
		}
		if expect.ExpectedLinkPreference != "" {
			if configuration.Link == nil || configuration.Link.DisplayPreference == nil {
				return fmt.Errorf("remote %s.link.display_preference missing", expect.Address)
			}
			if string(configuration.Link.DisplayPreference.Preference) != expect.ExpectedLinkPreference {
				return fmt.Errorf(
					"remote %s.link.display_preference.preference mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedLinkPreference,
					string(configuration.Link.DisplayPreference.Preference),
				)
			}
			if string(configuration.Link.DisplayPreference.Value) != expect.ExpectedLinkValue {
				return fmt.Errorf(
					"remote %s.link.display_preference.value mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedLinkValue,
					string(configuration.Link.DisplayPreference.Value),
				)
			}
		}

		return nil
	}
}

func verifyPaymentMethodConfigurationDestroyInactive(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	configuration, err := retrievePaymentMethodConfiguration(
		client,
		state,
		"stripe_payment_method_configuration.test",
	)
	if err != nil {
		return err
	}
	if configuration.Active {
		return fmt.Errorf(
			"expected stripe_payment_method_configuration.test to be inactive after destroy",
		)
	}

	return nil
}

func verifyPaymentMethodDomain(expect paymentMethodDomainExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		domain, err := retrievePaymentMethodDomain(client, state, expect.Address)
		if err != nil {
			return err
		}
		expectedDomainName, err := runner.ResourceAttribute(state, expect.Address, "domain_name")
		if err != nil {
			return err
		}
		if domain.DomainName != expectedDomainName {
			return fmt.Errorf(
				"remote %s.domain_name mismatch: expected %q, got %q",
				expect.Address,
				expectedDomainName,
				domain.DomainName,
			)
		}
		if domain.Enabled != expect.ExpectedEnabled {
			return fmt.Errorf(
				"remote %s.enabled mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedEnabled,
				domain.Enabled,
			)
		}

		return nil
	}
}

func verifyPaymentMethodDomainDestroyStillExists(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	_, err := retrievePaymentMethodDomain(client, state, "stripe_payment_method_domain.test")
	return err
}

func verifyProductFeature(expect productFeatureExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		productID, err := runner.ResourcePrimaryID(state, expect.ExpectedProductAddress)
		if err != nil {
			return err
		}
		feature, err := retrieveProductFeature(client, productID, state, expect.Address)
		if err != nil {
			return err
		}
		expectedEntitlementFeatureID, err := runner.ResourcePrimaryID(
			state,
			expect.ExpectedEntitlementFeatureAddress,
		)
		if err != nil {
			return err
		}
		if feature.EntitlementFeature == nil || feature.EntitlementFeature.ID != expectedEntitlementFeatureID {
			actualFeatureID := ""
			if feature.EntitlementFeature != nil {
				actualFeatureID = feature.EntitlementFeature.ID
			}
			return fmt.Errorf(
				"remote %s.entitlement_feature mismatch: expected %q, got %q",
				expect.Address,
				expectedEntitlementFeatureID,
				actualFeatureID,
			)
		}

		return nil
	}
}

func verifyProductFeatureDestroyMissing(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	id, err := runner.ResourcePrimaryID(state, "stripe_product_feature.test")
	if err != nil {
		return err
	}
	productID, err := runner.ResourcePrimaryID(state, "stripe_product.product")
	if err != nil {
		return err
	}

	params := &stripe.ProductFeatureRetrieveParams{
		Product: stripe.String(productID),
	}
	feature, err := client.V1ProductFeatures.Retrieve(context.Background(), id, params)
	if err == nil && feature != nil && feature.Deleted {
		return nil
	}
	return expectRemoteMissing("stripe_product_feature.test", id, err)
}

func verifyTestHelpersTestClock(expect testHelpersTestClockExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		start := time.Now()
		deadline := time.Now().Add(20 * time.Minute)
		var (
			testClock *stripe.TestHelpersTestClock
			err       error
		)
		for {
			testClock, err = retrieveTestHelpersTestClock(client, state, expect.Address)
			if err == nil {
				break
			}
			if time.Now().After(deadline) {
				return err
			}
			time.Sleep(2 * time.Second)
		}
		fmt.Printf("test clock stabilized for %s after %s\n", expect.Address, time.Since(start))
		if testClock.Name != expect.ExpectedName {
			return fmt.Errorf(
				"remote %s.name mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedName,
				testClock.Name,
			)
		}
		if testClock.FrozenTime != expect.ExpectedFrozenTime {
			return fmt.Errorf(
				"remote %s.frozen_time mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedFrozenTime,
				testClock.FrozenTime,
			)
		}
		if string(testClock.Status) != expect.ExpectedStatus {
			return fmt.Errorf(
				"remote %s.status mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedStatus,
				string(testClock.Status),
			)
		}
		if expect.ExpectedCustomerAddress != "" {
			expectedCustomerID, err := runner.ResourcePrimaryID(state, expect.ExpectedCustomerAddress)
			if err != nil {
				return err
			}
			customer, err := client.V1Customers.Retrieve(context.Background(), expectedCustomerID, nil)
			if err != nil {
				return err
			}
			if customer.TestClock == nil || customer.TestClock.ID != testClock.ID {
				actualTestClockID := ""
				if customer.TestClock != nil {
					actualTestClockID = customer.TestClock.ID
				}
				return fmt.Errorf(
					"remote %s.customer.test_clock mismatch: expected %q, got %q",
					expect.Address,
					testClock.ID,
					actualTestClockID,
				)
			}
		}

		return nil
	}
}

func verifyTestHelpersTestClockDestroyMissing(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	id, err := runner.ResourcePrimaryID(state, "stripe_test_helpers_test_clock.test")
	if err != nil {
		return err
	}

	testClock, err := client.V1TestHelpersTestClocks.Retrieve(context.Background(), id, nil)
	if err == nil && testClock != nil && testClock.Deleted {
		return nil
	}
	return expectRemoteMissing("stripe_test_helpers_test_clock.test", id, err)
}

func verifyApplePayDomain(expect applePayDomainExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		domain, err := retrieveApplePayDomain(client, state, expect.Address)
		if err != nil {
			return err
		}
		expectedDomainName, err := runner.ResourceAttribute(state, expect.Address, "domain_name")
		if err != nil {
			return err
		}
		if domain.DomainName != expectedDomainName {
			return fmt.Errorf(
				"remote %s.domain_name mismatch: expected %q, got %q",
				expect.Address,
				expectedDomainName,
				domain.DomainName,
			)
		}

		return nil
	}
}

func verifyApplePayDomainDestroyMissing(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	id, err := runner.ResourcePrimaryID(state, "stripe_apple_pay_domain.test")
	if err != nil {
		return err
	}

	_, err = client.V1ApplePayDomains.Retrieve(context.Background(), id, nil)
	return expectRemoteMissing("stripe_apple_pay_domain.test", id, err)
}

func retrieveCreditNote(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.CreditNote, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	creditNote, err := client.V1CreditNotes.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return creditNote, nil
}

func retrieveCustomerBalanceTransaction(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.CustomerBalanceTransaction, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}
	customerID, err := runner.ResourceAttribute(state, address, "customer")
	if err != nil {
		return nil, err
	}

	params := &stripe.CustomerBalanceTransactionRetrieveParams{
		Customer: stripe.String(customerID),
	}
	transaction, err := client.V1CustomerBalanceTransactions.Retrieve(
		context.Background(),
		id,
		params,
	)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return transaction, nil
}

func retrieveFileLink(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.FileLink, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	fileLink, err := client.V1FileLinks.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return fileLink, nil
}

func retrieveTaxID(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.TaxID, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}
	customerID, err := runner.ResourceAttribute(state, address, "customer")
	if err != nil {
		return nil, err
	}

	params := &stripe.TaxIDRetrieveParams{
		Customer: stripe.String(customerID),
	}
	taxID, err := client.V1TaxIDs.Retrieve(context.Background(), id, params)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return taxID, nil
}

func retrieveBillingPortalConfiguration(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.BillingPortalConfiguration, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	configuration, err := client.V1BillingPortalConfigurations.Retrieve(
		context.Background(),
		id,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return configuration, nil
}

func retrievePaymentMethodConfiguration(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.PaymentMethodConfiguration, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	configuration, err := client.V1PaymentMethodConfigurations.Retrieve(
		context.Background(),
		id,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return configuration, nil
}

func retrievePaymentMethodDomain(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.PaymentMethodDomain, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	domain, err := client.V1PaymentMethodDomains.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return domain, nil
}

func retrieveProductFeature(
	client *stripe.Client,
	productID string,
	state *terraform.State,
	address string,
) (*stripe.ProductFeature, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	params := &stripe.ProductFeatureRetrieveParams{
		Product: stripe.String(productID),
	}
	feature, err := client.V1ProductFeatures.Retrieve(context.Background(), id, params)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return feature, nil
}

func retrieveTestHelpersTestClock(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.TestHelpersTestClock, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	testClock, err := client.V1TestHelpersTestClocks.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return testClock, nil
}

func retrieveApplePayDomain(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.ApplePayDomain, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	domain, err := client.V1ApplePayDomains.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return domain, nil
}
