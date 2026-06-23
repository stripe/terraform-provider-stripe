# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name  = "acctest-invoice-item-period-customer-{{RAND}}"
  email = "sdk-codegen+invoice-item-period-{{RAND}}@example.com"
}

resource "stripe_invoice_item" "test" {
  customer            = stripe_customer.customer.id
  currency            = "usd"
  description         = "sdk-codegen invoice item period update"
  quantity            = 2
  unit_amount_decimal = 2200

  period = {
    start = 1893628800
    end   = 1893715200
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "invoice_item_period_quantity"
    phase = "update"
  }
}
