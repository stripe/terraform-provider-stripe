# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_terminal_location" "test" {
  display_name = "SDK Codegen Terminal Location Create"
  phone        = "+15555550141"

  address = {
    line1       = "100 Market St"
    city        = "San Francisco"
    state       = "CA"
    postal_code = "94105"
    country     = "US"
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "terminal_location_basic"
    phase = "create"
  }
}
