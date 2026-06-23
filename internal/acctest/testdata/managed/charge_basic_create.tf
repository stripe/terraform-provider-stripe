# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_charge" "test" {
  amount        = 2200
  currency      = "usd"
  description   = "sdk-codegen charge basic"
  receipt_email = "sdk-codegen+charge-create@example.com"
  source        = "tok_visa"

  metadata = {
    suite = "sdk-codegen"
    case  = "charge_basic"
    phase = "create"
  }
}
