terraform {
  required_version = ">= 1.15.4"
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

module "billing" {
  source    = "./billing"
  providers = { stripe = stripe.billing }
}
