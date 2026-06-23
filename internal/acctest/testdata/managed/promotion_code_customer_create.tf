# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name  = "acctest-promo-customer-{{RAND}}"
  email = "sdk-codegen+promo-{{RAND}}@example.com"

  metadata = {
    suite = "sdk-codegen"
    case  = "promotion_code_customer"
  }
}

resource "stripe_coupon" "coupon" {
  duration    = "once"
  percent_off = 25
}

resource "stripe_promotion_code" "test" {
  code     = "SDKCUST-{{RAND}}"
  active   = true
  customer = stripe_customer.customer.id

  promotion {
    type   = "coupon"
    coupon = stripe_coupon.coupon.id
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "promotion_code_customer"
    phase = "create"
  }
}
