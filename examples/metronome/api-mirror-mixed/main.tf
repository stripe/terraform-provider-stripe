//
// File generated from our OpenAPI spec
//

terraform {
  required_providers {
    stripe = { source = "stripe/stripe" }
  }
}

provider "stripe" {}

variable "markup_percent" {
  description = "Markup on the upstream cost reported by each OpenRouter event."
  type        = number
  default     = 30
}

variable "rate_start" {
  description = "Stable ISO 8601 effective start for both rates, e.g. 2026-10-02T16:00:00Z."
  type        = string
}

resource "stripe_metronome_billable_metric" "openrouter_cost" {
  name             = "OpenRouter upstream USD cost"
  event_type       = "openrouter_usage"
  aggregation_type = "SUM"
  aggregation_key  = "cost_usd"
}

resource "stripe_metronome_product" "usage" {
  name               = "OpenRouter upstream cost"
  type               = "USAGE"
  billable_metric_id = stripe_metronome_billable_metric.openrouter_cost.id
}

resource "stripe_metronome_product" "subscription" {
  name = "OpenRouter platform subscription"
  type = "SUBSCRIPTION"
}

# This FIXED product labels the separate prepaid commit and recurring credit.
# It is not needed for the subscription seat charge or metered usage.
resource "stripe_metronome_product" "balance" {
  name = "OpenRouter prepaid balance"
  type = "FIXED"
}

# A generated read data source can inspect any operation in the pinned API.
data "stripe_metronome_get_product_v1" "usage" {
  request = { id = stripe_metronome_product.usage.id }
}

# The first rate charges 130 cents per $1 of measured upstream cost at 30%.
# The second rate charges $10 per month for one subscription seat.
resource "stripe_metronome_rate_card" "mixed" {
  name = "OpenRouter mixed plan, ${var.markup_percent}% cost markup"
  rates = [
    {
      product_id  = stripe_metronome_product.usage.id
      rate_type   = "FLAT"
      entitled    = true
      price       = 100 + var.markup_percent
      starting_at = var.rate_start
    },
    {
      product_id        = stripe_metronome_product.subscription.id
      rate_type         = "FLAT"
      entitled          = true
      price             = 1000
      billing_frequency = "MONTHLY"
      starting_at       = var.rate_start
    }
  ]
}

# This request is shaped exactly like Metronome's createPackage payload:
# a $10 subscription, a $20 monthly commit with $25 of usage access, plus
# a separate $5 recurring credit. Usage beyond $30 of balance is billed at
# the usage rate above. The invoice and access amounts are deliberately distinct.
resource "stripe_metronome_package" "mixed" {
  request = {
    name             = "OpenRouter mixed subscription with overage"
    contract_name    = "OpenRouter mixed"
    uniqueness_key   = "openrouter:mixed:markup-${var.markup_percent}:v1"
    rate_card_id     = stripe_metronome_rate_card.mixed.id
    billing_provider = "stripe"
    delivery_method  = "direct_to_billing_provider"
    usage_statement_schedule = {
      frequency = "MONTHLY"
      day       = "CONTRACT_START"
    }
    subscriptions = [{
      name                = "Platform subscription"
      subscription_rate   = { product_id = stripe_metronome_product.subscription.id, billing_frequency = "MONTHLY" }
      collection_schedule = "ADVANCE"
      proration           = { is_prorated = false, invoice_behavior = "BILL_ON_NEXT_COLLECTION_DATE" }
      initial_quantity    = 1
    }]
    recurring_commits = [{
      name                   = "Monthly prepaid commit"
      product_id             = stripe_metronome_product.balance.id
      priority               = 10
      applicable_product_ids = [stripe_metronome_product.usage.id]
      access_amount          = { unit_price = 2500, quantity = 1, credit_type_id = stripe_metronome_rate_card.mixed.credit_type_id }
      invoice_amount         = { unit_price = 2000, quantity = 1, credit_type_id = stripe_metronome_rate_card.mixed.credit_type_id }
      starting_at_offset     = { value = 0, unit = "DAYS" }
      commit_duration        = { value = 1, unit = "PERIODS" }
      recurrence_frequency   = "MONTHLY"
      proration              = "NONE"
    }]
    recurring_credits = [{
      name                   = "Monthly bonus credit"
      product_id             = stripe_metronome_product.balance.id
      priority               = 20
      applicable_product_ids = [stripe_metronome_product.usage.id]
      access_amount          = { unit_price = 500, quantity = 1, credit_type_id = stripe_metronome_rate_card.mixed.credit_type_id }
      starting_at_offset     = { value = 0, unit = "DAYS" }
      commit_duration        = { value = 1, unit = "PERIODS" }
      recurrence_frequency   = "MONTHLY"
      proration              = "NONE"
    }]
  }
}

output "rate_card_id" { value = stripe_metronome_rate_card.mixed.id }
output "package_id" { value = stripe_metronome_package.mixed.id }
