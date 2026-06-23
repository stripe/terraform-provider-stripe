# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_payment_method_configuration" "test" {
  active = false
  name   = "SDK Codegen PM Configuration Update"

  card = {
    display_preference = {
      preference = "off"
    }
  }
}
