// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import (
	"testing"

	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

const legacyUpgradeProviderVersion = "0.2.2"

func runBaseManagedLegacyUpgradeCase(
	t *testing.T,
	name string,
	surface string,
	resourceAddress string,
	createTemplate string,
	verifyCreate runner.StateVerifier,
	verifyDestroy runner.StateVerifier,
) {
	t.Helper()

	runner.RunManagedLegacyUpgradeCase(t, runner.ManagedCase{
		Definition: runner.CaseDefinition{
			Name:    name,
			Surface: surface,
			Group:   "base",
			Kind:    "resource",
		},
		ResourceAddress: resourceAddress,
		CreateTemplate:  createTemplate,
		VerifyCreate:    verifyCreate,
		VerifyDestroy:   verifyDestroy,
	}, legacyUpgradeProviderVersion)
}

func TestAccManagedSubscriptionBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"subscription_basic",
		"stripe_subscription",
		"stripe_subscription.test",
		"subscription_basic_create.tf",
		"subscription_basic_update.tf",
		false,
		verifySubscription(subscriptionExpectations{
			Address:                 "stripe_subscription.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedDescription:     "sdk-codegen subscription basic",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "subscription_basic",
				"phase": "create",
			},
			ExpectedCollectionMethod: "send_invoice",
			ExpectedStatus:           "active",
			ExpectedPriceAddress:     "stripe_price.price",
			ExpectedQuantity:         1,
		}),
		verifySubscription(subscriptionExpectations{
			Address:                 "stripe_subscription.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedDescription:     "sdk-codegen subscription basic updated",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "subscription_basic",
				"phase": "update",
			},
			ExpectedCollectionMethod: "send_invoice",
			ExpectedStatus:           "active",
			ExpectedPriceAddress:     "stripe_price.price",
			ExpectedQuantity:         2,
		}),
		verifySubscriptionDestroyStillExists,
	)
}

func TestAccManagedSubscriptionTrialSettings(t *testing.T) {
	runBaseManagedCase(
		t,
		"subscription_trial_settings",
		"stripe_subscription",
		"stripe_subscription.test",
		"subscription_trial_settings_create.tf",
		"subscription_trial_settings_update.tf",
		false,
		verifySubscription(subscriptionExpectations{
			Address:                 "stripe_subscription.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedDescription:     "sdk-codegen subscription trial create",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "subscription_trial_settings",
				"phase": "create",
			},
			ExpectedCollectionMethod:          "charge_automatically",
			ExpectedStatus:                    "trialing",
			ExpectedTrialEndAfterStart:        true,
			ExpectedTrialMissingPaymentMethod: "pause",
			ExpectedPriceAddress:              "stripe_price.price",
			ExpectedQuantity:                  1,
		}),
		verifySubscription(subscriptionExpectations{
			Address:                 "stripe_subscription.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedDescription:     "sdk-codegen subscription trial update",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "subscription_trial_settings",
				"phase": "update",
			},
			ExpectedCollectionMethod:          "charge_automatically",
			ExpectedStatus:                    "trialing",
			ExpectedTrialEndAfterStart:        true,
			ExpectedTrialMissingPaymentMethod: "pause",
			ExpectedPriceAddress:              "stripe_price.price",
			ExpectedQuantity:                  2,
		}),
		verifySubscriptionDestroyStillExists,
	)
}

func TestAccManagedSubscriptionItemBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"subscription_item_basic",
		"stripe_subscription_item",
		"stripe_subscription_item.test",
		"subscription_item_basic_create.tf",
		"subscription_item_basic_update.tf",
		false,
		verifySubscriptionItem(subscriptionItemExpectations{
			Address:                     "stripe_subscription_item.test",
			ExpectedSubscriptionAddress: "stripe_subscription.subscription",
			ExpectedPriceAddress:        "stripe_price.addon",
			ExpectedQuantity:            2,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "subscription_item_basic",
				"phase": "create",
			},
		}),
		verifySubscriptionItem(subscriptionItemExpectations{
			Address:                     "stripe_subscription_item.test",
			ExpectedSubscriptionAddress: "stripe_subscription.subscription",
			ExpectedPriceAddress:        "stripe_price.addon",
			ExpectedQuantity:            4,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "subscription_item_basic",
				"phase": "update",
			},
		}),
		verifySubscriptionItemDestroyMissing,
	)
}

func TestAccManagedSubscriptionItemTaxRates(t *testing.T) {
	runBaseManagedCase(
		t,
		"subscription_item_tax_rates",
		"stripe_subscription_item",
		"stripe_subscription_item.test",
		"subscription_item_tax_rates_create.tf",
		"subscription_item_tax_rates_update.tf",
		false,
		verifySubscriptionItem(subscriptionItemExpectations{
			Address:                     "stripe_subscription_item.test",
			ExpectedSubscriptionAddress: "stripe_subscription.subscription",
			ExpectedPriceAddress:        "stripe_price.addon",
			ExpectedQuantity:            2,
			ExpectedTaxRateAddresses:    []string{"stripe_tax_rate.create"},
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "subscription_item_tax_rates",
				"phase": "create",
			},
		}),
		verifySubscriptionItem(subscriptionItemExpectations{
			Address:                     "stripe_subscription_item.test",
			ExpectedSubscriptionAddress: "stripe_subscription.subscription",
			ExpectedPriceAddress:        "stripe_price.addon",
			ExpectedQuantity:            4,
			ExpectedTaxRateAddresses:    []string{"stripe_tax_rate.update"},
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "subscription_item_tax_rates",
				"phase": "update",
			},
		}),
		verifySubscriptionItemDestroyMissing,
	)
}

func TestAccManagedInvoiceBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"invoice_basic",
		"stripe_invoice",
		"stripe_invoice.test",
		"invoice_basic_create.tf",
		"invoice_basic_update.tf",
		true,
		verifyInvoice(invoiceExpectations{
			Address:                 "stripe_invoice.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedDescription:     "sdk-codegen invoice basic",
			ExpectedStatus:          "draft",
			ExpectedAutoAdvance:     false,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "invoice_basic",
				"phase": "create",
			},
		}),
		verifyInvoice(invoiceExpectations{
			Address:                 "stripe_invoice.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedDescription:     "sdk-codegen invoice basic updated",
			ExpectedStatus:          "draft",
			ExpectedAutoAdvance:     false,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "invoice_basic",
				"phase": "update",
			},
		}),
		verifyInvoiceDestroyMissing,
	)
}

func TestAccManagedInvoiceSendInvoiceFields(t *testing.T) {
	runBaseManagedCase(
		t,
		"invoice_send_invoice_fields",
		"stripe_invoice",
		"stripe_invoice.test",
		"invoice_send_invoice_fields_create.tf",
		"invoice_send_invoice_fields_update.tf",
		true,
		verifyInvoice(invoiceExpectations{
			Address:                  "stripe_invoice.test",
			ExpectedCustomerAddress:  "stripe_customer.customer",
			ExpectedDescription:      "sdk-codegen invoice send invoice create",
			ExpectedStatus:           "draft",
			ExpectedAutoAdvance:      false,
			ExpectedCollectionMethod: "send_invoice",
			ExpectedDueDate:          1893456000,
			ExpectedFooter:           "SDK Codegen invoice footer create",
			ExpectedCustomFields: []invoiceCustomFieldExpectation{
				{Name: "order_id", Value: "INV-CREATE"},
				{Name: "region", Value: "NA"},
			},
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "invoice_send_invoice_fields",
				"phase": "create",
			},
		}),
		verifyInvoice(invoiceExpectations{
			Address:                  "stripe_invoice.test",
			ExpectedCustomerAddress:  "stripe_customer.customer",
			ExpectedDescription:      "sdk-codegen invoice send invoice update",
			ExpectedStatus:           "draft",
			ExpectedAutoAdvance:      false,
			ExpectedCollectionMethod: "send_invoice",
			ExpectedDueDate:          1896134400,
			ExpectedFooter:           "SDK Codegen invoice footer update",
			ExpectedCustomFields: []invoiceCustomFieldExpectation{
				{Name: "order_id", Value: "INV-UPDATE"},
				{Name: "region", Value: "EU"},
			},
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "invoice_send_invoice_fields",
				"phase": "update",
			},
		}),
		verifyInvoiceDestroyMissing,
	)
}

func TestAccManagedSetupIntentBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"setup_intent_basic",
		"stripe_setup_intent",
		"stripe_setup_intent.test",
		"setup_intent_basic_create.tf",
		"setup_intent_basic_update.tf",
		true,
		verifySetupIntent(setupIntentExpectations{
			Address:                    "stripe_setup_intent.test",
			ExpectedCustomerAddress:    "stripe_customer.customer",
			ExpectedDescription:        "sdk-codegen setup intent basic",
			ExpectedStatus:             "requires_payment_method",
			ExpectedUsage:              "off_session",
			ExpectedPaymentMethodTypes: []string{"card"},
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "setup_intent_basic",
				"phase": "create",
			},
		}),
		verifySetupIntent(setupIntentExpectations{
			Address:                    "stripe_setup_intent.test",
			ExpectedCustomerAddress:    "stripe_customer.customer",
			ExpectedDescription:        "sdk-codegen setup intent basic updated",
			ExpectedStatus:             "requires_payment_method",
			ExpectedUsage:              "on_session",
			ExpectedPaymentMethodTypes: []string{"card"},
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "setup_intent_basic",
				"phase": "update",
			},
		}),
		verifySetupIntentDestroyStillExists,
	)
}

func TestAccManagedSetupIntentAutomaticPaymentMethods(t *testing.T) {
	runBaseManagedCase(
		t,
		"setup_intent_automatic_payment_methods",
		"stripe_setup_intent",
		"stripe_setup_intent.test",
		"setup_intent_automatic_payment_methods_create.tf",
		"",
		false,
		verifySetupIntent(setupIntentExpectations{
			Address:                                "stripe_setup_intent.test",
			ExpectedDescription:                    "sdk-codegen setup intent automatic payment methods",
			ExpectedStatus:                         "requires_payment_method",
			ExpectedUsage:                          "off_session",
			ExpectedAutomaticPaymentMethodsEnabled: true,
			ExpectedAutomaticPaymentMethodsAllowRedirects: "never",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "setup_intent_automatic_payment_methods",
			},
		}),
		nil,
		verifySetupIntentDestroyStillExists,
	)
}

func TestAccManagedPaymentMethodBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"payment_method_basic",
		"stripe_payment_method",
		"stripe_payment_method.test",
		"payment_method_basic_create.tf",
		"",
		false,
		verifyPaymentMethod(paymentMethodExpectations{
			Address:              "stripe_payment_method.test",
			ExpectedType:         "card",
			ExpectedLast4:        "4242",
			ExpectedExpMonth:     12,
			ExpectedExpYear:      2035,
			ExpectedBillingName:  "SDK Codegen Card Create",
			ExpectedBillingEmail: "sdk-codegen+pm-create@example.com",
			ExpectedBillingPhone: "+15555550121",
			ExpectedAddressLine1: "100 Market St",
			ExpectedCity:         "San Francisco",
			ExpectedState:        "CA",
			ExpectedPostalCode:   "94105",
			ExpectedCountry:      "US",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "payment_method_basic",
				"phase": "create",
			},
		}),
		nil,
		verifyPaymentMethodDestroyStillExists,
	)
}

func TestAccManagedPaymentMethodAllowRedisplay(t *testing.T) {
	runBaseManagedCase(
		t,
		"payment_method_allow_redisplay",
		"stripe_payment_method",
		"stripe_payment_method.test",
		"payment_method_allow_redisplay_create.tf",
		"",
		false,
		verifyPaymentMethod(paymentMethodExpectations{
			Address:                "stripe_payment_method.test",
			ExpectedType:           "card",
			ExpectedAllowRedisplay: "always",
			ExpectedLast4:          "4242",
			ExpectedExpMonth:       11,
			ExpectedExpYear:        2036,
			ExpectedBillingName:    "SDK Codegen Card Redisplay",
			ExpectedBillingEmail:   "sdk-codegen+pm-redisplay@example.com",
			ExpectedBillingPhone:   "+15555550123",
			ExpectedAddressLine1:   "300 Howard St",
			ExpectedCity:           "San Francisco",
			ExpectedState:          "CA",
			ExpectedPostalCode:     "94105",
			ExpectedCountry:        "US",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "payment_method_allow_redisplay",
			},
		}),
		nil,
		verifyPaymentMethodDestroyStillExists,
	)
}

func TestAccManagedPaymentLinkBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"payment_link_basic",
		"stripe_payment_link",
		"stripe_payment_link.test",
		"payment_link_basic_create.tf",
		"payment_link_basic_update.tf",
		true,
		verifyPaymentLink(paymentLinkExpectations{
			Address:                  "stripe_payment_link.test",
			ExpectedActive:           true,
			ExpectedCustomerCreation: "always",
			ExpectedInactiveMessage:  "sdk-codegen payment link create message",
			ExpectedRedirectURL:      "https://example.com/sdk-codegen/payment-link/create",
			ExpectedPriceAddress:     "stripe_price.price",
			ExpectedQuantity:         1,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "payment_link_basic",
				"phase": "create",
			},
		}),
		verifyPaymentLink(paymentLinkExpectations{
			Address:                  "stripe_payment_link.test",
			ExpectedActive:           true,
			ExpectedCustomerCreation: "always",
			ExpectedInactiveMessage:  "sdk-codegen payment link update message",
			ExpectedRedirectURL:      "https://example.com/sdk-codegen/payment-link/update",
			ExpectedPriceAddress:     "stripe_price.price",
			ExpectedQuantity:         2,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "payment_link_basic",
				"phase": "update",
			},
		}),
		verifyPaymentLinkDestroyInactive,
	)
}

func TestAccManagedPaymentLinkCheckoutControls(t *testing.T) {
	runBaseManagedCase(
		t,
		"payment_link_checkout_controls",
		"stripe_payment_link",
		"stripe_payment_link.test",
		"payment_link_checkout_controls_create.tf",
		"",
		false,
		verifyPaymentLink(paymentLinkExpectations{
			Address:                          "stripe_payment_link.test",
			ExpectedActive:                   true,
			ExpectedInactiveMessage:          "sdk-codegen payment link controls create",
			ExpectedRedirectURL:              "https://example.com/sdk-codegen/payment-link/controls/create",
			ExpectedPriceAddress:             "stripe_price.price",
			ExpectedQuantity:                 1,
			ExpectedAllowPromotionCodes:      true,
			ExpectedBillingAddressCollection: "required",
			ExpectedPaymentMethodCollection:  "if_required",
			ExpectedSubmitType:               "donate",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "payment_link_checkout_controls",
				"phase": "create",
			},
		}),
		nil,
		verifyPaymentLinkDestroyInactive,
	)
}

func TestAccManagedQuoteBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"quote_basic",
		"stripe_quote",
		"stripe_quote.test",
		"quote_basic_create.tf",
		"quote_basic_update.tf",
		true,
		verifyQuote(quoteExpectations{
			Address:                 "stripe_quote.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedDescription:     "sdk-codegen quote basic",
			ExpectedHeader:          "SDK Codegen Quote Create",
			ExpectedStatus:          "draft",
			ExpectedPriceAddress:    "stripe_price.price",
			ExpectedQuantity:        1,
			ExpectedAmountSubtotal:  2200,
			ExpectedAmountTotal:     2200,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "quote_basic",
				"phase": "create",
			},
		}),
		verifyQuote(quoteExpectations{
			Address:                 "stripe_quote.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedDescription:     "sdk-codegen quote basic updated",
			ExpectedHeader:          "SDK Codegen Quote Update",
			ExpectedStatus:          "draft",
			ExpectedPriceAddress:    "stripe_price.price",
			ExpectedQuantity:        2,
			ExpectedAmountSubtotal:  4400,
			ExpectedAmountTotal:     4400,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "quote_basic",
				"phase": "update",
			},
		}),
		verifyQuoteDestroyStillExists,
	)
}

func TestAccManagedQuoteSendInvoiceFields(t *testing.T) {
	runBaseManagedCase(
		t,
		"quote_send_invoice_fields",
		"stripe_quote",
		"stripe_quote.test",
		"quote_send_invoice_fields_create.tf",
		"quote_send_invoice_fields_update.tf",
		true,
		verifyQuote(quoteExpectations{
			Address:                             "stripe_quote.test",
			ExpectedCustomerAddress:             "stripe_customer.customer",
			ExpectedDescription:                 "sdk-codegen quote send invoice create",
			ExpectedHeader:                      "SDK Codegen Quote Send Invoice Create",
			ExpectedFooter:                      "SDK Codegen quote footer create",
			ExpectedStatus:                      "draft",
			ExpectedCollectionMethod:            "send_invoice",
			ExpectedExpiresAt:                   1893456000,
			ExpectedInvoiceSettingsDaysUntilDue: 14,
			ExpectedPriceAddress:                "stripe_price.price",
			ExpectedQuantity:                    1,
			ExpectedAmountSubtotal:              2200,
			ExpectedAmountTotal:                 2200,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "quote_send_invoice_fields",
				"phase": "create",
			},
		}),
		verifyQuote(quoteExpectations{
			Address:                             "stripe_quote.test",
			ExpectedCustomerAddress:             "stripe_customer.customer",
			ExpectedDescription:                 "sdk-codegen quote send invoice update",
			ExpectedHeader:                      "SDK Codegen Quote Send Invoice Update",
			ExpectedFooter:                      "SDK Codegen quote footer update",
			ExpectedStatus:                      "draft",
			ExpectedCollectionMethod:            "send_invoice",
			ExpectedExpiresAt:                   1896134400,
			ExpectedInvoiceSettingsDaysUntilDue: 30,
			ExpectedPriceAddress:                "stripe_price.price",
			ExpectedQuantity:                    2,
			ExpectedAmountSubtotal:              4400,
			ExpectedAmountTotal:                 4400,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "quote_send_invoice_fields",
				"phase": "update",
			},
		}),
		verifyQuoteDestroyStillExists,
	)
}

func TestAccManagedChargeBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"charge_basic",
		"stripe_charge",
		"stripe_charge.test",
		"charge_basic_create.tf",
		"charge_basic_update.tf",
		false,
		verifyCharge(chargeExpectations{
			Address:              "stripe_charge.test",
			ExpectedAmount:       2200,
			ExpectedCurrency:     "usd",
			ExpectedDescription:  "sdk-codegen charge basic",
			ExpectedReceiptEmail: "sdk-codegen+charge-create@example.com",
			ExpectedStatus:       "succeeded",
			ExpectedPaid:         true,
			ExpectedCaptured:     true,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "charge_basic",
				"phase": "create",
			},
		}),
		verifyCharge(chargeExpectations{
			Address:              "stripe_charge.test",
			ExpectedAmount:       2200,
			ExpectedCurrency:     "usd",
			ExpectedDescription:  "sdk-codegen charge basic updated",
			ExpectedReceiptEmail: "sdk-codegen+charge-update@example.com",
			ExpectedStatus:       "succeeded",
			ExpectedPaid:         true,
			ExpectedCaptured:     true,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "charge_basic",
				"phase": "update",
			},
		}),
		verifyChargeDestroyStillExists,
	)
}

func TestAccManagedChargeCustomerShipping(t *testing.T) {
	runBaseManagedCase(
		t,
		"charge_customer_shipping",
		"stripe_charge",
		"stripe_charge.test",
		"charge_customer_shipping_create.tf",
		"charge_customer_shipping_update.tf",
		false,
		verifyCharge(chargeExpectations{
			Address:                 "stripe_charge.test",
			ExpectedAmount:          3400,
			ExpectedCurrency:        "usd",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedDescription:     "sdk-codegen charge customer shipping create",
			ExpectedReceiptEmail:    "sdk-codegen+charge-shipping-create@example.com",
			ExpectedStatus:          "succeeded",
			ExpectedPaid:            true,
			ExpectedCaptured:        true,
			ExpectedShippingName:    "SDK Codegen Charge Create",
			ExpectedShippingPhone:   "+15555550161",
			ExpectedAddressLine1:    "100 Market St",
			ExpectedCity:            "San Francisco",
			ExpectedState:           "CA",
			ExpectedPostalCode:      "94105",
			ExpectedCountry:         "US",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "charge_customer_shipping",
				"phase": "create",
			},
		}),
		verifyCharge(chargeExpectations{
			Address:                 "stripe_charge.test",
			ExpectedAmount:          3400,
			ExpectedCurrency:        "usd",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedDescription:     "sdk-codegen charge customer shipping update",
			ExpectedReceiptEmail:    "sdk-codegen+charge-shipping-update@example.com",
			ExpectedStatus:          "succeeded",
			ExpectedPaid:            true,
			ExpectedCaptured:        true,
			ExpectedShippingName:    "SDK Codegen Charge Update",
			ExpectedShippingPhone:   "+15555550162",
			ExpectedAddressLine1:    "200 Mission St",
			ExpectedCity:            "San Francisco",
			ExpectedState:           "CA",
			ExpectedPostalCode:      "94105",
			ExpectedCountry:         "US",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "charge_customer_shipping",
				"phase": "update",
			},
		}),
		verifyChargeDestroyStillExists,
	)
}

func TestAccManagedSourceBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"source_basic",
		"stripe_source",
		"stripe_source.test",
		"source_basic_create.tf",
		"source_basic_update.tf",
		false,
		verifySource(sourceExpectations{
			Address:            "stripe_source.test",
			ExpectedType:       "card",
			ExpectedStatus:     "chargeable",
			ExpectedOwnerName:  "SDK Codegen Source Create",
			ExpectedOwnerEmail: "sdk-codegen+source-create@example.com",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "source_basic",
				"phase": "create",
			},
		}),
		verifySource(sourceExpectations{
			Address:            "stripe_source.test",
			ExpectedType:       "card",
			ExpectedStatus:     "chargeable",
			ExpectedOwnerName:  "SDK Codegen Source Update",
			ExpectedOwnerEmail: "sdk-codegen+source-update@example.com",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "source_basic",
				"phase": "update",
			},
		}),
		verifySourceDestroyStillExists,
	)
}

func TestAccManagedShippingRateBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"shipping_rate_basic",
		"stripe_shipping_rate",
		"stripe_shipping_rate.test",
		"shipping_rate_basic_create.tf",
		"shipping_rate_basic_update.tf",
		true,
		verifyShippingRate(shippingRateExpectations{
			Address:                     "stripe_shipping_rate.test",
			ExpectedDisplayName:         "SDK Codegen Shipping Create",
			ExpectedActive:              true,
			ExpectedFixedAmount:         500,
			ExpectedFixedAmountCurrency: "usd",
			ExpectedMinimumUnit:         "business_day",
			ExpectedMinimumValue:        1,
			ExpectedMaximumUnit:         "business_day",
			ExpectedMaximumValue:        3,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "shipping_rate_basic",
				"phase": "create",
			},
		}),
		verifyShippingRate(shippingRateExpectations{
			Address:                     "stripe_shipping_rate.test",
			ExpectedDisplayName:         "SDK Codegen Shipping Update",
			ExpectedActive:              false,
			ExpectedFixedAmount:         500,
			ExpectedFixedAmountCurrency: "usd",
			ExpectedMinimumUnit:         "business_day",
			ExpectedMinimumValue:        1,
			ExpectedMaximumUnit:         "business_day",
			ExpectedMaximumValue:        3,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "shipping_rate_basic",
				"phase": "update",
			},
		}),
		verifyShippingRateDestroyInactive,
	)
}

func TestAccManagedShippingRateTaxBehavior(t *testing.T) {
	runBaseManagedCase(
		t,
		"shipping_rate_tax_behavior",
		"stripe_shipping_rate",
		"stripe_shipping_rate.test",
		"shipping_rate_tax_behavior_create.tf",
		"shipping_rate_tax_behavior_update.tf",
		true,
		verifyShippingRate(shippingRateExpectations{
			Address:                     "stripe_shipping_rate.test",
			ExpectedDisplayName:         "SDK Codegen Shipping Tax Create",
			ExpectedActive:              true,
			ExpectedFixedAmount:         700,
			ExpectedFixedAmountCurrency: "usd",
			ExpectedMinimumUnit:         "day",
			ExpectedMinimumValue:        2,
			ExpectedMaximumUnit:         "day",
			ExpectedMaximumValue:        5,
			ExpectedTaxBehavior:         "exclusive",
			ExpectedTaxCode:             "txcd_92010001",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "shipping_rate_tax_behavior",
				"phase": "create",
			},
		}),
		verifyShippingRate(shippingRateExpectations{
			Address:                     "stripe_shipping_rate.test",
			ExpectedDisplayName:         "SDK Codegen Shipping Tax Update",
			ExpectedActive:              true,
			ExpectedFixedAmount:         900,
			ExpectedFixedAmountCurrency: "usd",
			ExpectedMinimumUnit:         "business_day",
			ExpectedMinimumValue:        1,
			ExpectedMaximumUnit:         "business_day",
			ExpectedMaximumValue:        4,
			ExpectedTaxBehavior:         "inclusive",
			ExpectedTaxCode:             "txcd_92010001",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "shipping_rate_tax_behavior",
				"phase": "update",
			},
		}),
		verifyShippingRateDestroyInactive,
	)
}

func TestAccManagedShippingRateLegacyUpgrade(t *testing.T) {
	// Legacy provider v0.2.2 does not accept the current schema shape for this case.
	// Use a legacy-only create fixture so the upgrade path exercises state migration.
	runBaseManagedLegacyUpgradeCase(
		t,
		"shipping_rate_legacy_upgrade",
		"stripe_shipping_rate",
		"stripe_shipping_rate.test",
		"shipping_rate_legacy_upgrade_create.tf",
		verifyShippingRate(shippingRateExpectations{
			Address:                     "stripe_shipping_rate.test",
			ExpectedDisplayName:         "SDK Codegen Shipping Create",
			ExpectedActive:              true,
			ExpectedFixedAmount:         500,
			ExpectedFixedAmountCurrency: "usd",
			ExpectedMinimumUnit:         "business_day",
			ExpectedMinimumValue:        1,
			ExpectedMaximumUnit:         "business_day",
			ExpectedMaximumValue:        3,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "shipping_rate_basic",
				"phase": "create",
			},
		}),
		verifyShippingRateDestroyInactive,
	)
}

func TestAccManagedTaxRateBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"tax_rate_basic",
		"stripe_tax_rate",
		"stripe_tax_rate.test",
		"tax_rate_basic_create.tf",
		"tax_rate_basic_update.tf",
		true,
		verifyTaxRate(taxRateExpectations{
			Address:              "stripe_tax_rate.test",
			ExpectedDisplayName:  "SDK Codegen Tax Rate",
			ExpectedDescription:  "sdk-codegen tax rate create",
			ExpectedCountry:      "US",
			ExpectedState:        "CA",
			ExpectedJurisdiction: "California",
			ExpectedInclusive:    false,
			ExpectedPercentage:   8.25,
			ExpectedTaxType:      "",
			ExpectedActive:       true,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "tax_rate_basic",
				"phase": "create",
			},
		}),
		verifyTaxRate(taxRateExpectations{
			Address:              "stripe_tax_rate.test",
			ExpectedDisplayName:  "SDK Codegen Tax Rate",
			ExpectedDescription:  "sdk-codegen tax rate update",
			ExpectedCountry:      "US",
			ExpectedState:        "CA",
			ExpectedJurisdiction: "California",
			ExpectedInclusive:    false,
			ExpectedPercentage:   8.25,
			ExpectedTaxType:      "",
			ExpectedActive:       false,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "tax_rate_basic",
				"phase": "update",
			},
		}),
		verifyTaxRateDestroyInactive,
	)
}

func TestAccManagedTaxRateLegacyUpgrade(t *testing.T) {
	// Legacy provider v0.2.2 does not accept the current schema shape for this case.
	// Use a legacy-only create fixture so the upgrade path exercises state migration.
	runBaseManagedLegacyUpgradeCase(
		t,
		"tax_rate_legacy_upgrade",
		"stripe_tax_rate",
		"stripe_tax_rate.test",
		"tax_rate_legacy_upgrade_create.tf",
		verifyTaxRate(taxRateExpectations{
			Address:              "stripe_tax_rate.test",
			ExpectedDisplayName:  "SDK Codegen Tax Rate",
			ExpectedDescription:  "sdk-codegen tax rate create",
			ExpectedCountry:      "US",
			ExpectedState:        "CA",
			ExpectedJurisdiction: "California",
			ExpectedInclusive:    false,
			ExpectedPercentage:   8.25,
			ExpectedTaxType:      "",
			ExpectedActive:       true,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "tax_rate_basic",
				"phase": "create",
			},
		}),
		verifyTaxRateDestroyInactive,
	)
}

func TestAccManagedTaxRateTaxType(t *testing.T) {
	runBaseManagedCase(
		t,
		"tax_rate_tax_type",
		"stripe_tax_rate",
		"stripe_tax_rate.test",
		"tax_rate_tax_type_create.tf",
		"tax_rate_tax_type_update.tf",
		true,
		verifyTaxRate(taxRateExpectations{
			Address:              "stripe_tax_rate.test",
			ExpectedDisplayName:  "SDK Codegen Sales Tax",
			ExpectedDescription:  "sdk-codegen sales tax create",
			ExpectedCountry:      "US",
			ExpectedState:        "NY",
			ExpectedJurisdiction: "New York",
			ExpectedInclusive:    true,
			ExpectedPercentage:   4.5,
			ExpectedTaxType:      "sales_tax",
			ExpectedActive:       true,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "tax_rate_tax_type",
				"phase": "create",
			},
		}),
		verifyTaxRate(taxRateExpectations{
			Address:              "stripe_tax_rate.test",
			ExpectedDisplayName:  "SDK Codegen Sales Tax",
			ExpectedDescription:  "sdk-codegen sales tax update",
			ExpectedCountry:      "US",
			ExpectedState:        "NY",
			ExpectedJurisdiction: "New York",
			ExpectedInclusive:    true,
			ExpectedPercentage:   4.5,
			ExpectedTaxType:      "sales_tax",
			ExpectedActive:       false,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "tax_rate_tax_type",
				"phase": "update",
			},
		}),
		verifyTaxRateDestroyInactive,
	)
}

func TestAccManagedTaxRegistrationBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"tax_registration_basic",
		"stripe_tax_registration",
		"stripe_tax_registration.test",
		"tax_registration_basic_create.tf",
		"tax_registration_basic_update.tf",
		true,
		verifyTaxRegistration(taxRegistrationExpectations{
			Address:           "stripe_tax_registration.test",
			ExpectedCountry:   "CA",
			ExpectedType:      "province_standard",
			ExpectedProvince:  "QC",
			ExpectedExpiresAt: 0,
		}),
		verifyTaxRegistration(taxRegistrationExpectations{
			Address:           "stripe_tax_registration.test",
			ExpectedCountry:   "CA",
			ExpectedType:      "province_standard",
			ExpectedProvince:  "QC",
			ExpectedExpiresAt: -1,
		}),
		verifyTaxRegistrationDestroyCleanup,
	)
}

func TestAccManagedTaxRegistrationCanadaSimplified(t *testing.T) {
	runBaseManagedCase(
		t,
		"tax_registration_canada_simplified",
		"stripe_tax_registration",
		"stripe_tax_registration.test",
		"tax_registration_canada_simplified_create.tf",
		"tax_registration_canada_simplified_update.tf",
		true,
		verifyTaxRegistration(taxRegistrationExpectations{
			Address:           "stripe_tax_registration.test",
			ExpectedCountry:   "CA",
			ExpectedType:      "simplified",
			ExpectedProvince:  "",
			ExpectedExpiresAt: 0,
		}),
		verifyTaxRegistration(taxRegistrationExpectations{
			Address:           "stripe_tax_registration.test",
			ExpectedCountry:   "CA",
			ExpectedType:      "simplified",
			ExpectedProvince:  "",
			ExpectedExpiresAt: -1,
		}),
		verifyTaxRegistrationDestroyCleanup,
	)
}

func TestAccManagedFileUploadBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"file_upload_basic",
		"stripe_file",
		"stripe_file.test",
		"file_upload_basic_create.tf",
		"",
		false,
		verifyFile(fileExpectations{
			Address:          "stripe_file.test",
			ExpectedPurpose:  "dispute_evidence",
			ExpectedFilename: "file_upload_fixture.pdf",
		}),
		nil,
		verifyFileDestroyStillExists,
	)
}

func TestAccManagedFileUploadFileLinkData(t *testing.T) {
	runBaseManagedCase(
		t,
		"file_upload_file_link_data",
		"stripe_file",
		"stripe_file.test",
		"file_upload_file_link_data_create.tf",
		"",
		false,
		verifyFile(fileExpectations{
			Address:                   "stripe_file.test",
			ExpectedPurpose:           "dispute_evidence",
			ExpectedFilename:          "file_upload_fixture.pdf",
			ExpectedFileLinkExpiresAt: 1893456000,
			ExpectedFileLinkMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "file_upload_file_link_data",
			},
		}),
		nil,
		verifyFileDestroyStillExists,
	)
}

func TestAccManagedTreasuryFinancialAccountBasic(t *testing.T) {
	runManagedCaseForGroup(
		t,
		"treasury",
		nil,
		"treasury_financial_account_basic",
		"stripe_treasury_financial_account",
		"stripe_treasury_financial_account.test",
		"treasury_financial_account_basic_create.tf",
		"treasury_financial_account_basic_update.tf",
		true,
		verifyTreasuryFinancialAccount(treasuryFinancialAccountExpectations{
			Address:                     "stripe_treasury_financial_account.test",
			ExpectedNickname:            "SDK Codegen Treasury Account Create",
			ExpectedSupportedCurrencies: []string{"usd"},
			ExpectedFeaturePaths:        []string{"financial_addresses.aba"},
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "treasury_financial_account_basic",
				"phase": "create",
			},
		}),
		verifyTreasuryFinancialAccount(treasuryFinancialAccountExpectations{
			Address:                     "stripe_treasury_financial_account.test",
			ExpectedNickname:            "SDK Codegen Treasury Account Update",
			ExpectedSupportedCurrencies: []string{"usd"},
			ExpectedFeaturePaths:        []string{"financial_addresses.aba"},
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "treasury_financial_account_basic",
				"phase": "update",
			},
		}),
		verifyTreasuryFinancialAccountDestroyStillExists,
	)
}

func TestAccManagedTreasuryFinancialAccountMoneyMovementFeatures(t *testing.T) {
	runManagedCaseForGroup(
		t,
		"treasury",
		nil,
		"treasury_financial_account_money_movement_features",
		"stripe_treasury_financial_account",
		"stripe_treasury_financial_account.test",
		"treasury_financial_account_money_movement_features_create.tf",
		"treasury_financial_account_money_movement_features_update.tf",
		true,
		verifyTreasuryFinancialAccount(treasuryFinancialAccountExpectations{
			Address:                     "stripe_treasury_financial_account.test",
			ExpectedNickname:            "SDK Codegen Treasury Movement Create",
			ExpectedSupportedCurrencies: []string{"usd"},
			ExpectedFeaturePaths: []string{
				"financial_addresses.aba",
				"inbound_transfers.ach",
				"outbound_transfers.ach",
			},
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "treasury_financial_account_money_movement_features",
				"phase": "create",
			},
		}),
		verifyTreasuryFinancialAccount(treasuryFinancialAccountExpectations{
			Address:                     "stripe_treasury_financial_account.test",
			ExpectedNickname:            "SDK Codegen Treasury Movement Update",
			ExpectedSupportedCurrencies: []string{"usd"},
			ExpectedFeaturePaths: []string{
				"financial_addresses.aba",
				"inbound_transfers.ach",
				"outbound_payments.ach",
				"outbound_transfers.ach",
			},
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "treasury_financial_account_money_movement_features",
				"phase": "update",
			},
		}),
		verifyTreasuryFinancialAccountDestroyStillExists,
	)
}
