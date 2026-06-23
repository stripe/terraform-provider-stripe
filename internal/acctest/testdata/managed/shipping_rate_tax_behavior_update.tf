# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_shipping_rate" "test" {
  display_name = "SDK Codegen Shipping Tax Update"
  active       = true
  type         = "fixed_amount"
  tax_behavior = "inclusive"
  tax_code     = "txcd_92010001"

  fixed_amount {
    amount   = 900
    currency = "usd"
  }

  delivery_estimate {
    minimum {
      unit  = "business_day"
      value = 1
    }
    maximum {
      unit  = "business_day"
      value = 4
    }
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "shipping_rate_tax_behavior"
    phase = "update"
  }
}
