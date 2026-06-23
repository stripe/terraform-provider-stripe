# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_charge" "funding" {
  amount      = 500
  currency    = "usd"
  description = "sdk-codegen payout funding"
  source      = "tok_visa"
}

action "stripe_payout" "test" {
  config {
    amount      = 101
    currency    = "usd"
    description = "sdk-codegen payout basic"
    method      = "standard"
    source_type = "card"

    metadata = {
      suite = "sdk-codegen"
      case  = "payout_basic"
    }
  }
}
