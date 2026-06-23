# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_charge" "charge" {
  amount      = 400
  currency    = "usd"
  description = "sdk-codegen transfer funding"
  source      = "tok_visa"
}
