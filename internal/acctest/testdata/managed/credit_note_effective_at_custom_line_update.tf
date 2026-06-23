# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_credit_note" "test" {
  invoice      = "{{CREDIT_NOTE_INVOICE}}"
  effective_at = {{CREDIT_NOTE_INVOICE_FINALIZED_AT}} + 2
  amount       = 150
  memo         = "sdk-codegen credit note custom line update"
  reason       = "product_unsatisfactory"

  metadata = {
    suite = "sdk-codegen"
    case  = "credit_note_effective_at_custom_line"
    phase = "update"
  }
}
