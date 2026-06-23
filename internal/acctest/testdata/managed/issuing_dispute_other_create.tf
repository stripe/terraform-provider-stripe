# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_issuing_dispute" "test" {
  transaction = "{{ISSUING_DISPUTE_OTHER_TRANSACTION}}"

  evidence = {
    reason = "other"
    other = {
      explanation         = "sdk-codegen issuing dispute other create"
      product_description = "SDK Codegen other dispute create"
      product_type        = "service"
    }
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "issuing_dispute_other"
    phase = "create"
  }
}
