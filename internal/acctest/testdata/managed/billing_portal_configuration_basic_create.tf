# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "product" {
  name = "acctest-portal-product-{{RAND}}"
}

resource "stripe_price" "monthly" {
  product = stripe_product.product.id
  currency = "usd"
  unit_amount = 1500
  recurring { interval = "month" }
}

resource "stripe_price" "annual" {
  product = stripe_product.product.id
  currency = "usd"
  unit_amount = 15000
  recurring { interval = "year" }
}

resource "stripe_billing_portal_configuration" "test" {
  active             = true
  name               = "SDK Codegen Portal Create"
  default_return_url = "https://example.com/sdk-codegen/portal/create"

  business_profile = {
    headline             = "SDK Codegen portal create headline"
    privacy_policy_url   = "https://example.com/privacy"
    terms_of_service_url = "https://example.com/terms"
  }

  features = {
    subscription_update = {
      enabled = true
      default_allowed_updates = ["price"]
      products = [{ product = stripe_product.product.id, prices = [stripe_price.monthly.id] }]
    }
    customer_update = {
      enabled         = true
      allowed_updates = ["email", "tax_id"]
    }
    invoice_history = {
      enabled = true
    }
    payment_method_update = {
      enabled = true
    }
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "billing_portal_configuration_basic"
    phase = "create"
  }
}
