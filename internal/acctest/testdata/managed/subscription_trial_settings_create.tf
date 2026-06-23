# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name  = "acctest-subscription-trial-customer-{{RAND}}"
  email = "sdk-codegen+subscription-trial-{{RAND}}@example.com"
}

resource "stripe_product" "product" {
  name = "acctest-subscription-trial-product-{{RAND}}"
}

resource "stripe_price" "price" {
  product     = stripe_product.product.id
  currency    = "usd"
  unit_amount = 1800

  recurring {
    interval = "month"
  }
}

resource "stripe_subscription" "test" {
  customer          = stripe_customer.customer.id
  collection_method = "charge_automatically"
  description       = "sdk-codegen subscription trial create"
  trial_period_days = 14

  trial_settings = {
    end_behavior = {
      missing_payment_method = "pause"
    }
  }

  items = [
    {
      price    = stripe_price.price.id
      quantity = 1
    },
  ]

  metadata = {
    suite = "sdk-codegen"
    case  = "subscription_trial_settings"
    phase = "create"
  }
}
