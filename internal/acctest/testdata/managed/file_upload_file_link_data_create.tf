# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_file" "test" {
  purpose = "dispute_evidence"
  file    = "{{FILE_UPLOAD_FIXTURE_PATH}}"

  file_link_data = {
    create     = true
    expires_at = 1893456000
    metadata = {
      suite = "sdk-codegen"
      case  = "file_upload_file_link_data"
    }
  }
}
