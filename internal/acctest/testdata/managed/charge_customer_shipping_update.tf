# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name        = "acctest-charge-customer-{{RAND}}"
  email       = "sdk-codegen+charge-customer-{{RAND}}@example.com"
  description = "sdk-codegen charge customer"
  source      = "tok_visa"
}

resource "stripe_charge" "test" {
  amount        = 3400
  currency      = "usd"
  customer      = stripe_customer.customer.id
  description   = "sdk-codegen charge customer shipping update"
  receipt_email = "sdk-codegen+charge-shipping-update@example.com"
  shipping = {
    name  = "SDK Codegen Charge Update"
    phone = "+15555550162"
    address = {
      line1       = "200 Mission St"
      city        = "San Francisco"
      state       = "CA"
      postal_code = "94105"
      country     = "US"
    }
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "charge_customer_shipping"
    phase = "update"
  }
}
