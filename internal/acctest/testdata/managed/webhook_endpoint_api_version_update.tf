# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_webhook_endpoint" "test" {
  url         = "https://example.com/sdk-codegen/version-updated/{{RAND}}"
  api_version = "2025-08-27.basil"

  enabled_events = [
    "customer.created",
    "customer.deleted",
    "customer.updated",
  ]

  description = "sdk-codegen versioned webhook updated"

  metadata = {
    suite = "sdk-codegen"
    case  = "webhook_endpoint_api_version"
    phase = "update"
  }
}
