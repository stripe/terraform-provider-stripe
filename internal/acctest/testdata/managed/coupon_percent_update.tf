# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_coupon" "test" {
  duration    = "once"
  percent_off = 25
  name        = "SDK Percent Upd {{RAND}}"

  metadata = {
    suite = "sdk-codegen"
    case  = "coupon_percent_off"
    phase = "update"
  }
}
