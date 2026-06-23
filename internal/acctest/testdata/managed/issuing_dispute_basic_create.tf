# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_issuing_dispute" "test" {
  transaction = "{{ISSUING_DISPUTE_TRANSACTION}}"

  evidence = {
    reason = "fraudulent"
    fraudulent = {
      explanation = "sdk-codegen issuing dispute create"
    }
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "issuing_dispute_basic"
    phase = "create"
  }
}
