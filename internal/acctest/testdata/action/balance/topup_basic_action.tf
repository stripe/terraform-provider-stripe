# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
action "stripe_topup" "test" {
  config {
    amount         = 5100
    currency       = "usd"
    description    = "sdk-codegen topup basic"
    transfer_group = "sdk-codegen-topup"

    metadata = {
      suite = "sdk-codegen"
      case  = "topup_basic"
    }
  }
}
