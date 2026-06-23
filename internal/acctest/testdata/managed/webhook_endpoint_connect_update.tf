# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_webhook_endpoint" "test" {
  url     = "https://example.com/sdk-codegen/connect/{{RAND}}"
  connect = true

  enabled_events = [
    "checkout.session.completed",
    "price.created",
    "product.updated",
  ]

  description = "sdk-codegen connect webhook updated"

  metadata = {
    suite = "sdk-codegen"
    case  = "webhook_endpoint_connect"
    phase = "update"
  }
}
