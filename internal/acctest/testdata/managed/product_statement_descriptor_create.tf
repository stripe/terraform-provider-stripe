# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "test" {
  name        = "acctest-product-stmt-{{RAND}}"
  type        = "service"
  tax_code    = "txcd_10000000"
  description = "Statement descriptor acceptance coverage"

  statement_descriptor = "ACCTEST STRIPE SDKG"

  metadata = {
    suite = "sdk-codegen"
    case  = "product_statement_descriptor"
  }
}
