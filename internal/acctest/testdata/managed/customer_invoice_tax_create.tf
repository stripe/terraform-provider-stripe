# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "test" {
  name  = "acctest-customer-inv-{{RAND}}"
  email = "sdk-codegen+inv-{{RAND}}@example.com"

  invoice_prefix = upper(substr(md5("{{RAND}}"), 0, 12))
  tax_exempt     = "none"

  metadata = {
    suite = "sdk-codegen"
    case  = "customer_invoice_tax"
  }
}
