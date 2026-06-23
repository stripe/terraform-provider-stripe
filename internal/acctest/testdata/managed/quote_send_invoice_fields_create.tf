# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name  = "acctest-quote-send-customer-{{RAND}}"
  email = "sdk-codegen+quote-send-{{RAND}}@example.com"
}

resource "stripe_product" "product" {
  name = "acctest-quote-send-product-{{RAND}}"
}

resource "stripe_price" "price" {
  product     = stripe_product.product.id
  currency    = "usd"
  unit_amount = 2200
}

resource "stripe_quote" "test" {
  customer          = stripe_customer.customer.id
  collection_method = "send_invoice"
  expires_at        = 1893456000
  description       = "sdk-codegen quote send invoice create"
  header            = "SDK Codegen Quote Send Invoice Create"
  footer            = "SDK Codegen quote footer create"

  invoice_settings = {
    days_until_due = 14
  }

  line_items = [
    {
      price    = stripe_price.price.id
      quantity = 1
    },
  ]

  metadata = {
    suite = "sdk-codegen"
    case  = "quote_send_invoice_fields"
    phase = "create"
  }
}
