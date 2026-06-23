# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_terminal_configuration" "test" {
  name = "SDK Codegen Terminal Tips Create"

  offline = {
    enabled = true
  }

  cellular = {
    enabled = true
  }

  reboot_window = {
    start_hour = 5
    end_hour   = 6
  }

  tipping = {
    usd = {
      fixed_amounts       = [100, 200, 300]
      percentages         = [15, 20, 25]
      smart_tip_threshold = 2000
    }
  }
}
