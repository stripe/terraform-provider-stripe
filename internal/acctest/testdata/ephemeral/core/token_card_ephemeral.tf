# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
ephemeral "stripe_token" "test" {
  card = {
    number    = "4242424242424242"
    exp_month = 12
    exp_year  = 2030
    cvc       = "123"
  }
}
