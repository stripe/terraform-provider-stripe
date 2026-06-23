# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_charge" "charge" {
  amount      = 400
  currency    = "usd"
  description = "sdk-codegen transfer funding"
  source      = "tok_visa"
}

action "stripe_transfer" "test" {
  config {
    amount             = 110
    currency           = "usd"
    destination        = "{{STRIPE_ACCOUNT}}"
    source_transaction = stripe_charge.charge.id
    description        = "sdk-codegen transfer basic"
    transfer_group     = "sdk-codegen-transfer"

    metadata = {
      suite = "sdk-codegen"
      case  = "transfer_basic"
    }
  }
}
