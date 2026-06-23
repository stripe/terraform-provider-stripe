# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name  = "acctest-invoice-customer-{{RAND}}"
  email = "sdk-codegen+invoice-{{RAND}}@example.com"
}

resource "stripe_invoice" "test" {
  customer     = stripe_customer.customer.id
  auto_advance = false
  description  = "sdk-codegen invoice basic updated"

  metadata = {
    suite = "sdk-codegen"
    case  = "invoice_basic"
    phase = "update"
  }
}
