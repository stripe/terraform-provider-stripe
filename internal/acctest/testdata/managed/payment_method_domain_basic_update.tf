# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_payment_method_domain" "test" {
  domain_name = "{{PAYMENT_METHOD_DOMAIN}}"
  enabled     = false
}
