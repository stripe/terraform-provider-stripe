# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_webhook_endpoint" "test" {
  url = "https://example.com/sdk-codegen/base/{{RAND}}"

  enabled_events = [
    "customer.created",
  ]

  description = "sdk-codegen base webhook"

  metadata = {
    suite = "sdk-codegen"
    case  = "webhook_endpoint_url_update"
  }
}
