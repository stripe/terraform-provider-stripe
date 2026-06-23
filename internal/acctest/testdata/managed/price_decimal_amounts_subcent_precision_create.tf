# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "product" {
  name = "acctest-price-decimal-subcent-product-{{RAND}}"
}

resource "stripe_price" "test" {
  product             = stripe_product.product.id
  currency            = "usd"
  unit_amount_decimal = "0.0015000000"

  recurring {
    interval = "month"
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "price_decimal_amounts_subcent_precision"
  }
}
