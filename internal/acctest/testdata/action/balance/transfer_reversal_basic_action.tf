# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
action "stripe_transfer_reversal" "test" {
  config {
    transfer = "{{TRANSFER_ID}}"
    amount   = 40

    metadata = {
      suite = "sdk-codegen"
      case  = "transfer_reversal_basic"
    }
  }
}
