//
// File generated from our OpenAPI spec
//

package metronomeclient

import (
	"context"
	"errors"
	metronome "github.com/Metronome-Industries/metronome-go/v3"
	"github.com/Metronome-Industries/metronome-go/v3/option"
	"net/url"
)

func (c *Client) invoke(ctx context.Context, operation Operation, path string, pathParams map[string]string, opts []option.RequestOption) error {
	switch operation.ID {
	case "addCustomFieldKey-v1":
		if operation.Method != "POST" || operation.Path != "/v1/customFields/addKey" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.V1.CustomFields.AddKey(ctx, metronome.V1CustomFieldAddKeyParams{}, opts...)
	case "addManualBalanceLedgerEntry-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/addManualBalanceLedgerEntry" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.V1.Contracts.AddManualBalanceEntry(ctx, metronome.V1ContractAddManualBalanceEntryParams{}, opts...)
	case "addRate-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contract-pricing/rate-cards/addRate" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.RateCards.Rates.Add(ctx, metronome.V1ContractRateCardRateAddParams{}, opts...)
		return err
	case "addRates-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contract-pricing/rate-cards/addRates" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.RateCards.Rates.AddMany(ctx, metronome.V1ContractRateCardRateAddManyParams{}, opts...)
		return err
	case "amendContract-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/amend" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.Amend(ctx, metronome.V1ContractAmendParams{}, opts...)
		return err
	case "approveApprovalRequest-v1":
		if operation.Method != "POST" || operation.Path != "/v1/approval_requests/{id}/approve" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "archiveAlert-v1":
		if operation.Method != "POST" || operation.Path != "/v1/alerts/archive" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Alerts.Archive(ctx, metronome.V1AlertArchiveParams{}, opts...)
		return err
	case "archiveBillableMetric-v1":
		if operation.Method != "POST" || operation.Path != "/v1/billable-metrics/archive" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.BillableMetrics.Archive(ctx, metronome.V1BillableMetricArchiveParams{}, opts...)
		return err
	case "archiveCommit-v2":
		if operation.Method != "POST" || operation.Path != "/v2/contracts/commits/archive" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "archiveContract-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/archive" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.Archive(ctx, metronome.V1ContractArchiveParams{}, opts...)
		return err
	case "archiveCredit-v2":
		if operation.Method != "POST" || operation.Path != "/v2/contracts/credits/archive" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "archiveCreditType-v1":
		if operation.Method != "POST" || operation.Path != "/v1/credit-types/archive" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.PricingUnits.Archive(ctx, metronome.V1PricingUnitArchiveParams{}, opts...)
		return err
	case "archiveCustomer-v1":
		if operation.Method != "POST" || operation.Path != "/v1/customers/archive" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Customers.Archive(ctx, metronome.V1CustomerArchiveParams{}, opts...)
		return err
	case "archiveCustomerBillingProviderConfigurations-v1":
		if operation.Method != "POST" || operation.Path != "/v1/archiveCustomerBillingProviderConfigurations" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Customers.ArchiveBillingConfigurations(ctx, metronome.V1CustomerArchiveBillingConfigurationsParams{}, opts...)
		return err
	case "archiveCustomerRevenueSystemConfigurations-v1":
		if operation.Method != "POST" || operation.Path != "/v1/archiveCustomerRevenueSystemConfigurations" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "archiveNotificationConfig-v2":
		if operation.Method != "POST" || operation.Path != "/v2/notifications/archive" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V2.Notifications.Offset.Archive(ctx, metronome.V2NotificationOffsetArchiveParams{}, opts...)
		return err
	case "archivePackage-v1":
		if operation.Method != "POST" || operation.Path != "/v1/packages/archive" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Packages.Archive(ctx, metronome.V1PackageArchiveParams{}, opts...)
		return err
	case "archiveProductListItem-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contract-pricing/products/archive" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.Products.Archive(ctx, metronome.V1ContractProductArchiveParams{}, opts...)
		return err
	case "archiveRateCard-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contract-pricing/rate-cards/archive" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.RateCards.Archive(ctx, metronome.V1ContractRateCardArchiveParams{}, opts...)
		return err
	case "chargeSeats-v1":
		if operation.Method != "POST" || operation.Path != "/v1/customers/{customer_id}/invoices/invoice_seats" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "createAlert-v1":
		if operation.Method != "POST" || operation.Path != "/v1/alerts/create" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Alerts.New(ctx, metronome.V1AlertNewParams{}, opts...)
		return err
	case "createBillableMetric-v1":
		if operation.Method != "POST" || operation.Path != "/v1/billable-metrics" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "createBillableMetricV1-v1":
		if operation.Method != "POST" || operation.Path != "/v1/billable-metrics/create" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.BillableMetrics.New(ctx, metronome.V1BillableMetricNewParams{}, opts...)
		return err
	case "createContract-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/create" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.New(ctx, metronome.V1ContractNewParams{}, opts...)
		return err
	case "createCreditType-v1":
		if operation.Method != "POST" || operation.Path != "/v1/credit-types/create" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.PricingUnits.New(ctx, metronome.V1PricingUnitNewParams{}, opts...)
		return err
	case "createCustomer-v1":
		if operation.Method != "POST" || operation.Path != "/v1/customers" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Customers.New(ctx, metronome.V1CustomerNewParams{}, opts...)
		return err
	case "createCustomerCommit-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/customerCommits/create" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Customers.Commits.New(ctx, metronome.V1CustomerCommitNewParams{}, opts...)
		return err
	case "createCustomerCredit-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/customerCredits/create" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Customers.Credits.New(ctx, metronome.V1CustomerCreditNewParams{}, opts...)
		return err
	case "createCustomerWithContract-v1":
		if operation.Method != "POST" || operation.Path != "/v1/composite/createCustomerWithContract" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "createHistoricalContractUsageInvoices-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/createHistoricalInvoices" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.NewHistoricalInvoices(ctx, metronome.V1ContractNewHistoricalInvoicesParams{}, opts...)
		return err
	case "createNotificationConfig-v2":
		if operation.Method != "POST" || operation.Path != "/v2/notifications/create" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V2.Notifications.Offset.New(ctx, metronome.V2NotificationOffsetNewParams{}, opts...)
		return err
	case "createPackage-v1":
		if operation.Method != "POST" || operation.Path != "/v1/packages/create" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Packages.New(ctx, metronome.V1PackageNewParams{}, opts...)
		return err
	case "createProduct-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contract-pricing/products/create" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.Products.New(ctx, metronome.V1ContractProductNewParams{}, opts...)
		return err
	case "createRateCard-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contract-pricing/rate-cards/create" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.RateCards.New(ctx, metronome.V1ContractRateCardNewParams{}, opts...)
		return err
	case "deleteCustomFields-v1":
		if operation.Method != "POST" || operation.Path != "/v1/customFields/deleteValues" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.V1.CustomFields.DeleteValues(ctx, metronome.V1CustomFieldDeleteValuesParams{}, opts...)
	case "deleteStripeBillingSettings-v1":
		if operation.Method != "DELETE" || operation.Path != "/v1/client/billing-config/stripe" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "denyApprovalRequest-v1":
		if operation.Method != "POST" || operation.Path != "/v1/approval_requests/{id}/deny" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "disableCommitTrueup-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/commits/disableTrueup" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "disableCustomFieldKey-v1":
		if operation.Method != "POST" || operation.Path != "/v1/customFields/removeKey" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.V1.CustomFields.RemoveKey(ctx, metronome.V1CustomFieldRemoveKeyParams{}, opts...)
	case "editCommit-v2":
		if operation.Method != "POST" || operation.Path != "/v2/contracts/commits/edit" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V2.Contracts.EditCommit(ctx, metronome.V2ContractEditCommitParams{}, opts...)
		return err
	case "editContract-v2":
		if operation.Method != "POST" || operation.Path != "/v2/contracts/edit" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V2.Contracts.Edit(ctx, metronome.V2ContractEditParams{}, opts...)
		return err
	case "editCredit-v2":
		if operation.Method != "POST" || operation.Path != "/v2/contracts/credits/edit" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V2.Contracts.EditCredit(ctx, metronome.V2ContractEditCreditParams{}, opts...)
		return err
	case "editNotificationConfig-v2":
		if operation.Method != "POST" || operation.Path != "/v2/notifications/edit" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V2.Notifications.Offset.Edit(ctx, metronome.V2NotificationOffsetEditParams{}, opts...)
		return err
	case "embeddableDashboard-v1":
		if operation.Method != "POST" || operation.Path != "/v1/dashboards/getEmbeddableUrl" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Dashboards.GetEmbeddableURL(ctx, metronome.V1DashboardGetEmbeddableURLParams{}, opts...)
		return err
	case "getApprovalRequest-v1":
		if operation.Method != "GET" || operation.Path != "/v1/approval_requests/{id}" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "getAuditLogs-v1":
		if operation.Method != "GET" || operation.Path != "/v1/auditLogs" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "getBillableMetric-v1":
		if operation.Method != "GET" || operation.Path != "/v1/billable-metrics/{billable_metric_id}" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.BillableMetrics.Get(ctx, metronome.V1BillableMetricGetParams{BillableMetricID: url.PathEscape(pathParams["billable_metric_id"])}, opts...)
		return err
	case "getContract-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/get" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.Get(ctx, metronome.V1ContractGetParams{}, opts...)
		return err
	case "getContract-v2":
		if operation.Method != "POST" || operation.Path != "/v2/contracts/get" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V2.Contracts.Get(ctx, metronome.V2ContractGetParams{}, opts...)
		return err
	case "getContractEditHistory-v2":
		if operation.Method != "POST" || operation.Path != "/v2/contracts/getEditHistory" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V2.Contracts.GetEditHistory(ctx, metronome.V2ContractGetEditHistoryParams{}, opts...)
		return err
	case "getContractNamedSchedule-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/getNamedSchedule" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.RateCards.NamedSchedules.Get(ctx, metronome.V1ContractRateCardNamedScheduleGetParams{}, opts...)
		return err
	case "getContractRateSchedule-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/getContractRateSchedule" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.GetRateSchedule(ctx, metronome.V1ContractGetRateScheduleParams{}, opts...)
		return err
	case "getCustomer-v1":
		if operation.Method != "GET" || operation.Path != "/v1/customers/{customer_id}" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Customers.Get(ctx, metronome.V1CustomerGetParams{CustomerID: url.PathEscape(pathParams["customer_id"])}, opts...)
		return err
	case "getCustomerAlert-v1":
		if operation.Method != "POST" || operation.Path != "/v1/customer-alerts/get" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Customers.Alerts.Get(ctx, metronome.V1CustomerAlertGetParams{}, opts...)
		return err
	case "getCustomerBillingProviderConfigurations-v1":
		if operation.Method != "POST" || operation.Path != "/v1/getCustomerBillingProviderConfigurations" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Customers.GetBillingConfigurations(ctx, metronome.V1CustomerGetBillingConfigurationsParams{}, opts...)
		return err
	case "getCustomerNamedSchedule-v1":
		if operation.Method != "POST" || operation.Path != "/v1/customers/getNamedSchedule" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Customers.NamedSchedules.Get(ctx, metronome.V1CustomerNamedScheduleGetParams{}, opts...)
		return err
	case "getCustomerRevenueSystemConfigurations-v1":
		if operation.Method != "POST" || operation.Path != "/v1/getCustomerRevenueSystemConfigurations" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "getInvoice-v1":
		if operation.Method != "GET" || operation.Path != "/v1/customers/{customer_id}/invoices/{invoice_id}" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Customers.Invoices.Get(ctx, metronome.V1CustomerInvoiceGetParams{CustomerID: url.PathEscape(pathParams["customer_id"]), InvoiceID: url.PathEscape(pathParams["invoice_id"])}, opts...)
		return err
	case "getInvoicePdf-v1":
		if operation.Method != "GET" || operation.Path != "/v1/customers/{customer_id}/invoices/{invoice_id}/pdf" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Customers.Invoices.GetPdf(ctx, metronome.V1CustomerInvoiceGetPdfParams{CustomerID: url.PathEscape(pathParams["customer_id"]), InvoiceID: url.PathEscape(pathParams["invoice_id"])}, opts...)
		return err
	case "getNetBalance-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/customerBalances/getNetBalance" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.GetNetBalance(ctx, metronome.V1ContractGetNetBalanceParams{}, opts...)
		return err
	case "getNotificationConfig-v2":
		if operation.Method != "POST" || operation.Path != "/v2/notifications/get" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V2.Notifications.Offset.Get(ctx, metronome.V2NotificationOffsetGetParams{}, opts...)
		return err
	case "getPackage-v1":
		if operation.Method != "POST" || operation.Path != "/v1/packages/get" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Packages.Get(ctx, metronome.V1PackageGetParams{}, opts...)
		return err
	case "getPagedUsage-v1":
		if operation.Method != "POST" || operation.Path != "/v1/usage/groups" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "getProduct-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contract-pricing/products/get" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.Products.Get(ctx, metronome.V1ContractProductGetParams{}, opts...)
		return err
	case "getPurchasedSeats-v1":
		if operation.Method != "GET" || operation.Path != "/v1/customers/{customer_id}/purchasedSeats" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "getRateCard-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contract-pricing/rate-cards/get" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.RateCards.Get(ctx, metronome.V1ContractRateCardGetParams{}, opts...)
		return err
	case "getRateCardNamedSchedule-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contract-pricing/rate-cards/getNamedSchedule" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.NamedSchedules.Get(ctx, metronome.V1ContractNamedScheduleGetParams{}, opts...)
		return err
	case "getRateSchedule-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contract-pricing/rate-cards/getRateSchedule" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.RateCards.GetRateSchedule(ctx, metronome.V1ContractRateCardGetRateScheduleParams{}, opts...)
		return err
	case "getRates-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contract-pricing/rate-cards/getRates" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "getSeatsUsage-v1":
		if operation.Method != "POST" || operation.Path != "/v1/usage/seats" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "getServices-v1":
		if operation.Method != "GET" || operation.Path != "/v1/services" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Services.List(ctx, opts...)
		return err
	case "getSubscriptionQuantityHistory-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/getSubscriptionQuantityHistory" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.GetSubscriptionQuantityHistory(ctx, metronome.V1ContractGetSubscriptionQuantityHistoryParams{}, opts...)
		return err
	case "getSubscriptionSeatsHistory-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/getSubscriptionSeatsHistory" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.GetSubscriptionSeatsHistory(ctx, metronome.V1ContractGetSubscriptionSeatsHistoryParams{}, opts...)
		return err
	case "getSubscriptionStatus-v1":
		if operation.Method != "POST" || operation.Path != "/v1/subscriptions/getStatus" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "getUsageBatch-v1":
		if operation.Method != "POST" || operation.Path != "/v1/usage" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "getWebhookDelivery-v1":
		if operation.Method != "GET" || operation.Path != "/v1/webhooks/{webhook_id}/triggered_notifications/{triggered_notification_id}" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "ingest-v1":
		if operation.Method != "POST" || operation.Path != "/v1/ingest" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.V1.Usage.Ingest(ctx, metronome.V1UsageIngestParams{}, opts...)
	case "integrationCloudwatchLogger-v1":
		if operation.Method != "POST" || operation.Path != "/v1/integrations/log" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listAllBillableMetrics-v1":
		if operation.Method != "GET" || operation.Path != "/v1/billable-metrics" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listApprovalRequests-v1":
		if operation.Method != "GET" || operation.Path != "/v1/approval_requests" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listBillableMetrics-v1":
		if operation.Method != "GET" || operation.Path != "/v1/customers/{customer_id}/billable-metrics" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listBreakdownInvoices-v1":
		if operation.Method != "GET" || operation.Path != "/v1/customers/{customer_id}/invoices/breakdowns" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listConfiguredBillingProviders-v1":
		if operation.Method != "POST" || operation.Path != "/v1/listConfiguredBillingProviders" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Settings.BillingProviders.List(ctx, metronome.V1SettingBillingProviderListParams{}, opts...)
		return err
	case "listContracts-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/list" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listContracts-v2":
		if operation.Method != "POST" || operation.Path != "/v2/contracts/list" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listContractsNamedSchedules-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/listNamedSchedules" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listContractsOnPackage-v1":
		if operation.Method != "POST" || operation.Path != "/v1/packages/listContractsOnPackage" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listCreditTypes-v1":
		if operation.Method != "GET" || operation.Path != "/v1/credit-types/list" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listCustomFieldKeys-v1":
		if operation.Method != "POST" || operation.Path != "/v1/customFields/listKeys" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listCustomerAlerts-v1":
		if operation.Method != "POST" || operation.Path != "/v1/customer-alerts/list" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listCustomerBalances-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/customerBalances/list" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listCustomerCommits-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/customerCommits/list" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listCustomerCredits-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/customerCredits/list" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listCustomers-v1":
		if operation.Method != "GET" || operation.Path != "/v1/customers" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listHistoricalBalances-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/historicalBalances/list" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listInvoices-v1":
		if operation.Method != "GET" || operation.Path != "/v1/customers/{customer_id}/invoices" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listOffsetNotificationConfigs-v2":
		if operation.Method != "POST" || operation.Path != "/v2/notifications/offset/list" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listPackages-v1":
		if operation.Method != "POST" || operation.Path != "/v1/packages/list" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listProducts-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contract-pricing/products/list" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listRateCards-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contract-pricing/rate-cards/list" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listSeatBalances-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/seatBalances/list" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.ListSeatBalances(ctx, metronome.V1ContractListSeatBalancesParams{}, opts...)
		return err
	case "listSpendBreakdownInvoices-v1":
		if operation.Method != "POST" || operation.Path != "/v1/customers/{customer_id}/invoices/spend-breakdowns" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "listSystemNotificationConfigs-v2":
		if operation.Method != "POST" || operation.Path != "/v2/notifications/system/list" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V2.Notifications.System.List(ctx, opts...)
		return err
	case "listWebhookDeliveries-v1":
		if operation.Method != "POST" || operation.Path != "/v1/webhooks/{webhook_id}/triggered_notifications" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "moveRateCardProducts-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contract-pricing/rate-cards/moveRateCardProducts" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.RateCards.ProductOrders.Update(ctx, metronome.V1ContractRateCardProductOrderUpdateParams{}, opts...)
		return err
	case "moveToHistory-v2":
		if operation.Method != "POST" || operation.Path != "/v2/contracts/commits/moveToHistory" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "previewCustomerEvents-v1":
		if operation.Method != "POST" || operation.Path != "/v1/customers/{customer_id}/previewEvents" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Customers.PreviewEvents(ctx, metronome.V1CustomerPreviewEventsParams{CustomerID: url.PathEscape(pathParams["customer_id"])}, opts...)
		return err
	case "regenerateInvoice-v1":
		if operation.Method != "POST" || operation.Path != "/v1/invoices/regenerate" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Invoices.Regenerate(ctx, metronome.V1InvoiceRegenerateParams{}, opts...)
		return err
	case "releaseExternalPaymentGateThresholdCommit-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/commits/threshold-billing/release" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "resetCustomerAlerts-v1":
		if operation.Method != "POST" || operation.Path != "/v1/customer-alerts/reset" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.V1.Customers.Alerts.Reset(ctx, metronome.V1CustomerAlertResetParams{}, opts...)
	case "scheduleProServicesInvoice-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/scheduleProServicesInvoice" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.ScheduleProServicesInvoice(ctx, metronome.V1ContractScheduleProServicesInvoiceParams{}, opts...)
		return err
	case "searchEvents-v1":
		if operation.Method != "POST" || operation.Path != "/v1/events/search" || operation.Read != true {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Usage.Search(ctx, metronome.V1UsageSearchParams{}, opts...)
		return err
	case "setCustomFields-v1":
		if operation.Method != "POST" || operation.Path != "/v1/customFields/setValues" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.V1.CustomFields.SetValues(ctx, metronome.V1CustomFieldSetValuesParams{}, opts...)
	case "setCustomerBillableStatus-v1":
		if operation.Method != "POST" || operation.Path != "/v1/customers/setBillableStatus" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "setCustomerBillingProviderConfigurations-v1":
		if operation.Method != "POST" || operation.Path != "/v1/setCustomerBillingProviderConfigurations" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Customers.SetBillingConfigurations(ctx, metronome.V1CustomerSetBillingConfigurationsParams{}, opts...)
		return err
	case "setCustomerName-v1":
		if operation.Method != "POST" || operation.Path != "/v1/customers/{customer_id}/setName" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Customers.SetName(ctx, metronome.V1CustomerSetNameParams{CustomerID: url.PathEscape(pathParams["customer_id"])}, opts...)
		return err
	case "setCustomerRevenueSystemConfigurations-v1":
		if operation.Method != "POST" || operation.Path != "/v1/setCustomerRevenueSystemConfigurations" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "setIngestAliases-v1":
		if operation.Method != "POST" || operation.Path != "/v1/customers/{customer_id}/setIngestAliases" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.V1.Customers.SetIngestAliases(ctx, metronome.V1CustomerSetIngestAliasesParams{CustomerID: url.PathEscape(pathParams["customer_id"])}, opts...)
	case "setRateCardProductsOrder-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contract-pricing/rate-cards/setRateCardProductsOrder" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.RateCards.ProductOrders.Set(ctx, metronome.V1ContractRateCardProductOrderSetParams{}, opts...)
		return err
	case "setUpBillingProvider-v1":
		if operation.Method != "POST" || operation.Path != "/v1/setUpBillingProvider" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Settings.BillingProviders.New(ctx, metronome.V1SettingBillingProviderNewParams{}, opts...)
		return err
	case "setUsageFilter-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/setUsageFilter" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.V1.Contracts.SetUsageFilter(ctx, metronome.V1ContractSetUsageFilterParams{}, opts...)
	case "submitApprovalRequests-v1":
		if operation.Method != "POST" || operation.Path != "/v1/approval_requests" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "updateActiveRechargeConfig-v1":
		if operation.Method != "POST" || operation.Path != "/v1/threshold-billing/update-active-recharge-config" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "updateBillableMetric-v1":
		if operation.Method != "PUT" || operation.Path != "/v1/billable-metrics/{billable_metric_id}" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.BillableMetrics.Update(ctx, metronome.V1BillableMetricUpdateParams{BillableMetricID: url.PathEscape(pathParams["billable_metric_id"])}, opts...)
		return err
	case "updateCommitEndDate-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/customerCommits/updateEndDate" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Customers.Commits.UpdateEndDate(ctx, metronome.V1CustomerCommitUpdateEndDateParams{}, opts...)
		return err
	case "updateContractEndDate-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/updateEndDate" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.UpdateEndDate(ctx, metronome.V1ContractUpdateEndDateParams{}, opts...)
		return err
	case "updateContractNamedSchedule-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/updateNamedSchedule" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.V1.Contracts.RateCards.NamedSchedules.Update(ctx, metronome.V1ContractRateCardNamedScheduleUpdateParams{}, opts...)
	case "updateCreditEndDate-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/customerCredits/updateEndDate" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Customers.Credits.UpdateEndDate(ctx, metronome.V1CustomerCreditUpdateEndDateParams{}, opts...)
		return err
	case "updateCustomerConfig-v1":
		if operation.Method != "POST" || operation.Path != "/v1/customers/{customer_id}/updateConfig" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.V1.Customers.UpdateConfig(ctx, metronome.V1CustomerUpdateConfigParams{CustomerID: url.PathEscape(pathParams["customer_id"])}, opts...)
	case "updateCustomerNamedSchedule-v1":
		if operation.Method != "POST" || operation.Path != "/v1/customers/updateNamedSchedule" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.V1.Customers.NamedSchedules.Update(ctx, metronome.V1CustomerNamedScheduleUpdateParams{}, opts...)
	case "updateInvoiceIssueDate-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contracts/updateInvoiceIssueDate" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.Execute(ctx, operation.Method, path, nil, nil, opts...)
	case "updateProduct-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contract-pricing/products/update" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.Products.Update(ctx, metronome.V1ContractProductUpdateParams{}, opts...)
		return err
	case "updateRateCard-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contract-pricing/rate-cards/update" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Contracts.RateCards.Update(ctx, metronome.V1ContractRateCardUpdateParams{}, opts...)
		return err
	case "updateRateCardNamedSchedule-v1":
		if operation.Method != "POST" || operation.Path != "/v1/contract-pricing/rate-cards/updateNamedSchedule" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		return c.SDK.V1.Contracts.NamedSchedules.Update(ctx, metronome.V1ContractNamedScheduleUpdateParams{}, opts...)
	case "voidInvoice-v1":
		if operation.Method != "POST" || operation.Path != "/v1/invoices/void" || operation.Read != false {
			return errors.New("Metronome operation does not match its reviewed binding")
		}
		_, err := c.SDK.V1.Invoices.Void(ctx, metronome.V1InvoiceVoidParams{}, opts...)
		return err
	default:
		return errors.New("Metronome operation has no reviewed SDK binding")
	}
}
