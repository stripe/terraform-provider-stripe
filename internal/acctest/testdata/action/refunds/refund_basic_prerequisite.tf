# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_charge" "charge" {
  amount      = 550
  currency    = "usd"
  description = "sdk-codegen refund funding"
  source      = "tok_visa"
}
