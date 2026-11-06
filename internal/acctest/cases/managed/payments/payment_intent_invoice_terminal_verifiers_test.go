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

type paymentIntentShippingExpectation struct {
	Name           string
	Phone          string
	Carrier        string
	TrackingNumber string
	AddressLine1   string
	City           string
	State          string
	PostalCode     string
	Country        string
}

type paymentIntentExpectations struct {
	Address                                string
	ExpectedAmount                         int64
	ExpectedCurrency                       string
	ExpectedDescription                    string
	ExpectedMetadata                       map[string]string
	CheckAutomaticPaymentMethods           bool
	ExpectedAutomaticPaymentMethodsEnabled bool
	ExpectedReceiptEmail                   string
	ExpectedSetupFutureUsage               string
	ExpectedStatementDescriptorSuffix      string
	ExpectedCustomerAddress                string
	ExpectedShipping                       *paymentIntentShippingExpectation
	CheckTipAmount                         bool
	ExpectedTipAmount                      int64
}

type periodExpectation struct {
	Start int64
	End   int64
}

type invoiceItemExpectations struct {
	Address                     string
	ExpectedCustomerAddress     string
	ExpectedInvoiceAddress      string
	ExpectedAmount              int64
	ExpectedCurrency            string
	ExpectedDescription         string
	ExpectedDiscountable        *bool
	ExpectedMetadata            map[string]string
	ExpectedQuantity            *int64
	ExpectedPeriod              *periodExpectation
	ExpectedPricingPriceAddress string
}

type subscriptionScheduleExpectations struct {
	Address                                 string
	ExpectedCustomerAddress                 string
	ExpectedEndBehavior                     string
	ExpectedDefaultSettingsCollectionMethod string
	ExpectedDefaultSettingsDescription      string
	ExpectedDefaultSettingsDaysUntilDue     int64
	ExpectedMetadata                        map[string]string
	ExpectedTopLevelStartDate               int64
	ExpectedFirstPriceAddress               string
	ExpectedFirstQuantity                   int64
}

type terminalReaderExpectations struct {
	Address                 string
	ExpectedLabel           string
	ExpectedMetadata        map[string]string
	ExpectedLocationFromEnv bool
	CheckDeviceDetails      bool
}

type identityVerificationSessionEphemeralExpectations struct {
	ClientReferenceID                string
	ExpectedType                     string
	ExpectedMetadata                 map[string]string
	ExpectedEmail                    string
	ExpectedPhone                    string
	ExpectedAllowedTypes             []string
	ExpectedRequireIDNumber          bool
	ExpectedRequireLiveCapture       bool
	ExpectedRequireMatchingSelfie    bool
	ExpectedRequireEmailVerification bool
	ExpectedRequirePhoneVerification bool
}

type forwardingRequestActionExpectations struct {
	CaseName             string
	ExpectedBody         string
	ExpectedReplacements []string
	ExpectedHeaders      map[string]string
}

func verifyPaymentIntent(expect paymentIntentExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		paymentIntent, err := retrievePaymentIntent(client, state, expect.Address)
		if err != nil {
			return err
		}
		if paymentIntent.Amount != expect.ExpectedAmount {
			return fmt.Errorf(
				"remote %s.amount mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedAmount,
				paymentIntent.Amount,
			)
		}
		if string(paymentIntent.Currency) != expect.ExpectedCurrency {
			return fmt.Errorf(
				"remote %s.currency mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedCurrency,
				string(paymentIntent.Currency),
			)
		}
		if expect.ExpectedDescription != "" && paymentIntent.Description != expect.ExpectedDescription {
			return fmt.Errorf(
				"remote %s.description mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedDescription,
				paymentIntent.Description,
			)
		}
		if err := expectMetadataSubset(
			expect.Address+".metadata",
			paymentIntent.Metadata,
			expect.ExpectedMetadata,
		); err != nil {
			return err
		}
		if expect.CheckAutomaticPaymentMethods {
			if paymentIntent.AutomaticPaymentMethods == nil {
				return fmt.Errorf("remote %s.automatic_payment_methods missing", expect.Address)
			}
			if paymentIntent.AutomaticPaymentMethods.Enabled !=
				expect.ExpectedAutomaticPaymentMethodsEnabled {
				return fmt.Errorf(
					"remote %s.automatic_payment_methods.enabled mismatch: expected %t, got %t",
					expect.Address,
					expect.ExpectedAutomaticPaymentMethodsEnabled,
					paymentIntent.AutomaticPaymentMethods.Enabled,
				)
			}
		}
		if expect.ExpectedReceiptEmail != "" &&
			paymentIntent.ReceiptEmail != expect.ExpectedReceiptEmail {
			return fmt.Errorf(
				"remote %s.receipt_email mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedReceiptEmail,
				paymentIntent.ReceiptEmail,
			)
		}
		if expect.ExpectedSetupFutureUsage != "" &&
			string(paymentIntent.SetupFutureUsage) != expect.ExpectedSetupFutureUsage {
			return fmt.Errorf(
				"remote %s.setup_future_usage mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedSetupFutureUsage,
				string(paymentIntent.SetupFutureUsage),
			)
		}
		if expect.ExpectedStatementDescriptorSuffix != "" &&
			paymentIntent.StatementDescriptorSuffix != expect.ExpectedStatementDescriptorSuffix {
			return fmt.Errorf(
				"remote %s.statement_descriptor_suffix mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedStatementDescriptorSuffix,
				paymentIntent.StatementDescriptorSuffix,
			)
		}
		if expect.ExpectedCustomerAddress != "" {
			expectedCustomerID, err := runner.ResourcePrimaryID(state, expect.ExpectedCustomerAddress)
			if err != nil {
				return err
			}
			if paymentIntent.Customer == nil || paymentIntent.Customer.ID != expectedCustomerID {
				actualCustomerID := ""
				if paymentIntent.Customer != nil {
					actualCustomerID = paymentIntent.Customer.ID
				}
				return fmt.Errorf(
					"remote %s.customer mismatch: expected %q, got %q",
					expect.Address,
					expectedCustomerID,
					actualCustomerID,
				)
			}
		}
		if expect.ExpectedShipping != nil {
			if paymentIntent.Shipping == nil || paymentIntent.Shipping.Address == nil {
				return fmt.Errorf("remote %s.shipping missing", expect.Address)
			}
			if paymentIntent.Shipping.Name != expect.ExpectedShipping.Name {
				return fmt.Errorf(
					"remote %s.shipping.name mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedShipping.Name,
					paymentIntent.Shipping.Name,
				)
			}
			if paymentIntent.Shipping.Phone != expect.ExpectedShipping.Phone {
				return fmt.Errorf(
					"remote %s.shipping.phone mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedShipping.Phone,
					paymentIntent.Shipping.Phone,
				)
			}
			if paymentIntent.Shipping.Carrier != expect.ExpectedShipping.Carrier {
				return fmt.Errorf(
					"remote %s.shipping.carrier mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedShipping.Carrier,
					paymentIntent.Shipping.Carrier,
				)
			}
			if paymentIntent.Shipping.TrackingNumber != expect.ExpectedShipping.TrackingNumber {
				return fmt.Errorf(
					"remote %s.shipping.tracking_number mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedShipping.TrackingNumber,
					paymentIntent.Shipping.TrackingNumber,
				)
			}
			if paymentIntent.Shipping.Address.Line1 != expect.ExpectedShipping.AddressLine1 {
				return fmt.Errorf(
					"remote %s.shipping.address.line1 mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedShipping.AddressLine1,
					paymentIntent.Shipping.Address.Line1,
				)
			}
			if paymentIntent.Shipping.Address.City != expect.ExpectedShipping.City {
				return fmt.Errorf(
					"remote %s.shipping.address.city mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedShipping.City,
					paymentIntent.Shipping.Address.City,
				)
			}
			if paymentIntent.Shipping.Address.State != expect.ExpectedShipping.State {
				return fmt.Errorf(
					"remote %s.shipping.address.state mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedShipping.State,
					paymentIntent.Shipping.Address.State,
				)
			}
			if paymentIntent.Shipping.Address.PostalCode != expect.ExpectedShipping.PostalCode {
				return fmt.Errorf(
					"remote %s.shipping.address.postal_code mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedShipping.PostalCode,
					paymentIntent.Shipping.Address.PostalCode,
				)
			}
			if paymentIntent.Shipping.Address.Country != expect.ExpectedShipping.Country {
				return fmt.Errorf(
					"remote %s.shipping.address.country mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedShipping.Country,
					paymentIntent.Shipping.Address.Country,
				)
			}
		}
		if expect.CheckTipAmount {
			if paymentIntent.AmountDetails == nil || paymentIntent.AmountDetails.Tip == nil {
				return fmt.Errorf("remote %s.amount_details.tip missing", expect.Address)
			}
			if paymentIntent.AmountDetails.Tip.Amount != expect.ExpectedTipAmount {
				return fmt.Errorf(
					"remote %s.amount_details.tip.amount mismatch: expected %d, got %d",
					expect.Address,
					expect.ExpectedTipAmount,
					paymentIntent.AmountDetails.Tip.Amount,
				)
			}
		}

		return nil
	}
}

func verifyPaymentIntentDestroyCanceled(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	paymentIntent, err := retrievePaymentIntent(client, state, "stripe_payment_intent.test")
	if err != nil {
		return err
	}
	if string(paymentIntent.Status) != "canceled" && string(paymentIntent.Status) != "requires_payment_method" {
		return fmt.Errorf(
			"expected stripe_payment_intent.test to be canceled or left requires_payment_method after destroy, got %q",
			string(paymentIntent.Status),
		)
	}
	if string(paymentIntent.Status) == "canceled" && paymentIntent.CanceledAt == 0 {
		return fmt.Errorf("expected stripe_payment_intent.test canceled_at after destroy")
	}

	return nil
}

func verifyInvoiceItem(expect invoiceItemExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		invoiceItem, err := retrieveInvoiceItem(client, state, expect.Address)
		if err != nil {
			return err
		}
		if expect.ExpectedCustomerAddress != "" {
			expectedCustomerID, err := runner.ResourcePrimaryID(state, expect.ExpectedCustomerAddress)
			if err != nil {
				return err
			}
			if invoiceItem.Customer == nil || invoiceItem.Customer.ID != expectedCustomerID {
				actualCustomerID := ""
				if invoiceItem.Customer != nil {
					actualCustomerID = invoiceItem.Customer.ID
				}
				return fmt.Errorf(
					"remote %s.customer mismatch: expected %q, got %q",
					expect.Address,
					expectedCustomerID,
					actualCustomerID,
				)
			}
		}
		if expect.ExpectedInvoiceAddress != "" {
			expectedInvoiceID, err := runner.ResourcePrimaryID(state, expect.ExpectedInvoiceAddress)
			if err != nil {
				return err
			}
			actualInvoiceID := ""
			if invoiceItem.Invoice != nil {
				actualInvoiceID = invoiceItem.Invoice.ID
			}
			if actualInvoiceID != expectedInvoiceID {
				return fmt.Errorf(
					"remote %s.invoice mismatch: expected %q, got %q",
					expect.Address,
					expectedInvoiceID,
					actualInvoiceID,
				)
			}
		}
		if expect.ExpectedAmount != 0 && invoiceItem.Amount != expect.ExpectedAmount {
			return fmt.Errorf(
				"remote %s.amount mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedAmount,
				invoiceItem.Amount,
			)
		}
		if expect.ExpectedCurrency != "" && string(invoiceItem.Currency) != expect.ExpectedCurrency {
			return fmt.Errorf(
				"remote %s.currency mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedCurrency,
				string(invoiceItem.Currency),
			)
		}
		if expect.ExpectedDescription != "" &&
			invoiceItem.Description != expect.ExpectedDescription {
			return fmt.Errorf(
				"remote %s.description mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedDescription,
				invoiceItem.Description,
			)
		}
		if expect.ExpectedDiscountable != nil && invoiceItem.Discountable != *expect.ExpectedDiscountable {
			return fmt.Errorf(
				"remote %s.discountable mismatch: expected %t, got %t",
				expect.Address,
				*expect.ExpectedDiscountable,
				invoiceItem.Discountable,
			)
		}
		if err := expectMetadataSubset(
			expect.Address+".metadata",
			invoiceItem.Metadata,
			expect.ExpectedMetadata,
		); err != nil {
			return err
		}
		if expect.ExpectedQuantity != nil && invoiceItem.Quantity != *expect.ExpectedQuantity {
			return fmt.Errorf(
				"remote %s.quantity mismatch: expected %d, got %d",
				expect.Address,
				*expect.ExpectedQuantity,
				invoiceItem.Quantity,
			)
		}
		if expect.ExpectedPeriod != nil {
			if invoiceItem.Period == nil {
				return fmt.Errorf("remote %s.period missing", expect.Address)
			}
			if invoiceItem.Period.Start != expect.ExpectedPeriod.Start {
				return fmt.Errorf(
					"remote %s.period.start mismatch: expected %d, got %d",
					expect.Address,
					expect.ExpectedPeriod.Start,
					invoiceItem.Period.Start,
				)
			}
			if invoiceItem.Period.End != expect.ExpectedPeriod.End {
				return fmt.Errorf(
					"remote %s.period.end mismatch: expected %d, got %d",
					expect.Address,
					expect.ExpectedPeriod.End,
					invoiceItem.Period.End,
				)
			}
		}
		if expect.ExpectedPricingPriceAddress != "" {
			expectedPriceID, err := runner.ResourcePrimaryID(state, expect.ExpectedPricingPriceAddress)
			if err != nil {
				return err
			}
			if invoiceItem.Pricing == nil || invoiceItem.Pricing.PriceDetails == nil {
				return fmt.Errorf("remote %s.pricing.price_details missing", expect.Address)
			}
			actualPriceID := ""
			if invoiceItem.Pricing.PriceDetails.Price != nil {
				actualPriceID = invoiceItem.Pricing.PriceDetails.Price.ID
			}
			if actualPriceID != expectedPriceID {
				return fmt.Errorf(
					"remote %s.pricing.price_details.price mismatch: expected %q, got %q",
					expect.Address,
					expectedPriceID,
					actualPriceID,
				)
			}
		}

		return nil
	}
}

func verifyInvoiceItemDestroyMissing(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	id, err := runner.ResourcePrimaryID(state, "stripe_invoice_item.test")
	if err != nil {
		return err
	}

	_, err = client.V1InvoiceItems.Retrieve(context.Background(), id, nil)
	return expectRemoteMissing("stripe_invoice_item.test", id, err)
}

func verifySubscriptionSchedule(expect subscriptionScheduleExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		schedule, err := retrieveSubscriptionSchedule(client, state, expect.Address)
		if err != nil {
			return err
		}
		expectedCustomerID, err := runner.ResourcePrimaryID(state, expect.ExpectedCustomerAddress)
		if err != nil {
			return err
		}
		if schedule.Customer == nil || schedule.Customer.ID != expectedCustomerID {
			actualCustomerID := ""
			if schedule.Customer != nil {
				actualCustomerID = schedule.Customer.ID
			}
			return fmt.Errorf(
				"remote %s.customer mismatch: expected %q, got %q",
				expect.Address,
				expectedCustomerID,
				actualCustomerID,
			)
		}
		if string(schedule.EndBehavior) != expect.ExpectedEndBehavior {
			return fmt.Errorf(
				"remote %s.end_behavior mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedEndBehavior,
				string(schedule.EndBehavior),
			)
		}
		if err := expectMetadataSubset(
			expect.Address+".metadata",
			schedule.Metadata,
			expect.ExpectedMetadata,
		); err != nil {
			return err
		}
		if expect.ExpectedDefaultSettingsCollectionMethod != "" {
			if schedule.DefaultSettings == nil {
				return fmt.Errorf("remote %s.default_settings missing", expect.Address)
			}
			if schedule.DefaultSettings.CollectionMethod == nil {
				return fmt.Errorf("remote %s.default_settings.collection_method missing", expect.Address)
			}
			if string(*schedule.DefaultSettings.CollectionMethod) != expect.ExpectedDefaultSettingsCollectionMethod {
				return fmt.Errorf(
					"remote %s.default_settings.collection_method mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedDefaultSettingsCollectionMethod,
					string(*schedule.DefaultSettings.CollectionMethod),
				)
			}
		}
		if expect.ExpectedDefaultSettingsDescription != "" {
			if schedule.DefaultSettings == nil {
				return fmt.Errorf("remote %s.default_settings missing", expect.Address)
			}
			if schedule.DefaultSettings.Description != expect.ExpectedDefaultSettingsDescription {
				return fmt.Errorf(
					"remote %s.default_settings.description mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedDefaultSettingsDescription,
					schedule.DefaultSettings.Description,
				)
			}
		}
		if expect.ExpectedDefaultSettingsDaysUntilDue != 0 {
			if schedule.DefaultSettings == nil || schedule.DefaultSettings.InvoiceSettings == nil {
				return fmt.Errorf("remote %s.default_settings.invoice_settings missing", expect.Address)
			}
			if schedule.DefaultSettings.InvoiceSettings.DaysUntilDue != expect.ExpectedDefaultSettingsDaysUntilDue {
				return fmt.Errorf(
					"remote %s.default_settings.invoice_settings.days_until_due mismatch: expected %d, got %d",
					expect.Address,
					expect.ExpectedDefaultSettingsDaysUntilDue,
					schedule.DefaultSettings.InvoiceSettings.DaysUntilDue,
				)
			}
		}
		if len(schedule.Phases) == 0 || schedule.Phases[0] == nil {
			return fmt.Errorf("remote %s.phases[0] missing", expect.Address)
		}
		if schedule.Phases[0].StartDate != expect.ExpectedTopLevelStartDate {
			return fmt.Errorf(
				"remote %s.phases[0].start_date mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedTopLevelStartDate,
				schedule.Phases[0].StartDate,
			)
		}
		expectedStateStartDate, err := runner.ResourceAttribute(state, expect.Address, "start_date")
		if err == nil && expectedStateStartDate != fmt.Sprintf("%d", expect.ExpectedTopLevelStartDate) {
			return fmt.Errorf(
				"terraform state %s.start_date mismatch: expected %d, got %s",
				expect.Address,
				expect.ExpectedTopLevelStartDate,
				expectedStateStartDate,
			)
		}
		expectedPriceID, err := runner.ResourcePrimaryID(state, expect.ExpectedFirstPriceAddress)
		if err != nil {
			return err
		}
		if len(schedule.Phases[0].Items) == 0 || schedule.Phases[0].Items[0] == nil {
			return fmt.Errorf("remote %s.phases[0].items[0] missing", expect.Address)
		}
		if schedule.Phases[0].Items[0].Price == nil ||
			schedule.Phases[0].Items[0].Price.ID != expectedPriceID {
			actualPriceID := ""
			if schedule.Phases[0].Items[0].Price != nil {
				actualPriceID = schedule.Phases[0].Items[0].Price.ID
			}
			return fmt.Errorf(
				"remote %s.phases[0].items[0].price mismatch: expected %q, got %q",
				expect.Address,
				expectedPriceID,
				actualPriceID,
			)
		}
		if schedule.Phases[0].Items[0].Quantity != expect.ExpectedFirstQuantity {
			return fmt.Errorf(
				"remote %s.phases[0].items[0].quantity mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedFirstQuantity,
				schedule.Phases[0].Items[0].Quantity,
			)
		}

		return nil
	}
}

func verifySubscriptionScheduleDestroyCanceled(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	schedule, err := retrieveSubscriptionSchedule(
		client,
		state,
		"stripe_subscription_schedule.test",
	)
	if err != nil {
		return err
	}
	if string(schedule.Status) != "canceled" {
		return fmt.Errorf(
			"expected stripe_subscription_schedule.test to be canceled after destroy, got %q",
			string(schedule.Status),
		)
	}
	if schedule.CanceledAt == 0 {
		return fmt.Errorf("expected stripe_subscription_schedule.test canceled_at after destroy")
	}

	return nil
}

func verifyTerminalReader(expect terminalReaderExpectations) runner.StateVerifier {
	return func(
		env runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		reader, err := retrieveTerminalReader(client, state, expect.Address)
		if err != nil {
			return err
		}
		if expect.ExpectedLabel != "" && reader.Label != expect.ExpectedLabel {
			return fmt.Errorf(
				"remote %s.label mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedLabel,
				reader.Label,
			)
		}
		if err := expectMetadataSubset(
			expect.Address+".metadata",
			reader.Metadata,
			expect.ExpectedMetadata,
		); err != nil {
			return err
		}
		if expect.ExpectedLocationFromEnv {
			if reader.Location == nil || reader.Location.ID != env.TerminalLocation {
				actualLocationID := ""
				if reader.Location != nil {
					actualLocationID = reader.Location.ID
				}
				return fmt.Errorf(
					"remote %s.location mismatch: expected %q, got %q",
					expect.Address,
					env.TerminalLocation,
					actualLocationID,
				)
			}
		}
		if expect.CheckDeviceDetails {
			if reader.DeviceType == "" {
				return fmt.Errorf("remote %s.device_type unexpectedly empty", expect.Address)
			}
			if string(reader.Status) == "" {
				return fmt.Errorf("remote %s.status unexpectedly empty", expect.Address)
			}
			if reader.SerialNumber == "" {
				return fmt.Errorf("remote %s.serial_number unexpectedly empty", expect.Address)
			}
		}

		return nil
	}
}

func verifyTerminalReaderDestroyMissing(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	id, err := runner.ResourcePrimaryID(state, "stripe_terminal_reader.test")
	if err != nil {
		return err
	}

	reader, err := client.V1TerminalReaders.Retrieve(context.Background(), id, nil)
	if err == nil && reader != nil && reader.Deleted {
		return nil
	}
	return expectRemoteMissing("stripe_terminal_reader.test", id, err)
}

func verifyTokenEphemeralState(
	_ runner.TestEnv,
	_ *stripe.Client,
	state *terraform.State,
) error {
	if err := runner.ExpectResourceAbsent(state, "ephemeral.stripe_token.test"); err != nil {
		return err
	}
	for _, outputName := range []string{"token_id", "token_type", "token_reference", "token_bank_country"} {
		if err := runner.ExpectOutputAbsent(state, outputName); err != nil {
			return err
		}
	}

	return nil
}

func verifyIdentityVerificationSessionEphemeral(
	expect identityVerificationSessionEphemeralExpectations,
) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		if err := runner.ExpectResourceAbsent(
			state,
			"ephemeral.stripe_identity_verification_session.test",
		); err != nil {
			return err
		}
		for _, outputName := range []string{
			"verification_session_client_reference_id",
			"verification_session_type",
			"verification_session_url",
			"verification_session_client_secret",
			"verification_session_provided_email",
			"verification_session_provided_phone",
		} {
			if err := runner.ExpectOutputAbsent(state, outputName); err != nil {
				return err
			}
		}

		session, err := findIdentityVerificationSession(
			client,
			expect.ClientReferenceID,
		)
		if err != nil {
			return err
		}
		if string(session.Type) != expect.ExpectedType {
			return fmt.Errorf(
				"identity verification session %s type mismatch: expected %q, got %q",
				expect.ClientReferenceID,
				expect.ExpectedType,
				string(session.Type),
			)
		}
		if err := expectMetadataSubset(
			expect.ClientReferenceID+".metadata",
			session.Metadata,
			expect.ExpectedMetadata,
		); err != nil {
			return err
		}
		if expect.ExpectedEmail != "" {
			if session.ProvidedDetails == nil || session.ProvidedDetails.Email != expect.ExpectedEmail {
				actualEmail := ""
				if session.ProvidedDetails != nil {
					actualEmail = session.ProvidedDetails.Email
				}
				return fmt.Errorf(
					"identity verification session %s provided_details.email mismatch: expected %q, got %q",
					expect.ClientReferenceID,
					expect.ExpectedEmail,
					actualEmail,
				)
			}
		}
		if expect.ExpectedPhone != "" {
			if session.ProvidedDetails == nil || session.ProvidedDetails.Phone != expect.ExpectedPhone {
				actualPhone := ""
				if session.ProvidedDetails != nil {
					actualPhone = session.ProvidedDetails.Phone
				}
				return fmt.Errorf(
					"identity verification session %s provided_details.phone mismatch: expected %q, got %q",
					expect.ClientReferenceID,
					expect.ExpectedPhone,
					actualPhone,
				)
			}
		}
		if len(expect.ExpectedAllowedTypes) > 0 {
			if session.Options == nil || session.Options.Document == nil {
				return fmt.Errorf(
					"identity verification session %s document options missing",
					expect.ClientReferenceID,
				)
			}
			actualAllowedTypes := make([]string, 0, len(session.Options.Document.AllowedTypes))
			for _, allowedType := range session.Options.Document.AllowedTypes {
				actualAllowedTypes = append(actualAllowedTypes, string(allowedType))
			}
			if err := expectRemoteStringList(
				expect.ClientReferenceID+".options.document.allowed_types",
				actualAllowedTypes,
				expect.ExpectedAllowedTypes,
			); err != nil {
				return err
			}
			if session.Options.Document.RequireIDNumber != expect.ExpectedRequireIDNumber {
				return fmt.Errorf(
					"identity verification session %s require_id_number mismatch: expected %t, got %t",
					expect.ClientReferenceID,
					expect.ExpectedRequireIDNumber,
					session.Options.Document.RequireIDNumber,
				)
			}
			if session.Options.Document.RequireLiveCapture != expect.ExpectedRequireLiveCapture {
				return fmt.Errorf(
					"identity verification session %s require_live_capture mismatch: expected %t, got %t",
					expect.ClientReferenceID,
					expect.ExpectedRequireLiveCapture,
					session.Options.Document.RequireLiveCapture,
				)
			}
			if session.Options.Document.RequireMatchingSelfie !=
				expect.ExpectedRequireMatchingSelfie {
				return fmt.Errorf(
					"identity verification session %s require_matching_selfie mismatch: expected %t, got %t",
					expect.ClientReferenceID,
					expect.ExpectedRequireMatchingSelfie,
					session.Options.Document.RequireMatchingSelfie,
				)
			}
		}
		if expect.ExpectedRequireEmailVerification {
			if session.Options == nil || session.Options.Email == nil || !session.Options.Email.RequireVerification {
				return fmt.Errorf(
					"identity verification session %s expected email verification",
					expect.ClientReferenceID,
				)
			}
		}
		if expect.ExpectedRequirePhoneVerification {
			if session.Options == nil || session.Options.Phone == nil || !session.Options.Phone.RequireVerification {
				return fmt.Errorf(
					"identity verification session %s expected phone verification",
					expect.ClientReferenceID,
				)
			}
		}

		return nil
	}
}

func verifyForwardingRequestAction(
	expect forwardingRequestActionExpectations,
) runner.StateVerifier {
	return func(
		env runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		if err := runner.ExpectResourceAbsent(state, "action.stripe_forwarding_request.test"); err != nil {
			return err
		}

		request, err := findForwardingRequest(client, expect.CaseName)
		if err != nil {
			return err
		}
		if request.URL != env.ForwardingDestination {
			return fmt.Errorf(
				"forwarding request %s url mismatch: expected %q, got %q",
				expect.CaseName,
				env.ForwardingDestination,
				request.URL,
			)
		}
		if request.PaymentMethod != env.ForwardingPaymentMethod {
			return fmt.Errorf(
				"forwarding request %s payment_method mismatch: expected %q, got %q",
				expect.CaseName,
				env.ForwardingPaymentMethod,
				request.PaymentMethod,
			)
		}
		if request.RequestDetails == nil {
			return fmt.Errorf("forwarding request %s request_details missing", expect.CaseName)
		}
		if request.RequestDetails.Body != expect.ExpectedBody {
			return fmt.Errorf(
				"forwarding request %s body mismatch: expected %q, got %q",
				expect.CaseName,
				expect.ExpectedBody,
				request.RequestDetails.Body,
			)
		}
		if string(request.RequestDetails.HTTPMethod) != "POST" {
			return fmt.Errorf(
				"forwarding request %s http_method mismatch: expected %q, got %q",
				expect.CaseName,
				"POST",
				string(request.RequestDetails.HTTPMethod),
			)
		}
		actualReplacements := make([]string, 0, len(request.Replacements))
		for _, replacement := range request.Replacements {
			actualReplacements = append(actualReplacements, string(replacement))
		}
		if err := expectRemoteStringList(
			expect.CaseName+".replacements",
			actualReplacements,
			expect.ExpectedReplacements,
		); err != nil {
			return err
		}
		actualHeaders := map[string]string{}
		for _, header := range request.RequestDetails.Headers {
			if header == nil {
				continue
			}
			actualHeaders[header.Name] = header.Value
		}
		for key, expectedValue := range expect.ExpectedHeaders {
			actualValue, ok := actualHeaders[key]
			if !ok {
				return fmt.Errorf(
					"forwarding request %s missing header %q",
					expect.CaseName,
					key,
				)
			}
			if actualValue != expectedValue {
				return fmt.Errorf(
					"forwarding request %s header %q mismatch: expected %q, got %q",
					expect.CaseName,
					key,
					expectedValue,
					actualValue,
				)
			}
		}
		if err := expectMetadataSubset(
			expect.CaseName+".metadata",
			request.Metadata,
			map[string]string{
				"suite": "sdk-codegen",
				"case":  expect.CaseName,
			},
		); err != nil {
			return err
		}

		return nil
	}
}

func retrievePaymentIntent(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.PaymentIntent, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	paymentIntent, err := client.V1PaymentIntents.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return paymentIntent, nil
}

func retrieveInvoiceItem(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.InvoiceItem, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	invoiceItem, err := client.V1InvoiceItems.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return invoiceItem, nil
}

func retrieveSubscriptionSchedule(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.SubscriptionSchedule, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	schedule, err := client.V1SubscriptionSchedules.Retrieve(
		context.Background(),
		id,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return schedule, nil
}

func retrieveTerminalReader(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.TerminalReader, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	reader, err := client.V1TerminalReaders.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return reader, nil
}

func findIdentityVerificationSession(
	client *stripe.Client,
	clientReferenceID string,
) (*stripe.IdentityVerificationSession, error) {
	params := &stripe.IdentityVerificationSessionListParams{
		ClientReferenceID: stripe.String(clientReferenceID),
	}
	params.Limit = stripe.Int64(10)

	for session, err := range client.V1IdentityVerificationSessions.List(
		context.Background(),
		params,
	).All(context.Background()) {
		if err != nil {
			return nil, fmt.Errorf(
				"list identity verification sessions for %s: %w",
				clientReferenceID,
				err,
			)
		}
		if session != nil && session.ClientReferenceID == clientReferenceID {
			return session, nil
		}
	}

	return nil, fmt.Errorf(
		"identity verification session with client_reference_id %q not found",
		clientReferenceID,
	)
}

func findForwardingRequest(
	client *stripe.Client,
	caseName string,
) (*stripe.ForwardingRequest, error) {
	params := &stripe.ForwardingRequestListParams{}
	params.Limit = stripe.Int64(25)

	for request, err := range client.V1ForwardingRequests.List(
		context.Background(),
		params,
	).All(context.Background()) {
		if err != nil {
			return nil, fmt.Errorf("list forwarding requests for %s: %w", caseName, err)
		}
		if request == nil {
			continue
		}
		if request.Metadata["suite"] == "sdk-codegen" && request.Metadata["case"] == caseName {
			return request, nil
		}
	}

	return nil, fmt.Errorf("forwarding request for case %q not found", caseName)
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
