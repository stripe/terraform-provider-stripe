# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_product" "test" {
  name       = "acctest-product-marketing-{{RAND}}"
  unit_label = "seat"

  marketing_features {
    name = "Priority support"
  }

  marketing_features {
    name = "Usage insights"
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "product_marketing_features"
  }
}
