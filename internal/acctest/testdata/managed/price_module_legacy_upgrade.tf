# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
module "test" {
  source = "{{MODULE_PATH:managed/modules/price_module_legacy_upgrade}}"
  rand   = "{{RAND}}"
}

output "price_product_id" {
  value = module.test.price_product_id
}

output "price_recurring_interval" {
  value = module.test.price_recurring_interval
}

output "price_tier_unit_amount" {
  value = module.test.price_tier_unit_amount
}
