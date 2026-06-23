# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_billing_meter" "meter" {
  display_name = "SDK Codegen Alert Meter {{RAND}}"
  event_name   = "sdk_codegen_alert_meter_{{RAND}}"

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

resource "stripe_billing_alert" "test" {
  alert_type = "usage_threshold"
  title      = "Codex Billing Alert"

  usage_threshold = {
    gte        = 100
    meter      = stripe_billing_meter.meter.id
    recurrence = "one_time"
  }
}
