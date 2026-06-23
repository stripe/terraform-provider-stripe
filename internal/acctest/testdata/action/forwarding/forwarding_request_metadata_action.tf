# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
action "stripe_forwarding_request" "test" {
  config {
    payment_method = "{{FORWARDING_PAYMENT_METHOD}}"
    replacements   = ["card_number", "card_expiry", "card_cvc", "cardholder_name"]
    url            = "{{FORWARDING_DESTINATION}}"

    request = {
      body = "{\"amount\":{\"currency\":\"USD\",\"value\":1000},\"reference\":\"forwarding_request_metadata\",\"card\":{\"number\":\"\",\"exp_month\":\"\",\"exp_year\":\"\",\"cvc\":\"\",\"name\":\"\"}}"
      headers = [
        {
          name  = "Authorization"
          value = "Bearer eyJhbGciOiJIUzI1NiJ9.Zm9yd2FyZGluZy1hcGktZGVtbw.2qoK37CNBmMjMDRERSYUSE-YrjsTgGhHnxMeqOxjrAg"
        },
      ]
    }

    metadata = {
      suite  = "sdk-codegen"
      case   = "forwarding_request_metadata"
      cohort = "top10"
    }
  }
}
