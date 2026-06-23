# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_payment_method_configuration" "test" {
  active = true
  name   = "SDK Codegen PM Configuration Wallets"

  card = {
    display_preference = {
      preference = "on"
    }
  }

  apple_pay = {
    display_preference = {
      preference = "off"
    }
  }

  link = {
    display_preference = {
      preference = "off"
    }
  }
}
