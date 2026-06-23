# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "product" {
  name = "acctest-payment-link-product-{{RAND}}"
}

resource "stripe_price" "price" {
  product     = stripe_product.product.id
  currency    = "usd"
  unit_amount = 2200
}

resource "stripe_payment_link" "test" {
  active            = true
  customer_creation = "always"
  inactive_message  = "sdk-codegen payment link create message"

  after_completion = {
    type = "redirect"
    redirect = {
      url = "https://example.com/sdk-codegen/payment-link/create"
    }
  }

  line_items = [
    {
      price    = stripe_price.price.id
      quantity = 1
    },
  ]

  metadata = {
    suite = "sdk-codegen"
    case  = "payment_link_basic"
    phase = "create"
  }
}
