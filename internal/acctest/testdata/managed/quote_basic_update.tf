# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name  = "acctest-quote-customer-{{RAND}}"
  email = "sdk-codegen+quote-{{RAND}}@example.com"
}

resource "stripe_product" "product" {
  name = "acctest-quote-product-{{RAND}}"
}

resource "stripe_price" "price" {
  product     = stripe_product.product.id
  currency    = "usd"
  unit_amount = 2200
}

resource "stripe_quote" "test" {
  customer    = stripe_customer.customer.id
  description = "sdk-codegen quote basic updated"
  header      = "SDK Codegen Quote Update"

  line_items = [
    {
      price    = stripe_price.price.id
      quantity = 2
    },
  ]

  metadata = {
    suite = "sdk-codegen"
    case  = "quote_basic"
    phase = "update"
  }
}
