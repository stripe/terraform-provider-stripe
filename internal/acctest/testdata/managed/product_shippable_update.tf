# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "test" {
  name      = "acctest-product-good-{{RAND}}"
  type      = "good"
  shippable = true

  package_dimensions {
    height = 4.25
    length = 8.5
    weight = 16
    width  = 2.75
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "product_shippable"
    phase = "update"
  }
}
