# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_shipping_rate" "test" {
  display_name = "SDK Codegen Shipping Currency Options"
  active       = true
  type         = "fixed_amount"

  fixed_amount {
    amount   = 500
    currency = "usd"

    currency_options {
      key          = "eur"
      amount       = 450
      tax_behavior = "inclusive"
    }
  }

  delivery_estimate {
    minimum {
      unit  = "business_day"
      value = 1
    }
    maximum {
      unit  = "business_day"
      value = 3
    }
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "shipping_rate_currency_options"
    phase = "create"
  }
}
