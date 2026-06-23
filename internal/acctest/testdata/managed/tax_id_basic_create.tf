# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name        = "acctest-tax-id-customer-{{RAND}}"
  email       = "sdk-codegen+tax-id-{{RAND}}@example.com"
  description = "sdk-codegen tax id customer"
}

resource "stripe_tax_id" "test" {
  customer = stripe_customer.customer.id
  type     = "us_ein"
  value    = "12-3456789"
}
