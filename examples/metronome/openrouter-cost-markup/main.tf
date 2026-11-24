//
// File generated from our OpenAPI spec
//

terraform {
  required_providers {
    stripe = {
      source = "stripe/stripe"
    }
  }
}

provider "stripe" {}

variable "markup_percent" {
  description = "Whole percentage markup on the event's upstream cost. 20 bills $1.20 for $1.00 of cost."
  type        = number
  default     = 20
  validation {
    condition     = var.markup_percent >= 0 && floor(var.markup_percent) == var.markup_percent
    error_message = "markup_percent must be a nonnegative whole percentage."
  }
}

# Each openrouter_usage event contributes its cost_usd property to usage.
# The input property is a decimal string, e.g. "0.0000064".
resource "stripe_metronome_billable_metric" "upstream_cost" {
  name             = "OpenRouter upstream cost USD"
  event_type       = "openrouter_usage"
  aggregation_type = "SUM"
  aggregation_key  = "cost_usd"
}

resource "stripe_metronome_product" "openrouter_usage" {
  name               = "OpenRouter upstream cost"
  type               = "USAGE"
  billable_metric_id = stripe_metronome_billable_metric.upstream_cost.id
}

# The monthly Pro grant is booked against this separate fixed product.
resource "stripe_metronome_product" "prepaid_balance" {
  name = "OpenRouter prepaid balance"
  type = "FIXED"
}

# The SUM metric's unit is one USD of upstream cost. Metronome prices in cents,
# so $1 of upstream cost is 100 cents before markup. At 20%, the rate is 120
# cents per $1 upstream cost; fractional event costs accumulate as usage.
resource "stripe_metronome_rate_card" "openrouter_markup" {
  name             = "OpenRouter cost plus ${var.markup_percent}% markup"
  usage_product_id = stripe_metronome_product.openrouter_usage.id
  price_cents      = 100 + var.markup_percent
}

# Usage-only customers pay the same markup on each matched event.
# Billing through Stripe needs the account integration connected before
# customer invoicing.
resource "stripe_metronome_package" "metered" {
  name           = "OpenRouter metered ${var.markup_percent}% markup"
  contract_name  = "OpenRouter metered"
  rate_card_id   = stripe_metronome_rate_card.openrouter_markup.id
  uniqueness_key = "package:openrouter:metered:markup-${var.markup_percent}:v1"
  stripe_billing = true
}

# Pro charges $20 monthly, grants $25 of balance for this usage product, and
# bills further usage at the same marked-up rate after the grant is spent.
resource "stripe_metronome_package" "pro" {
  name           = "OpenRouter Pro ${var.markup_percent}% markup"
  contract_name  = "OpenRouter Pro"
  rate_card_id   = stripe_metronome_rate_card.openrouter_markup.id
  uniqueness_key = "package:openrouter:pro:markup-${var.markup_percent}:v1"
  stripe_billing = true
  recurring_commit = {
    name                   = "Pro monthly balance"
    fixed_product_id       = stripe_metronome_product.prepaid_balance.id
    credit_type_id         = stripe_metronome_rate_card.openrouter_markup.credit_type_id
    monthly_price_cents    = 2000
    monthly_grant_cents    = 2500
    applicable_product_ids = [stripe_metronome_product.openrouter_usage.id]
    grant_duration_periods = 1
    rollover_fraction      = 0
    priority               = 10
  }
}

output "billable_metric_id" {
  value = stripe_metronome_billable_metric.upstream_cost.id
}

output "usage_product_id" {
  value = stripe_metronome_product.openrouter_usage.id
}

output "prepaid_balance_product_id" {
  value = stripe_metronome_product.prepaid_balance.id
}

output "rate_card_id" {
  value = stripe_metronome_rate_card.openrouter_markup.id
}

output "metered_package_id" {
  value = stripe_metronome_package.metered.id
}

output "pro_package_id" {
  value = stripe_metronome_package.pro.id
}
