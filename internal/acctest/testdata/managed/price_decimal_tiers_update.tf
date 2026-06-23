# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "product" {
  name = "acctest-price-decimal-tiers-product-{{RAND}}"
}

resource "stripe_price" "test" {
  product        = stripe_product.product.id
  currency       = "usd"
  billing_scheme = "tiered"
  tiers_mode     = "graduated"

  recurring {
    interval = "month"
  }

  tiers {
    unit_amount_decimal = "950.00"
    up_to               = "10"
  }

  tiers {
    flat_amount_decimal = "175.0000"
    up_to               = "inf"
  }

  currency_options {
    key = "eur"

    tiers = [
      {
        unit_amount_decimal = "850.00"
        up_to               = "10"
      },
      {
        flat_amount_decimal = "140.0000"
        up_to               = "inf"
      },
    ]
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "price_decimal_tiers_update"
  }
}
