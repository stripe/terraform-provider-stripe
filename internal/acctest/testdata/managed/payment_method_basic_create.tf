# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_payment_method" "test" {
  type = "card"

  card = {
    number   = "4242424242424242"
    exp_month = 12
    exp_year  = 2035
    cvc       = "123"
  }

  billing_details = {
    name  = "SDK Codegen Card Create"
    email = "sdk-codegen+pm-create@example.com"
    phone = "+15555550121"
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
    case  = "payment_method_basic"
    phase = "create"
  }
}
