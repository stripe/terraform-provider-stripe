# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
action "stripe_treasury_outbound_transfer" "test" {
  config {
    amount            = 65
    currency          = "usd"
    financial_account = "{{TREASURY_SOURCE_FINANCIAL_ACCOUNT}}"
    description       = "sdk-codegen treasury outbound transfer basic"

    destination_payment_method_data = {
      type              = "financial_account"
      financial_account = "{{TREASURY_DESTINATION_FINANCIAL_ACCOUNT}}"
    }

    metadata = {
      suite = "sdk-codegen"
      case  = "treasury_outbound_transfer_basic"
    }
  }
}
