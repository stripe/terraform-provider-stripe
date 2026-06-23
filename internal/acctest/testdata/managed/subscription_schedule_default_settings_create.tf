# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name  = "acctest-subscription-schedule-default-customer-{{RAND}}"
  email = "sdk-codegen+subsched-default-{{RAND}}@example.com"
}

resource "stripe_product" "product" {
  name = "acctest-subscription-schedule-default-product-{{RAND}}"
}

resource "stripe_price" "price" {
  product     = stripe_product.product.id
  currency    = "usd"
  unit_amount = 1500

  recurring {
    interval = "month"
  }
}

resource "stripe_subscription_schedule" "test" {
  customer     = stripe_customer.customer.id
  start_date   = 1894060800
  end_behavior = "release"

  default_settings = {
    collection_method = "send_invoice"
    description       = "sdk-codegen schedule default settings create"
    invoice_settings = {
      days_until_due = 14
    }
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "subscription_schedule_default_settings"
    phase = "create"
  }

  phases = [
    {
      iterations = 1
      items = [
        {
          price    = stripe_price.price.id
          quantity = 1
        },
      ]
    },
  ]
}
