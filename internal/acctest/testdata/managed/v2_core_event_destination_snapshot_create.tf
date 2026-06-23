# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_v2_core_event_destination" "eventbridge" {
  name                 = "acctest-eventbridge-snapshot-{{RAND}}"
  description          = "sdk-codegen eventbridge snapshot destination"
  type                 = "amazon_eventbridge"
  event_payload        = "snapshot"
  snapshot_api_version = "2025-07-30.basil"

  enabled_events = [
    "customer.created",
  ]

  amazon_eventbridge {
    aws_account_id = "111122223333"
    aws_region     = "us-east-1"
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "v2_core_event_destination_snapshot"
    kind  = "eventbridge"
  }
}

resource "stripe_v2_core_event_destination" "webhook" {
  name                 = "acctest-webhook-snapshot-destination-{{RAND}}"
  description          = "sdk-codegen webhook snapshot destination"
  type                 = "webhook_endpoint"
  event_payload        = "snapshot"
  snapshot_api_version = "2025-07-30.basil"
  include              = ["webhook_endpoint.signing_secret", "webhook_endpoint.url"]

  enabled_events = [
    "customer.created",
  ]

  webhook_endpoint {
    url = "https://example.com/sdk-codegen/event-destination-snapshot/{{RAND}}"
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "v2_core_event_destination_snapshot"
    kind  = "webhook"
  }
}
