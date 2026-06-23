# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_coupon" "test" {
  duration           = "repeating"
  duration_in_months = 3
  percent_off        = 15
  name               = "SDK Repeat {{RAND}}"
  max_redemptions    = 5

  metadata = {
    suite = "sdk-codegen"
    case  = "coupon_repeating"
  }
}
