# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
action "stripe_treasury_debit_reversal" "test" {
  config {
    received_debit = "{{TREASURY_RECEIVED_DEBIT}}"

    metadata = {
      suite             = "sdk-codegen"
      case              = "treasury_debit_reversal_basic"
      financial_account = "{{TREASURY_SOURCE_FINANCIAL_ACCOUNT}}"
    }
  }
}
