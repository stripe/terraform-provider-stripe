# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_customer" "customer" {
  name  = "acctest-subscription-customer-{{RAND}}"
  email = "sdk-codegen+subscription-{{RAND}}@example.com"
}

resource "stripe_product" "product" {
  name = "acctest-subscription-product-{{RAND}}"
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
  collection_method = "send_invoice"
  days_until_due    = 7
  description       = "sdk-codegen subscription basic updated"

  items = [
    {
      price    = stripe_price.price.id
      quantity = 2
    },
  ]

  metadata = {
    suite = "sdk-codegen"
    case  = "subscription_basic"
    phase = "update"
  }
}
