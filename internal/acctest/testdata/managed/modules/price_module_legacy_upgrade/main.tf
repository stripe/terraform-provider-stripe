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

resource "stripe_product" "product" {
  name = "acctest-price-tiered-product-${var.rand}"
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
    unit_amount = 900
    up_to       = "10"
  }

  tiers {
    unit_amount = 700
    up_to       = "inf"
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "price_module_legacy_upgrade"
  }
}

output "price_product_id" {
  value = stripe_price.test.product
}

output "price_recurring_interval" {
  value = stripe_price.test.recurring[0].interval
}

output "price_tier_unit_amount" {
  value = stripe_price.test.tiers[0].unit_amount
}
