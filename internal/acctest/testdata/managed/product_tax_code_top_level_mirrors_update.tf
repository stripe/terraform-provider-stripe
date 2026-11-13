# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
resource "stripe_product" "test" {
  name = "acctest-tax-code-top-level-mirrors-{{RAND}}"
  tax_code = "txcd_20030000"
}
