# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name = "SDK Codegen Credit Grant Alt Customer {{RAND}}"
}

resource "stripe_billing_credit_grant" "test" {
  customer = stripe_customer.customer.id
  category = "promotional"
  name     = "SDK Codegen Credit Grant Alt Updated"
  priority = 90

  amount = {
    type = "monetary"
    monetary = {
      currency = "usd"
      value    = 900
    }
  }

  applicability_config = {
    scope = {
      price_type = "metered"
    }
  }

  expires_at = 1896134400

  metadata = {
    suite = "sdk-codegen"
    case  = "billing_credit_grant_amount_priority_update"
    phase = "update"
  }
}
