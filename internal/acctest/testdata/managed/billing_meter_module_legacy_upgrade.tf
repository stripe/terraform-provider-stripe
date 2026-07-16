# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
module "test" {
  source = "{{MODULE_PATH:managed/modules/billing_meter_module_legacy_upgrade}}"
  rand   = "{{RAND}}"
}

output "customer_mapping_event_payload_key" {
  value = module.test.customer_mapping_event_payload_key
}

output "default_aggregation_formula" {
  value = module.test.default_aggregation_formula
}

output "value_settings_event_payload_key" {
  value = module.test.value_settings_event_payload_key
}
