# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
resource "stripe_product" "test" {
  name = "acctest-tax-code-mirrors-{{RAND}}"
  tax_details = {
    tax_code = "txcd_20030000"
  }
}
