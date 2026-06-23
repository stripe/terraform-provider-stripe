# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_billing_portal_configuration" "test" {
  active             = true
  name               = "SDK Codegen Portal Login Cancel Create"
  default_return_url = "https://example.com/sdk-codegen/portal/login-cancel/create"

  business_profile = {
    headline             = "SDK Codegen portal login cancel create headline"
    privacy_policy_url   = "https://example.com/login-cancel/privacy"
    terms_of_service_url = "https://example.com/login-cancel/terms"
  }

  features = {
    customer_update = {
      enabled         = true
      allowed_updates = ["email", "phone"]
    }
    invoice_history = {
      enabled = true
    }
    payment_method_update = {
      enabled = true
    }
    subscription_cancel = {
      enabled            = true
      mode               = "at_period_end"
      proration_behavior = "none"
      cancellation_reason = {
        enabled = true
        options = ["missing_features", "too_expensive"]
      }
    }
  }

  login_page = {
    enabled = true
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "billing_portal_configuration_login_cancel"
    phase = "create"
  }
}
