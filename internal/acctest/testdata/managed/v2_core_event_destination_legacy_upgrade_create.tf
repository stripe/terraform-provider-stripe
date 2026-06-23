# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_v2_core_event_destination" "eventbridge" {
  name          = "acctest-eventbridge-{{RAND}}"
  description   = "sdk-codegen eventbridge destination"
  type          = "amazon_eventbridge"
  event_payload = "thin"

  enabled_events = [
    "v1.billing.meter.error_report_triggered",
  ]

  amazon_eventbridge {
    aws_account_id = "111122223333"
    aws_region     = "us-east-1"
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "v2_core_event_destination_basic"
    kind  = "eventbridge"
  }
}

resource "stripe_v2_core_event_destination" "webhook" {
  name          = "acctest-webhook-destination-{{RAND}}"
  description   = "sdk-codegen webhook destination"
  type          = "webhook_endpoint"
  event_payload = "thin"

  enabled_events = [
    "v1.billing.meter.error_report_triggered",
  ]

  webhook_endpoint {
    url = "https://example.com/sdk-codegen/event-destination/{{RAND}}"
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "v2_core_event_destination_basic"
    kind  = "webhook"
  }
}
