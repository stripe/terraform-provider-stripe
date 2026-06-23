# File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
# Handwritten Terraform acceptance source emitted by sdk-codegen.
ephemeral "stripe_checkout_session" "test" {
  success_url = "https://example.com/success?session_id={CHECKOUT_SESSION_ID}"
  cancel_url  = "https://example.com/cancel"
  mode        = "payment"

  line_items = [
    {
      quantity = 1
      price_data = {
        currency    = "usd"
        unit_amount = 2000
        product_data = {
          name = "SDK Codegen Checkout Session"
        }
      }
    }
  ]
}
