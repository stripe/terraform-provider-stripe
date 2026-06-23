// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import (
	"context"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	stripe "github.com/stripe/stripe-go/v86"

	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

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

		session, err := findIdentityVerificationSession(client, expect.ClientReferenceID)
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
			if session.Options.Document.RequireMatchingSelfie != expect.ExpectedRequireMatchingSelfie {
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
				return fmt.Errorf("identity verification session %s expected email verification", expect.ClientReferenceID)
			}
		}
		if expect.ExpectedRequirePhoneVerification {
			if session.Options == nil || session.Options.Phone == nil || !session.Options.Phone.RequireVerification {
				return fmt.Errorf("identity verification session %s expected phone verification", expect.ClientReferenceID)
			}
		}

		return nil
	}
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
			return nil, fmt.Errorf("list identity verification sessions for %s: %w", clientReferenceID, err)
		}
		if session != nil && session.ClientReferenceID == clientReferenceID {
			return session, nil
		}
	}

	return nil, fmt.Errorf("identity verification session with client_reference_id %q not found", clientReferenceID)
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
			return fmt.Errorf("%s metadata[%q] mismatch: expected %q, got %q", name, key, expectedValue, actualValue)
		}
	}

	return nil
}
