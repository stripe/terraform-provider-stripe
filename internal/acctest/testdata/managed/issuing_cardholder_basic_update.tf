# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_issuing_cardholder" "test" {
  type         = "individual"
  name         = "SDK Cardholder Create"
  email        = "sdk-codegen+cardholder-update@example.com"
  phone_number = "+15555550152"
  status       = "inactive"

  billing = {
    address = {
      line1       = "100 Main St"
      city        = "San Francisco"
      postal_code = "94105"
      country     = "US"
      state       = "CA"
    }
  }

  individual = {
    first_name = "Grace"
    last_name  = "Hopper"
    dob = {
      day   = 9
      month = 12
      year  = 1985
    }
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "issuing_cardholder_basic"
    phase = "update"
  }
}
