// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import (
	"testing"

	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

func runManagedCaseForGroup(
	t *testing.T,
	group string,
	requiredEnv []string,
	name string,
	surface string,
	resourceAddress string,
	createTemplate string,
	updateTemplate string,
	importStable bool,
	verifyCreate runner.StateVerifier,
	verifyUpdate runner.StateVerifier,
	verifyDestroy runner.StateVerifier,
) {
	t.Helper()

	runner.RunManagedCase(t, runner.ManagedCase{
		Definition: runner.CaseDefinition{
			Name:        name,
			Surface:     surface,
			Group:       group,
			Kind:        "resource",
			RequiredEnv: requiredEnv,
		},
		ResourceAddress: resourceAddress,
		CreateTemplate:  createTemplate,
		UpdateTemplate:  updateTemplate,
		ImportStable:    importStable,
		VerifyCreate:    verifyCreate,
		VerifyUpdate:    verifyUpdate,
		VerifyDestroy:   verifyDestroy,
	})
}

func TestAccManagedPaymentIntentBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"payment_intent_basic",
		"stripe_payment_intent",
		"stripe_payment_intent.test",
		"payment_intent_basic_create.tf",
		"payment_intent_basic_update.tf",
		true,
		verifyPaymentIntent(paymentIntentExpectations{
			Address:             "stripe_payment_intent.test",
			ExpectedAmount:      2500,
			ExpectedCurrency:    "usd",
			ExpectedDescription: "sdk-codegen payment intent basic",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "payment_intent_basic",
				"phase": "create",
			},
			CheckAutomaticPaymentMethods:           true,
			ExpectedAutomaticPaymentMethodsEnabled: true,
		}),
		verifyPaymentIntent(paymentIntentExpectations{
			Address:             "stripe_payment_intent.test",
			ExpectedAmount:      2500,
			ExpectedCurrency:    "usd",
			ExpectedDescription: "sdk-codegen payment intent basic updated",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "payment_intent_basic",
				"phase": "update",
			},
			CheckAutomaticPaymentMethods:           true,
			ExpectedAutomaticPaymentMethodsEnabled: true,
		}),
		verifyPaymentIntentDestroyCanceled,
	)
}

func TestAccManagedPaymentIntentShipping(t *testing.T) {
	runBaseManagedCase(
		t,
		"payment_intent_shipping",
		"stripe_payment_intent",
		"stripe_payment_intent.test",
		"payment_intent_shipping_create.tf",
		"payment_intent_shipping_update.tf",
		true,
		verifyPaymentIntent(paymentIntentExpectations{
			Address:          "stripe_payment_intent.test",
			ExpectedAmount:   3600,
			ExpectedCurrency: "usd",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "payment_intent_shipping",
				"phase": "create",
			},
			CheckAutomaticPaymentMethods:           true,
			ExpectedAutomaticPaymentMethodsEnabled: true,
			ExpectedReceiptEmail:                   "sdk-codegen+shipping@example.com",
			ExpectedStatementDescriptorSuffix:      "SHIPA",
			ExpectedShipping: &paymentIntentShippingExpectation{
				Name:           "Shipping Create",
				Phone:          "+15555550101",
				Carrier:        "UPS",
				TrackingNumber: "TRACK-CREATE-001",
				AddressLine1:   "100 Market St",
				City:           "San Francisco",
				State:          "CA",
				PostalCode:     "94105",
				Country:        "US",
			},
		}),
		verifyPaymentIntent(paymentIntentExpectations{
			Address:          "stripe_payment_intent.test",
			ExpectedAmount:   3600,
			ExpectedCurrency: "usd",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "payment_intent_shipping",
				"phase": "update",
			},
			CheckAutomaticPaymentMethods:           true,
			ExpectedAutomaticPaymentMethodsEnabled: true,
			ExpectedReceiptEmail:                   "sdk-codegen+shipping-updated@example.com",
			ExpectedStatementDescriptorSuffix:      "SHIPB",
			ExpectedShipping: &paymentIntentShippingExpectation{
				Name:           "Shipping Update",
				Phone:          "+15555550102",
				Carrier:        "FedEx",
				TrackingNumber: "TRACK-UPDATE-002",
				AddressLine1:   "200 Mission St",
				City:           "San Francisco",
				State:          "CA",
				PostalCode:     "94105",
				Country:        "US",
			},
		}),
		verifyPaymentIntentDestroyCanceled,
	)
}

func TestAccManagedPaymentIntentSetupFutureUsage(t *testing.T) {
	runBaseManagedCase(
		t,
		"payment_intent_setup_future_usage",
		"stripe_payment_intent",
		"stripe_payment_intent.test",
		"payment_intent_setup_future_usage_create.tf",
		"payment_intent_setup_future_usage_update.tf",
		true,
		verifyPaymentIntent(paymentIntentExpectations{
			Address:             "stripe_payment_intent.test",
			ExpectedAmount:      4200,
			ExpectedCurrency:    "usd",
			ExpectedDescription: "sdk-codegen setup future usage create",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "payment_intent_setup_future_usage",
				"phase": "create",
			},
			CheckAutomaticPaymentMethods:           true,
			ExpectedAutomaticPaymentMethodsEnabled: true,
			ExpectedSetupFutureUsage:               "off_session",
			ExpectedCustomerAddress:                "stripe_customer.customer",
		}),
		verifyPaymentIntent(paymentIntentExpectations{
			Address:             "stripe_payment_intent.test",
			ExpectedAmount:      4200,
			ExpectedCurrency:    "usd",
			ExpectedDescription: "sdk-codegen setup future usage update",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "payment_intent_setup_future_usage",
				"phase": "update",
			},
			CheckAutomaticPaymentMethods:           true,
			ExpectedAutomaticPaymentMethodsEnabled: true,
			ExpectedSetupFutureUsage:               "on_session",
			ExpectedCustomerAddress:                "stripe_customer.customer",
		}),
		verifyPaymentIntentDestroyCanceled,
	)
}

func TestAccManagedInvoiceItemBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"invoice_item_basic",
		"stripe_invoice_item",
		"stripe_invoice_item.test",
		"invoice_item_basic_create.tf",
		"invoice_item_basic_update.tf",
		true,
		verifyInvoiceItem(invoiceItemExpectations{
			Address:                 "stripe_invoice_item.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedAmount:          1200,
			ExpectedCurrency:        "usd",
			ExpectedDescription:     "sdk-codegen invoice item basic",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "invoice_item_basic",
				"phase": "create",
			},
		}),
		verifyInvoiceItem(invoiceItemExpectations{
			Address:                 "stripe_invoice_item.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedAmount:          1800,
			ExpectedCurrency:        "usd",
			ExpectedDescription:     "sdk-codegen invoice item basic updated",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "invoice_item_basic",
				"phase": "update",
			},
		}),
		verifyInvoiceItemDestroyMissing,
	)
}

func TestAccManagedInvoiceItemPricingPriceRegression(t *testing.T) {
	quantityCreate := int64(3)
	quantityUpdate := int64(5)
	runBaseManagedCase(
		t,
		"invoice_item_pricing_price_regression",
		"stripe_invoice_item",
		"stripe_invoice_item.test",
		"invoice_item_pricing_price_create.tf",
		"invoice_item_pricing_price_update.tf",
		true,
		verifyInvoiceItem(invoiceItemExpectations{
			Address:                 "stripe_invoice_item.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedDescription:     "sdk-codegen invoice item pricing create",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "invoice_item_pricing_price_regression",
				"phase": "create",
			},
			ExpectedQuantity:            &quantityCreate,
			ExpectedPricingPriceAddress: "stripe_price.price",
		}),
		verifyInvoiceItem(invoiceItemExpectations{
			Address:                 "stripe_invoice_item.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedDescription:     "sdk-codegen invoice item pricing update",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "invoice_item_pricing_price_regression",
				"phase": "update",
			},
			ExpectedQuantity:            &quantityUpdate,
			ExpectedPricingPriceAddress: "stripe_price.price",
		}),
		verifyInvoiceItemDestroyMissing,
	)
}

func TestAccManagedInvoiceItemPeriodQuantity(t *testing.T) {
	quantityCreate := int64(1)
	quantityUpdate := int64(2)
	runBaseManagedCase(
		t,
		"invoice_item_period_quantity",
		"stripe_invoice_item",
		"stripe_invoice_item.test",
		"invoice_item_period_quantity_create.tf",
		"invoice_item_period_quantity_update.tf",
		true,
		verifyInvoiceItem(invoiceItemExpectations{
			Address:                 "stripe_invoice_item.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedAmount:          2200,
			ExpectedCurrency:        "usd",
			ExpectedDescription:     "sdk-codegen invoice item period create",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "invoice_item_period_quantity",
				"phase": "create",
			},
			ExpectedQuantity: &quantityCreate,
			ExpectedPeriod: &periodExpectation{
				Start: 1893456000,
				End:   1893542400,
			},
		}),
		verifyInvoiceItem(invoiceItemExpectations{
			Address:                 "stripe_invoice_item.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedAmount:          4400,
			ExpectedCurrency:        "usd",
			ExpectedDescription:     "sdk-codegen invoice item period update",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "invoice_item_period_quantity",
				"phase": "update",
			},
			ExpectedQuantity: &quantityUpdate,
			ExpectedPeriod: &periodExpectation{
				Start: 1893628800,
				End:   1893715200,
			},
		}),
		verifyInvoiceItemDestroyMissing,
	)
}

func TestAccManagedInvoiceItemAttachedInvoice(t *testing.T) {
	discountable := false
	runBaseManagedCase(
		t,
		"invoice_item_attached_invoice",
		"stripe_invoice_item",
		"stripe_invoice_item.test",
		"invoice_item_attached_invoice_create.tf",
		"invoice_item_attached_invoice_update.tf",
		true,
		verifyInvoiceItem(invoiceItemExpectations{
			Address:                 "stripe_invoice_item.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedInvoiceAddress:  "stripe_invoice.invoice",
			ExpectedAmount:          1300,
			ExpectedCurrency:        "usd",
			ExpectedDescription:     "sdk-codegen invoice item attached create",
			ExpectedDiscountable:    &discountable,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "invoice_item_attached_invoice",
				"phase": "create",
			},
		}),
		verifyInvoiceItem(invoiceItemExpectations{
			Address:                 "stripe_invoice_item.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedInvoiceAddress:  "stripe_invoice.invoice",
			ExpectedAmount:          2100,
			ExpectedCurrency:        "usd",
			ExpectedDescription:     "sdk-codegen invoice item attached update",
			ExpectedDiscountable:    &discountable,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "invoice_item_attached_invoice",
				"phase": "update",
			},
		}),
		verifyInvoiceItemDestroyMissing,
	)
}

func TestAccManagedSubscriptionScheduleBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"subscription_schedule_basic",
		"stripe_subscription_schedule",
		"stripe_subscription_schedule.test",
		"subscription_schedule_basic_create.tf",
		"subscription_schedule_basic_update.tf",
		true,
		verifySubscriptionSchedule(subscriptionScheduleExpectations{
			Address:                 "stripe_subscription_schedule.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedEndBehavior:     "release",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "subscription_schedule_basic",
				"phase": "create",
			},
			ExpectedTopLevelStartDate: 1893456000,
			ExpectedFirstPriceAddress: "stripe_price.price",
			ExpectedFirstQuantity:     1,
		}),
		verifySubscriptionSchedule(subscriptionScheduleExpectations{
			Address:                 "stripe_subscription_schedule.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedEndBehavior:     "release",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "subscription_schedule_basic",
				"phase": "update",
			},
			ExpectedTopLevelStartDate: 1893456000,
			ExpectedFirstPriceAddress: "stripe_price.price",
			ExpectedFirstQuantity:     2,
		}),
		verifySubscriptionScheduleDestroyCanceled,
	)
}

func TestAccManagedSubscriptionScheduleDefaultSettings(t *testing.T) {
	runBaseManagedCase(
		t,
		"subscription_schedule_default_settings",
		"stripe_subscription_schedule",
		"stripe_subscription_schedule.test",
		"subscription_schedule_default_settings_create.tf",
		"subscription_schedule_default_settings_update.tf",
		true,
		verifySubscriptionSchedule(subscriptionScheduleExpectations{
			Address:                                 "stripe_subscription_schedule.test",
			ExpectedCustomerAddress:                 "stripe_customer.customer",
			ExpectedEndBehavior:                     "release",
			ExpectedDefaultSettingsCollectionMethod: "send_invoice",
			ExpectedDefaultSettingsDescription:      "sdk-codegen schedule default settings create",
			ExpectedDefaultSettingsDaysUntilDue:     14,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "subscription_schedule_default_settings",
				"phase": "create",
			},
			ExpectedTopLevelStartDate: 1894060800,
			ExpectedFirstPriceAddress: "stripe_price.price",
			ExpectedFirstQuantity:     1,
		}),
		verifySubscriptionSchedule(subscriptionScheduleExpectations{
			Address:                                 "stripe_subscription_schedule.test",
			ExpectedCustomerAddress:                 "stripe_customer.customer",
			ExpectedEndBehavior:                     "release",
			ExpectedDefaultSettingsCollectionMethod: "send_invoice",
			ExpectedDefaultSettingsDescription:      "sdk-codegen schedule default settings update",
			ExpectedDefaultSettingsDaysUntilDue:     30,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "subscription_schedule_default_settings",
				"phase": "update",
			},
			ExpectedTopLevelStartDate: 1894060800,
			ExpectedFirstPriceAddress: "stripe_price.price",
			ExpectedFirstQuantity:     2,
		}),
		verifySubscriptionScheduleDestroyCanceled,
	)
}

func TestAccManagedSubscriptionSchedulePhaseQuantity(t *testing.T) {
	runBaseManagedCaseWithImportIgnore(
		t,
		"subscription_schedule_phase_quantity",
		"stripe_subscription_schedule",
		"stripe_subscription_schedule.test",
		"subscription_schedule_phase_quantity_create.tf",
		"subscription_schedule_phase_quantity_update.tf",
		true,
		[]string{
			"phases.0.duration.%",
			"phases.0.duration.interval",
			"phases.0.duration.interval_count",
		},
		verifySubscriptionSchedule(subscriptionScheduleExpectations{
			Address:                 "stripe_subscription_schedule.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedEndBehavior:     "cancel",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "subscription_schedule_phase_quantity",
				"phase": "create",
			},
			ExpectedTopLevelStartDate: 1893628800,
			ExpectedFirstPriceAddress: "stripe_price.price",
			ExpectedFirstQuantity:     1,
		}),
		verifySubscriptionSchedule(subscriptionScheduleExpectations{
			Address:                 "stripe_subscription_schedule.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedEndBehavior:     "cancel",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "subscription_schedule_phase_quantity",
				"phase": "update",
			},
			ExpectedTopLevelStartDate: 1893628800,
			ExpectedFirstPriceAddress: "stripe_price.price",
			ExpectedFirstQuantity:     3,
		}),
		verifySubscriptionScheduleDestroyCanceled,
	)
}

func TestAccManagedSubscriptionScheduleStartDateRegression(t *testing.T) {
	runBaseManagedCase(
		t,
		"subscription_schedule_start_date_regression",
		"stripe_subscription_schedule",
		"stripe_subscription_schedule.test",
		"subscription_schedule_start_date_regression_create.tf",
		"",
		true,
		verifySubscriptionSchedule(subscriptionScheduleExpectations{
			Address:                 "stripe_subscription_schedule.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedEndBehavior:     "release",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "subscription_schedule_start_date_regression",
			},
			ExpectedTopLevelStartDate: 1893888000,
			ExpectedFirstPriceAddress: "stripe_price.price",
			ExpectedFirstQuantity:     1,
		}),
		nil,
		verifySubscriptionScheduleDestroyCanceled,
	)
}

func TestAccManagedTerminalReaderBasic(t *testing.T) {
	runManagedCaseForGroup(
		t,
		"terminal",
		[]string{
			"STRIPE_TERMINAL_READER_REGISTRATION_CODE",
			"STRIPE_TERMINAL_LOCATION",
		},
		"terminal_reader_basic",
		"stripe_terminal_reader",
		"stripe_terminal_reader.test",
		"terminal_reader_basic_create.tf",
		"terminal_reader_basic_update.tf",
		true,
		verifyTerminalReader(terminalReaderExpectations{
			Address:       "stripe_terminal_reader.test",
			ExpectedLabel: "SDK Codegen Reader Basic",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "terminal_reader_basic",
				"phase": "create",
			},
			ExpectedLocationFromEnv: true,
		}),
		verifyTerminalReader(terminalReaderExpectations{
			Address:       "stripe_terminal_reader.test",
			ExpectedLabel: "SDK Codegen Reader Basic Updated",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "terminal_reader_basic",
				"phase": "update",
			},
			ExpectedLocationFromEnv: true,
		}),
		verifyTerminalReaderDestroyMissing,
	)
}

func TestAccManagedTerminalReaderLabelUpdate(t *testing.T) {
	runManagedCaseForGroup(
		t,
		"terminal",
		[]string{
			"STRIPE_TERMINAL_READER_REGISTRATION_CODE_UPDATE",
			"STRIPE_TERMINAL_LOCATION",
		},
		"terminal_reader_label_update",
		"stripe_terminal_reader",
		"stripe_terminal_reader.test",
		"terminal_reader_label_update_create.tf",
		"terminal_reader_label_update_update.tf",
		true,
		verifyTerminalReader(terminalReaderExpectations{
			Address:       "stripe_terminal_reader.test",
			ExpectedLabel: "SDK Codegen Reader Update",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "terminal_reader_label_update",
				"phase": "create",
			},
			ExpectedLocationFromEnv: true,
		}),
		verifyTerminalReader(terminalReaderExpectations{
			Address:       "stripe_terminal_reader.test",
			ExpectedLabel: "SDK Codegen Reader Update Final",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "terminal_reader_label_update",
				"phase": "update",
			},
			ExpectedLocationFromEnv: true,
		}),
		verifyTerminalReaderDestroyMissing,
	)
}

func TestAccManagedTerminalReaderRegistrationCodeRegression(t *testing.T) {
	runManagedCaseForGroup(
		t,
		"terminal",
		[]string{
			"STRIPE_TERMINAL_READER_REGISTRATION_CODE_REGRESSION",
			"STRIPE_TERMINAL_LOCATION",
		},
		"terminal_reader_registration_code_regression",
		"stripe_terminal_reader",
		"stripe_terminal_reader.test",
		"terminal_reader_registration_code_regression_create.tf",
		"",
		true,
		verifyTerminalReader(terminalReaderExpectations{
			Address:       "stripe_terminal_reader.test",
			ExpectedLabel: "SDK Codegen Reader Regression",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "terminal_reader_registration_code_regression",
			},
		}),
		nil,
		verifyTerminalReaderDestroyMissing,
	)
}

func TestAccManagedTerminalReaderMetadataUpdate(t *testing.T) {
	runManagedCaseForGroup(
		t,
		"terminal",
		[]string{
			"STRIPE_TERMINAL_READER_REGISTRATION_CODE",
			"STRIPE_TERMINAL_LOCATION",
		},
		"terminal_reader_metadata_update",
		"stripe_terminal_reader",
		"stripe_terminal_reader.test",
		"terminal_reader_metadata_update_create.tf",
		"terminal_reader_metadata_update_update.tf",
		true,
		verifyTerminalReader(terminalReaderExpectations{
			Address:       "stripe_terminal_reader.test",
			ExpectedLabel: "SDK Codegen Reader Metadata",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "terminal_reader_metadata_update",
				"phase": "create",
			},
			ExpectedLocationFromEnv: true,
			CheckDeviceDetails:      true,
		}),
		verifyTerminalReader(terminalReaderExpectations{
			Address:       "stripe_terminal_reader.test",
			ExpectedLabel: "SDK Codegen Reader Metadata Updated",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "terminal_reader_metadata_update",
				"phase": "update",
			},
			ExpectedLocationFromEnv: true,
			CheckDeviceDetails:      true,
		}),
		verifyTerminalReaderDestroyMissing,
	)
}
