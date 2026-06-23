# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_billing_meter" "test" {
  display_name = "SDK Codegen Meter {{RAND}}"
  event_name   = "sdk_codegen_meter_{{RAND}}"

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
