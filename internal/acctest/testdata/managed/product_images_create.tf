# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "test" {
  name = "acctest-product-images-{{RAND}}"
  url  = "https://example.com/sdk-codegen/product/{{RAND}}"

  images = [
    "https://example.com/sdk-codegen/product-a.png",
  ]

  metadata = {
    suite = "sdk-codegen"
    case  = "product_images_url"
  }
}
