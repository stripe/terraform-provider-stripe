# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name  = "acctest-subscription-schedule-quantity-customer-{{RAND}}"
  email = "sdk-codegen+subsched-quantity-{{RAND}}@example.com"
}

resource "stripe_product" "product" {
  name = "acctest-subscription-schedule-quantity-product-{{RAND}}"
}

resource "stripe_price" "price" {
  product     = stripe_product.product.id
  currency    = "usd"
  unit_amount = 2000

  recurring {
    interval = "month"
  }
}

resource "stripe_subscription_schedule" "test" {
  customer     = stripe_customer.customer.id
  start_date   = 1893628800
  end_behavior = "cancel"

  metadata = {
    suite = "sdk-codegen"
    case  = "subscription_schedule_phase_quantity"
    phase = "create"
  }

  phases = [
    {
      iterations = 2
      duration = {
        interval       = "month"
        interval_count = 1
      }
      items = [
        {
          price    = stripe_price.price.id
          quantity = 1
        },
      ]
    },
  ]
}
