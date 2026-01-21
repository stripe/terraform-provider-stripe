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
  # api_key = "api_key"
}


resource "stripe_product" "product" {
  name = "My Product 23"
}

resource "stripe_price" "price" {
  product = stripe_product.product.id
  unit_amount = 1000
  currency = "usd"
}

resource "stripe_webhook_endpoint" "webhook_endpoint" {
  url = "https://example.com/webhook"
  enabled_events = [
    "checkout.session.completed",
  ]
}

output "price_id" {
  value = stripe_price.price.id
}

output "product_id" {
  value = stripe_product.product.id
}

output "webhook_endpoint_id" {
  value = stripe_webhook_endpoint.webhook_endpoint.id
}