# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_coupon" "test" {
  duration   = "once"
  amount_off = 500
  currency   = "usd"
  name       = "SDK Amount Upd {{RAND}}"
  redeem_by  = 1893456000

  metadata = {
    suite = "sdk-codegen"
    case  = "coupon_amount_off"
    phase = "update"
  }
}
