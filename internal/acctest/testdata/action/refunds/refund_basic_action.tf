# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_charge" "charge" {
  amount      = 550
  currency    = "usd"
  description = "sdk-codegen refund funding"
  source      = "tok_visa"
}

action "stripe_refund" "test" {
  config {
    charge = stripe_charge.charge.id
    amount = 220
    reason = "requested_by_customer"

    metadata = {
      suite = "sdk-codegen"
      case  = "refund_basic"
    }
  }
}
