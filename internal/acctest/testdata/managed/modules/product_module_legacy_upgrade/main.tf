# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
terraform {
  required_providers {
    stripe = {
      source = "stripe/stripe"
    }
  }
}

variable "rand" {
  type = string
}

resource "stripe_product" "test" {
  name        = "acctest-product-default-price-${var.rand}"
  description = "sdk-codegen acceptance product default price create"

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
    case  = "product_module_legacy_upgrade"
  }
}

output "default_price_currency" {
  value = stripe_product.test.default_price_data[0].currency
}

output "default_price_interval" {
  value = stripe_product.test.default_price_data[0].recurring[0].interval
}
