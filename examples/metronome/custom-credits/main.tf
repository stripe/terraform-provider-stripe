//
// File generated from our OpenAPI spec
//

terraform {
  required_version = ">= 1.16.5"
  required_providers {
    stripe = { source = "stripe/stripe" }
  }
}

provider "stripe" {}

resource "stripe_metronome_credit_type" "ai" {
  name = "AI credits"
}

resource "stripe_metronome_billable_metric" "upstream_cost" {
  name             = "AI upstream cost"
  event_type       = "openrouter_usage"
  aggregation_type = "SUM"
  aggregation_key  = "cost_usd"
}

resource "stripe_metronome_product" "usage" {
  name               = "AI usage"
  type               = "USAGE"
  billable_metric_id = stripe_metronome_billable_metric.upstream_cost.id
}

resource "stripe_metronome_rate_card" "ai" {
  name = "AI credit pricing"
  # Defaults to USD cents for invoices.
  credit_type_conversions = [{
    custom_credit_type_id  = stripe_metronome_credit_type.ai.id
    fiat_per_custom_credit = 2
  }]
  rates = [{
    product_id     = stripe_metronome_product.usage.id
    credit_type_id = stripe_metronome_credit_type.ai.id
    rate_type      = "FLAT"
    entitled       = true
    price          = 65
    starting_at    = "2026-01-01T00:00:00.000Z"
  }]
}
