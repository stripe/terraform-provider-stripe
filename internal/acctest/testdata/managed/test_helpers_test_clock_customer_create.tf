# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name        = "acctest-clock-customer-{{RAND}}"
  email       = "sdk-codegen+clock-customer-{{RAND}}@example.com"
  description = "sdk-codegen test clock customer"
}

resource "stripe_test_helpers_test_clock" "test" {
  frozen_time = 1896134400
  name        = "SDK Codegen Test Clock Customer"
  customer    = stripe_customer.customer.id
}
