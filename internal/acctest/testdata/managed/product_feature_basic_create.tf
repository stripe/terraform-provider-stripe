# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "product" {
  name        = "SDK Codegen Feature Product {{RAND}}"
  description = "sdk-codegen product feature attachment"
}

resource "stripe_entitlements_feature" "feature" {
  name       = "SDK Codegen Feature {{RAND}}"
  lookup_key = "sdk-codegen-feature-{{RAND}}"

  metadata = {
    suite = "sdk-codegen"
    case  = "product_feature_basic"
  }
}

resource "stripe_product_feature" "test" {
  product             = stripe_product.product.id
  entitlement_feature = stripe_entitlements_feature.feature.id
}
