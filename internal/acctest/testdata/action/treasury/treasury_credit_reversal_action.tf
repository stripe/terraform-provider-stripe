# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
action "stripe_treasury_credit_reversal" "test" {
  config {
    received_credit = "{{TREASURY_RECEIVED_CREDIT}}"

    metadata = {
      suite             = "sdk-codegen"
      case              = "treasury_credit_reversal_action"
      financial_account = "{{TREASURY_SOURCE_FINANCIAL_ACCOUNT}}"
    }
  }
}
