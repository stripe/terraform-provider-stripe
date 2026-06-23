# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "scope" {
  name = "acctest-coupon-product-{{RAND}}"
}

resource "stripe_coupon" "test" {
  duration    = "forever"
  percent_off = 10
  name        = "SDK Applies {{RAND}}"

  applies_to {
    products = [stripe_product.scope.id]
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "coupon_applies_to"
  }
}
