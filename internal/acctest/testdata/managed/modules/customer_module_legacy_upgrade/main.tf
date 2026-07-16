# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
terraform {
  required_providers {
    stripe = {
      source = "stripe/stripe"
    }
  }
}

variable "rand" {
  type = string
}

resource "stripe_customer" "test" {
  name        = "acctest-customer-${var.rand}"
  email       = "sdk-codegen+${var.rand}@example.com"
  description = "sdk-codegen acceptance customer"

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
    case  = "customer_module_legacy_upgrade"
  }
}

output "address_city" {
  value = stripe_customer.test.address[0].city
}

output "invoice_footer" {
  value = stripe_customer.test.invoice_settings[0].footer
}

output "shipping_address_city" {
  value = stripe_customer.test.shipping[0].address[0].city
}
