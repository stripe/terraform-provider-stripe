# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_billing_portal_configuration" "test" {
  active             = false
  name               = "SDK Codegen Portal Update"
  default_return_url = "https://example.com/sdk-codegen/portal/update"

  business_profile = {
    headline             = "SDK Codegen portal update headline"
    privacy_policy_url   = "https://example.com/privacy"
    terms_of_service_url = "https://example.com/terms"
  }

  features = {
    customer_update = {
      enabled         = true
      allowed_updates = ["address", "email"]
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
    phase = "update"
  }
}
