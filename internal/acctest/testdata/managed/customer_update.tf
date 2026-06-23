# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "test" {
  name = "acctest-customer-{{RAND}}"
  email = "sdk-codegen+{{RAND}}@example.com"
  description = "updated by sdk-codegen acceptance"

  metadata = {
    suite = "sdk-codegen"
    case  = "customer_basic"
    phase = "update"
  }
}
