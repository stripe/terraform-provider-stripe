# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_tax_rate" "test" {
  display_name = "SDK Codegen Sales Tax"
  description  = "sdk-codegen sales tax update"
  country      = "US"
  state        = "NY"
  jurisdiction = "New York"
  inclusive    = true
  percentage   = 4.5
  tax_type     = "sales_tax"
  active       = false

  metadata = {
    suite = "sdk-codegen"
    case  = "tax_rate_tax_type"
    phase = "update"
  }
}
