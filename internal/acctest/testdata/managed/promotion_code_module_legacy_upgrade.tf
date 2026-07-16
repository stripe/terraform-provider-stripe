# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
module "test" {
  source = "{{MODULE_PATH:managed/modules/promotion_code_module_legacy_upgrade}}"
  rand   = "{{RAND}}"
}

output "promotion_coupon_id" {
  value = module.test.promotion_coupon_id
}

output "minimum_amount_currency" {
  value = module.test.minimum_amount_currency
}
