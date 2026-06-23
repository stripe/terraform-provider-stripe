# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "product" {
  name = "acctest-price-decimal-amounts-product-{{RAND}}"
}

resource "stripe_price" "test" {
  product             = stripe_product.product.id
  currency            = "usd"
  unit_amount_decimal = "1700.00"

  recurring {
    interval = "month"
  }

  currency_options {
    key                 = "eur"
    unit_amount_decimal = "1400.2500"
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "price_decimal_amounts_update"
  }
}
