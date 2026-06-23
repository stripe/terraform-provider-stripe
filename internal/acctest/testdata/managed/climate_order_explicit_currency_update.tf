# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_climate_order" "test" {
  product     = "climsku_frontier_offtake_portfolio_2027"
  metric_tons = 0.1
  currency    = "usd"

  beneficiary = {
    public_name = "SDK Codegen Climate Currency Update"
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "climate_order_explicit_currency"
    phase = "update"
  }
}
