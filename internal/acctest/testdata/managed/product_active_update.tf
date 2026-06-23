# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "test" {
  name   = "acctest-product-active-{{RAND}}"
  active = false

  metadata = {
    suite = "sdk-codegen"
    case  = "product_active_toggle"
    phase = "update"
  }
}
