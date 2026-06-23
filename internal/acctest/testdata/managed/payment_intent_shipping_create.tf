# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_payment_intent" "test" {
  amount        = 3600
  currency      = "usd"
  receipt_email = "sdk-codegen+shipping@example.com"
  statement_descriptor_suffix = "SHIPA"

  automatic_payment_methods = {
    enabled = true
  }

  shipping = {
    name            = "Shipping Create"
    phone           = "+15555550101"
    carrier         = "UPS"
    tracking_number = "TRACK-CREATE-001"
    address = {
      line1       = "100 Market St"
      city        = "San Francisco"
      state       = "CA"
      postal_code = "94105"
      country     = "US"
    }
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "payment_intent_shipping"
    phase = "create"
  }
}
