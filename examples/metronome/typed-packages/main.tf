//
// File generated from our OpenAPI spec
//

terraform {
  required_version = ">= 1.16.5"
  required_providers {
    stripe = {
      source = "stripe/stripe"
    }
  }
}

provider "stripe" {}

variable "rate_card_id" { type = string }
variable "usage_product_id" { type = string }
variable "fixed_product_id" { type = string }
variable "subscription_product_id" { type = string }
variable "credit_type_id" { type = string }

resource "stripe_metronome_package" "usage_only" {
  package = {
    name           = "Usage only"
    uniqueness_key = "package:typed-example-usage-v1"
    rate_card_id   = var.rate_card_id
    usage_statement_schedule = {
      frequency = "MONTHLY"
      day       = "CONTRACT_START"
    }
  }
}

resource "stripe_metronome_package" "subscription" {
  package = {
    name             = "Seat subscription"
    uniqueness_key   = "package:typed-example-seats-v1"
    rate_card_id     = var.rate_card_id
    billing_provider = "stripe"
    delivery_method  = "direct_to_billing_provider"
    subscriptions = [{
      subscription_rate = {
        product_id        = var.subscription_product_id
        billing_frequency = "MONTHLY"
      }
      collection_schedule = "ADVANCE"
      initial_quantity    = 5
      proration = {
        is_prorated      = false
        invoice_behavior = "BILL_IMMEDIATELY"
      }
    }]
  }
}

resource "stripe_metronome_package" "monthly_credit" {
  package = {
    name           = "Monthly credit"
    uniqueness_key = "package:typed-example-credit-v1"
    rate_card_id   = var.rate_card_id
    recurring_credits = [{
      name       = "Monthly grant"
      product_id = var.fixed_product_id
      priority   = 10
      access_amount = {
        unit_price     = 2500
        quantity       = 1
        credit_type_id = var.credit_type_id
      }
      starting_at_offset     = { value = 0, unit = "DAYS" }
      commit_duration        = { value = 1, unit = "PERIODS" }
      applicable_product_ids = [var.usage_product_id]
      recurrence_frequency   = "MONTHLY"
    }]
  }
}

resource "stripe_metronome_package" "mixed" {
  package = {
    name             = "Subscription with overage"
    uniqueness_key   = "package:typed-example-mixed-v1"
    rate_card_id     = var.rate_card_id
    billing_provider = "stripe"
    delivery_method  = "direct_to_billing_provider"
    recurring_commits = [{
      name       = "Monthly prepaid balance"
      product_id = var.fixed_product_id
      priority   = 10
      access_amount = {
        unit_price     = 2500
        quantity       = 1
        credit_type_id = var.credit_type_id
      }
      invoice_amount = {
        unit_price     = 2000
        quantity       = 1
        credit_type_id = var.credit_type_id
      }
      starting_at_offset     = { value = 0, unit = "DAYS" }
      commit_duration        = { value = 1, unit = "PERIODS" }
      applicable_product_ids = [var.usage_product_id]
      recurrence_frequency   = "MONTHLY"
    }]
  }
}
