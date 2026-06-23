# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_billing_portal_configuration" "test" {
  active             = false
  name               = "SDK Codegen Portal Login Cancel Update"
  default_return_url = "https://example.com/sdk-codegen/portal/login-cancel/update"

  business_profile = {
    headline             = "SDK Codegen portal login cancel update headline"
    privacy_policy_url   = "https://example.com/login-cancel/privacy-update"
    terms_of_service_url = "https://example.com/login-cancel/terms-update"
  }

  features = {
    customer_update = {
      enabled         = true
      allowed_updates = ["address", "shipping"]
    }
    invoice_history = {
      enabled = true
    }
    payment_method_update = {
      enabled = true
    }
    subscription_cancel = {
      enabled            = true
      mode               = "immediately"
      proration_behavior = "create_prorations"
      cancellation_reason = {
        enabled = true
        options = ["too_complex", "unused"]
      }
    }
  }

  login_page = {
    enabled = false
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "billing_portal_configuration_login_cancel"
    phase = "update"
  }
}
