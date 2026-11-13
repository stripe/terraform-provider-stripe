//
// File generated from our OpenAPI spec
//

terraform {
  required_providers {
    stripe = { source = "stripe/stripe" }
  }
}

provider "stripe" {}

variable "rate_start" {
  description = "Stable UTC effective rate start."
  type        = string
}

resource "stripe_metronome_billable_metric" "requests" {
  name       = "Composite example requests"
  event_type = "composite_example_request"
}

resource "stripe_metronome_product" "usage" {
  name               = "Composite example usage"
  type               = "USAGE"
  billable_metric_id = stripe_metronome_billable_metric.requests.id
}

resource "stripe_metronome_product" "fee" {
  name                  = "Composite example fee"
  type                  = "COMPOSITE"
  composite_product_ids = [stripe_metronome_product.usage.id]
  composite_tags        = ["compute"]
}

resource "stripe_metronome_rate_card" "usd" {
  name = "Composite example rates"
  rates = [
    {
      product_id  = stripe_metronome_product.usage.id
      rate_type   = "FLAT"
      entitled    = true
      price       = 1
      starting_at = var.rate_start
    },
    {
      product_id  = stripe_metronome_product.fee.id
      rate_type   = "PERCENTAGE"
      entitled    = true
      price       = 0.1
      starting_at = var.rate_start
    }
  ]
}

resource "stripe_metronome_package" "metered" {
  name           = "Composite example metered"
  contract_name  = "Composite example metered"
  uniqueness_key = "package:example:composite-percentage:v1"
  rate_card_id   = stripe_metronome_rate_card.usd.id
  stripe_billing = false
}
