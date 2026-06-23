// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import "testing"

func TestAccEphemeralFinancialConnectionsSessionBasic(t *testing.T) {
	runBaseEphemeralCase(
		t,
		"financial_connections_session_basic",
		"stripe_financial_connections_session",
		"ephemeral/financial_connections/financial_connections_session_basic_ephemeral.tf",
		nil,
		verifyFinancialConnectionsSessionEphemeral,
	)
}
