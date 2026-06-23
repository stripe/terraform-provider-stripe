# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_tax_registration" "test" {
  active_from = 1896000000 + ({{RAND}} % 50000000)
  country     = "CA"

  country_options = {
    ca = {
      type = "province_standard"
      province_standard = {
        province = "QC"
      }
    }
  }
}
