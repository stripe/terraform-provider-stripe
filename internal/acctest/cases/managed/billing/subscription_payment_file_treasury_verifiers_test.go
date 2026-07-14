// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	stripe "github.com/stripe/stripe-go/v86"

	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

type subscriptionExpectations struct {
	Address                           string
	ExpectedCustomerAddress           string
	ExpectedDescription               string
	ExpectedMetadata                  map[string]string
	ExpectedCollectionMethod          string
	ExpectedStatus                    string
	ExpectedDaysUntilDue              int64
	ExpectedTrialEndAfterStart        bool
	ExpectedTrialMissingPaymentMethod string
	ExpectedPriceAddress              string
	ExpectedQuantity                  int64
}

type subscriptionItemExpectations struct {
	Address                     string
	ExpectedSubscriptionAddress string
	ExpectedPriceAddress        string
	ExpectedQuantity            int64
	ExpectedTaxRateAddresses    []string
	ExpectedMetadata            map[string]string
}

type invoiceExpectations struct {
	Address                  string
	ExpectedCustomerAddress  string
	ExpectedDescription      string
	ExpectedStatus           string
	ExpectedAutoAdvance      bool
	ExpectedCollectionMethod string
	ExpectedDueDate          int64
	ExpectedFooter           string
	ExpectedCustomFields     []invoiceCustomFieldExpectation
	ExpectedMetadata         map[string]string
}

type invoiceCustomFieldExpectation struct {
	Name  string
	Value string
}

type setupIntentExpectations struct {
	Address                                       string
	ExpectedCustomerAddress                       string
	ExpectedDescription                           string
	ExpectedStatus                                string
	ExpectedUsage                                 string
	ExpectedPaymentMethodTypes                    []string
	ExpectedAutomaticPaymentMethodsEnabled        bool
	ExpectedAutomaticPaymentMethodsAllowRedirects string
	ExpectedMetadata                              map[string]string
}

type paymentMethodExpectations struct {
	Address                 string
	ExpectedCustomerAddress string
	ExpectedType            string
	ExpectedAllowRedisplay  string
	ExpectedLast4           string
	ExpectedExpMonth        int64
	ExpectedExpYear         int64
	ExpectedBillingName     string
	ExpectedBillingEmail    string
	ExpectedBillingPhone    string
	ExpectedAddressLine1    string
	ExpectedCity            string
	ExpectedState           string
	ExpectedPostalCode      string
	ExpectedCountry         string
	ExpectedMetadata        map[string]string
}

type paymentLinkExpectations struct {
	Address                          string
	ExpectedActive                   bool
	ExpectedCustomerCreation         string
	ExpectedInactiveMessage          string
	ExpectedRedirectURL              string
	ExpectedPriceAddress             string
	ExpectedQuantity                 int64
	ExpectedAllowPromotionCodes      bool
	ExpectedBillingAddressCollection string
	ExpectedPaymentMethodCollection  string
	ExpectedSubmitType               string
	ExpectedMetadata                 map[string]string
}

type quoteExpectations struct {
	Address                             string
	ExpectedCustomerAddress             string
	ExpectedDescription                 string
	ExpectedHeader                      string
	ExpectedFooter                      string
	ExpectedStatus                      string
	ExpectedCollectionMethod            string
	ExpectedExpiresAt                   int64
	ExpectedInvoiceSettingsDaysUntilDue int64
	ExpectedPriceAddress                string
	ExpectedQuantity                    int64
	ExpectedAmountSubtotal              int64
	ExpectedAmountTotal                 int64
	ExpectedMetadata                    map[string]string
}

type chargeExpectations struct {
	Address                 string
	ExpectedAmount          int64
	ExpectedCurrency        string
	ExpectedCustomerAddress string
	ExpectedDescription     string
	ExpectedReceiptEmail    string
	ExpectedStatus          string
	ExpectedPaid            bool
	ExpectedCaptured        bool
	ExpectedShippingName    string
	ExpectedShippingPhone   string
	ExpectedAddressLine1    string
	ExpectedCity            string
	ExpectedState           string
	ExpectedPostalCode      string
	ExpectedCountry         string
	ExpectedMetadata        map[string]string
}

type sourceExpectations struct {
	Address                 string
	ExpectedCustomerAddress string
	ExpectedType            string
	ExpectedStatus          string
	ExpectedOwnerName       string
	ExpectedOwnerEmail      string
	ExpectedOwnerPhone      string
	ExpectedAddressLine1    string
	ExpectedCity            string
	ExpectedState           string
	ExpectedPostalCode      string
	ExpectedCountry         string
	ExpectedMetadata        map[string]string
}

type shippingRateExpectations struct {
	Address                     string
	ExpectedDisplayName         string
	ExpectedActive              bool
	ExpectedFixedAmount         int64
	ExpectedFixedAmountCurrency string
	ExpectedMinimumUnit         string
	ExpectedMinimumValue        int64
	ExpectedMaximumUnit         string
	ExpectedMaximumValue        int64
	ExpectedTaxBehavior         string
	ExpectedTaxCode             string
	StateStrings                []stateStringExpectation
	ExpectedMetadata            map[string]string
}

type stateStringExpectation struct {
	Attribute string
	Expected  string
}

type taxRateExpectations struct {
	Address              string
	ExpectedDisplayName  string
	ExpectedDescription  string
	ExpectedCountry      string
	ExpectedState        string
	ExpectedJurisdiction string
	ExpectedInclusive    bool
	ExpectedPercentage   float64
	ExpectedTaxType      string
	ExpectedActive       bool
	ExpectedMetadata     map[string]string
}

type taxRegistrationExpectations struct {
	Address           string
	ExpectedCountry   string
	ExpectedType      string
	ExpectedProvince  string
	ExpectedExpiresAt int64
}

type fileExpectations struct {
	Address                   string
	ExpectedPurpose           string
	ExpectedFilename          string
	ExpectedFileLinkExpiresAt int64
	ExpectedFileLinkMetadata  map[string]string
}

type treasuryFinancialAccountExpectations struct {
	Address                     string
	ExpectedNickname            string
	ExpectedSupportedCurrencies []string
	ExpectedFeaturePaths        []string
	ExpectedMetadata            map[string]string
}

func verifySubscription(expect subscriptionExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		subscription, err := retrieveSubscription(client, state, expect.Address)
		if err != nil {
			return err
		}
		expectedCustomerID, err := runner.ResourcePrimaryID(state, expect.ExpectedCustomerAddress)
		if err != nil {
			return err
		}
		if subscription.Customer == nil || subscription.Customer.ID != expectedCustomerID {
			actualCustomerID := ""
			if subscription.Customer != nil {
				actualCustomerID = subscription.Customer.ID
			}
			return fmt.Errorf(
				"remote %s.customer mismatch: expected %q, got %q",
				expect.Address,
				expectedCustomerID,
				actualCustomerID,
			)
		}
		if subscription.Description != expect.ExpectedDescription {
			return fmt.Errorf(
				"remote %s.description mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedDescription,
				subscription.Description,
			)
		}
		if string(subscription.CollectionMethod) != expect.ExpectedCollectionMethod {
			return fmt.Errorf(
				"remote %s.collection_method mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedCollectionMethod,
				string(subscription.CollectionMethod),
			)
		}
		if string(subscription.Status) != expect.ExpectedStatus {
			return fmt.Errorf(
				"remote %s.status mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedStatus,
				string(subscription.Status),
			)
		}
		if expect.ExpectedDaysUntilDue != 0 && subscription.DaysUntilDue != expect.ExpectedDaysUntilDue {
			return fmt.Errorf(
				"remote %s.days_until_due mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedDaysUntilDue,
				subscription.DaysUntilDue,
			)
		}
		if expect.ExpectedTrialEndAfterStart {
			if subscription.TrialStart == 0 || subscription.TrialEnd == 0 {
				return fmt.Errorf(
					"remote %s trial window missing: start=%d end=%d",
					expect.Address,
					subscription.TrialStart,
					subscription.TrialEnd,
				)
			}
			if subscription.TrialEnd <= subscription.TrialStart {
				return fmt.Errorf(
					"remote %s trial window invalid: start=%d end=%d",
					expect.Address,
					subscription.TrialStart,
					subscription.TrialEnd,
				)
			}
		}
		if expect.ExpectedTrialMissingPaymentMethod != "" {
			if subscription.TrialSettings == nil || subscription.TrialSettings.EndBehavior == nil {
				return fmt.Errorf("remote %s.trial_settings.end_behavior missing", expect.Address)
			}
			if string(subscription.TrialSettings.EndBehavior.MissingPaymentMethod) != expect.ExpectedTrialMissingPaymentMethod {
				return fmt.Errorf(
					"remote %s.trial_settings.end_behavior.missing_payment_method mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedTrialMissingPaymentMethod,
					string(subscription.TrialSettings.EndBehavior.MissingPaymentMethod),
				)
			}
		}
		if err := expectMetadataSubset(expect.Address+".metadata", subscription.Metadata, expect.ExpectedMetadata); err != nil {
			return err
		}
		if subscription.Items == nil || len(subscription.Items.Data) == 0 || subscription.Items.Data[0] == nil {
			return fmt.Errorf("remote %s.items[0] missing", expect.Address)
		}
		expectedPriceID, err := runner.ResourcePrimaryID(state, expect.ExpectedPriceAddress)
		if err != nil {
			return err
		}
		if subscription.Items.Data[0].Price == nil || subscription.Items.Data[0].Price.ID != expectedPriceID {
			actualPriceID := ""
			if subscription.Items.Data[0].Price != nil {
				actualPriceID = subscription.Items.Data[0].Price.ID
			}
			return fmt.Errorf(
				"remote %s.items[0].price mismatch: expected %q, got %q",
				expect.Address,
				expectedPriceID,
				actualPriceID,
			)
		}
		if subscription.Items.Data[0].Quantity != expect.ExpectedQuantity {
			return fmt.Errorf(
				"remote %s.items[0].quantity mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedQuantity,
				subscription.Items.Data[0].Quantity,
			)
		}

		return nil
	}
}

func verifySubscriptionDestroyStillExists(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	subscription, err := retrieveSubscription(client, state, "stripe_subscription.test")
	if err != nil {
		return err
	}
	if string(subscription.Status) != "canceled" {
		return fmt.Errorf(
			"expected stripe_subscription.test to be canceled after destroy, got %q",
			string(subscription.Status),
		)
	}

	return nil
}

func verifySubscriptionItem(expect subscriptionItemExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		item, err := retrieveSubscriptionItem(client, state, expect.Address)
		if err != nil {
			return err
		}
		expectedSubscriptionID, err := runner.ResourcePrimaryID(state, expect.ExpectedSubscriptionAddress)
		if err != nil {
			return err
		}
		if item.Subscription != expectedSubscriptionID {
			return fmt.Errorf(
				"remote %s.subscription mismatch: expected %q, got %q",
				expect.Address,
				expectedSubscriptionID,
				item.Subscription,
			)
		}
		expectedPriceID, err := runner.ResourcePrimaryID(state, expect.ExpectedPriceAddress)
		if err != nil {
			return err
		}
		if item.Price == nil || item.Price.ID != expectedPriceID {
			actualPriceID := ""
			if item.Price != nil {
				actualPriceID = item.Price.ID
			}
			return fmt.Errorf(
				"remote %s.price mismatch: expected %q, got %q",
				expect.Address,
				expectedPriceID,
				actualPriceID,
			)
		}
		if item.Quantity != expect.ExpectedQuantity {
			return fmt.Errorf(
				"remote %s.quantity mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedQuantity,
				item.Quantity,
			)
		}
		if len(expect.ExpectedTaxRateAddresses) > 0 {
			actualTaxRateIDs := make([]string, 0, len(item.TaxRates))
			for _, taxRate := range item.TaxRates {
				if taxRate != nil {
					actualTaxRateIDs = append(actualTaxRateIDs, taxRate.ID)
				}
			}
			expectedTaxRateIDs := make([]string, 0, len(expect.ExpectedTaxRateAddresses))
			for _, address := range expect.ExpectedTaxRateAddresses {
				expectedTaxRateID, err := runner.ResourcePrimaryID(state, address)
				if err != nil {
					return err
				}
				expectedTaxRateIDs = append(expectedTaxRateIDs, expectedTaxRateID)
			}
			if err := expectRemoteStringList(expect.Address+".tax_rates", actualTaxRateIDs, expectedTaxRateIDs); err != nil {
				return err
			}
		}
		if err := expectMetadataSubset(expect.Address+".metadata", item.Metadata, expect.ExpectedMetadata); err != nil {
			return err
		}

		return nil
	}
}

func verifySubscriptionItemDestroyMissing(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	id, err := runner.ResourcePrimaryID(state, "stripe_subscription_item.test")
	if err != nil {
		return err
	}

	_, err = client.V1SubscriptionItems.Retrieve(context.Background(), id, nil)
	return expectRemoteMissing("stripe_subscription_item.test", id, err)
}

func verifyInvoice(expect invoiceExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		invoice, err := retrieveInvoice(client, state, expect.Address)
		if err != nil {
			return err
		}
		expectedCustomerID, err := runner.ResourcePrimaryID(state, expect.ExpectedCustomerAddress)
		if err != nil {
			return err
		}
		if invoice.Customer == nil || invoice.Customer.ID != expectedCustomerID {
			actualCustomerID := ""
			if invoice.Customer != nil {
				actualCustomerID = invoice.Customer.ID
			}
			return fmt.Errorf(
				"remote %s.customer mismatch: expected %q, got %q",
				expect.Address,
				expectedCustomerID,
				actualCustomerID,
			)
		}
		if invoice.Description != expect.ExpectedDescription {
			return fmt.Errorf(
				"remote %s.description mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedDescription,
				invoice.Description,
			)
		}
		if string(invoice.Status) != expect.ExpectedStatus {
			return fmt.Errorf(
				"remote %s.status mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedStatus,
				string(invoice.Status),
			)
		}
		if invoice.AutoAdvance != expect.ExpectedAutoAdvance {
			return fmt.Errorf(
				"remote %s.auto_advance mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedAutoAdvance,
				invoice.AutoAdvance,
			)
		}
		if expect.ExpectedCollectionMethod != "" &&
			string(invoice.CollectionMethod) != expect.ExpectedCollectionMethod {
			return fmt.Errorf(
				"remote %s.collection_method mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedCollectionMethod,
				string(invoice.CollectionMethod),
			)
		}
		if expect.ExpectedDueDate != 0 && invoice.DueDate != expect.ExpectedDueDate {
			return fmt.Errorf(
				"remote %s.due_date mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedDueDate,
				invoice.DueDate,
			)
		}
		if expect.ExpectedFooter != "" && invoice.Footer != expect.ExpectedFooter {
			return fmt.Errorf(
				"remote %s.footer mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedFooter,
				invoice.Footer,
			)
		}
		if len(expect.ExpectedCustomFields) > 0 {
			if len(invoice.CustomFields) != len(expect.ExpectedCustomFields) {
				return fmt.Errorf(
					"remote %s.custom_fields length mismatch: expected %d, got %d",
					expect.Address,
					len(expect.ExpectedCustomFields),
					len(invoice.CustomFields),
				)
			}
			for i, expectedCustomField := range expect.ExpectedCustomFields {
				actualCustomField := invoice.CustomFields[i]
				if actualCustomField == nil {
					return fmt.Errorf("remote %s.custom_fields[%d] missing", expect.Address, i)
				}
				if actualCustomField.Name != expectedCustomField.Name {
					return fmt.Errorf(
						"remote %s.custom_fields[%d].name mismatch: expected %q, got %q",
						expect.Address,
						i,
						expectedCustomField.Name,
						actualCustomField.Name,
					)
				}
				if actualCustomField.Value != expectedCustomField.Value {
					return fmt.Errorf(
						"remote %s.custom_fields[%d].value mismatch: expected %q, got %q",
						expect.Address,
						i,
						expectedCustomField.Value,
						actualCustomField.Value,
					)
				}
			}
		}
		if err := expectMetadataSubset(expect.Address+".metadata", invoice.Metadata, expect.ExpectedMetadata); err != nil {
			return err
		}

		return nil
	}
}

func verifyInvoiceDestroyMissing(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	id, err := runner.ResourcePrimaryID(state, "stripe_invoice.test")
	if err != nil {
		return err
	}

	_, err = client.V1Invoices.Retrieve(context.Background(), id, nil)
	return expectRemoteMissing("stripe_invoice.test", id, err)
}

func verifySetupIntent(expect setupIntentExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		setupIntent, err := retrieveSetupIntent(client, state, expect.Address)
		if err != nil {
			return err
		}
		if expect.ExpectedCustomerAddress != "" {
			expectedCustomerID, err := runner.ResourcePrimaryID(state, expect.ExpectedCustomerAddress)
			if err != nil {
				return err
			}
			if setupIntent.Customer == nil || setupIntent.Customer.ID != expectedCustomerID {
				actualCustomerID := ""
				if setupIntent.Customer != nil {
					actualCustomerID = setupIntent.Customer.ID
				}
				return fmt.Errorf(
					"remote %s.customer mismatch: expected %q, got %q",
					expect.Address,
					expectedCustomerID,
					actualCustomerID,
				)
			}
		}
		if setupIntent.Description != expect.ExpectedDescription {
			return fmt.Errorf(
				"remote %s.description mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedDescription,
				setupIntent.Description,
			)
		}
		if string(setupIntent.Status) != expect.ExpectedStatus {
			return fmt.Errorf(
				"remote %s.status mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedStatus,
				string(setupIntent.Status),
			)
		}
		if string(setupIntent.Usage) != expect.ExpectedUsage {
			return fmt.Errorf(
				"remote %s.usage mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedUsage,
				string(setupIntent.Usage),
			)
		}
		if expect.ExpectedPaymentMethodTypes != nil {
			if err := expectRemoteStringList(expect.Address+".payment_method_types", setupIntent.PaymentMethodTypes, expect.ExpectedPaymentMethodTypes); err != nil {
				return err
			}
		}
		if expect.ExpectedAutomaticPaymentMethodsAllowRedirects != "" {
			if setupIntent.AutomaticPaymentMethods == nil {
				return fmt.Errorf("remote %s.automatic_payment_methods missing", expect.Address)
			}
			if setupIntent.AutomaticPaymentMethods.Enabled != expect.ExpectedAutomaticPaymentMethodsEnabled {
				return fmt.Errorf(
					"remote %s.automatic_payment_methods.enabled mismatch: expected %t, got %t",
					expect.Address,
					expect.ExpectedAutomaticPaymentMethodsEnabled,
					setupIntent.AutomaticPaymentMethods.Enabled,
				)
			}
			if string(setupIntent.AutomaticPaymentMethods.AllowRedirects) != expect.ExpectedAutomaticPaymentMethodsAllowRedirects {
				return fmt.Errorf(
					"remote %s.automatic_payment_methods.allow_redirects mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedAutomaticPaymentMethodsAllowRedirects,
					string(setupIntent.AutomaticPaymentMethods.AllowRedirects),
				)
			}
		}
		if err := expectMetadataSubset(expect.Address+".metadata", setupIntent.Metadata, expect.ExpectedMetadata); err != nil {
			return err
		}

		return nil
	}
}

func verifySetupIntentDestroyStillExists(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	_, err := retrieveSetupIntent(client, state, "stripe_setup_intent.test")
	return err
}

func verifyPaymentMethod(expect paymentMethodExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		paymentMethod, err := retrievePaymentMethod(client, state, expect.Address)
		if err != nil {
			return err
		}
		if expect.ExpectedCustomerAddress != "" {
			expectedCustomerID, err := runner.ResourcePrimaryID(state, expect.ExpectedCustomerAddress)
			if err != nil {
				return err
			}
			if paymentMethod.Customer == nil || paymentMethod.Customer.ID != expectedCustomerID {
				actualCustomerID := ""
				if paymentMethod.Customer != nil {
					actualCustomerID = paymentMethod.Customer.ID
				}
				return fmt.Errorf(
					"remote %s.customer mismatch: expected %q, got %q",
					expect.Address,
					expectedCustomerID,
					actualCustomerID,
				)
			}
		}
		if string(paymentMethod.Type) != expect.ExpectedType {
			return fmt.Errorf(
				"remote %s.type mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedType,
				string(paymentMethod.Type),
			)
		}
		if expect.ExpectedAllowRedisplay != "" && string(paymentMethod.AllowRedisplay) != expect.ExpectedAllowRedisplay {
			return fmt.Errorf(
				"remote %s.allow_redisplay mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedAllowRedisplay,
				string(paymentMethod.AllowRedisplay),
			)
		}
		if paymentMethod.Card == nil {
			return fmt.Errorf("remote %s.card missing", expect.Address)
		}
		if paymentMethod.Card.Last4 != expect.ExpectedLast4 {
			return fmt.Errorf(
				"remote %s.card.last4 mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedLast4,
				paymentMethod.Card.Last4,
			)
		}
		if paymentMethod.Card.ExpMonth != expect.ExpectedExpMonth {
			return fmt.Errorf(
				"remote %s.card.exp_month mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedExpMonth,
				paymentMethod.Card.ExpMonth,
			)
		}
		if paymentMethod.Card.ExpYear != expect.ExpectedExpYear {
			return fmt.Errorf(
				"remote %s.card.exp_year mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedExpYear,
				paymentMethod.Card.ExpYear,
			)
		}
		if paymentMethod.BillingDetails == nil || paymentMethod.BillingDetails.Address == nil {
			return fmt.Errorf("remote %s.billing_details missing", expect.Address)
		}
		if paymentMethod.BillingDetails.Name != expect.ExpectedBillingName {
			return fmt.Errorf(
				"remote %s.billing_details.name mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedBillingName,
				paymentMethod.BillingDetails.Name,
			)
		}
		if paymentMethod.BillingDetails.Email != expect.ExpectedBillingEmail {
			return fmt.Errorf(
				"remote %s.billing_details.email mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedBillingEmail,
				paymentMethod.BillingDetails.Email,
			)
		}
		if paymentMethod.BillingDetails.Phone != expect.ExpectedBillingPhone {
			return fmt.Errorf(
				"remote %s.billing_details.phone mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedBillingPhone,
				paymentMethod.BillingDetails.Phone,
			)
		}
		if paymentMethod.BillingDetails.Address.Line1 != expect.ExpectedAddressLine1 {
			return fmt.Errorf(
				"remote %s.billing_details.address.line1 mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedAddressLine1,
				paymentMethod.BillingDetails.Address.Line1,
			)
		}
		if paymentMethod.BillingDetails.Address.City != expect.ExpectedCity {
			return fmt.Errorf(
				"remote %s.billing_details.address.city mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedCity,
				paymentMethod.BillingDetails.Address.City,
			)
		}
		if paymentMethod.BillingDetails.Address.State != expect.ExpectedState {
			return fmt.Errorf(
				"remote %s.billing_details.address.state mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedState,
				paymentMethod.BillingDetails.Address.State,
			)
		}
		if paymentMethod.BillingDetails.Address.PostalCode != expect.ExpectedPostalCode {
			return fmt.Errorf(
				"remote %s.billing_details.address.postal_code mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedPostalCode,
				paymentMethod.BillingDetails.Address.PostalCode,
			)
		}
		if paymentMethod.BillingDetails.Address.Country != expect.ExpectedCountry {
			return fmt.Errorf(
				"remote %s.billing_details.address.country mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedCountry,
				paymentMethod.BillingDetails.Address.Country,
			)
		}
		if err := expectMetadataSubset(expect.Address+".metadata", paymentMethod.Metadata, expect.ExpectedMetadata); err != nil {
			return err
		}

		return nil
	}
}

func verifyPaymentMethodDestroyStillExists(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	_, err := retrievePaymentMethod(client, state, "stripe_payment_method.test")
	return err
}

func verifyPaymentLink(expect paymentLinkExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		paymentLink, err := retrievePaymentLink(client, state, expect.Address)
		if err != nil {
			return err
		}
		if paymentLink.Active != expect.ExpectedActive {
			return fmt.Errorf(
				"remote %s.active mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedActive,
				paymentLink.Active,
			)
		}
		if expect.ExpectedCustomerCreation != "" &&
			string(paymentLink.CustomerCreation) != expect.ExpectedCustomerCreation {
			return fmt.Errorf(
				"remote %s.customer_creation mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedCustomerCreation,
				string(paymentLink.CustomerCreation),
			)
		}
		if paymentLink.InactiveMessage != expect.ExpectedInactiveMessage {
			return fmt.Errorf(
				"remote %s.inactive_message mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedInactiveMessage,
				paymentLink.InactiveMessage,
			)
		}
		if paymentLink.AllowPromotionCodes != expect.ExpectedAllowPromotionCodes {
			return fmt.Errorf(
				"remote %s.allow_promotion_codes mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedAllowPromotionCodes,
				paymentLink.AllowPromotionCodes,
			)
		}
		if expect.ExpectedBillingAddressCollection != "" &&
			string(paymentLink.BillingAddressCollection) != expect.ExpectedBillingAddressCollection {
			return fmt.Errorf(
				"remote %s.billing_address_collection mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedBillingAddressCollection,
				string(paymentLink.BillingAddressCollection),
			)
		}
		if expect.ExpectedPaymentMethodCollection != "" &&
			string(paymentLink.PaymentMethodCollection) != expect.ExpectedPaymentMethodCollection {
			return fmt.Errorf(
				"remote %s.payment_method_collection mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedPaymentMethodCollection,
				string(paymentLink.PaymentMethodCollection),
			)
		}
		if expect.ExpectedSubmitType != "" && string(paymentLink.SubmitType) != expect.ExpectedSubmitType {
			return fmt.Errorf(
				"remote %s.submit_type mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedSubmitType,
				string(paymentLink.SubmitType),
			)
		}
		if paymentLink.URL == "" {
			return fmt.Errorf("remote %s.url missing", expect.Address)
		}
		if paymentLink.AfterCompletion == nil || paymentLink.AfterCompletion.Redirect == nil {
			return fmt.Errorf("remote %s.after_completion.redirect missing", expect.Address)
		}
		if paymentLink.AfterCompletion.Redirect.URL != expect.ExpectedRedirectURL {
			return fmt.Errorf(
				"remote %s.after_completion.redirect.url mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedRedirectURL,
				paymentLink.AfterCompletion.Redirect.URL,
			)
		}
		if err := expectMetadataSubset(expect.Address+".metadata", paymentLink.Metadata, expect.ExpectedMetadata); err != nil {
			return err
		}
		if paymentLink.LineItems == nil || len(paymentLink.LineItems.Data) == 0 || paymentLink.LineItems.Data[0] == nil {
			return fmt.Errorf("remote %s.line_items[0] missing", expect.Address)
		}
		if paymentLink.LineItems.Data[0].Quantity != expect.ExpectedQuantity {
			return fmt.Errorf(
				"remote %s.line_items[0].quantity mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedQuantity,
				paymentLink.LineItems.Data[0].Quantity,
			)
		}
		expectedPriceID, err := runner.ResourcePrimaryID(state, expect.ExpectedPriceAddress)
		if err != nil {
			return err
		}
		if paymentLink.LineItems.Data[0].Price == nil || paymentLink.LineItems.Data[0].Price.ID != expectedPriceID {
			actualPriceID := ""
			if paymentLink.LineItems.Data[0].Price != nil {
				actualPriceID = paymentLink.LineItems.Data[0].Price.ID
			}
			return fmt.Errorf(
				"remote %s.line_items[0].price mismatch: expected %q, got %q",
				expect.Address,
				expectedPriceID,
				actualPriceID,
			)
		}

		return nil
	}
}

func verifyPaymentLinkDestroyInactive(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	paymentLink, err := retrievePaymentLink(client, state, "stripe_payment_link.test")
	if err != nil {
		return err
	}
	if paymentLink.Active {
		return fmt.Errorf("expected stripe_payment_link.test to be inactive after destroy")
	}

	return nil
}

func verifyQuote(expect quoteExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		quote, err := retrieveQuote(client, state, expect.Address)
		if err != nil {
			return err
		}
		expectedCustomerID, err := runner.ResourcePrimaryID(state, expect.ExpectedCustomerAddress)
		if err != nil {
			return err
		}
		if quote.Customer == nil || quote.Customer.ID != expectedCustomerID {
			actualCustomerID := ""
			if quote.Customer != nil {
				actualCustomerID = quote.Customer.ID
			}
			return fmt.Errorf(
				"remote %s.customer mismatch: expected %q, got %q",
				expect.Address,
				expectedCustomerID,
				actualCustomerID,
			)
		}
		if quote.Description != expect.ExpectedDescription {
			return fmt.Errorf(
				"remote %s.description mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedDescription,
				quote.Description,
			)
		}
		if quote.Header != expect.ExpectedHeader {
			return fmt.Errorf(
				"remote %s.header mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedHeader,
				quote.Header,
			)
		}
		if expect.ExpectedFooter != "" && quote.Footer != expect.ExpectedFooter {
			return fmt.Errorf(
				"remote %s.footer mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedFooter,
				quote.Footer,
			)
		}
		if string(quote.Status) != expect.ExpectedStatus {
			return fmt.Errorf(
				"remote %s.status mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedStatus,
				string(quote.Status),
			)
		}
		if expect.ExpectedCollectionMethod != "" &&
			string(quote.CollectionMethod) != expect.ExpectedCollectionMethod {
			return fmt.Errorf(
				"remote %s.collection_method mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedCollectionMethod,
				string(quote.CollectionMethod),
			)
		}
		if expect.ExpectedExpiresAt != 0 && quote.ExpiresAt != expect.ExpectedExpiresAt {
			return fmt.Errorf(
				"remote %s.expires_at mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedExpiresAt,
				quote.ExpiresAt,
			)
		}
		if expect.ExpectedInvoiceSettingsDaysUntilDue != 0 {
			if quote.InvoiceSettings == nil {
				return fmt.Errorf("remote %s.invoice_settings missing", expect.Address)
			}
			if quote.InvoiceSettings.DaysUntilDue != expect.ExpectedInvoiceSettingsDaysUntilDue {
				return fmt.Errorf(
					"remote %s.invoice_settings.days_until_due mismatch: expected %d, got %d",
					expect.Address,
					expect.ExpectedInvoiceSettingsDaysUntilDue,
					quote.InvoiceSettings.DaysUntilDue,
				)
			}
		}
		if quote.AmountSubtotal != expect.ExpectedAmountSubtotal {
			return fmt.Errorf(
				"remote %s.amount_subtotal mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedAmountSubtotal,
				quote.AmountSubtotal,
			)
		}
		if quote.AmountTotal != expect.ExpectedAmountTotal {
			return fmt.Errorf(
				"remote %s.amount_total mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedAmountTotal,
				quote.AmountTotal,
			)
		}
		if err := expectMetadataSubset(expect.Address+".metadata", quote.Metadata, expect.ExpectedMetadata); err != nil {
			return err
		}
		if quote.LineItems == nil || len(quote.LineItems.Data) == 0 || quote.LineItems.Data[0] == nil {
			return fmt.Errorf("remote %s.line_items[0] missing", expect.Address)
		}
		if quote.LineItems.Data[0].Quantity != expect.ExpectedQuantity {
			return fmt.Errorf(
				"remote %s.line_items[0].quantity mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedQuantity,
				quote.LineItems.Data[0].Quantity,
			)
		}
		expectedPriceID, err := runner.ResourcePrimaryID(state, expect.ExpectedPriceAddress)
		if err != nil {
			return err
		}
		if quote.LineItems.Data[0].Price == nil || quote.LineItems.Data[0].Price.ID != expectedPriceID {
			actualPriceID := ""
			if quote.LineItems.Data[0].Price != nil {
				actualPriceID = quote.LineItems.Data[0].Price.ID
			}
			return fmt.Errorf(
				"remote %s.line_items[0].price mismatch: expected %q, got %q",
				expect.Address,
				expectedPriceID,
				actualPriceID,
			)
		}

		return nil
	}
}

func verifyQuoteDestroyStillExists(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	_, err := retrieveQuote(client, state, "stripe_quote.test")
	return err
}

func verifyCharge(expect chargeExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		charge, err := retrieveCharge(client, state, expect.Address)
		if err != nil {
			return err
		}
		if charge.Amount != expect.ExpectedAmount {
			return fmt.Errorf(
				"remote %s.amount mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedAmount,
				charge.Amount,
			)
		}
		if string(charge.Currency) != expect.ExpectedCurrency {
			return fmt.Errorf(
				"remote %s.currency mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedCurrency,
				string(charge.Currency),
			)
		}
		if expect.ExpectedCustomerAddress != "" {
			expectedCustomerID, err := runner.ResourcePrimaryID(state, expect.ExpectedCustomerAddress)
			if err != nil {
				return err
			}
			actualCustomerID := ""
			if charge.Customer != nil {
				actualCustomerID = charge.Customer.ID
			}
			if actualCustomerID != expectedCustomerID {
				return fmt.Errorf(
					"remote %s.customer mismatch: expected %q, got %q",
					expect.Address,
					expectedCustomerID,
					actualCustomerID,
				)
			}
		}
		if charge.Description != expect.ExpectedDescription {
			return fmt.Errorf(
				"remote %s.description mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedDescription,
				charge.Description,
			)
		}
		if charge.ReceiptEmail != expect.ExpectedReceiptEmail {
			return fmt.Errorf(
				"remote %s.receipt_email mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedReceiptEmail,
				charge.ReceiptEmail,
			)
		}
		if string(charge.Status) != expect.ExpectedStatus {
			return fmt.Errorf(
				"remote %s.status mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedStatus,
				string(charge.Status),
			)
		}
		if charge.Paid != expect.ExpectedPaid {
			return fmt.Errorf(
				"remote %s.paid mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedPaid,
				charge.Paid,
			)
		}
		if charge.Captured != expect.ExpectedCaptured {
			return fmt.Errorf(
				"remote %s.captured mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedCaptured,
				charge.Captured,
			)
		}
		if expect.ExpectedShippingName != "" {
			if charge.Shipping == nil {
				return fmt.Errorf("remote %s.shipping missing", expect.Address)
			}
			if charge.Shipping.Name != expect.ExpectedShippingName {
				return fmt.Errorf(
					"remote %s.shipping.name mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedShippingName,
					charge.Shipping.Name,
				)
			}
			if charge.Shipping.Phone != expect.ExpectedShippingPhone {
				return fmt.Errorf(
					"remote %s.shipping.phone mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedShippingPhone,
					charge.Shipping.Phone,
				)
			}
			if charge.Shipping.Address == nil {
				return fmt.Errorf("remote %s.shipping.address missing", expect.Address)
			}
			if charge.Shipping.Address.Line1 != expect.ExpectedAddressLine1 {
				return fmt.Errorf(
					"remote %s.shipping.address.line1 mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedAddressLine1,
					charge.Shipping.Address.Line1,
				)
			}
			if charge.Shipping.Address.City != expect.ExpectedCity {
				return fmt.Errorf(
					"remote %s.shipping.address.city mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedCity,
					charge.Shipping.Address.City,
				)
			}
			if charge.Shipping.Address.State != expect.ExpectedState {
				return fmt.Errorf(
					"remote %s.shipping.address.state mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedState,
					charge.Shipping.Address.State,
				)
			}
			if charge.Shipping.Address.PostalCode != expect.ExpectedPostalCode {
				return fmt.Errorf(
					"remote %s.shipping.address.postal_code mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedPostalCode,
					charge.Shipping.Address.PostalCode,
				)
			}
			if charge.Shipping.Address.Country != expect.ExpectedCountry {
				return fmt.Errorf(
					"remote %s.shipping.address.country mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedCountry,
					charge.Shipping.Address.Country,
				)
			}
		}
		if err := expectMetadataSubset(expect.Address+".metadata", charge.Metadata, expect.ExpectedMetadata); err != nil {
			return err
		}

		return nil
	}
}

func verifyChargeDestroyStillExists(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	_, err := retrieveCharge(client, state, "stripe_charge.test")
	return err
}

func verifySource(expect sourceExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		source, err := retrieveSource(client, state, expect.Address)
		if err != nil {
			return err
		}
		if source.Type != expect.ExpectedType {
			return fmt.Errorf(
				"remote %s.type mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedType,
				source.Type,
			)
		}
		if string(source.Status) != expect.ExpectedStatus {
			return fmt.Errorf(
				"remote %s.status mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedStatus,
				string(source.Status),
			)
		}
		if expect.ExpectedCustomerAddress != "" {
			expectedCustomerID, err := runner.ResourcePrimaryID(state, expect.ExpectedCustomerAddress)
			if err != nil {
				return err
			}
			if source.Customer != expectedCustomerID {
				return fmt.Errorf(
					"remote %s.customer mismatch: expected %q, got %q",
					expect.Address,
					expectedCustomerID,
					source.Customer,
				)
			}
		}
		if source.Owner == nil {
			return fmt.Errorf("remote %s.owner missing", expect.Address)
		}
		if source.Owner.Name != expect.ExpectedOwnerName {
			return fmt.Errorf(
				"remote %s.owner.name mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedOwnerName,
				source.Owner.Name,
			)
		}
		if source.Owner.Email != expect.ExpectedOwnerEmail {
			return fmt.Errorf(
				"remote %s.owner.email mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedOwnerEmail,
				source.Owner.Email,
			)
		}
		if expect.ExpectedOwnerPhone != "" && source.Owner.Phone != expect.ExpectedOwnerPhone {
			return fmt.Errorf(
				"remote %s.owner.phone mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedOwnerPhone,
				source.Owner.Phone,
			)
		}
		if expect.ExpectedAddressLine1 != "" {
			if source.Owner.Address == nil {
				return fmt.Errorf("remote %s.owner.address missing", expect.Address)
			}
			if source.Owner.Address.Line1 != expect.ExpectedAddressLine1 {
				return fmt.Errorf(
					"remote %s.owner.address.line1 mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedAddressLine1,
					source.Owner.Address.Line1,
				)
			}
			if source.Owner.Address.City != expect.ExpectedCity {
				return fmt.Errorf(
					"remote %s.owner.address.city mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedCity,
					source.Owner.Address.City,
				)
			}
			if source.Owner.Address.State != expect.ExpectedState {
				return fmt.Errorf(
					"remote %s.owner.address.state mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedState,
					source.Owner.Address.State,
				)
			}
			if source.Owner.Address.PostalCode != expect.ExpectedPostalCode {
				return fmt.Errorf(
					"remote %s.owner.address.postal_code mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedPostalCode,
					source.Owner.Address.PostalCode,
				)
			}
			if source.Owner.Address.Country != expect.ExpectedCountry {
				return fmt.Errorf(
					"remote %s.owner.address.country mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedCountry,
					source.Owner.Address.Country,
				)
			}
		}
		if err := expectMetadataSubset(expect.Address+".metadata", source.Metadata, expect.ExpectedMetadata); err != nil {
			return err
		}

		return nil
	}
}

func verifySourceDestroyStillExists(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	_, err := retrieveSource(client, state, "stripe_source.test")
	return err
}

func verifyShippingRate(expect shippingRateExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		shippingRate, err := retrieveShippingRate(client, state, expect.Address)
		if err != nil {
			return err
		}
		if shippingRate.DisplayName != expect.ExpectedDisplayName {
			return fmt.Errorf(
				"remote %s.display_name mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedDisplayName,
				shippingRate.DisplayName,
			)
		}
		if shippingRate.Active != expect.ExpectedActive {
			return fmt.Errorf(
				"remote %s.active mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedActive,
				shippingRate.Active,
			)
		}
		for _, stateString := range expect.StateStrings {
			if err := expectStateString(
				state,
				expect.Address,
				stateString.Attribute,
				stateString.Expected,
			); err != nil {
				return err
			}
		}
		if shippingRate.FixedAmount == nil {
			return fmt.Errorf("remote %s.fixed_amount missing", expect.Address)
		}
		if shippingRate.FixedAmount.Amount != expect.ExpectedFixedAmount {
			return fmt.Errorf(
				"remote %s.fixed_amount.amount mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedFixedAmount,
				shippingRate.FixedAmount.Amount,
			)
		}
		if string(shippingRate.FixedAmount.Currency) != expect.ExpectedFixedAmountCurrency {
			return fmt.Errorf(
				"remote %s.fixed_amount.currency mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedFixedAmountCurrency,
				string(shippingRate.FixedAmount.Currency),
			)
		}
		if shippingRate.DeliveryEstimate == nil ||
			shippingRate.DeliveryEstimate.Minimum == nil ||
			shippingRate.DeliveryEstimate.Maximum == nil {
			return fmt.Errorf("remote %s.delivery_estimate missing", expect.Address)
		}
		if string(shippingRate.DeliveryEstimate.Minimum.Unit) != expect.ExpectedMinimumUnit {
			return fmt.Errorf(
				"remote %s.delivery_estimate.minimum.unit mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedMinimumUnit,
				string(shippingRate.DeliveryEstimate.Minimum.Unit),
			)
		}
		if shippingRate.DeliveryEstimate.Minimum.Value != expect.ExpectedMinimumValue {
			return fmt.Errorf(
				"remote %s.delivery_estimate.minimum.value mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedMinimumValue,
				shippingRate.DeliveryEstimate.Minimum.Value,
			)
		}
		if string(shippingRate.DeliveryEstimate.Maximum.Unit) != expect.ExpectedMaximumUnit {
			return fmt.Errorf(
				"remote %s.delivery_estimate.maximum.unit mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedMaximumUnit,
				string(shippingRate.DeliveryEstimate.Maximum.Unit),
			)
		}
		if shippingRate.DeliveryEstimate.Maximum.Value != expect.ExpectedMaximumValue {
			return fmt.Errorf(
				"remote %s.delivery_estimate.maximum.value mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedMaximumValue,
				shippingRate.DeliveryEstimate.Maximum.Value,
			)
		}
		if err := expectMetadataSubset(expect.Address+".metadata", shippingRate.Metadata, expect.ExpectedMetadata); err != nil {
			return err
		}
		if expect.ExpectedTaxBehavior != "" &&
			string(shippingRate.TaxBehavior) != expect.ExpectedTaxBehavior {
			return fmt.Errorf(
				"remote %s.tax_behavior mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedTaxBehavior,
				string(shippingRate.TaxBehavior),
			)
		}
		if expect.ExpectedTaxCode != "" {
			if shippingRate.TaxCode == nil {
				return fmt.Errorf("remote %s.tax_code missing", expect.Address)
			}
			if shippingRate.TaxCode.ID != expect.ExpectedTaxCode {
				return fmt.Errorf(
					"remote %s.tax_code mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedTaxCode,
					shippingRate.TaxCode.ID,
				)
			}
		}

		return nil
	}
}

func verifyShippingRateDestroyInactive(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	shippingRate, err := retrieveShippingRate(client, state, "stripe_shipping_rate.test")
	if err != nil {
		return err
	}
	if shippingRate.Active {
		return fmt.Errorf("expected stripe_shipping_rate.test to be inactive after destroy")
	}

	return nil
}

func expectStateString(
	state *terraform.State,
	address string,
	attribute string,
	expected string,
) error {
	actual, err := runner.ResourceAttribute(state, address, attribute)
	if err != nil {
		return err
	}
	if actual != expected {
		return fmt.Errorf(
			"state %s.%s mismatch: expected %q, got %q",
			address,
			attribute,
			expected,
			actual,
		)
	}

	return nil
}

func verifyTaxRate(expect taxRateExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		taxRate, err := retrieveTaxRate(client, state, expect.Address)
		if err != nil {
			return err
		}
		if taxRate.DisplayName != expect.ExpectedDisplayName {
			return fmt.Errorf(
				"remote %s.display_name mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedDisplayName,
				taxRate.DisplayName,
			)
		}
		if taxRate.Description != expect.ExpectedDescription {
			return fmt.Errorf(
				"remote %s.description mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedDescription,
				taxRate.Description,
			)
		}
		if taxRate.Country != expect.ExpectedCountry {
			return fmt.Errorf(
				"remote %s.country mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedCountry,
				taxRate.Country,
			)
		}
		if taxRate.State != expect.ExpectedState {
			return fmt.Errorf(
				"remote %s.state mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedState,
				taxRate.State,
			)
		}
		if taxRate.Jurisdiction != expect.ExpectedJurisdiction {
			return fmt.Errorf(
				"remote %s.jurisdiction mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedJurisdiction,
				taxRate.Jurisdiction,
			)
		}
		if taxRate.Inclusive != expect.ExpectedInclusive {
			return fmt.Errorf(
				"remote %s.inclusive mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedInclusive,
				taxRate.Inclusive,
			)
		}
		if err := expectFloat(expect.Address+".percentage", taxRate.Percentage, expect.ExpectedPercentage); err != nil {
			return err
		}
		if expect.ExpectedTaxType != "" && string(taxRate.TaxType) != expect.ExpectedTaxType {
			return fmt.Errorf(
				"remote %s.tax_type mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedTaxType,
				string(taxRate.TaxType),
			)
		}
		if taxRate.Active != expect.ExpectedActive {
			return fmt.Errorf(
				"remote %s.active mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedActive,
				taxRate.Active,
			)
		}
		if err := expectMetadataSubset(expect.Address+".metadata", taxRate.Metadata, expect.ExpectedMetadata); err != nil {
			return err
		}

		return nil
	}
}

func verifyTaxRateDestroyInactive(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	taxRate, err := retrieveTaxRate(client, state, "stripe_tax_rate.test")
	if err != nil {
		return err
	}
	if taxRate.Active {
		return fmt.Errorf("expected stripe_tax_rate.test to be inactive after destroy")
	}

	return nil
}

func verifyTaxRegistration(expect taxRegistrationExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		registration, err := retrieveTaxRegistration(client, state, expect.Address)
		if err != nil {
			return err
		}
		if registration.Country != expect.ExpectedCountry {
			return fmt.Errorf(
				"remote %s.country mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedCountry,
				registration.Country,
			)
		}
		if registration.CountryOptions == nil || registration.CountryOptions.Ca == nil {
			return fmt.Errorf("remote %s.country_options.ca missing", expect.Address)
		}
		if string(registration.CountryOptions.Ca.Type) != expect.ExpectedType {
			return fmt.Errorf(
				"remote %s.country_options.ca.type mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedType,
				string(registration.CountryOptions.Ca.Type),
			)
		}
		if expect.ExpectedProvince == "" {
			if registration.CountryOptions.Ca.ProvinceStandard != nil {
				return fmt.Errorf(
					"remote %s.country_options.ca.province_standard mismatch: expected absent, got present",
					expect.Address,
				)
			}
		} else {
			if registration.CountryOptions.Ca.ProvinceStandard == nil {
				return fmt.Errorf("remote %s.country_options.ca.province_standard missing", expect.Address)
			}
			if registration.CountryOptions.Ca.ProvinceStandard.Province != expect.ExpectedProvince {
				return fmt.Errorf(
					"remote %s.country_options.ca.province_standard.province mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedProvince,
					registration.CountryOptions.Ca.ProvinceStandard.Province,
				)
			}
		}
		if registration.ActiveFrom == 0 {
			return fmt.Errorf("remote %s.active_from missing", expect.Address)
		}
		if expect.ExpectedExpiresAt == 0 {
			if registration.ExpiresAt != 0 {
				return fmt.Errorf(
					"remote %s.expires_at mismatch: expected %d, got %d",
					expect.Address,
					expect.ExpectedExpiresAt,
					registration.ExpiresAt,
				)
			}
		} else if expect.ExpectedExpiresAt < 0 {
			if registration.ExpiresAt == 0 {
				return fmt.Errorf("remote %s.expires_at missing", expect.Address)
			}
		} else if registration.ExpiresAt != expect.ExpectedExpiresAt {
			return fmt.Errorf(
				"remote %s.expires_at mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedExpiresAt,
				registration.ExpiresAt,
			)
		}
		if string(registration.Status) == "" {
			return fmt.Errorf("remote %s.status missing", expect.Address)
		}

		return nil
	}
}

func verifyTaxRegistrationDestroyCleanup(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	registration, err := retrieveTaxRegistration(client, state, "stripe_tax_registration.test")
	if err != nil {
		return err
	}

	// Tax registrations are not deleted by the provider's destroy path.
	// End the test-created registration in a deterministic future slot so
	// repeated reruns do not collide on province.
	cleanupActiveFrom := time.Now().Unix() + 3600 + stableStringModulo(registration.ID, 31536000)
	cleanupExpiresAt := cleanupActiveFrom + 60
	params := &stripe.TaxRegistrationUpdateParams{
		ActiveFrom: stripe.Int64(cleanupActiveFrom),
		ExpiresAt:  stripe.Int64(cleanupExpiresAt),
	}

	_, err = client.V1TaxRegistrations.Update(
		context.Background(),
		registration.ID,
		params,
	)
	return err
}

func stableStringModulo(value string, mod int64) int64 {
	if mod <= 0 {
		return 0
	}
	var acc int64
	for i := 0; i < len(value); i++ {
		acc = (acc*31 + int64(value[i])) % mod
	}
	return acc
}

func verifyFile(expect fileExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		file, err := retrieveFile(client, state, expect.Address)
		if err != nil {
			return err
		}
		if string(file.Purpose) != expect.ExpectedPurpose {
			return fmt.Errorf(
				"remote %s.purpose mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedPurpose,
				string(file.Purpose),
			)
		}
		if file.Filename != expect.ExpectedFilename {
			return fmt.Errorf(
				"remote %s.filename mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedFilename,
				file.Filename,
			)
		}
		if file.Size == 0 {
			return fmt.Errorf("remote %s.size missing", expect.Address)
		}
		if expect.ExpectedFileLinkExpiresAt != 0 || len(expect.ExpectedFileLinkMetadata) > 0 {
			var fileLink *stripe.FileLink
			for nextFileLink, err := range client.V1FileLinks.List(context.Background(), &stripe.FileLinkListParams{
				File: stripe.String(file.ID),
			}).All(context.Background()) {
				if err != nil {
					return err
				}
				fileLink = nextFileLink
				break
			}
			if fileLink == nil {
				return fmt.Errorf("remote %s file_link missing", expect.Address)
			}
			if expect.ExpectedFileLinkExpiresAt != 0 && fileLink.ExpiresAt != expect.ExpectedFileLinkExpiresAt {
				return fmt.Errorf(
					"remote %s.file_link.expires_at mismatch: expected %d, got %d",
					expect.Address,
					expect.ExpectedFileLinkExpiresAt,
					fileLink.ExpiresAt,
				)
			}
			if err := expectMetadataSubset(expect.Address+".file_link.metadata", fileLink.Metadata, expect.ExpectedFileLinkMetadata); err != nil {
				return err
			}
		}

		return nil
	}
}

func verifyFileDestroyStillExists(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	_, err := retrieveFile(client, state, "stripe_file.test")
	return err
}

func verifyTreasuryFinancialAccount(
	expect treasuryFinancialAccountExpectations,
) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		financialAccount, err := retrieveTreasuryFinancialAccount(client, state, expect.Address)
		if err != nil {
			return err
		}
		if financialAccount.Nickname != expect.ExpectedNickname {
			return fmt.Errorf(
				"remote %s.nickname mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedNickname,
				financialAccount.Nickname,
			)
		}
		actualCurrencies := make([]string, 0, len(financialAccount.SupportedCurrencies))
		for _, currency := range financialAccount.SupportedCurrencies {
			actualCurrencies = append(actualCurrencies, string(currency))
		}
		if err := expectRemoteStringList(
			expect.Address+".supported_currencies",
			actualCurrencies,
			expect.ExpectedSupportedCurrencies,
		); err != nil {
			return err
		}
		if string(financialAccount.Status) == "" {
			return fmt.Errorf("remote %s.status missing", expect.Address)
		}
		if len(expect.ExpectedFeaturePaths) > 0 {
			availableFeatures := map[string]bool{}
			for _, feature := range financialAccount.ActiveFeatures {
				availableFeatures[string(feature)] = true
			}
			for _, feature := range financialAccount.PendingFeatures {
				availableFeatures[string(feature)] = true
			}
			for _, feature := range financialAccount.RestrictedFeatures {
				availableFeatures[string(feature)] = true
			}
			for _, expectedFeature := range expect.ExpectedFeaturePaths {
				if !availableFeatures[expectedFeature] {
					return fmt.Errorf(
						"remote %s missing requested feature path %q in active/pending/restricted features",
						expect.Address,
						expectedFeature,
					)
				}
			}
		}
		if err := expectMetadataSubset(expect.Address+".metadata", financialAccount.Metadata, expect.ExpectedMetadata); err != nil {
			return err
		}

		return nil
	}
}

func verifyTreasuryFinancialAccountDestroyStillExists(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	account, err := retrieveTreasuryFinancialAccount(
		client,
		state,
		"stripe_treasury_financial_account.test",
	)
	if err != nil {
		return err
	}

	_, err = client.V1TreasuryFinancialAccounts.Close(
		context.Background(),
		account.ID,
		nil,
	)
	if err != nil {
		return fmt.Errorf("close stripe_treasury_financial_account.test (%s): %w", account.ID, err)
	}

	return nil
}

func retrieveSubscription(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.Subscription, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	params := &stripe.SubscriptionRetrieveParams{}
	params.AddExpand("items.data.price")
	subscription, err := client.V1Subscriptions.Retrieve(context.Background(), id, params)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return subscription, nil
}

func retrieveSubscriptionItem(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.SubscriptionItem, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	item, err := client.V1SubscriptionItems.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return item, nil
}

func retrieveInvoice(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.Invoice, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	invoice, err := client.V1Invoices.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return invoice, nil
}

func retrieveSetupIntent(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.SetupIntent, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	setupIntent, err := client.V1SetupIntents.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return setupIntent, nil
}

func retrievePaymentMethod(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.PaymentMethod, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	paymentMethod, err := client.V1PaymentMethods.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return paymentMethod, nil
}

func retrievePaymentLink(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.PaymentLink, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	params := &stripe.PaymentLinkRetrieveParams{}
	params.AddExpand("line_items")
	params.AddExpand("line_items.data.price")
	paymentLink, err := client.V1PaymentLinks.Retrieve(context.Background(), id, params)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return paymentLink, nil
}

func retrieveQuote(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.Quote, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	params := &stripe.QuoteRetrieveParams{}
	params.AddExpand("line_items")
	params.AddExpand("line_items.data.price")
	quote, err := client.V1Quotes.Retrieve(context.Background(), id, params)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return quote, nil
}

func retrieveCharge(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.Charge, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	charge, err := client.V1Charges.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return charge, nil
}

func retrieveSource(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.Source, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	source, err := client.V1Sources.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return source, nil
}

func retrieveShippingRate(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.ShippingRate, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	shippingRate, err := client.V1ShippingRates.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return shippingRate, nil
}

func retrieveTaxRate(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.TaxRate, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	taxRate, err := client.V1TaxRates.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return taxRate, nil
}

func retrieveTaxRegistration(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.TaxRegistration, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	registration, err := client.V1TaxRegistrations.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return registration, nil
}

func retrieveFile(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.File, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	file, err := client.V1Files.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return file, nil
}

func retrieveTreasuryFinancialAccount(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.TreasuryFinancialAccount, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	financialAccount, err := client.V1TreasuryFinancialAccounts.Retrieve(
		context.Background(),
		id,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return financialAccount, nil
}
