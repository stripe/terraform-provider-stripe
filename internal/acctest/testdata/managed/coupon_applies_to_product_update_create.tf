# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "scope_a" {
  name = "acctest-coupon-product-a-{{RAND}}"
}

resource "stripe_product" "scope_b" {
  name = "acctest-coupon-product-b-{{RAND}}"
}

resource "stripe_coupon" "test" {
  duration    = "forever"
  percent_off = 10
  name        = "SDK ApplProd1 {{RAND}}"

  applies_to {
    products = [stripe_product.scope_a.id]
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "coupon_applies_to_product_update"
  }
}
