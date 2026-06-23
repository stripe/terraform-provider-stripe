# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_payment_intent" "test" {
  amount      = 2500
  currency    = "usd"
  description = "sdk-codegen payment intent basic updated"

  automatic_payment_methods = {
    enabled = true
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "payment_intent_basic"
    phase = "update"
  }
}
