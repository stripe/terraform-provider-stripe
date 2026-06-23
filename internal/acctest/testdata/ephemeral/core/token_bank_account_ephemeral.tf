# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
ephemeral "stripe_token" "test" {
  bank_account = {
    country             = "US"
    currency            = "usd"
    account_holder_name = "Jenny Rosen"
    account_holder_type = "individual"
    routing_number      = "110000000"
    account_number      = "000123456789"
  }
}
