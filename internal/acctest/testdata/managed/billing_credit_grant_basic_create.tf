# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name = "SDK Codegen Credit Grant Customer {{RAND}}"
}

resource "stripe_billing_credit_grant" "test" {
  customer = stripe_customer.customer.id
  category = "promotional"
  name     = "SDK Codegen Credit Grant"
  priority = 50

  amount = {
    type = "monetary"
    monetary = {
      currency = "usd"
      value    = 500
    }
  }

  applicability_config = {
    scope = {
      price_type = "metered"
    }
  }

  expires_at = 1893456000

  metadata = {
    suite = "sdk-codegen"
    case  = "billing_credit_grant_basic"
    phase = "create"
  }
}
