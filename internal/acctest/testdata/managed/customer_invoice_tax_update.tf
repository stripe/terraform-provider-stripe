# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "test" {
  name  = "acctest-customer-inv-{{RAND}}"
  email = "sdk-codegen+inv-{{RAND}}@example.com"

  address {
    line1       = "100 Market St"
    city        = "San Francisco"
    state       = "CA"
    postal_code = "94105"
    country     = "US"
  }

  invoice_prefix = upper(substr(md5("invupd{{RAND}}"), 0, 12))

  invoice_settings {
    footer = "sdk-codegen customer invoice footer"
  }

  tax_exempt     = "exempt"

  metadata = {
    suite = "sdk-codegen"
    case  = "customer_invoice_tax"
    phase = "update"
  }
}
