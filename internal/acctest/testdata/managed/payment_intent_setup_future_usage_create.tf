# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name  = "acctest-payment-intent-customer-{{RAND}}"
  email = "sdk-codegen+pi-customer-{{RAND}}@example.com"
}

resource "stripe_payment_intent" "test" {
  amount             = 4200
  currency           = "usd"
  customer           = stripe_customer.customer.id
  description        = "sdk-codegen setup future usage create"
  setup_future_usage = "off_session"

  automatic_payment_methods = {
    enabled = true
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "payment_intent_setup_future_usage"
    phase = "create"
  }
}
