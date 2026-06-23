// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import "testing"

func TestAccEphemeralReportingReportRunBasic(t *testing.T) {
	runBaseEphemeralCase(
		t,
		"reporting_report_run_basic",
		"stripe_reporting_report_run",
		"ephemeral/reporting/reporting_report_run_basic_ephemeral.tf",
		nil,
		verifyReportingReportRunEphemeral,
	)
}
