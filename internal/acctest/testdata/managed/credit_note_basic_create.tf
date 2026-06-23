# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_credit_note" "test" {
  invoice = "{{CREDIT_NOTE_INVOICE}}"
  amount  = 100
  reason  = "duplicate"
  memo    = "sdk-codegen credit note create"

  metadata = {
    suite = "sdk-codegen"
    case  = "credit_note_basic"
    phase = "create"
  }
}
