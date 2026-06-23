# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
action "stripe_fee_refund" "test" {
  config {
    fee    = "{{APPLICATION_FEE_ID}}"
    amount = 55

    metadata = {
      suite = "sdk-codegen"
      case  = "fee_refund_action_regression"
    }
  }
}
