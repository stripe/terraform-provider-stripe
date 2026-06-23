# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_coupon" "coupon" {
  duration    = "once"
  percent_off = 30
}

resource "stripe_promotion_code" "test" {
  code   = "SDKPROMO-{{RAND}}"
  active = true

  promotion {
    type   = "coupon"
    coupon = stripe_coupon.coupon.id
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "promotion_code_basic"
  }
}
