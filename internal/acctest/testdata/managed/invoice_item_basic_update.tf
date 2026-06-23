# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name  = "acctest-invoice-item-customer-{{RAND}}"
  email = "sdk-codegen+invoice-item-{{RAND}}@example.com"
}

resource "stripe_invoice_item" "test" {
  customer    = stripe_customer.customer.id
  amount      = 1800
  currency    = "usd"
  description = "sdk-codegen invoice item basic updated"

  metadata = {
    suite = "sdk-codegen"
    case  = "invoice_item_basic"
    phase = "update"
  }
}
