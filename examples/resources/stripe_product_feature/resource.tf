resource "stripe_product" "growth" {
  name = "Growth"
}

resource "stripe_entitlements_feature" "api_access" {
  lookup_key = "api_access"
  name       = "API access"
}

resource "stripe_product_feature" "growth_api_access" {
  product             = stripe_product.growth.id
  entitlement_feature = stripe_entitlements_feature.api_access.id
}
