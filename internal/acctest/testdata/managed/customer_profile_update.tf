# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "test" {
  name        = "acctest-customer-profile-{{RAND}}"
  email       = "sdk-codegen+profile-{{RAND}}@example.com"
  description = "sdk-codegen profile customer updated"
  phone       = "+15555550101"

  metadata = {
    suite = "sdk-codegen"
    case  = "customer_profile"
    phase = "update"
  }
}
