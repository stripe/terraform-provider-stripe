# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_tax_registration" "test" {
  active_from = 1899000000 + ({{RAND}} % 50000000)
  country     = "CA"

  country_options = {
    ca = {
      type = "simplified"
    }
  }
}
