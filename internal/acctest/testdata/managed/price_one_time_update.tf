# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "product" {
  name = "acctest-price-one-time-product-{{RAND}}"
}

resource "stripe_price" "test" {
  product     = stripe_product.product.id
  currency    = "usd"
  unit_amount = 4800
  lookup_key  = "acctest-price-onetime-{{RAND}}"
  nickname    = "SDK setup fee updated"

  metadata = {
    suite = "sdk-codegen"
    case  = "price_one_time"
    phase = "update"
  }
}
