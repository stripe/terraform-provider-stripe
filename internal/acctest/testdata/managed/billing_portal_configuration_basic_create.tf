# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
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
