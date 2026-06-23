# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_shipping_rate" "test" {
  display_name = "SDK Codegen Shipping Tax Create"
  active       = true
  type         = "fixed_amount"
  tax_behavior = "exclusive"
  tax_code     = "txcd_92010001"

  fixed_amount {
    amount   = 700
    currency = "usd"
  }

  delivery_estimate {
    minimum {
      unit  = "day"
      value = 2
    }
    maximum {
      unit  = "day"
      value = 5
    }
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "shipping_rate_tax_behavior"
    phase = "create"
  }
}
