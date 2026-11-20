terraform {
  required_providers {
    stripe = { source = "stripe/stripe" }
  }
}

provider "stripe" {
  metronome = {
    base_url     = "{{API_URL}}"
    ca_cert_file = "{{CA_CERT_FILE}}"
  }
}

data "stripe_metronome_get_product_v1" "lookup" {
  request = { id = "product" }
}

output "result" {
  value     = data.stripe_metronome_get_product_v1.lookup.response_json
  sensitive = true
}

action "stripe_metronome_set_customer_name_v1" "rename" {
  config {
    path_params = { customer_id = "acct 123" }
    request     = { name = "New name" }
  }
}
