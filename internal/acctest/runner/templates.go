// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package runner

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"time"
)

func RenderManagedConfig(env TestEnv, templateName string, extras map[string]string) string {
	return renderConfig(env, templateName, extras)
}

func RenderEphemeralConfig(env TestEnv, templateName string, extras map[string]string) string {
	return renderConfig(env, templateName, extras)
}

func RenderActionConfig(env TestEnv, templateName string, extras map[string]string) string {
	return renderConfig(env, templateName, extras)
}

func renderConfig(env TestEnv, templateName string, extras map[string]string) string {
	replacements := map[string]string{
		"{{API_KEY}}":                                      env.APIKey,
		"{{STRIPE_ACCOUNT}}":                               env.StripeAccount,
		"{{ISSUING_ACCOUNT}}":                              env.IssuingAccount,
		"{{ISSUING_FINANCIAL_ACCOUNT}}":                    env.IssuingFinancialAccount,
		"{{TREASURY_ACCOUNT}}":                             env.TreasuryAccount,
		"{{CREDIT_NOTE_INVOICE}}":                          env.CreditNoteInvoice,
		"{{CREDIT_NOTE_INVOICE_FINALIZED_AT}}":             env.CreditNoteInvoiceFinalizedAt,
		"{{APPLE_PAY_DOMAIN}}":                             env.ApplePayDomain,
		"{{PAYMENT_METHOD_DOMAIN}}":                        env.PaymentMethodDomain,
		"{{ISSUING_DISPUTE_TRANSACTION}}":                  env.IssuingDisputeTransaction,
		"{{ISSUING_DISPUTE_OTHER_TRANSACTION}}":            env.IssuingDisputeOtherTransaction,
		"{{ISSUING_PHYSICAL_BUNDLE}}":                      env.IssuingPhysicalBundle,
		"{{ISSUING_CARD_LOGO}}":                            env.IssuingCardLogo,
		"{{FORWARDING_DESTINATION}}":                       env.ForwardingDestination,
		"{{FORWARDING_PAYMENT_METHOD}}":                    env.ForwardingPaymentMethod,
		"{{TREASURY_SOURCE_FINANCIAL_ACCOUNT}}":            env.TreasurySourceFinancialAccount,
		"{{TREASURY_DESTINATION_FINANCIAL_ACCOUNT}}":       env.TreasuryDestinationFinancialAccount,
		"{{TREASURY_ORIGIN_PAYMENT_METHOD}}":               env.TreasuryOriginPaymentMethod,
		"{{TREASURY_RECEIVED_CREDIT}}":                     env.TreasuryReceivedCredit,
		"{{TREASURY_RECEIVED_DEBIT}}":                      env.TreasuryReceivedDebit,
		"{{TERMINAL_LOCATION}}":                            env.TerminalLocation,
		"{{TERMINAL_READER_REGISTRATION_CODE}}":            env.TerminalReaderRegistrationCode,
		"{{TERMINAL_READER_REGISTRATION_CODE_UPDATE}}":     env.TerminalReaderRegistrationCodeUpdate,
		"{{TERMINAL_READER_REGISTRATION_CODE_REGRESSION}}": env.TerminalReaderRegistrationCodeRegression,
		"{{FILE_UPLOAD_FIXTURE_PATH}}":                     mustResolveTemplatePath("file_upload_fixture.pdf"),
		"{{RAND}}":                                         fmt.Sprintf("%d", time.Now().UnixNano()),
	}
	for key, value := range extras {
		replacements[key] = value
	}

	template := loadTemplate(templateName)
	rendered := template
	for key, value := range replacements {
		rendered = strings.ReplaceAll(rendered, key, value)
	}

	return providerConfig(env) + "\n" + rendered
}

func providerConfig(env TestEnv) string {
	if env.StripeAccount == "" {
		return fmt.Sprintf("provider \"stripe\" {\n  api_key = %q\n}\n", env.APIKey)
	}

	return fmt.Sprintf(
		"provider \"stripe\" {\n  api_key = %q\n  stripe_account = %q\n}\n",
		env.APIKey,
		env.StripeAccount,
	)
}

func loadTemplate(templateName string) string {
	templatePath := testdataPath(templateName)
	contents, err := os.ReadFile(templatePath)
	if err == nil {
		return string(contents)
	}

	resolvedPath, resolveErr := resolveTemplatePath(templateName)
	if resolveErr != nil {
		panic(err)
	}

	contents, err = os.ReadFile(resolvedPath)
	if err != nil {
		panic(err)
	}

	return string(contents)
}

func testdataPath(name string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to resolve acctest template path")
	}

	return filepath.Join(filepath.Dir(filename), "..", "testdata", name)
}

func resolveTemplatePath(templateName string) (string, error) {
	if strings.Contains(templateName, "/") || strings.Contains(templateName, string(filepath.Separator)) {
		return "", os.ErrNotExist
	}

	testdataRoot := testdataPath("")
	matches := make([]string, 0, 1)
	walkErr := filepath.Walk(testdataRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if info.Name() != templateName {
			return nil
		}
		matches = append(matches, path)
		return nil
	})
	if walkErr != nil {
		return "", walkErr
	}
	if len(matches) == 0 {
		return "", os.ErrNotExist
	}

	slices.Sort(matches)
	return matches[0], nil
}

func mustResolveTemplatePath(templateName string) string {
	resolvedPath, err := resolveTemplatePath(templateName)
	if err == nil {
		return resolvedPath
	}

	return testdataPath(templateName)
}
