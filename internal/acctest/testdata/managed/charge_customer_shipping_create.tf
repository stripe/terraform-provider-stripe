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
  description   = "sdk-codegen charge customer shipping create"
  receipt_email = "sdk-codegen+charge-shipping-create@example.com"
  shipping = {
    name  = "SDK Codegen Charge Create"
    phone = "+15555550161"
    address = {
      line1       = "100 Market St"
      city        = "San Francisco"
      state       = "CA"
      postal_code = "94105"
      country     = "US"
    }
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "charge_customer_shipping"
    phase = "create"
  }
}
