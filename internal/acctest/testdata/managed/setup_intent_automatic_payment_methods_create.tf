# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_setup_intent" "test" {
  description = "sdk-codegen setup intent automatic payment methods"
  usage       = "off_session"

  automatic_payment_methods = {
    enabled         = true
    allow_redirects = "never"
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "setup_intent_automatic_payment_methods"
  }
}
