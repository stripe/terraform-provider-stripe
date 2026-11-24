//
// File generated from our OpenAPI spec
//

package metronome

// Permanently disable a threshold notification and remove it from active monitoring across all customers. Archived threshold notifications stop evaluating immediately and can optionally release their uniqueness key for reuse in future threshold notification configurations.
//
// ### Use this endpoint to:
// - Decommission threshold notifications that are no longer needed
// - Clean up test or deprecated threshold notification configurations
// - Free up uniqueness keys for reuse with new threshold notifications
// - Stop threshold notification evaluations without losing historical configuration data
// - Disable outdated monitoring rules during pricing model transitions
//
// ### Key response fields:
// - data: Object containing the archived threshold notification's ID
//
// ### Usage guidelines:
// - Irreversible for evaluation: Archived threshold notifications cannot be re-enabled; create a new threshold notification to resume monitoring
// - Uniqueness key handling: Set `release_uniqueness_key` : `true` to reuse the key in future threshold notifications
// - Immediate effect: Threshold notification evaluation stops instantly across all customers
// - Historical preservation: Archive operation maintains threshold notification history and configuration for compliance and auditing
type V1AlertsArchiveAlertParams struct {
	Params `form:"*"`
	// The Metronome ID of the threshold notification
	ID *string `form:"id" json:"id"`
	// If true, resets the uniqueness key on this threshold notification so it can be re-used
	ReleaseUniquenessKey *bool `form:"release_uniqueness_key" json:"release_uniqueness_key,omitempty"`
}

// A list of custom field filters for threshold notification types that support advanced filtering. Only present for contract invoices.
type V1AlertsCreateAlertCustomFieldFilterParams struct {
	Entity *string `form:"entity" json:"entity"`
	Key    *string `form:"key" json:"key"`
	Value  *string `form:"value" json:"value"`
}

// Only present for `spend_threshold_reached` notifications. Scope notification to a specific group key on individual line items.
type V1AlertsCreateAlertGroupValueParams struct {
	Key   *string `form:"key" json:"key"`
	Value *string `form:"value" json:"value,omitempty"`
}

// Required for `low_remaining_seat_balance_reached` notifications. The alert is scoped to this seat group key-value pair.
type V1AlertsCreateAlertSeatFilterParams struct {
	// The seat group key (e.g., "seat_id", "user_id")
	SeatGroupKey *string `form:"seat_group_key" json:"seat_group_key"`
	// Optional seat identifier the alert is scoped to.
	SeatGroupValue *string `form:"seat_group_value" json:"seat_group_value,omitempty"`
}

// A list of custom field filters for notification types that support advanced filtering
type V1AlertsCreateAlertAlertSpecifierCustomFieldFilterParams struct {
	Entity *string `form:"entity" json:"entity"`
	Key    *string `form:"key" json:"key"`
	Value  *string `form:"value" json:"value,omitempty"`
}

// A list of custom field filters for notification types that support advanced filtering
type V1AlertsCreateAlertAlertSpecifierExcludeCustomFieldFilterParams struct {
	Entity *string `form:"entity" json:"entity"`
	Key    *string `form:"key" json:"key"`
	Value  *string `form:"value" json:"value"`
}

// If provided, the specifier will not apply to balances that matches the inclusion criteria and any of the excluding values.
type V1AlertsCreateAlertAlertSpecifierExcludeParams struct {
	// A list of custom field filters for notification types that support advanced filtering
	CustomFieldFilters []*V1AlertsCreateAlertAlertSpecifierExcludeCustomFieldFilterParams `form:"custom_field_filters" json:"custom_field_filters,omitempty"`
}

// Can be used only with `low_remaining_contract_credit_and_commit_balance_reached` and `low_remaining_contract_credit_and_commit_percentage_reached` notifications. Defines the commits and credits used to calculate the remaining balance or percentage.
type V1AlertsCreateAlertAlertSpecifierParams struct {
	// A list of custom field filters for notification types that support advanced filtering
	CustomFieldFilters []*V1AlertsCreateAlertAlertSpecifierCustomFieldFilterParams `form:"custom_field_filters" json:"custom_field_filters,omitempty"`
	// If provided, the specifier will not apply to balances that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1AlertsCreateAlertAlertSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
}

// Create a new threshold notification to monitor customer spending, balances, and billing metrics in real-time. Metronome's notification system provides industry-leading speed with immediate evaluation capabilities, enabling you to proactively manage customer accounts and prevent revenue leakage.
//
// This endpoint creates configurable threshold notifications that continuously monitor various billing thresholds including spend limits, credit balances, commitment utilization, and invoice totals. Threshold notifications can be configured globally for all customers or targeted to specific customer accounts.
//
// ### Use this endpoint to:
// - Proactively monitor customer spending patterns to prevent unexpected overages or credit exhaustion
// - Automate notifications when customers approach commitment limits or credit thresholds
// - Enable real-time intervention for accounts at risk of churn or payment issues
// - Scale billing operations by automating threshold-based workflows and notifications
//
// ### Key response fields:
// A successful response returns a CustomerAlert object containing:
//
// - The threshold notification configuration with its unique ID and current status
// - The customer's evaluation status (ok, in_alarm, or evaluating)
// - Threshold notification metadata including type, threshold values, and update timestamps
//
// ### Usage guidelines:
// - Immediate evaluation: Set `evaluate_on_create` : `true` (default) for instant evaluation against existing customers
// - Uniqueness constraints: Each threshold notification must have a unique `uniqueness_key` within your organization. Use `release_uniqueness_key` : `true` when archiving to reuse keys
// - Notification type requirements: Different threshold notification types require specific fields (e.g., `billable_metric_id` for usage notifications, `credit_type_id` for credit-based threshold notifications)
// - Webhook delivery: Threshold notifications trigger webhook notifications for real-time integration with your systems. Configure webhook endpoints before creating threshold notifications
// - Performance at scale: Metronome's event-driven architecture processes threshold notification evaluations in real-time as usage events stream in, unlike competitors who rely on periodic polling or batch evaluation cycles
type V1AlertsCreateAlertParams struct {
	Params `form:"*"`
	// Filters the notification to commits/credits with this access type. Only supported for `low_remaining_commit_balance_reached`, `low_remaining_commit_percentage_reached`, `low_remaining_contract_credit_and_commit_balance_reached`, `low_remaining_contract_credit_and_commit_percentage_reached`, `low_remaining_contract_credit_balance_reached`, `low_remaining_contract_credit_percentage_reached`, and `low_remaining_seat_balance_reached` notifications. Credit type cannot be specified if using QUANTITY access type.
	AccessType *string `form:"access_type" json:"access_type,omitempty"`
	// Can be used only with `low_remaining_contract_credit_and_commit_balance_reached` and `low_remaining_contract_credit_and_commit_percentage_reached` notifications. Defines the commits and credits used to calculate the remaining balance or percentage.
	AlertSpecifiers []*V1AlertsCreateAlertAlertSpecifierParams `form:"alert_specifiers" json:"alert_specifiers,omitempty"`
	// Type of the threshold notification
	AlertType *string `form:"alert_type" json:"alert_type"`
	// For threshold notifications of type `usage_threshold_reached`, specifies which billable metric to track the usage for.
	BillableMetricID *string `form:"billable_metric_id" json:"billable_metric_id,omitempty"`
	// An array of strings, representing a way to filter the credit grant this threshold notification applies to, by looking at the credit_grant_type field on the credit grant. This field is only defined for CreditPercentage and CreditBalance notifications
	CreditGrantTypeFilters []*string `form:"credit_grant_type_filters" json:"credit_grant_type_filters,omitempty"`
	// ID of the credit's currency, defaults to USD. If the specific notification type requires a pricing unit/currency, find the ID in the [Metronome app](https://app.metronome.com/offering/pricing-units).
	CreditTypeID *string `form:"credit_type_id" json:"credit_type_id,omitempty"`
	// If provided, will create this threshold notification for this specific customer. To create a notification for all customers, do not specify a `customer_id`.
	CustomerID *string `form:"customer_id" json:"customer_id,omitempty"`
	// A list of custom field filters for threshold notification types that support advanced filtering. Only present for contract invoices.
	CustomFieldFilters []*V1AlertsCreateAlertCustomFieldFilterParams `form:"custom_field_filters" json:"custom_field_filters,omitempty"`
	// If true, the threshold notification will evaluate immediately on customers that already meet the notification threshold. If false, it will only evaluate on future customers that trigger the threshold. Defaults to true.
	EvaluateOnCreate *bool `form:"evaluate_on_create" json:"evaluate_on_create,omitempty"`
	// Only present for `spend_threshold_reached` notifications. Scope notification to a specific group key on individual line items.
	GroupValues        []*V1AlertsCreateAlertGroupValueParams `form:"group_values" json:"group_values,omitempty"`
	InvoiceTypesFilter []*string                              `form:"invoice_types_filter" json:"invoice_types_filter,omitempty"`
	// Name of the threshold notification
	Name *string `form:"name" json:"name"`
	// Required for `low_remaining_seat_balance_reached` notifications. The alert is scoped to this seat group key-value pair.
	SeatFilter *V1AlertsCreateAlertSeatFilterParams `form:"seat_filter" json:"seat_filter,omitempty"`
	// Threshold value of the notification policy.  Depending upon the notification type, this number may represent a financial amount, the days remaining, or a percentage reached.
	Threshold     *Number `form:"threshold" json:"threshold"`
	UniquenessKey *string `form:"uniqueness_key" json:"uniqueness_key,omitempty"`
}

// Only present for `spend_threshold_reached` notifications. Retrieve the notification for a specific group key-value pair.
type V1AlertsGetCustomerAlertGroupValueParams struct {
	Key   *string `form:"key" json:"key"`
	Value *string `form:"value" json:"value"`
}

// Only allowed for `low_remaining_seat_balance_reached` notifications. This filters alerts by the seat group key-value pair.
type V1AlertsGetCustomerAlertSeatFilterParams struct {
	// The seat group key (e.g., "seat_id", "user_id")
	SeatGroupKey *string `form:"seat_group_key" json:"seat_group_key"`
	// The specific seat identifier to filter by
	SeatGroupValue *string `form:"seat_group_value" json:"seat_group_value"`
}

// A list of custom field filters for notification types that support advanced filtering
type V1AlertsGetCustomerAlertAlertSpecifierCustomFieldFilterParams struct {
	Entity *string `form:"entity" json:"entity"`
	Key    *string `form:"key" json:"key"`
	Value  *string `form:"value" json:"value"`
}

// A list of custom field filters for notification types that support advanced filtering
type V1AlertsGetCustomerAlertAlertSpecifierExcludeCustomFieldFilterParams struct {
	Entity *string `form:"entity" json:"entity"`
	Key    *string `form:"key" json:"key"`
	Value  *string `form:"value" json:"value"`
}

// If provided, the specifier will not apply to balances that matches the inclusion criteria and any of the excluding values.
type V1AlertsGetCustomerAlertAlertSpecifierExcludeParams struct {
	// A list of custom field filters for notification types that support advanced filtering
	CustomFieldFilters []*V1AlertsGetCustomerAlertAlertSpecifierExcludeCustomFieldFilterParams `form:"custom_field_filters" json:"custom_field_filters,omitempty"`
}

// Can be used only with `low_remaining_contract_credit_and_commit_balance_reached` and `low_remaining_contract_credit_and_commit_percentage_reached` notifications. Used to filter the alert by the custom field key-value pair.
type V1AlertsGetCustomerAlertAlertSpecifierParams struct {
	// A list of custom field filters for notification types that support advanced filtering
	CustomFieldFilters []*V1AlertsGetCustomerAlertAlertSpecifierCustomFieldFilterParams `form:"custom_field_filters" json:"custom_field_filters"`
	// If provided, the specifier will not apply to balances that matches the inclusion criteria and any of the excluding values.
	Exclude []*V1AlertsGetCustomerAlertAlertSpecifierExcludeParams `form:"exclude" json:"exclude,omitempty"`
}

// Used to filter the alert by the custom field key-value pair.
type V1AlertsGetCustomerAlertCustomFieldFilterParams struct {
	Entity *string `form:"entity" json:"entity"`
	Key    *string `form:"key" json:"key"`
	Value  *string `form:"value" json:"value"`
}

// Retrieve the real-time evaluation status for a specific threshold notification-customer pair. This endpoint provides instant visibility into whether a customer has triggered a threshold notification condition, enabling you to monitor account health and take proactive action based on current threshold notification states.
//
// ### Use this endpoint to:
// - Check if a specific customer is currently violating an threshold notification (`in_alarm` status)
// - Verify threshold notification configuration details and threshold values for a customer
// - Monitor the evaluation state of newly created or recently modified threshold notification
// - Build dashboards or automated workflows that respond to specific threshold notification conditions
// - Validate threshold notification behavior before deploying to production customers
// - Integrate threshold notification status checks into customer support tools or admin interfaces
//
// ### Key response fields:
// A CustomerAlert object containing:
//
// - `customer_status`: The current evaluation state
//
// - `ok` - Customer is within acceptable thresholds
// - `in_alarm` - Customer has breached the threshold for the notification
// - `evaluating` - Notification is currently being evaluated (typically during initial setup)
// - `null` - Notification has been archived
// - `triggered_by`: Additional context about what caused the notification to trigger (when applicable)
// - `updated_at`: Timestamp of when the `customer_status` was last updated
// - alert: Complete threshold notification configuration including:
//   - Notification ID, name, and type
//   - Current threshold values and credit type information
//   - Notification status (enabled, disabled, or archived)
//   - Any applied filters (credit grant types, custom fields, group values)
//
// ### Usage guidelines:
// - Customer status: Returns the current evaluation state, not historical data. For threshold notification history, use webhook notifications or event logs
// - Required parameters: Both customer_id and alert_id must be valid UUIDs that exist in your organization
// - Archived notifications: Returns null for customer_status if the notification has been archived, but still includes the notification configuration details
// - Performance considerations: This endpoint queries live evaluation state, making it ideal for real-time monitoring but not for bulk status checks
// - Integration patterns: Best used for on-demand status checks in response to user actions or as part of targeted monitoring workflows
// - Error handling: Returns 404 if either the customer or alert_id doesn't exist or isn't accessible to your organization
type V1AlertsGetCustomerAlertParams struct {
	Params `form:"*"`
	// The Metronome ID of the threshold notification
	AlertID *string `form:"alert_id" json:"alert_id"`
	// Can be used only with `low_remaining_contract_credit_and_commit_balance_reached` and `low_remaining_contract_credit_and_commit_percentage_reached` notifications. Used to filter the alert by the custom field key-value pair.
	AlertSpecifiers []*V1AlertsGetCustomerAlertAlertSpecifierParams `form:"alert_specifiers" json:"alert_specifiers,omitempty"`
	// The Metronome ID of the customer
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// Used to filter the alert by the custom field key-value pair.
	CustomFieldFilters []*V1AlertsGetCustomerAlertCustomFieldFilterParams `form:"custom_field_filters" json:"custom_field_filters,omitempty"`
	// Only present for `spend_threshold_reached` notifications. Retrieve the notification for a specific group key-value pair.
	GroupValues []*V1AlertsGetCustomerAlertGroupValueParams `form:"group_values" json:"group_values,omitempty"`
	// When parallel threshold notifications are enabled during migration, this flag denotes whether to fetch notifications for plans or contracts.
	PlansOrContracts *string `form:"plans_or_contracts" json:"plans_or_contracts,omitempty"`
	// Only allowed for `low_remaining_seat_balance_reached` notifications. This filters alerts by the seat group key-value pair.
	SeatFilter *V1AlertsGetCustomerAlertSeatFilterParams `form:"seat_filter" json:"seat_filter,omitempty"`
	// Indicates that this API request was triggered by a webhook notification with the provided ID.
	WebhookNotificationID *string `form:"webhook_notification_id" json:"webhook_notification_id,omitempty"`
}

// Retrieve all threshold notification configurations and their current statuses for a specific customer in a single API call. This endpoint provides a comprehensive view of all threshold notification monitoring a customer account.
//
// ### Use this endpoint to:
// - Display all active threshold notifications for a customer in dashboards or admin panels
// - Quickly identify which threshold notifications a customer is currently triggering
// - Audit threshold notification coverage for specific accounts
// - Filter threshold notifications by status (enabled, disabled, or archived)
//
// ### Key response fields:
// - data: Array of CustomerAlert objects, each containing:
//   - Current evaluation status (`ok`, `in_alarm`, `evaluating`, or `null`)
//   - Complete threshold notification configuration and threshold details
//   - Threshold notification metadata including type, name, and last update time
//
// - next_page: Pagination cursor for retrieving additional results
//
// ### Usage guidelines:
// - Default behavior: Returns only enabled threshold notifications unless `alert_statuses` filter is specified
// - Pagination: Use the `next_page` cursor to retrieve all results for customers with many notifications
// - Performance: Efficiently retrieves multiple threshold notification statuses in a single request instead of making individual calls
// - Filtering: Pass the `alert_statuses` array to include disabled or archived threshold notifications in results
type V1AlertsListCustomerAlertsParams struct {
	Params `form:"*"`
	// Optionally filter by threshold notification status. If absent, only enabled notifications will be returned.
	AlertStatuses []*string `form:"alert_statuses" json:"alert_statuses,omitempty"`
	// The Metronome ID of the customer
	CustomerID *string `form:"customer_id" json:"customer_id"`
	// Cursor that indicates where the next page of results should start.
	NextPage *string `form:"next_page" json:"next_page,omitempty"`
}

// Force an immediate re-evaluation of a specific threshold notification for a customer, clearing any previous state and triggering a fresh assessment against current thresholds. This endpoint ensures threshold notification accuracy after configuration changes or data corrections.
//
// ### Use this endpoint to:
// - Clear false positive threshold notifications after fixing data issues
// - Re-evaluate threshold notifications after adjusting customer balances or credits
// - Test threshold notification behavior during development and debugging
// - Resolve stuck threshold notification that may be in an incorrect state
// - Trigger immediate evaluation after threshold modifications
//
// ### Key response fields:
// - 200 Success: Confirmation that the threshold notification has been reset and re-evaluation initiated
// - No response body is returned - the operation completes asynchronously
//
// ### Usage guidelines:
// - Immediate effect: Triggers re-evaluation instantly, which may result in new webhook notifications if thresholds are breached
// - State clearing: Removes any cached evaluation state, ensuring a fresh assessment
// - Use sparingly: Intended for exceptional cases, not routine operations
// - Asynchronous processing: The reset completes immediately, but re-evaluation happens in the background
type V1AlertsResetCustomerAlertsParams struct {
	Params `form:"*"`
	// The Metronome ID of the threshold notification
	AlertID *string `form:"alert_id" json:"alert_id"`
	// The Metronome ID of the customer
	CustomerID *string `form:"customer_id" json:"customer_id"`
}
