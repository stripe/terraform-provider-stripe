# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
action "stripe_treasury_inbound_transfer" "test" {
  config {
    amount                = 210
    currency              = "usd"
    financial_account     = "{{TREASURY_SOURCE_FINANCIAL_ACCOUNT}}"
    origin_payment_method = "{{TREASURY_ORIGIN_PAYMENT_METHOD}}"
    description           = "sdk-codegen treasury inbound transfer basic"

    metadata = {
      suite = "sdk-codegen"
      case  = "treasury_inbound_transfer_basic"
    }
  }
}
