# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_issuing_cardholder" "cardholder" {
  type         = "individual"
  name         = "SDK Cardholder Card"
  email        = "sdk-codegen+issuing-cardholder@example.com"
  phone_number = "+15555550161"
  status       = "active"

  billing = {
    address = {
      line1       = "100 Main St"
      city        = "San Francisco"
      postal_code = "94105"
      country     = "US"
      state       = "CA"
    }
  }

  individual = {
    first_name = "Grace"
    last_name  = "Hopper"
    dob = {
      day   = 9
      month = 12
      year  = 1985
    }
  }
}

resource "stripe_issuing_card" "test" {
  cardholder = stripe_issuing_cardholder.cardholder.id
  financial_account = "{{ISSUING_FINANCIAL_ACCOUNT}}"
  currency   = "usd"
  type       = "virtual"
  status     = "inactive"

  metadata = {
    suite = "sdk-codegen"
    case  = "issuing_card_basic"
    phase = "update"
  }
}
