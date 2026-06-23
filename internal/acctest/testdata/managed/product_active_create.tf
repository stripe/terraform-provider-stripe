# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "test" {
  name   = "acctest-product-active-{{RAND}}"
  active = true

  metadata = {
    suite = "sdk-codegen"
    case  = "product_active_toggle"
  }
}
