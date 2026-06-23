# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "product" {
  name = "acctest-plan-product-{{RAND}}"
}

resource "stripe_plan" "test" {
  product  = stripe_product.product.id
  currency = "usd"
  amount   = 1200
  interval = "month"
  nickname = "SDK Codegen Plan Create"

  metadata = {
    suite = "sdk-codegen"
    case  = "plan_basic"
    phase = "create"
  }
}
