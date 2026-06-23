// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import "testing"

func TestAccEphemeralTaxCalculationBasic(t *testing.T) {
	runBaseEphemeralCase(
		t,
		"tax_calculation_basic",
		"stripe_tax_calculation",
		"ephemeral/tax/tax_calculation_basic_ephemeral.tf",
		nil,
		verifyTaxCalculationEphemeral,
	)
}
