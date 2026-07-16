# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
terraform {
  required_providers {
    stripe = {
      source = "stripe/stripe"
    }
  }
}

variable "rand" {
  type = string
}

resource "stripe_v2_core_event_destination" "eventbridge" {
  name          = "acctest-module-eventbridge-${var.rand}"
  description   = "sdk-codegen module eventbridge destination"
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
    case  = "v2_core_event_destination_module_legacy_upgrade"
    kind  = "eventbridge"
  }
}

resource "stripe_v2_core_event_destination" "webhook" {
  name          = "acctest-module-webhook-destination-${var.rand}"
  description   = "sdk-codegen module webhook destination"
  type          = "webhook_endpoint"
  event_payload = "thin"

  enabled_events = [
    "v1.billing.meter.error_report_triggered",
  ]

  webhook_endpoint {
    url = "https://example.com/sdk-codegen/module-event-destination"
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "v2_core_event_destination_module_legacy_upgrade"
    kind  = "webhook"
  }
}

output "eventbridge_aws_account_id" {
  value = stripe_v2_core_event_destination.eventbridge.amazon_eventbridge[0].aws_account_id
}

output "webhook_url" {
  value = stripe_v2_core_event_destination.webhook.webhook_endpoint[0].url
}
