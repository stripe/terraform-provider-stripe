// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package runner

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

type TestEnv struct {
	APIKey                                   string
	StripeAccount                            string
	IssuingAccount                           string
	IssuingFinancialAccount                  string
	TreasuryAccount                          string
	CreditNoteInvoice                        string
	CreditNoteInvoiceFinalizedAt             string
	ApplePayDomain                           string
	PaymentMethodDomain                      string
	IssuingDisputeTransaction                string
	IssuingDisputeOtherTransaction           string
	IssuingPhysicalBundle                    string
	IssuingCardLogo                          string
	ForwardingDestination                    string
	ForwardingPaymentMethod                  string
	TreasurySourceFinancialAccount           string
	TreasuryDestinationFinancialAccount      string
	TreasuryOriginPaymentMethod              string
	TreasuryReceivedCredit                   string
	TreasuryReceivedDebit                    string
	TerminalLocation                         string
	TerminalReaderRegistrationCode           string
	TerminalReaderRegistrationCodeUpdate     string
	TerminalReaderRegistrationCodeRegression string
}

func AcceptanceEnabled() bool {
	return os.Getenv("TF_ACC") != ""
}

func ResolveTestEnv(t *testing.T, group string, requiredEnv []string) TestEnv {
	t.Helper()

	env := TestEnv{
		APIKey:                                   os.Getenv("STRIPE_API_KEY"),
		StripeAccount:                            os.Getenv("STRIPE_ACCOUNT"),
		IssuingAccount:                           os.Getenv("STRIPE_ISSUING_ACCOUNT"),
		IssuingFinancialAccount:                  os.Getenv("STRIPE_ISSUING_FINANCIAL_ACCOUNT"),
		TreasuryAccount:                          os.Getenv("STRIPE_TREASURY_ACCOUNT"),
		CreditNoteInvoice:                        os.Getenv("STRIPE_CREDIT_NOTE_INVOICE"),
		CreditNoteInvoiceFinalizedAt:             os.Getenv("STRIPE_CREDIT_NOTE_INVOICE_FINALIZED_AT"),
		ApplePayDomain:                           os.Getenv("STRIPE_APPLE_PAY_DOMAIN"),
		PaymentMethodDomain:                      os.Getenv("STRIPE_PAYMENT_METHOD_DOMAIN"),
		IssuingDisputeTransaction:                os.Getenv("STRIPE_ISSUING_DISPUTE_TRANSACTION"),
		IssuingDisputeOtherTransaction:           os.Getenv("STRIPE_ISSUING_DISPUTE_OTHER_TRANSACTION"),
		IssuingPhysicalBundle:                    os.Getenv("STRIPE_ISSUING_PHYSICAL_BUNDLE"),
		IssuingCardLogo:                          os.Getenv("STRIPE_ISSUING_CARD_LOGO"),
		ForwardingDestination:                    os.Getenv("STRIPE_FORWARDING_DESTINATION"),
		ForwardingPaymentMethod:                  os.Getenv("STRIPE_FORWARDING_PAYMENT_METHOD"),
		TreasurySourceFinancialAccount:           os.Getenv("STRIPE_TREASURY_SOURCE_FINANCIAL_ACCOUNT"),
		TreasuryDestinationFinancialAccount:      os.Getenv("STRIPE_TREASURY_DESTINATION_FINANCIAL_ACCOUNT"),
		TreasuryOriginPaymentMethod:              os.Getenv("STRIPE_TREASURY_ORIGIN_PAYMENT_METHOD"),
		TreasuryReceivedCredit:                   os.Getenv("STRIPE_TREASURY_RECEIVED_CREDIT"),
		TreasuryReceivedDebit:                    os.Getenv("STRIPE_TREASURY_RECEIVED_DEBIT"),
		TerminalLocation:                         os.Getenv("STRIPE_TERMINAL_LOCATION"),
		TerminalReaderRegistrationCode:           os.Getenv("STRIPE_TERMINAL_READER_REGISTRATION_CODE"),
		TerminalReaderRegistrationCodeUpdate:     os.Getenv("STRIPE_TERMINAL_READER_REGISTRATION_CODE_UPDATE"),
		TerminalReaderRegistrationCodeRegression: os.Getenv("STRIPE_TERMINAL_READER_REGISTRATION_CODE_REGRESSION"),
	}

	switch group {
	case "base":
		// Base surfaces should run against the API key owner account.
		// Using STRIPE_ACCOUNT here unintentionally impersonates connected accounts
		// and causes endpoint/capability errors on many core resources.
		env.StripeAccount = ""
	case "forwarding":
		// Forwarding endpoints are platform-scoped and do not support connected
		// account impersonation.
		env.StripeAccount = ""
		if env.ForwardingDestination == "" {
			env.ForwardingDestination = "https://forwarding-api-demo.stripedemos.com/payments"
		}
	case "issuing":
		// Issuing fixtures and IDs are created in STRIPE_ISSUING_ACCOUNT.
		// Always target that account to avoid cross-account resource_missing failures.
		env.StripeAccount = env.IssuingAccount
	case "treasury":
		// Treasury fixtures and IDs are tied to STRIPE_TREASURY_ACCOUNT.
		env.StripeAccount = env.TreasuryAccount
	}

	missing := missingEnv(group, requiredEnv, env)
	if len(missing) == 0 {
		return env
	}

	message := fmt.Sprintf(
		"missing env for %s group: %s",
		group,
		strings.Join(missing, ", "),
	)
	if os.Getenv("STRIPE_TF_ACCTEST_REQUIRE_ENV") == "1" {
		t.Fatal(message)
	}
	t.Skip(message)
	return env
}

func missingEnv(group string, requiredEnv []string, env TestEnv) []string {
	missing := []string{}
	if env.APIKey == "" {
		missing = append(missing, "STRIPE_API_KEY")
	}

	switch group {
	case "base":
	case "connect":
		if env.StripeAccount == "" {
			missing = append(missing, "STRIPE_ACCOUNT")
		}
	case "issuing":
		if env.IssuingAccount == "" {
			missing = append(missing, "STRIPE_ISSUING_ACCOUNT")
		}
	case "treasury":
		if env.TreasuryAccount == "" {
			missing = append(missing, "STRIPE_TREASURY_ACCOUNT")
		}
	case "forwarding":
	}

	for _, envKey := range requiredEnv {
		if envValue(env, envKey) == "" {
			missing = append(missing, envKey)
		}
	}

	return missing
}

func envValue(env TestEnv, envKey string) string {
	switch envKey {
	case "STRIPE_API_KEY":
		return env.APIKey
	case "STRIPE_ACCOUNT":
		return env.StripeAccount
	case "STRIPE_ISSUING_ACCOUNT":
		return env.IssuingAccount
	case "STRIPE_ISSUING_FINANCIAL_ACCOUNT":
		return env.IssuingFinancialAccount
	case "STRIPE_TREASURY_ACCOUNT":
		return env.TreasuryAccount
	case "STRIPE_CREDIT_NOTE_INVOICE":
		return env.CreditNoteInvoice
	case "STRIPE_CREDIT_NOTE_INVOICE_FINALIZED_AT":
		return env.CreditNoteInvoiceFinalizedAt
	case "STRIPE_APPLE_PAY_DOMAIN":
		return env.ApplePayDomain
	case "STRIPE_PAYMENT_METHOD_DOMAIN":
		return env.PaymentMethodDomain
	case "STRIPE_ISSUING_DISPUTE_TRANSACTION":
		return env.IssuingDisputeTransaction
	case "STRIPE_ISSUING_DISPUTE_OTHER_TRANSACTION":
		return env.IssuingDisputeOtherTransaction
	case "STRIPE_ISSUING_PHYSICAL_BUNDLE":
		return env.IssuingPhysicalBundle
	case "STRIPE_ISSUING_CARD_LOGO":
		return env.IssuingCardLogo
	case "STRIPE_FORWARDING_DESTINATION":
		return env.ForwardingDestination
	case "STRIPE_FORWARDING_PAYMENT_METHOD":
		return env.ForwardingPaymentMethod
	case "STRIPE_TREASURY_SOURCE_FINANCIAL_ACCOUNT":
		return env.TreasurySourceFinancialAccount
	case "STRIPE_TREASURY_DESTINATION_FINANCIAL_ACCOUNT":
		return env.TreasuryDestinationFinancialAccount
	case "STRIPE_TREASURY_ORIGIN_PAYMENT_METHOD":
		return env.TreasuryOriginPaymentMethod
	case "STRIPE_TREASURY_RECEIVED_CREDIT":
		return env.TreasuryReceivedCredit
	case "STRIPE_TREASURY_RECEIVED_DEBIT":
		return env.TreasuryReceivedDebit
	case "STRIPE_TERMINAL_LOCATION":
		return env.TerminalLocation
	case "STRIPE_TERMINAL_READER_REGISTRATION_CODE":
		return env.TerminalReaderRegistrationCode
	case "STRIPE_TERMINAL_READER_REGISTRATION_CODE_UPDATE":
		return env.TerminalReaderRegistrationCodeUpdate
	case "STRIPE_TERMINAL_READER_REGISTRATION_CODE_REGRESSION":
		return env.TerminalReaderRegistrationCodeRegression
	default:
		return os.Getenv(envKey)
	}
}

func ApplyResolvedAccountEnv(t *testing.T, env TestEnv) {
	t.Helper()
	t.Setenv("STRIPE_ACCOUNT", env.StripeAccount)
	t.Setenv("STRIPE_ACCOUNT_CLASSIC", env.StripeAccount)
}
