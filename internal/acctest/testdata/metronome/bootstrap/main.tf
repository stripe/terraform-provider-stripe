terraform {
  required_version = ">= 1.16.5"
  required_providers {
    stripe = { source = "stripe/stripe" }
  }
}

provider "stripe" {
  alias = "billing"
  metronome = {
    base_url     = "{{API_URL}}"
    ca_cert_file = "{{CA_CERT_FILE}}"
  }
}

resource "stripe_metronome_account_bridge" "billing" {
  provider = stripe.billing
}

module "billing" {
  source     = "./billing"
  providers  = { stripe = stripe.billing }
  depends_on = [stripe_metronome_account_bridge.billing]
}
