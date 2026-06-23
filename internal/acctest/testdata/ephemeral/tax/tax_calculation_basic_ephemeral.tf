# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
ephemeral "stripe_tax_calculation" "test" {
  currency = "usd"

  customer_details = {
    address_source = "shipping"
    address = {
      line1       = "354 Oyster Point Blvd"
      city        = "South San Francisco"
      state       = "CA"
      postal_code = "94080"
      country     = "US"
    }
  }

  ship_from_details = {
    address = {
      line1       = "510 Townsend St"
      city        = "San Francisco"
      state       = "CA"
      postal_code = "94103"
      country     = "US"
    }
  }

  line_items = [
    {
      amount       = 1099
      quantity     = 1
      reference    = "sku_sdk_codegen"
      tax_behavior = "exclusive"
      tax_code     = "txcd_99999999"
    },
  ]
}
