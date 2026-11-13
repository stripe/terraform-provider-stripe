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

variable "monthly_price_cents" {
  description = "Monthly platform subscription price in USD cents."
  type        = number
  default     = 2000
  validation {
    condition     = var.monthly_price_cents >= 0 && floor(var.monthly_price_cents) == var.monthly_price_cents
    error_message = "monthly_price_cents must be a nonnegative whole number."
  }
}

variable "rate_start" {
  description = "Stable UTC hour when subscription pricing becomes effective."
  type        = string
  default     = "2026-01-01T00:00:00.000Z"
}

resource "stripe_metronome_product" "platform" {
  name = "Example subscription-only platform"
  type = "SUBSCRIPTION"
}

resource "stripe_metronome_rate_card" "subscription" {
  name = "Example subscription-only USD"
  rates = [{
    product_id        = stripe_metronome_product.platform.id
    rate_type         = "FLAT"
    entitled          = true
    price             = var.monthly_price_cents
    billing_frequency = "MONTHLY"
    starting_at       = var.rate_start
  }]
}

# This is a monthly subscription charge. There is no metered usage or grant.
# The application creates customer contracts from this package at runtime.
resource "stripe_metronome_package" "subscription" {
  request = {
    name             = "Example subscription-only"
    contract_name    = "Platform subscription"
    uniqueness_key   = "example:subscription-only:v1"
    rate_card_id     = stripe_metronome_rate_card.subscription.id
    billing_provider = "stripe"
    delivery_method  = "direct_to_billing_provider"
    usage_statement_schedule = {
      frequency = "MONTHLY"
      day       = "CONTRACT_START"
    }
    subscriptions = [{
      name                = "Platform subscription"
      subscription_rate   = { product_id = stripe_metronome_product.platform.id, billing_frequency = "MONTHLY" }
      collection_schedule = "ADVANCE"
      proration           = { is_prorated = false, invoice_behavior = "BILL_ON_NEXT_COLLECTION_DATE" }
      initial_quantity    = 1
    }]
  }
}

output "package_id" {
  value = stripe_metronome_package.subscription.id
}

output "subscription_product_id" {
  value = stripe_metronome_product.platform.id
}

output "rate_card_id" {
  value = stripe_metronome_rate_card.subscription.id
}
