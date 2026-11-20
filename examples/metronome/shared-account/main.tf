//
// File generated from our OpenAPI spec
//

terraform {
  required_version = ">= 1.15.4"
  required_providers {
    stripe = { source = "stripe/stripe" }
  }
}

# Supply the shared key through STRIPE_API_KEY.
provider "stripe" {}

resource "stripe_product" "service" {
  name = "Metered service"
}

module "billing" {
  source = "./billing"
}
