# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_shipping_rate" "test" {
  display_name = "SDK Codegen Shipping Create"
  type         = "fixed_amount"

  fixed_amount {
    amount   = 500
    currency = "usd"
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
    case  = "shipping_rate_basic"
    phase = "create"
  }
}
