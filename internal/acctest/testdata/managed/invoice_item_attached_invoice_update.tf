# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name  = "acctest-invoice-item-attached-customer-{{RAND}}"
  email = "sdk-codegen+invoice-item-attached-{{RAND}}@example.com"
}

resource "stripe_invoice" "invoice" {
  customer     = stripe_customer.customer.id
  auto_advance = false
  description  = "sdk-codegen invoice item attached invoice"
}

resource "stripe_invoice_item" "test" {
  customer     = stripe_customer.customer.id
  invoice      = stripe_invoice.invoice.id
  amount       = 2100
  currency     = "usd"
  description  = "sdk-codegen invoice item attached update"
  discountable = false

  metadata = {
    suite = "sdk-codegen"
    case  = "invoice_item_attached_invoice"
    phase = "update"
  }
}
