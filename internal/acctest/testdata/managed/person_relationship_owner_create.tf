# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_person" "test" {
  account             = "{{STRIPE_ACCOUNT}}"
  first_name          = "Katherine"
  last_name           = "Johnson"
  email               = "sdk-codegen+person-owner-create@example.com"
  phone               = "+15555550133"
  nationality         = "US"
  political_exposure  = "none"

  dob = {
    day   = 26
    month = 8
    year  = 1918
  }

  address = {
    line1       = "500 Howard St"
    line2       = "Suite 700"
    city        = "San Francisco"
    state       = "CA"
    postal_code = "94105"
    country     = "US"
  }

  relationship = {
    executive         = false
    director          = false
    representative    = false
    owner             = true
    percent_ownership = 25
    title             = "Managing Member"
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "person_relationship_owner"
    phase = "create"
  }
}
