# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
ephemeral "stripe_reporting_report_run" "test" {
  report_type = "balance.summary.1"

  parameters = {
    interval_start = 1778544000
    interval_end   = 1778630400
  }
}
