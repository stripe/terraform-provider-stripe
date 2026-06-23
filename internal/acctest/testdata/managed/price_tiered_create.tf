# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "product" {
  name = "acctest-price-tiered-product-{{RAND}}"
}

resource "stripe_price" "test" {
  product         = stripe_product.product.id
  currency        = "usd"
  billing_scheme  = "tiered"
  tiers_mode      = "graduated"

  recurring {
    interval = "month"
  }

  tiers {
    unit_amount = 900
    up_to       = "10"
  }

  tiers {
    unit_amount = 700
    up_to       = "inf"
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "price_tiered"
  }
}
