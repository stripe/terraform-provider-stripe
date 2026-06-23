# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_billing_meter" "test" {
  display_name = "SDK Codegen Count Meter Updated {{RAND}}"
  event_name   = "sdk_codegen_meter_count_{{RAND}}"

  default_aggregation {
    formula = "count"
  }

  value_settings {
    event_payload_key = "event_count"
  }

  customer_mapping {
    type              = "by_id"
    event_payload_key = "stripe_customer_id"
  }
}
