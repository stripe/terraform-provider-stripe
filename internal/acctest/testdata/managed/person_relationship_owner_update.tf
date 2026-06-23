# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_person" "test" {
  account             = "{{STRIPE_ACCOUNT}}"
  first_name          = "Katherine"
  last_name           = "Johnson"
  email               = "sdk-codegen+person-owner-update@example.com"
  phone               = "+15555550134"
  nationality         = "US"
  political_exposure  = "none"

  dob = {
    day   = 26
    month = 8
    year  = 1918
  }

  address = {
    line1       = "550 Howard St"
    line2       = "Floor 8"
    city        = "San Francisco"
    state       = "CA"
    postal_code = "94105"
    country     = "US"
  }

  relationship = {
    executive         = false
    director          = true
    representative    = false
    owner             = true
    percent_ownership = 35
    title             = "Principal Owner"
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "person_relationship_owner"
    phase = "update"
  }
}
