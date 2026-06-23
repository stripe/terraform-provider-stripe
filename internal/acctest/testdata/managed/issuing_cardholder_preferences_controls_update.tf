# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_issuing_cardholder" "test" {
  type         = "individual"
  name         = "SDK Pref Cardholder"
  email        = "sdk-codegen+cardholder-preferences-update@example.com"
  phone_number = "+15555550154"
  status       = "inactive"

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

  preferred_locales = ["en", "es"]

  spending_controls = {
    allowed_card_presences = ["not_present"]
    spending_limits_currency = "usd"
    spending_limits = [{
      amount   = 7000
      interval = "monthly"
    }]
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "issuing_cardholder_preferences_controls"
    phase = "update"
  }
}
