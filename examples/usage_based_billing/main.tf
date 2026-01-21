terraform {
  required_providers {
    stripe = {
      source  = "stripe/stripe"
      version = "0.1.0"
    }
  }
}

provider "stripe" {
  # API key is read from STRIPE_API_KEY environment variable
  # Alternatively, set it explicitly (not recommended for production)
  api_key = "api_key"
}

resource "stripe_billing_meter" "api_calls" {
  display_name = "API Calls"
  event_name   = "api_call"
  
  default_aggregation {
    formula = "sum"
  }
  
  value_settings {
    event_payload_key = "value"
  }
  
  customer_mapping {
    type             = "by_id"
    event_payload_key = "stripe_customer_id"
  }
}

output "meter_id" {
  value = stripe_billing_meter.api_calls.id
}

output "meter_event_name" {
  value = stripe_billing_meter.api_calls.event_name
}