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

variable "rate_card_id" {
  type        = string
  description = "An existing rate card ID. A package is not needed for this contract."
}

variable "start_at" {
  type        = string
  description = "A stable UTC hour when this contract begins."
}

resource "stripe_metronome_customer" "enterprise" {
  name                = "Direct contract example"
  ingest_aliases      = ["direct-contract-example"]
  deletion_protection = true
}

resource "stripe_metronome_contract" "enterprise" {
  customer_id              = stripe_metronome_customer.enterprise.id
  rate_card_id             = var.rate_card_id
  name                     = "Direct rate card contract"
  starting_at              = var.start_at
  uniqueness_key           = "direct-contract-example:v1"
  void_invoices_on_destroy = false
}

output "customer_id" {
  value = stripe_metronome_customer.enterprise.id
}

output "contract_id" {
  value = stripe_metronome_contract.enterprise.id
}
