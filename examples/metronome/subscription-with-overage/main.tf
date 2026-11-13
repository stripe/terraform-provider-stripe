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

variable "credit_type_id" {
  description = "ID of an existing Metronome USD_CENTS credit type."
  type        = string
}

variable "rate_start_at" {
  description = "UTC hour-boundary start for the usage rate, for example 2026-10-01T00:00:00.000Z."
  type        = string
}

resource "stripe_metronome_billable_metric" "api_requests" {
  name       = "Overage example API requests"
  event_type = "api_request"
}

resource "stripe_metronome_product" "api_requests" {
  name               = "Overage example API requests"
  type               = "USAGE"
  billable_metric_id = stripe_metronome_billable_metric.api_requests.id
}

resource "stripe_metronome_product" "prepaid_balance" {
  name = "Overage example prepaid balance"
  type = "FIXED"
}

# Usage costs one cent per request. After the monthly credit is spent,
# additional usage is billed at this rate.
resource "stripe_metronome_rate_card" "usd" {
  name             = "Overage example USD"
  credit_type_id   = var.credit_type_id
  usage_product_id = stripe_metronome_product.api_requests.id
  price_cents      = 1
  starting_at      = var.rate_start_at
}

resource "stripe_metronome_package" "free" {
  name           = "Overage example Free"
  contract_name  = "Free"
  rate_card_id   = stripe_metronome_rate_card.usd.id
  uniqueness_key = "package:example:overage:free:v1"
  stripe_billing = false
}

# $20 each month grants $25 of usage credit. At one cent per request, that
# covers 2,500 requests; 3,000 requests would add $5 of usage overage.
resource "stripe_metronome_package" "pro" {
  name           = "Overage example Pro"
  contract_name  = "Pro"
  rate_card_id   = stripe_metronome_rate_card.usd.id
  uniqueness_key = "package:example:overage:pro:v1"
  stripe_billing = true
  recurring_commit = {
    name                   = "Pro monthly balance"
    fixed_product_id       = stripe_metronome_product.prepaid_balance.id
    credit_type_id         = var.credit_type_id
    monthly_price_cents    = 2000
    monthly_grant_cents    = 2500
    applicable_product_ids = [stripe_metronome_product.api_requests.id]
    grant_duration_periods = 1
    rollover_fraction      = 0
    priority               = 10
  }
}

output "rate_card_id" {
  value = stripe_metronome_rate_card.usd.id
}

output "free_package_id" {
  value = stripe_metronome_package.free.id
}

output "pro_package_id" {
  value = stripe_metronome_package.pro.id
}
