# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name  = "acctest-setup-intent-customer-{{RAND}}"
  email = "sdk-codegen+setup-intent-{{RAND}}@example.com"
}

resource "stripe_setup_intent" "test" {
  customer             = stripe_customer.customer.id
  description          = "sdk-codegen setup intent basic"
  payment_method_types = ["card"]
  usage                = "off_session"

  metadata = {
    suite = "sdk-codegen"
    case  = "setup_intent_basic"
    phase = "create"
  }
}
