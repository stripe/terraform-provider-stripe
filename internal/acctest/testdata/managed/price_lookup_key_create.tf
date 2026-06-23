# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "product" {
  name = "acctest-price-lookup-product-{{RAND}}"
}

resource "stripe_price" "test" {
  product     = stripe_product.product.id
  currency    = "usd"
  unit_amount = 3600
  lookup_key  = "acctest-price-lookup-{{RAND}}"
  nickname    = "SDK annual plan"
  tax_behavior = "exclusive"

  lifecycle {
    create_before_destroy = true
  }

  recurring {
    interval = "year"
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "price_lookup_key"
    phase = "create"
  }
}
