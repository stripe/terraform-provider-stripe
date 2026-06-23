# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name  = "acctest-subscription-item-customer-{{RAND}}"
  email = "sdk-codegen+subscription-item-{{RAND}}@example.com"
}

resource "stripe_product" "base" {
  name = "acctest-subscription-item-base-{{RAND}}"
}

resource "stripe_price" "base" {
  product     = stripe_product.base.id
  currency    = "usd"
  unit_amount = 2000

  recurring {
    interval = "month"
  }
}

resource "stripe_product" "addon" {
  name = "acctest-subscription-item-addon-{{RAND}}"
}

resource "stripe_price" "addon" {
  product     = stripe_product.addon.id
  currency    = "usd"
  unit_amount = 600

  recurring {
    interval = "month"
  }
}

resource "stripe_subscription" "subscription" {
  customer          = stripe_customer.customer.id
  collection_method = "send_invoice"
  days_until_due    = 7

  items = [
    {
      price    = stripe_price.base.id
      quantity = 1
    },
  ]
}

resource "stripe_subscription_item" "test" {
  subscription = stripe_subscription.subscription.id
  price        = stripe_price.addon.id
  quantity     = 4

  metadata = {
    suite = "sdk-codegen"
    case  = "subscription_item_basic"
    phase = "update"
  }
}
