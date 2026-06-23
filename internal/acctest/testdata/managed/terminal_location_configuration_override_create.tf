# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_terminal_configuration" "override" {
  name = "SDK Codegen Terminal Override Config Create"

  offline = {
    enabled = true
  }

  cellular = {
    enabled = false
  }

  reboot_window = {
    start_hour = 9
    end_hour   = 10
  }
}

resource "stripe_terminal_location" "test" {
  display_name             = "SDK Codegen Terminal Override Create"
  phone                    = "+15555550143"
  configuration_overrides  = stripe_terminal_configuration.override.id

  address = {
    line1       = "300 Howard St"
    line2       = "Suite 400"
    city        = "San Francisco"
    state       = "CA"
    postal_code = "94105"
    country     = "US"
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "terminal_location_configuration_override"
    phase = "create"
  }
}
