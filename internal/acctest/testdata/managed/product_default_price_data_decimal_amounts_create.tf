# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "test" {
  name        = "acctest-product-default-price-data-decimal-amounts-{{RAND}}"
  description = "decimal amount create"

  default_price_data {
    currency            = "usd"
    unit_amount_decimal = "1500.0"

    recurring {
      interval = "month"
    }

    currency_options {
      key                 = "eur"
      unit_amount_decimal = "1300.5000"
    }
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "product_default_price_data_decimal_amounts"
  }
}
