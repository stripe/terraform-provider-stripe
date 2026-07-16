# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
module "test" {
  source = "{{MODULE_PATH:managed/modules/product_module_legacy_upgrade}}"
  rand   = "{{RAND}}"
}

output "default_price_currency" {
  value = module.test.default_price_currency
}

output "default_price_interval" {
  value = module.test.default_price_interval
}
