resource "stripe_entitlements_feature" "api_access" {
  lookup_key = "api_access"
  name       = "API access"
}

data "stripe_entitlements_feature" "api_access" {
  lookup_key = stripe_entitlements_feature.api_access.lookup_key
}
