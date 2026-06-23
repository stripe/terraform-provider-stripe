// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import (
	"testing"

	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

func TestAccEphemeralIdentityVerificationSessionBasic(t *testing.T) {
	runner.RunEphemeralCase(t, runner.EphemeralCase{
		Definition: runner.CaseDefinition{
			Name:    "identity_verification_session_basic",
			Surface: "stripe_identity_verification_session",
			Group:   "base",
			Kind:    "ephemeral",
		},
		ConfigTemplate: "ephemeral/identity/identity_verification_session_basic_ephemeral.tf",
		VerifyPersistedState: verifyIdentityVerificationSessionEphemeral(
			identityVerificationSessionEphemeralExpectations{
				ClientReferenceID: "acctest-identity-basic",
				ExpectedType:      "document",
				ExpectedMetadata: map[string]string{
					"suite": "sdk-codegen",
					"case":  "identity_verification_session_basic",
				},
				ExpectedAllowedTypes:       []string{"passport"},
				ExpectedRequireLiveCapture: true,
			},
		),
	})
}

func TestAccEphemeralIdentityVerificationSessionProvidedDetails(t *testing.T) {
	runner.RunEphemeralCase(t, runner.EphemeralCase{
		Definition: runner.CaseDefinition{
			Name:    "identity_verification_session_provided_details",
			Surface: "stripe_identity_verification_session",
			Group:   "base",
			Kind:    "ephemeral",
		},
		ConfigTemplate: "ephemeral/identity/identity_verification_session_provided_details_ephemeral.tf",
		VerifyPersistedState: verifyIdentityVerificationSessionEphemeral(
			identityVerificationSessionEphemeralExpectations{
				ClientReferenceID: "acctest-identity-provided-details",
				ExpectedType:      "document",
				ExpectedMetadata: map[string]string{
					"suite": "sdk-codegen",
					"case":  "identity_verification_session_provided_details",
				},
			},
		),
	})
}

func TestAccEphemeralIdentityVerificationSessionDocumentOptions(t *testing.T) {
	runner.RunEphemeralCase(t, runner.EphemeralCase{
		Definition: runner.CaseDefinition{
			Name:    "identity_verification_session_document_options",
			Surface: "stripe_identity_verification_session",
			Group:   "base",
			Kind:    "ephemeral",
		},
		ConfigTemplate: "ephemeral/identity/identity_verification_session_document_options_ephemeral.tf",
		VerifyPersistedState: verifyIdentityVerificationSessionEphemeral(
			identityVerificationSessionEphemeralExpectations{
				ClientReferenceID: "acctest-identity-document-options",
				ExpectedType:      "document",
				ExpectedMetadata: map[string]string{
					"suite": "sdk-codegen",
					"case":  "identity_verification_session_document_options",
				},
				ExpectedAllowedTypes:          []string{"driving_license", "passport"},
				ExpectedRequireIDNumber:       true,
				ExpectedRequireMatchingSelfie: true,
			},
		),
	})
}
