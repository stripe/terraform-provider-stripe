terraform {
  required_version = ">= 1.16.5"
  required_providers {
    stripe = { source = "stripe/stripe" }
  }
}

resource "stripe_metronome_credit_type" "credits" {
  name = "module credits"
}
