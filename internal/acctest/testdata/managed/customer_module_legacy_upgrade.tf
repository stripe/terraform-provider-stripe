# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
module "test" {
  source = "{{MODULE_PATH:managed/modules/customer_module_legacy_upgrade}}"
  rand   = "{{RAND}}"
}

output "address_city" {
  value = module.test.address_city
}

output "invoice_footer" {
  value = module.test.invoice_footer
}

output "shipping_address_city" {
  value = module.test.shipping_address_city
}
