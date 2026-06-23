# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_terminal_configuration" "test" {
  name = "SDK Codegen Terminal Configuration Update"

  offline = {
    enabled = false
  }

  cellular = {
    enabled = true
  }

  reboot_window = {
    start_hour = 3
    end_hour   = 4
  }
}
