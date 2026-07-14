# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "test" {
  name = "acctest-customer-{{RAND}}"
  email = "sdk-codegen+{{RAND}}@example.com"
  description = "updated by sdk-codegen acceptance"

  address {
    line1       = "100 Market St"
    city        = "San Francisco"
    state       = "CA"
    postal_code = "94105"
    country     = "US"
  }

  invoice_settings {
    footer = "sdk-codegen customer footer"
  }

  shipping {
    name  = "SDK Codegen Customer"
    phone = "+15555550101"

    address {
      line1       = "100 Market St"
      city        = "San Francisco"
      state       = "CA"
      postal_code = "94105"
      country     = "US"
    }
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "customer_basic"
    phase = "update"
  }
}
