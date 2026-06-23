# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
action "stripe_treasury_outbound_payment" "test" {
  config {
    amount            = 55
    currency          = "usd"
    financial_account = "{{TREASURY_SOURCE_FINANCIAL_ACCOUNT}}"
    description       = "sdk-codegen treasury outbound payment basic"

    end_user_details = {
      present    = true
      ip_address = "127.0.0.1"
    }

    destination_payment_method_data = {
      type = "us_bank_account"
      billing_details = {
        name = "SDK Codegen Treasury Payment"
      }
      us_bank_account = {
        account_holder_type = "individual"
        account_number      = "000123456789"
        routing_number      = "110000000"
      }
    }

    metadata = {
      suite = "sdk-codegen"
      case  = "treasury_outbound_payment_basic"
    }
  }
}
