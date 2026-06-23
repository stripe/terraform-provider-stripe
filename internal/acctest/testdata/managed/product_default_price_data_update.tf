# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "test" {
  name        = "acctest-product-default-price-{{RAND}}"
  description = "sdk-codegen acceptance product default price update"

  default_price_data {
    currency    = "usd"
    unit_amount = 1500

    recurring {
      interval       = "month"
      interval_count = 1
    }
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "product_default_price_data"
    phase = "update"
  }
}
