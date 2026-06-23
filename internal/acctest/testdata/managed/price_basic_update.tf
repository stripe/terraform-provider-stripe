# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "product" {
  name = "acctest-price-product-{{RAND}}"
}

resource "stripe_price" "test" {
  product     = stripe_product.product.id
  currency    = "usd"
  unit_amount = 1500

  recurring {
    interval = "month"
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "price_basic"
    phase = "update"
  }
}
