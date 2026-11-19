terraform {
  required_version = ">= 1.16.5"
  required_providers {
    stripe = { source = "stripe/stripe" }
  }
}

data "stripe_metronome_list_credit_types_v1" "existing" {}

resource "stripe_metronome_credit_type" "credits" {
  name = "module credits"
}
