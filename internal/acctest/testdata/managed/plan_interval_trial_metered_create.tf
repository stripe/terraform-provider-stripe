# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "product" {
  name = "acctest-plan-interval-product-{{RAND}}"
}

resource "stripe_plan" "test" {
  product           = stripe_product.product.id
  currency          = "usd"
  amount            = 3500
  interval          = "month"
  interval_count    = 3
  nickname          = "SDK Codegen Metered Trial Plan"
  trial_period_days = 14
  usage_type        = "licensed"

  metadata = {
    suite = "sdk-codegen"
    case  = "plan_interval_trial"
  }
}
