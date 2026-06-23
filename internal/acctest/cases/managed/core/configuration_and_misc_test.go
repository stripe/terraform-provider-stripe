// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import "testing"

func TestAccManagedCreditNoteBasic(t *testing.T) {
	runManagedCaseForGroup(
		t,
		"base",
		[]string{"STRIPE_CREDIT_NOTE_INVOICE"},
		"credit_note_basic",
		"stripe_credit_note",
		"stripe_credit_note.test",
		"credit_note_basic_create.tf",
		"credit_note_basic_update.tf",
		true,
		verifyCreditNote(creditNoteExpectations{
			Address:        "stripe_credit_note.test",
			ExpectedMemo:   "sdk-codegen credit note create",
			ExpectedStatus: "issued",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "credit_note_basic",
				"phase": "create",
			},
		}),
		verifyCreditNote(creditNoteExpectations{
			Address:        "stripe_credit_note.test",
			ExpectedMemo:   "sdk-codegen credit note update",
			ExpectedStatus: "issued",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "credit_note_basic",
				"phase": "update",
			},
		}),
		verifyCreditNoteDestroyStillExists,
	)
}

func TestAccManagedCreditNoteEffectiveAtCustomLine(t *testing.T) {
	runManagedCaseForGroup(
		t,
		"base",
		[]string{"STRIPE_CREDIT_NOTE_INVOICE", "STRIPE_CREDIT_NOTE_INVOICE_FINALIZED_AT"},
		"credit_note_effective_at_custom_line",
		"stripe_credit_note",
		"stripe_credit_note.test",
		"credit_note_effective_at_custom_line_create.tf",
		"credit_note_effective_at_custom_line_update.tf",
		true,
		verifyCreditNote(creditNoteExpectations{
			Address:        "stripe_credit_note.test",
			ExpectedAmount: 100,
			ExpectedMemo:   "sdk-codegen credit note custom line create",
			ExpectedReason: "order_change",
			ExpectedStatus: "issued",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "credit_note_effective_at_custom_line",
				"phase": "create",
			},
		}),
		verifyCreditNote(creditNoteExpectations{
			Address:        "stripe_credit_note.test",
			ExpectedAmount: 150,
			ExpectedMemo:   "sdk-codegen credit note custom line update",
			ExpectedReason: "product_unsatisfactory",
			ExpectedStatus: "issued",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "credit_note_effective_at_custom_line",
				"phase": "update",
			},
		}),
		verifyCreditNoteDestroyStillExists,
	)
}

func TestAccManagedCustomerBalanceTransactionBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"customer_balance_transaction_basic",
		"stripe_customer_balance_transaction",
		"stripe_customer_balance_transaction.test",
		"customer_balance_transaction_basic_create.tf",
		"customer_balance_transaction_basic_update.tf",
		false,
		verifyCustomerBalanceTransaction(customerBalanceTransactionExpectations{
			Address:                 "stripe_customer_balance_transaction.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedAmount:          -500,
			ExpectedCurrency:        "usd",
			ExpectedDescription:     "sdk-codegen customer balance create",
			ExpectedType:            "adjustment",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "customer_balance_transaction_basic",
				"phase": "create",
			},
		}),
		verifyCustomerBalanceTransaction(customerBalanceTransactionExpectations{
			Address:                 "stripe_customer_balance_transaction.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedAmount:          -500,
			ExpectedCurrency:        "usd",
			ExpectedDescription:     "sdk-codegen customer balance update",
			ExpectedType:            "adjustment",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "customer_balance_transaction_basic",
				"phase": "update",
			},
		}),
		verifyCustomerBalanceTransactionDestroyStillExists,
	)
}

func TestAccManagedFileLinkBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"file_link_basic",
		"stripe_file_link",
		"stripe_file_link.test",
		"file_link_basic_create.tf",
		"file_link_basic_update.tf",
		true,
		verifyFileLink(fileLinkExpectations{
			Address:             "stripe_file_link.test",
			ExpectedFileAddress: "stripe_file.file",
			ExpectedExpiresAt:   1893456000,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "file_link_basic",
				"phase": "create",
			},
		}),
		verifyFileLink(fileLinkExpectations{
			Address:             "stripe_file_link.test",
			ExpectedFileAddress: "stripe_file.file",
			ExpectedExpiresAt:   1893888000,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "file_link_basic",
				"phase": "update",
			},
		}),
		verifyFileLinkDestroyStillExists,
	)
}

func TestAccManagedFileLinkNoExpiry(t *testing.T) {
	runBaseManagedCase(
		t,
		"file_link_no_expiry",
		"stripe_file_link",
		"stripe_file_link.test",
		"file_link_no_expiry_create.tf",
		"",
		true,
		verifyFileLink(fileLinkExpectations{
			Address:             "stripe_file_link.test",
			ExpectedFileAddress: "stripe_file.file",
			ExpectedExpiresAt:   0,
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "file_link_no_expiry",
			},
		}),
		nil,
		verifyFileLinkDestroyStillExists,
	)
}

func TestAccManagedTaxIDBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"tax_id_basic",
		"stripe_tax_id",
		"stripe_tax_id.test",
		"tax_id_basic_create.tf",
		"",
		false,
		verifyTaxID(taxIDExpectations{
			Address:                 "stripe_tax_id.test",
			ExpectedCustomerAddress: "stripe_customer.customer",
			ExpectedType:            "us_ein",
			ExpectedValue:           "12-3456789",
			ExpectedCountry:         "US",
		}),
		nil,
		verifyTaxIDDestroyMissing,
	)
}

func TestAccManagedBillingPortalConfigurationBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"billing_portal_configuration_basic",
		"stripe_billing_portal_configuration",
		"stripe_billing_portal_configuration.test",
		"billing_portal_configuration_basic_create.tf",
		"billing_portal_configuration_basic_update.tf",
		true,
		verifyBillingPortalConfiguration(billingPortalConfigurationExpectations{
			Address:                  "stripe_billing_portal_configuration.test",
			ExpectedActive:           true,
			ExpectedName:             "SDK Codegen Portal Create",
			ExpectedDefaultReturnURL: "https://example.com/sdk-codegen/portal/create",
			ExpectedHeadline:         "SDK Codegen portal create headline",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "billing_portal_configuration_basic",
				"phase": "create",
			},
			ExpectedCustomerUpdateEnabled:      true,
			ExpectedAllowedUpdates:             []string{"email", "tax_id"},
			ExpectedInvoiceHistoryEnabled:      true,
			ExpectedPaymentMethodUpdateEnabled: true,
		}),
		verifyBillingPortalConfiguration(billingPortalConfigurationExpectations{
			Address:                  "stripe_billing_portal_configuration.test",
			ExpectedActive:           false,
			ExpectedName:             "SDK Codegen Portal Update",
			ExpectedDefaultReturnURL: "https://example.com/sdk-codegen/portal/update",
			ExpectedHeadline:         "SDK Codegen portal update headline",
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "billing_portal_configuration_basic",
				"phase": "update",
			},
			ExpectedCustomerUpdateEnabled:      true,
			ExpectedAllowedUpdates:             []string{"address", "email"},
			ExpectedInvoiceHistoryEnabled:      true,
			ExpectedPaymentMethodUpdateEnabled: true,
		}),
		verifyBillingPortalConfigurationDestroyInactive,
	)
}

func TestAccManagedBillingPortalConfigurationLoginCancel(t *testing.T) {
	runBaseManagedCase(
		t,
		"billing_portal_configuration_login_cancel",
		"stripe_billing_portal_configuration",
		"stripe_billing_portal_configuration.test",
		"billing_portal_configuration_login_cancel_create.tf",
		"billing_portal_configuration_login_cancel_update.tf",
		false,
		verifyBillingPortalConfiguration(billingPortalConfigurationExpectations{
			Address:                                     "stripe_billing_portal_configuration.test",
			ExpectedActive:                              true,
			ExpectedName:                                "SDK Codegen Portal Login Cancel Create",
			ExpectedDefaultReturnURL:                    "https://example.com/sdk-codegen/portal/login-cancel/create",
			ExpectedHeadline:                            "SDK Codegen portal login cancel create headline",
			CheckLoginPage:                              true,
			ExpectedLoginPageEnabled:                    true,
			ExpectedPrivacyPolicyURL:                    "https://example.com/login-cancel/privacy",
			ExpectedTermsOfServiceURL:                   "https://example.com/login-cancel/terms",
			CheckSubscriptionCancel:                     true,
			ExpectedSubscriptionCancelEnabled:           true,
			ExpectedSubscriptionCancelMode:              "at_period_end",
			ExpectedSubscriptionCancelProrationBehavior: "none",
			ExpectedCancellationReasonEnabled:           true,
			ExpectedCancellationReasonOptions:           []string{"missing_features", "too_expensive"},
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "billing_portal_configuration_login_cancel",
				"phase": "create",
			},
			ExpectedCustomerUpdateEnabled:      true,
			ExpectedAllowedUpdates:             []string{"email", "phone"},
			ExpectedInvoiceHistoryEnabled:      true,
			ExpectedPaymentMethodUpdateEnabled: true,
		}),
		verifyBillingPortalConfiguration(billingPortalConfigurationExpectations{
			Address:                                     "stripe_billing_portal_configuration.test",
			ExpectedActive:                              false,
			ExpectedName:                                "SDK Codegen Portal Login Cancel Update",
			ExpectedDefaultReturnURL:                    "https://example.com/sdk-codegen/portal/login-cancel/update",
			ExpectedHeadline:                            "SDK Codegen portal login cancel update headline",
			CheckLoginPage:                              true,
			ExpectedLoginPageEnabled:                    false,
			ExpectedPrivacyPolicyURL:                    "https://example.com/login-cancel/privacy-update",
			ExpectedTermsOfServiceURL:                   "https://example.com/login-cancel/terms-update",
			CheckSubscriptionCancel:                     true,
			ExpectedSubscriptionCancelEnabled:           true,
			ExpectedSubscriptionCancelMode:              "immediately",
			ExpectedSubscriptionCancelProrationBehavior: "create_prorations",
			ExpectedCancellationReasonEnabled:           true,
			ExpectedCancellationReasonOptions:           []string{"too_complex", "unused"},
			ExpectedMetadata: map[string]string{
				"suite": "sdk-codegen",
				"case":  "billing_portal_configuration_login_cancel",
				"phase": "update",
			},
			ExpectedCustomerUpdateEnabled:      true,
			ExpectedAllowedUpdates:             []string{"address", "shipping"},
			ExpectedInvoiceHistoryEnabled:      true,
			ExpectedPaymentMethodUpdateEnabled: true,
		}),
		verifyBillingPortalConfigurationDestroyInactive,
	)
}

func TestAccManagedPaymentMethodConfigurationBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"payment_method_configuration_basic",
		"stripe_payment_method_configuration",
		"stripe_payment_method_configuration.test",
		"payment_method_configuration_basic_create.tf",
		"payment_method_configuration_basic_update.tf",
		true,
		verifyPaymentMethodConfiguration(paymentMethodConfigurationExpectations{
			Address:                "stripe_payment_method_configuration.test",
			ExpectedActive:         true,
			ExpectedName:           "SDK Codegen PM Configuration Create",
			ExpectedCardPreference: "on",
			ExpectedCardValue:      "on",
		}),
		verifyPaymentMethodConfiguration(paymentMethodConfigurationExpectations{
			Address:                "stripe_payment_method_configuration.test",
			ExpectedActive:         false,
			ExpectedName:           "SDK Codegen PM Configuration Update",
			ExpectedCardPreference: "off",
			ExpectedCardValue:      "off",
		}),
		verifyPaymentMethodConfigurationDestroyInactive,
	)
}

func TestAccManagedPaymentMethodConfigurationWalletPreferences(t *testing.T) {
	runBaseManagedCase(
		t,
		"payment_method_configuration_wallet_preferences",
		"stripe_payment_method_configuration",
		"stripe_payment_method_configuration.test",
		"payment_method_configuration_wallet_preferences_create.tf",
		"",
		false,
		verifyPaymentMethodConfiguration(paymentMethodConfigurationExpectations{
			Address:                    "stripe_payment_method_configuration.test",
			ExpectedActive:             true,
			ExpectedName:               "SDK Codegen PM Configuration Wallets",
			ExpectedCardPreference:     "on",
			ExpectedCardValue:          "on",
			ExpectedApplePayPreference: "off",
			ExpectedApplePayValue:      "off",
			ExpectedLinkPreference:     "off",
			ExpectedLinkValue:          "off",
		}),
		nil,
		verifyPaymentMethodConfigurationDestroyInactive,
	)
}

func TestAccManagedPaymentMethodDomainBasic(t *testing.T) {
	runManagedCaseForGroup(
		t,
		"base",
		[]string{"STRIPE_PAYMENT_METHOD_DOMAIN"},
		"payment_method_domain_basic",
		"stripe_payment_method_domain",
		"stripe_payment_method_domain.test",
		"payment_method_domain_basic_create.tf",
		"payment_method_domain_basic_update.tf",
		true,
		verifyPaymentMethodDomain(paymentMethodDomainExpectations{
			Address:         "stripe_payment_method_domain.test",
			ExpectedEnabled: true,
		}),
		verifyPaymentMethodDomain(paymentMethodDomainExpectations{
			Address:         "stripe_payment_method_domain.test",
			ExpectedEnabled: false,
		}),
		verifyPaymentMethodDomainDestroyStillExists,
	)
}

func TestAccManagedProductFeatureBasic(t *testing.T) {
	runBaseManagedCase(
		t,
		"product_feature_basic",
		"stripe_product_feature",
		"stripe_product_feature.test",
		"product_feature_basic_create.tf",
		"",
		false,
		verifyProductFeature(productFeatureExpectations{
			Address:                           "stripe_product_feature.test",
			ExpectedProductAddress:            "stripe_product.product",
			ExpectedEntitlementFeatureAddress: "stripe_entitlements_feature.feature",
		}),
		nil,
		verifyProductFeatureDestroyMissing,
	)
}

func TestAccManagedApplePayDomainBasic(t *testing.T) {
	runManagedCaseForGroup(
		t,
		"base",
		[]string{"STRIPE_APPLE_PAY_DOMAIN"},
		"apple_pay_domain_basic",
		"stripe_apple_pay_domain",
		"stripe_apple_pay_domain.test",
		"apple_pay_domain_basic_create.tf",
		"",
		true,
		verifyApplePayDomain(applePayDomainExpectations{
			Address: "stripe_apple_pay_domain.test",
		}),
		nil,
		verifyApplePayDomainDestroyMissing,
	)
}
