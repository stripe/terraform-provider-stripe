# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_terminal_configuration" "test" {
  name = "SDK Codegen Terminal Configuration Create"

  offline = {
    enabled = true
  }

  cellular = {
    enabled = false
  }

  reboot_window = {
    start_hour = 1
    end_hour   = 2
  }
}
