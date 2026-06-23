# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_issuing_cardholder" "test" {
  type         = "individual"
  name         = "SDK Pref Cardholder"
  email        = "sdk-codegen+cardholder-preferences-create@example.com"
  phone_number = "+15555550153"
  status       = "active"

  billing = {
    address = {
      line1       = "500 Howard St"
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

  preferred_locales = ["en", "fr"]

  spending_controls = {
    allowed_card_presences = ["present"]
    spending_limits_currency = "usd"
    spending_limits = [{
      amount   = 5000
      interval = "weekly"
    }]
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "issuing_cardholder_preferences_controls"
    phase = "create"
  }
}
