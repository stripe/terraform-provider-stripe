# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_tax_rate" "test" {
  display_name = "SDK Codegen Tax Rate"
  description  = "sdk-codegen tax rate create"
  country      = "US"
  state        = "CA"
  jurisdiction = "California"
  inclusive    = false
  percentage   = 8.25

  metadata = {
    suite = "sdk-codegen"
    case  = "tax_rate_basic"
    phase = "create"
  }
}
