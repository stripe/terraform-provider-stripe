# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_person" "test" {
  account    = "{{STRIPE_ACCOUNT}}"
  first_name = "Ada"
  last_name  = "Lovelace"
  email      = "sdk-codegen+person-create@example.com"
  phone      = "+15555550131"

  dob = {
    day   = 10
    month = 12
    year  = 1990
  }

  address = {
    line1       = "100 Market St"
    city        = "San Francisco"
    state       = "CA"
    postal_code = "94105"
    country     = "US"
  }

  relationship = {
    executive      = true
    director       = true
    representative = false
    title          = "Chief Architect"
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "person_basic"
    phase = "create"
  }
}
