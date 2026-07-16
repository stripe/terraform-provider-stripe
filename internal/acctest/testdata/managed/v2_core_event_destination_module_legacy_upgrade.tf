# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
module "test" {
  source = "{{MODULE_PATH:managed/modules/v2_core_event_destination_module_legacy_upgrade}}"
  rand   = "{{RAND}}"
}

output "eventbridge_aws_account_id" {
  value = module.test.eventbridge_aws_account_id
}

output "webhook_url" {
  value = module.test.webhook_url
}
