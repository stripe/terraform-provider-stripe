# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "product" {
  name = "acctest-payment-link-controls-product-{{RAND}}"
}

resource "stripe_price" "price" {
  product     = stripe_product.product.id
  currency    = "usd"
  unit_amount = 2200

  recurring {
    interval = "month"
  }
}

resource "stripe_payment_link" "test" {
  active                     = true
  inactive_message           = "sdk-codegen payment link controls create"
  allow_promotion_codes      = true
  billing_address_collection = "required"
  payment_method_collection  = "if_required"
  submit_type                = "donate"

  after_completion = {
    type = "redirect"
    redirect = {
      url = "https://example.com/sdk-codegen/payment-link/controls/create"
    }
  }

  line_items = [
    {
      price    = stripe_price.price.id
      quantity = 1
    },
  ]

  metadata = {
    suite = "sdk-codegen"
    case  = "payment_link_checkout_controls"
    phase = "create"
  }
}
