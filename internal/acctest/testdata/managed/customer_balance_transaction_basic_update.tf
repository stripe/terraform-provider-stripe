# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name        = "acctest-balance-customer-{{RAND}}"
  email       = "sdk-codegen+balance-{{RAND}}@example.com"
  description = "sdk-codegen balance transaction customer"
}

resource "stripe_customer_balance_transaction" "test" {
  customer    = stripe_customer.customer.id
  amount      = -500
  currency    = "usd"
  description = "sdk-codegen customer balance update"

  metadata = {
    suite = "sdk-codegen"
    case  = "customer_balance_transaction_basic"
    phase = "update"
  }
}
