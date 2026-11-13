//
// File generated from our OpenAPI spec
//

terraform {
  required_version = ">= 1.16.5"
  required_providers {
    stripe = { source = "stripe/stripe" }
  }
}

# Supply the shared key through STRIPE_API_KEY.
provider "stripe" {}

resource "stripe_product" "service" {
  name = "Metered service"
}

resource "stripe_metronome_account_bridge" "billing" {}

module "billing" {
  source     = "./billing"
  depends_on = [stripe_metronome_account_bridge.billing]
}
