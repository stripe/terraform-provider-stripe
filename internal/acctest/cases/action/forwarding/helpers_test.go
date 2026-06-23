// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	stripe "github.com/stripe/stripe-go/v86"

	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

var forwardingPaymentMethod string

type forwardingRequestActionExpectations struct {
	CaseName             string
	ExpectedReference    string
	ExpectedReplacements []string
	ExpectedHeaders      map[string]string
}

type forwardingRequestBody struct {
	Amount    forwardingRequestAmount `json:"amount"`
	Card      forwardingRequestCard   `json:"card"`
	Reference string                  `json:"reference"`
}

type forwardingRequestAmount struct {
	Currency string `json:"currency"`
	Value    int64  `json:"value"`
}

type forwardingRequestCard struct {
	CVC      string `json:"cvc"`
	ExpMonth string `json:"exp_month"`
	ExpYear  string `json:"exp_year"`
	Name     string `json:"name"`
	Number   string `json:"number"`
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
		expectedPaymentMethod := expectedForwardingPaymentMethod(env)
		if request.PaymentMethod != expectedPaymentMethod {
			return fmt.Errorf(
				"forwarding request %s payment_method mismatch: expected %q, got %q",
				expect.CaseName,
				expectedPaymentMethod,
				request.PaymentMethod,
			)
		}
		if request.RequestDetails == nil {
			return fmt.Errorf("forwarding request %s request_details missing", expect.CaseName)
		}
		if err := verifyForwardingRequestBody(expect, request.RequestDetails.Body); err != nil {
			return err
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
				return fmt.Errorf("forwarding request %s missing header %q", expect.CaseName, key)
			}
			if key == "Authorization" && strings.HasPrefix(actualValue, "Bearer ") {
				continue
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

func verifyForwardingRequestBody(
	expect forwardingRequestActionExpectations,
	body string,
) error {
	var actual forwardingRequestBody
	if err := json.Unmarshal([]byte(body), &actual); err != nil {
		return fmt.Errorf("forwarding request %s body is not JSON: %w", expect.CaseName, err)
	}
	if actual.Reference != expect.ExpectedReference {
		return fmt.Errorf(
			"forwarding request %s reference mismatch: expected %q, got %q",
			expect.CaseName,
			expect.ExpectedReference,
			actual.Reference,
		)
	}
	if actual.Amount.Currency != "USD" || actual.Amount.Value != 1000 {
		return fmt.Errorf(
			"forwarding request %s amount mismatch: expected USD 1000, got %s %d",
			expect.CaseName,
			actual.Amount.Currency,
			actual.Amount.Value,
		)
	}
	if actual.Card.Number != "411111******1111" {
		return fmt.Errorf(
			"forwarding request %s card.number mismatch: expected %q, got %q",
			expect.CaseName,
			"411111******1111",
			actual.Card.Number,
		)
	}
	if actual.Card.ExpMonth != "03" || actual.Card.ExpYear != "2030" {
		return fmt.Errorf(
			"forwarding request %s card expiry mismatch: expected 03/2030, got %s/%s",
			expect.CaseName,
			actual.Card.ExpMonth,
			actual.Card.ExpYear,
		)
	}
	if actual.Card.CVC != "***" {
		return fmt.Errorf(
			"forwarding request %s card.cvc mismatch: expected %q, got %q",
			expect.CaseName,
			"***",
			actual.Card.CVC,
		)
	}
	if actual.Card.Name != "SDK Codegen Forwarding" {
		return fmt.Errorf(
			"forwarding request %s card.name mismatch: expected %q, got %q",
			expect.CaseName,
			"SDK Codegen Forwarding",
			actual.Card.Name,
		)
	}

	return nil
}

func prepareForwardingRequestConfig(
	t *testing.T,
	env runner.TestEnv,
	client *stripe.Client,
) map[string]string {
	t.Helper()

	if env.ForwardingPaymentMethod != "" {
		forwardingPaymentMethod = env.ForwardingPaymentMethod
	} else {
		paymentMethod, err := client.V1PaymentMethods.Create(
			context.Background(),
			&stripe.PaymentMethodCreateParams{
				BillingDetails: &stripe.PaymentMethodCreateBillingDetailsParams{
					Name: stripe.String("SDK Codegen Forwarding"),
				},
				Card: &stripe.PaymentMethodCreateCardParams{
					Number:   stripe.String("4111111111111111"),
					ExpMonth: stripe.Int64(3),
					ExpYear:  stripe.Int64(2030),
					CVC:      stripe.String("737"),
				},
				Type: stripe.String(string(stripe.PaymentMethodTypeCard)),
			},
		)
		if err != nil {
			t.Fatalf("create forwarding payment method: %v", err)
		}
		forwardingPaymentMethod = paymentMethod.ID
	}

	return map[string]string{
		"{{FORWARDING_PAYMENT_METHOD}}": forwardingPaymentMethod,
	}
}

func expectedForwardingPaymentMethod(env runner.TestEnv) string {
	if env.ForwardingPaymentMethod != "" {
		return env.ForwardingPaymentMethod
	}
	return forwardingPaymentMethod
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

func expectRemoteStringList(name string, actual []string, expected []string) error {
	if !reflect.DeepEqual(actual, expected) {
		return fmt.Errorf("remote %s mismatch: expected %v, got %v", name, expected, actual)
	}

	return nil
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
