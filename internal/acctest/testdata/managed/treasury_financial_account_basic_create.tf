# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_treasury_financial_account" "test" {
  supported_currencies = ["usd"]
  nickname             = "SDK Codegen Treasury Account Create"

  features = {
    financial_addresses = {
      aba = {
        requested = true
      }
    }
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "treasury_financial_account_basic"
    phase = "create"
  }
}
