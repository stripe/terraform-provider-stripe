// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import (
	"testing"

	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

func runForwardingActionCase(
	t *testing.T,
	name string,
	configTemplate string,
	configPreparer runner.ActionConfigPreparer,
	verify runner.StateVerifier,
) {
	t.Helper()

	runner.RunActionCase(t, runner.ActionCase{
		Definition: runner.CaseDefinition{
			Name:    name,
			Surface: "stripe_forwarding_request",
			Group:   "forwarding",
			Kind:    "action",
		},
		ConfigTemplate: configTemplate,
		ConfigPreparer: configPreparer,
		VerifyInvoke:   verify,
	})
}

func TestAccActionForwardingRequestBasic(t *testing.T) {
	runForwardingActionCase(
		t,
		"forwarding_request_basic",
		"action/forwarding/forwarding_request_basic_action.tf",
		prepareForwardingRequestConfig,
		verifyForwardingRequestAction(forwardingRequestActionExpectations{
			CaseName:             "forwarding_request_basic",
			ExpectedReference:    "forwarding_request_basic",
			ExpectedReplacements: []string{"card_number", "card_expiry", "card_cvc", "cardholder_name"},
			ExpectedHeaders: map[string]string{
				"Authorization": "Bearer eyJhbGciOiJIUzI1NiJ9.Zm9yd2FyZGluZy1hcGktZGVtbw.2qoK37CNBmMjMDRERSYUSE-YrjsTgGhHnxMeqOxjrAg",
			},
		}),
	)
}

func TestAccActionForwardingRequestMetadata(t *testing.T) {
	runForwardingActionCase(
		t,
		"forwarding_request_metadata",
		"action/forwarding/forwarding_request_metadata_action.tf",
		prepareForwardingRequestConfig,
		verifyForwardingRequestAction(forwardingRequestActionExpectations{
			CaseName:             "forwarding_request_metadata",
			ExpectedReference:    "forwarding_request_metadata",
			ExpectedReplacements: []string{"card_number", "card_expiry", "card_cvc", "cardholder_name"},
			ExpectedHeaders: map[string]string{
				"Authorization": "Bearer eyJhbGciOiJIUzI1NiJ9.Zm9yd2FyZGluZy1hcGktZGVtbw.2qoK37CNBmMjMDRERSYUSE-YrjsTgGhHnxMeqOxjrAg",
			},
		}),
	)
}

func TestAccActionForwardingRequestRegression(t *testing.T) {
	runForwardingActionCase(
		t,
		"forwarding_request_action_regression",
		"action/forwarding/forwarding_request_regression_action.tf",
		prepareForwardingRequestConfig,
		verifyForwardingRequestAction(forwardingRequestActionExpectations{
			CaseName:          "forwarding_request_action_regression",
			ExpectedReference: "forwarding_request_action_regression",
			ExpectedReplacements: []string{
				"card_number",
				"card_expiry",
				"card_cvc",
				"cardholder_name",
			},
			ExpectedHeaders: map[string]string{
				"Authorization": "Bearer eyJhbGciOiJIUzI1NiJ9.Zm9yd2FyZGluZy1hcGktZGVtbw.2qoK37CNBmMjMDRERSYUSE-YrjsTgGhHnxMeqOxjrAg",
			},
		}),
	)
}
