# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_webhook_endpoint" "test" {
  url         = "https://example.com/sdk-codegen/version/{{RAND}}"
  api_version = "2025-07-30.basil"

  enabled_events = [
    "customer.created",
    "customer.deleted",
  ]

  description = "sdk-codegen versioned webhook"

  metadata = {
    suite = "sdk-codegen"
    case  = "webhook_endpoint_api_version"
    phase = "create"
  }
}
