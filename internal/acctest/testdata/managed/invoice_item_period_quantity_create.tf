# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name  = "acctest-invoice-item-period-customer-{{RAND}}"
  email = "sdk-codegen+invoice-item-period-{{RAND}}@example.com"
}

resource "stripe_invoice_item" "test" {
  customer            = stripe_customer.customer.id
  currency            = "usd"
  description         = "sdk-codegen invoice item period create"
  quantity            = 1
  unit_amount_decimal = 2200

  period = {
    start = 1893456000
    end   = 1893542400
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "invoice_item_period_quantity"
    phase = "create"
  }
}
