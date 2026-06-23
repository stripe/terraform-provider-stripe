# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_payment_method" "test" {
  type            = "card"
  allow_redisplay = "always"

  card = {
    number    = "4242424242424242"
    exp_month = 11
    exp_year  = 2036
    cvc       = "123"
  }

  billing_details = {
    name  = "SDK Codegen Card Redisplay"
    email = "sdk-codegen+pm-redisplay@example.com"
    phone = "+15555550123"
    address = {
      line1       = "300 Howard St"
      city        = "San Francisco"
      state       = "CA"
      postal_code = "94105"
      country     = "US"
    }
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "payment_method_allow_redisplay"
  }
}
