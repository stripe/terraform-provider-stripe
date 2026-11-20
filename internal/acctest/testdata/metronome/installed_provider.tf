terraform {
  required_version = ">= 1.15.4"
  required_providers {
    stripe = {
      source  = "stripe/stripe"
      version = "{{VERSION}}"
    }
  }
}

provider "stripe" {}

resource "stripe_metronome_credit_type" "credits" {
  name = "validate without credentials"
}
