# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
resource "stripe_treasury_financial_account" "test" {
  supported_currencies = ["usd"]
  nickname             = "SDK Codegen Treasury Movement Update"

  features = {
    financial_addresses = {
      aba = {
        requested = true
      }
    }
    inbound_transfers = {
      ach = {
        requested = true
      }
    }
    outbound_payments = {
      ach = {
        requested = true
      }
    }
    outbound_transfers = {
      ach = {
        requested = true
      }
    }
  }

  metadata = {
    suite = "sdk-codegen"
    case  = "treasury_financial_account_money_movement_features"
    phase = "update"
  }
}
