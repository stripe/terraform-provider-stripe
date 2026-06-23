# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_terminal_configuration" "test" {
  name = "SDK Codegen Terminal Tips Update"

  offline = {
    enabled = false
  }

  cellular = {
    enabled = true
  }

  reboot_window = {
    start_hour = 7
    end_hour   = 8
  }

  tipping = {
    usd = {
      fixed_amounts       = [150, 250, 350]
      percentages         = [18, 22, 28]
      smart_tip_threshold = 3000
    }
  }
}
