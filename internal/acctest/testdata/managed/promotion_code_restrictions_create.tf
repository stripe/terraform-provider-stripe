# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_coupon" "coupon" {
  duration    = "forever"
  percent_off = 20
}

resource "stripe_promotion_code" "test" {
  code            = "SDKREST-{{RAND}}"
  active          = true
  max_redemptions = 10

  promotion {
    type   = "coupon"
    coupon = stripe_coupon.coupon.id
  }

  restrictions {
    minimum_amount          = 2000
    minimum_amount_currency = "usd"
    first_time_transaction  = true
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "promotion_code_restrictions"
  }
}
