# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "test" {
  name = "acctest-product-{{RAND}}"
  description = "sdk-codegen acceptance product"
  type = "service"
  tax_code = "txcd_10000000"

  metadata = {
    suite = "sdk-codegen"
    case  = "product_basic"
  }
}
