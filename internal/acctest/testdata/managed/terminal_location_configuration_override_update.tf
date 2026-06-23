# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_terminal_configuration" "override" {
  name = "SDK Codegen Terminal Override Config Update"

  offline = {
    enabled = false
  }

  cellular = {
    enabled = true
  }

  reboot_window = {
    start_hour = 11
    end_hour   = 12
  }
}

resource "stripe_terminal_location" "test" {
  display_name             = "SDK Codegen Terminal Override Update"
  phone                    = "+15555550144"
  configuration_overrides  = stripe_terminal_configuration.override.id

  address = {
    line1       = "400 Spear St"
    line2       = "Floor 9"
    city        = "San Francisco"
    state       = "CA"
    postal_code = "94105"
    country     = "US"
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "terminal_location_configuration_override"
    phase = "update"
  }
}
