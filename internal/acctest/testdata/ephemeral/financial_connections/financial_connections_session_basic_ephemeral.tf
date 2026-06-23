# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name = "SDK Codegen Financial Connections {{RAND}}"
}

ephemeral "stripe_financial_connections_session" "test" {
  account_holder = {
    type     = "customer"
    customer = stripe_customer.customer.id
  }

  permissions = ["balances", "ownership"]
  prefetch    = ["balances"]

  filters = {
    countries             = ["US"]
    account_subcategories = ["checking"]
  }
}
