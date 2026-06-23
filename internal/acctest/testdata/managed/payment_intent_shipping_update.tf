# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_payment_intent" "test" {
  amount        = 3600
  currency      = "usd"
  receipt_email = "sdk-codegen+shipping-updated@example.com"
  statement_descriptor_suffix = "SHIPB"

  automatic_payment_methods = {
    enabled = true
  }

  shipping = {
    name            = "Shipping Update"
    phone           = "+15555550102"
    carrier         = "FedEx"
    tracking_number = "TRACK-UPDATE-002"
    address = {
      line1       = "200 Mission St"
      city        = "San Francisco"
      state       = "CA"
      postal_code = "94105"
      country     = "US"
    }
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "payment_intent_shipping"
    phase = "update"
  }
}
