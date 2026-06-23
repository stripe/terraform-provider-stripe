# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name  = "acctest-invoice-item-price-customer-{{RAND}}"
  email = "sdk-codegen+invoice-item-price-{{RAND}}@example.com"
}

resource "stripe_product" "product" {
  name = "acctest-invoice-item-product-{{RAND}}"
}

resource "stripe_price" "price" {
  product     = stripe_product.product.id
  currency    = "usd"
  unit_amount = 900
}

resource "stripe_invoice_item" "test" {
  customer    = stripe_customer.customer.id
  description = "sdk-codegen invoice item pricing create"
  quantity    = 3

  pricing = {
    price = stripe_price.price.id
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "invoice_item_pricing_price_regression"
    phase = "create"
  }
}
