terraform {
  required_providers {
    stripe = {
      source  = "stripe/stripe"
      version = "0.1.0"
    }
  }
}

provider "stripe" {
  # API key is read from STRIPE_API_KEY
}

resource "stripe_product" "growth" {
  name        = "Growth"
  description = "Growth plan with API access"

  marketing_features {
    name = "API access"
  }
}

resource "stripe_entitlements_feature" "api_access" {
  lookup_key = "api_access"
  name       = "API access"
}

resource "stripe_product_feature" "growth_api_access" {
  product             = stripe_product.growth.id
  entitlement_feature = stripe_entitlements_feature.api_access.id
}

data "stripe_entitlements_feature" "api_access" {
  lookup_key = stripe_entitlements_feature.api_access.lookup_key
}

output "product_feature_id" {
  value = stripe_product_feature.growth_api_access.id
}

output "feature_lookup_key" {
  value = data.stripe_entitlements_feature.api_access.lookup_key
}
