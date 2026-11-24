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

variable "package_id" {
  type        = string
  description = "An existing package ID from a separate catalog Terraform root."
}

variable "start_at" {
  type        = string
  description = "A stable UTC hour when this contract begins."
}

resource "stripe_metronome_customer" "enterprise" {
  name                = "Example enterprise"
  ingest_aliases      = ["example-enterprise-customer"]
  deletion_protection = true
}

resource "stripe_metronome_contract" "enterprise" {
  customer_id              = stripe_metronome_customer.enterprise.id
  package_id               = var.package_id
  starting_at              = var.start_at
  uniqueness_key           = "example-enterprise-contract:v1"
  void_invoices_on_destroy = false
}

output "customer_id" {
  value = stripe_metronome_customer.enterprise.id
}

output "contract_id" {
  value = stripe_metronome_contract.enterprise.id
}
