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

resource "stripe_billing_meter" "test" {
  display_name = "SDK Codegen Meter ${var.rand}"
  event_name   = "sdk_codegen_meter_${var.rand}"

  default_aggregation {
    formula = "sum"
  }

  value_settings {
    event_payload_key = "usage_total"
  }

  customer_mapping {
    type              = "by_id"
    event_payload_key = "stripe_customer_id"
  }
}

output "customer_mapping_event_payload_key" {
  value = stripe_billing_meter.test.customer_mapping[0].event_payload_key
}

output "default_aggregation_formula" {
  value = stripe_billing_meter.test.default_aggregation[0].formula
}

output "value_settings_event_payload_key" {
  value = stripe_billing_meter.test.value_settings[0].event_payload_key
}
