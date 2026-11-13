//
// File generated from our OpenAPI spec
//

terraform {
  required_providers {
    stripe = { source = "stripe/stripe" }
  }
}

resource "stripe_metronome_credit_type" "usage" {
  name = "Usage credits"
}
