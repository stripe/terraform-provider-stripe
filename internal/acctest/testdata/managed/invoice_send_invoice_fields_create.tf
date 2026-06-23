# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name  = "acctest-invoice-send-customer-{{RAND}}"
  email = "sdk-codegen+invoice-send-{{RAND}}@example.com"
}

resource "stripe_invoice" "test" {
  customer          = stripe_customer.customer.id
  auto_advance      = false
  collection_method = "send_invoice"
  due_date          = 1893456000
  description       = "sdk-codegen invoice send invoice create"
  footer            = "SDK Codegen invoice footer create"

  custom_fields = [
    {
      name  = "order_id"
      value = "INV-CREATE"
    },
    {
      name  = "region"
      value = "NA"
    },
  ]

  metadata = {
    suite = "sdk-codegen"
    case  = "invoice_send_invoice_fields"
    phase = "create"
  }
}
