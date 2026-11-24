//
// File generated from our OpenAPI spec
//

package metronome

import "context"

type Client struct {
	V1CustomFields      *v1CustomFieldsService
	V1CreditsAndCommits *v1CreditsAndCommitsService
	V1RateCards         *v1RateCardsService
	V1Contracts         *v1ContractsService
	V1ApprovalRequests  *v1ApprovalRequestsService
	V1Alerts            *v1AlertsService
	V1BillableMetrics   *v1BillableMetricsService
	V2CreditsAndCommits *v2CreditsAndCommitsService
	V1Settings          *v1SettingsService
	V1Customers         *v1CustomersService
	V2Notifications     *v2NotificationsService
	V1Packages          *v1PackagesService
	V1Products          *v1ProductsService
	V1Invoices          *v1InvoicesService
	V1API               *v1APIService
	V2Contracts         *v2ContractsService
	V1Security          *v1SecurityService
	V1NamedSchedules    *v1NamedSchedulesService
	V1Usage             *v1UsageService
	V1Webhooks          *v1WebhooksService
}

func newClient(b Backend, key string) *Client {
	return &Client{
		V1CustomFields:      &v1CustomFieldsService{B: b, Key: key},
		V1CreditsAndCommits: &v1CreditsAndCommitsService{B: b, Key: key},
		V1RateCards:         &v1RateCardsService{B: b, Key: key},
		V1Contracts:         &v1ContractsService{B: b, Key: key},
		V1ApprovalRequests:  &v1ApprovalRequestsService{B: b, Key: key},
		V1Alerts:            &v1AlertsService{B: b, Key: key},
		V1BillableMetrics:   &v1BillableMetricsService{B: b, Key: key},
		V2CreditsAndCommits: &v2CreditsAndCommitsService{B: b, Key: key},
		V1Settings:          &v1SettingsService{B: b, Key: key},
		V1Customers:         &v1CustomersService{B: b, Key: key},
		V2Notifications:     &v2NotificationsService{B: b, Key: key},
		V1Packages:          &v1PackagesService{B: b, Key: key},
		V1Products:          &v1ProductsService{B: b, Key: key},
		V1Invoices:          &v1InvoicesService{B: b, Key: key},
		V1API:               &v1APIService{B: b, Key: key},
		V2Contracts:         &v2ContractsService{B: b, Key: key},
		V1Security:          &v1SecurityService{B: b, Key: key},
		V1NamedSchedules:    &v1NamedSchedulesService{B: b, Key: key},
		V1Usage:             &v1UsageService{B: b, Key: key},
		V1Webhooks:          &v1WebhooksService{B: b, Key: key},
	}
}

func (c *Client) invoke(ctx context.Context, operation Operation, pathParams map[string]string, raw *rawRequest) error {
	switch operation.ID {
	case "addCustomFieldKey-v1":
		_, err := c.V1CustomFields.AddCustomFieldKey(ctx, &V1CustomFieldsAddCustomFieldKeyParams{Params: Params{raw: raw}})
		return err
	case "addManualBalanceLedgerEntry-v1":
		_, err := c.V1CreditsAndCommits.AddManualBalanceLedgerEntry(ctx, &V1CreditsAndCommitsAddManualBalanceLedgerEntryParams{Params: Params{raw: raw}})
		return err
	case "addRate-v1":
		_, err := c.V1RateCards.AddRate(ctx, &V1RateCardsAddRateParams{Params: Params{raw: raw}})
		return err
	case "addRates-v1":
		_, err := c.V1RateCards.AddRates(ctx, &V1RateCardsAddRatesParams{Params: Params{raw: raw}})
		return err
	case "amendContract-v1":
		_, err := c.V1Contracts.AmendContract(ctx, &V1ContractsAmendContractParams{Params: Params{raw: raw}})
		return err
	case "approveApprovalRequest-v1":
		_, err := c.V1ApprovalRequests.ApproveApprovalRequest(ctx, pathParams["id"], &V1ApprovalRequestsApproveApprovalRequestParams{Params: Params{raw: raw}})
		return err
	case "archiveAlert-v1":
		_, err := c.V1Alerts.ArchiveAlert(ctx, &V1AlertsArchiveAlertParams{Params: Params{raw: raw}})
		return err
	case "archiveBillableMetric-v1":
		_, err := c.V1BillableMetrics.ArchiveBillableMetric(ctx, &V1BillableMetricsArchiveBillableMetricParams{Params: Params{raw: raw}})
		return err
	case "archiveCommit-v2":
		_, err := c.V2CreditsAndCommits.ArchiveCommit(ctx, &V2CreditsAndCommitsArchiveCommitParams{Params: Params{raw: raw}})
		return err
	case "archiveContract-v1":
		_, err := c.V1Contracts.ArchiveContract(ctx, &V1ContractsArchiveContractParams{Params: Params{raw: raw}})
		return err
	case "archiveCredit-v2":
		_, err := c.V2CreditsAndCommits.ArchiveCredit(ctx, &V2CreditsAndCommitsArchiveCreditParams{Params: Params{raw: raw}})
		return err
	case "archiveCreditType-v1":
		_, err := c.V1Settings.ArchiveCreditType(ctx, &V1SettingsArchiveCreditTypeParams{Params: Params{raw: raw}})
		return err
	case "archiveCustomer-v1":
		_, err := c.V1Customers.ArchiveCustomer(ctx, &V1CustomersArchiveCustomerParams{Params: Params{raw: raw}})
		return err
	case "archiveCustomerBillingProviderConfigurations-v1":
		_, err := c.V1Customers.ArchiveCustomerBillingProviderConfigurations(ctx, &V1CustomersArchiveCustomerBillingProviderConfigurationsParams{Params: Params{raw: raw}})
		return err
	case "archiveCustomerRevenueSystemConfigurations-v1":
		_, err := c.V1Customers.ArchiveCustomerRevenueSystemConfigurations(ctx, &V1CustomersArchiveCustomerRevenueSystemConfigurationsParams{Params: Params{raw: raw}})
		return err
	case "archiveNotificationConfig-v2":
		_, err := c.V2Notifications.ArchiveNotificationConfig(ctx, &V2NotificationsArchiveNotificationConfigParams{Params: Params{raw: raw}})
		return err
	case "archivePackage-v1":
		_, err := c.V1Packages.ArchivePackage(ctx, &V1PackagesArchivePackageParams{Params: Params{raw: raw}})
		return err
	case "archiveProductListItem-v1":
		_, err := c.V1Products.ArchiveProductListItem(ctx, &V1ProductsArchiveProductListItemParams{Params: Params{raw: raw}})
		return err
	case "archiveRateCard-v1":
		_, err := c.V1RateCards.ArchiveRateCard(ctx, &V1RateCardsArchiveRateCardParams{Params: Params{raw: raw}})
		return err
	case "chargeSeats-v1":
		_, err := c.V1Invoices.ChargeSeats(ctx, pathParams["customer_id"], &V1InvoicesChargeSeatsParams{Params: Params{raw: raw}})
		return err
	case "createAlert-v1":
		_, err := c.V1Alerts.CreateAlert(ctx, &V1AlertsCreateAlertParams{Params: Params{raw: raw}})
		return err
	case "createBillableMetric-v1":
		_, err := c.V1API.CreateBillableMetric(ctx, &V1APICreateBillableMetricParams{Params: Params{raw: raw}})
		return err
	case "createBillableMetricV1-v1":
		_, err := c.V1BillableMetrics.CreateBillableMetricV1(ctx, &V1BillableMetricsCreateBillableMetricV1Params{Params: Params{raw: raw}})
		return err
	case "createContract-v1":
		_, err := c.V1Contracts.CreateContract(ctx, &V1ContractsCreateContractParams{Params: Params{raw: raw}})
		return err
	case "createCreditType-v1":
		_, err := c.V1Settings.CreateCreditType(ctx, &V1SettingsCreateCreditTypeParams{Params: Params{raw: raw}})
		return err
	case "createCustomer-v1":
		_, err := c.V1Customers.CreateCustomer(ctx, &V1CustomersCreateCustomerParams{Params: Params{raw: raw}})
		return err
	case "createCustomerCommit-v1":
		_, err := c.V1CreditsAndCommits.CreateCustomerCommit(ctx, &V1CreditsAndCommitsCreateCustomerCommitParams{Params: Params{raw: raw}})
		return err
	case "createCustomerCredit-v1":
		_, err := c.V1CreditsAndCommits.CreateCustomerCredit(ctx, &V1CreditsAndCommitsCreateCustomerCreditParams{Params: Params{raw: raw}})
		return err
	case "createCustomerWithContract-v1":
		_, err := c.V1Contracts.CreateCustomerWithContract(ctx, &V1ContractsCreateCustomerWithContractParams{Params: Params{raw: raw}})
		return err
	case "createHistoricalContractUsageInvoices-v1":
		_, err := c.V1Contracts.CreateHistoricalContractUsageInvoices(ctx, &V1ContractsCreateHistoricalContractUsageInvoicesParams{Params: Params{raw: raw}})
		return err
	case "createNotificationConfig-v2":
		_, err := c.V2Notifications.CreateNotificationConfig(ctx, &V2NotificationsCreateNotificationConfigParams{Params: Params{raw: raw}})
		return err
	case "createPackage-v1":
		_, err := c.V1Contracts.CreatePackage(ctx, &V1ContractsCreatePackageParams{Params: Params{raw: raw}})
		return err
	case "createProduct-v1":
		_, err := c.V1Products.CreateProduct(ctx, &V1ProductsCreateProductParams{Params: Params{raw: raw}})
		return err
	case "createRateCard-v1":
		_, err := c.V1RateCards.CreateRateCard(ctx, &V1RateCardsCreateRateCardParams{Params: Params{raw: raw}})
		return err
	case "deleteCustomFields-v1":
		_, err := c.V1CustomFields.DeleteCustomFields(ctx, &V1CustomFieldsDeleteCustomFieldsParams{Params: Params{raw: raw}})
		return err
	case "denyApprovalRequest-v1":
		_, err := c.V1ApprovalRequests.DenyApprovalRequest(ctx, pathParams["id"], &V1ApprovalRequestsDenyApprovalRequestParams{Params: Params{raw: raw}})
		return err
	case "disableCommitTrueup-v1":
		_, err := c.V1CreditsAndCommits.DisableCommitTrueup(ctx, &V1CreditsAndCommitsDisableCommitTrueupParams{Params: Params{raw: raw}})
		return err
	case "disableCustomFieldKey-v1":
		_, err := c.V1CustomFields.DisableCustomFieldKey(ctx, &V1CustomFieldsDisableCustomFieldKeyParams{Params: Params{raw: raw}})
		return err
	case "editCommit-v2":
		_, err := c.V2CreditsAndCommits.EditCommit(ctx, &V2CreditsAndCommitsEditCommitParams{Params: Params{raw: raw}})
		return err
	case "editContract-v2":
		_, err := c.V2Contracts.EditContract(ctx, &V2ContractsEditContractParams{Params: Params{raw: raw}})
		return err
	case "editCredit-v2":
		_, err := c.V2CreditsAndCommits.EditCredit(ctx, &V2CreditsAndCommitsEditCreditParams{Params: Params{raw: raw}})
		return err
	case "editNotificationConfig-v2":
		_, err := c.V2Notifications.EditNotificationConfig(ctx, &V2NotificationsEditNotificationConfigParams{Params: Params{raw: raw}})
		return err
	case "embeddableDashboard-v1":
		_, err := c.V1Customers.EmbeddableDashboard(ctx, &V1CustomersEmbeddableDashboardParams{Params: Params{raw: raw}})
		return err
	case "getApprovalRequest-v1":
		_, err := c.V1ApprovalRequests.GetApprovalRequest(ctx, pathParams["id"], &V1ApprovalRequestsGetApprovalRequestParams{Params: Params{raw: raw}})
		return err
	case "getAuditLogs-v1":
		_, err := c.V1Security.GetAuditLogs(ctx, &V1SecurityGetAuditLogsParams{Params: Params{raw: raw}})
		return err
	case "getBillableMetric-v1":
		_, err := c.V1BillableMetrics.GetBillableMetric(ctx, pathParams["billable_metric_id"], &V1BillableMetricsGetBillableMetricParams{Params: Params{raw: raw}})
		return err
	case "getContract-v1":
		_, err := c.V1Contracts.GetContract(ctx, &V1ContractsGetContractParams{Params: Params{raw: raw}})
		return err
	case "getContract-v2":
		_, err := c.V2Contracts.GetContract(ctx, &V2ContractsGetContractParams{Params: Params{raw: raw}})
		return err
	case "getContractEditHistory-v2":
		_, err := c.V2Contracts.GetContractEditHistory(ctx, &V2ContractsGetContractEditHistoryParams{Params: Params{raw: raw}})
		return err
	case "getContractNamedSchedule-v1":
		_, err := c.V1NamedSchedules.GetContractNamedSchedule(ctx, &V1NamedSchedulesGetContractNamedScheduleParams{Params: Params{raw: raw}})
		return err
	case "getContractRateSchedule-v1":
		_, err := c.V1Contracts.GetContractRateSchedule(ctx, &V1ContractsGetContractRateScheduleParams{Params: Params{raw: raw}})
		return err
	case "getCustomer-v1":
		_, err := c.V1Customers.GetCustomer(ctx, pathParams["customer_id"], &V1CustomersGetCustomerParams{Params: Params{raw: raw}})
		return err
	case "getCustomerAlert-v1":
		_, err := c.V1Alerts.GetCustomerAlert(ctx, &V1AlertsGetCustomerAlertParams{Params: Params{raw: raw}})
		return err
	case "getCustomerBillingProviderConfigurations-v1":
		_, err := c.V1Customers.GetCustomerBillingProviderConfigurations(ctx, &V1CustomersGetCustomerBillingProviderConfigurationsParams{Params: Params{raw: raw}})
		return err
	case "getCustomerNamedSchedule-v1":
		_, err := c.V1NamedSchedules.GetCustomerNamedSchedule(ctx, &V1NamedSchedulesGetCustomerNamedScheduleParams{Params: Params{raw: raw}})
		return err
	case "getCustomerRevenueSystemConfigurations-v1":
		_, err := c.V1Customers.GetCustomerRevenueSystemConfigurations(ctx, &V1CustomersGetCustomerRevenueSystemConfigurationsParams{Params: Params{raw: raw}})
		return err
	case "getInvoice-v1":
		_, err := c.V1Invoices.GetInvoice(ctx, pathParams["invoice_id"], &V1InvoicesGetInvoiceParams{Params: Params{raw: raw}, CustomerID: String(pathParams["customer_id"])})
		return err
	case "getInvoicePdf-v1":
		_, err := c.V1Invoices.GetInvoicePDF(ctx, pathParams["invoice_id"], &V1InvoicesGetInvoicePDFParams{Params: Params{raw: raw}, CustomerID: String(pathParams["customer_id"])})
		return err
	case "getNetBalance-v1":
		_, err := c.V1CreditsAndCommits.GetNetBalance(ctx, &V1CreditsAndCommitsGetNetBalanceParams{Params: Params{raw: raw}})
		return err
	case "getNotificationConfig-v2":
		_, err := c.V2Notifications.GetNotificationConfig(ctx, &V2NotificationsGetNotificationConfigParams{Params: Params{raw: raw}})
		return err
	case "getPackage-v1":
		_, err := c.V1Packages.GetPackage(ctx, &V1PackagesGetPackageParams{Params: Params{raw: raw}})
		return err
	case "getPagedUsage-v1":
		_, err := c.V1Usage.GetPagedUsage(ctx, &V1UsageGetPagedUsageParams{Params: Params{raw: raw}})
		return err
	case "getProduct-v1":
		_, err := c.V1Products.GetProduct(ctx, &V1ProductsGetProductParams{Params: Params{raw: raw}})
		return err
	case "getRateCard-v1":
		_, err := c.V1RateCards.GetRateCard(ctx, &V1RateCardsGetRateCardParams{Params: Params{raw: raw}})
		return err
	case "getRateCardNamedSchedule-v1":
		_, err := c.V1NamedSchedules.GetRateCardNamedSchedule(ctx, &V1NamedSchedulesGetRateCardNamedScheduleParams{Params: Params{raw: raw}})
		return err
	case "getRates-v1":
		_, err := c.V1RateCards.GetRates(ctx, &V1RateCardsGetRatesParams{Params: Params{raw: raw}})
		return err
	case "getRateSchedule-v1":
		_, err := c.V1RateCards.GetRateSchedule(ctx, &V1RateCardsGetRateScheduleParams{Params: Params{raw: raw}})
		return err
	case "getServices-v1":
		_, err := c.V1Security.GetServices(ctx, &V1SecurityGetServicesParams{Params: Params{raw: raw}})
		return err
	case "getSubscriptionQuantityHistory-v1":
		_, err := c.V1Contracts.GetSubscriptionQuantityHistory(ctx, &V1ContractsGetSubscriptionQuantityHistoryParams{Params: Params{raw: raw}})
		return err
	case "getSubscriptionSeatsHistory-v1":
		_, err := c.V1Contracts.GetSubscriptionSeatsHistory(ctx, &V1ContractsGetSubscriptionSeatsHistoryParams{Params: Params{raw: raw}})
		return err
	case "getSubscriptionStatus-v1":
		_, err := c.V1Contracts.GetSubscriptionStatus(ctx, &V1ContractsGetSubscriptionStatusParams{Params: Params{raw: raw}})
		return err
	case "getUsageBatch-v1":
		_, err := c.V1Usage.GetUsageBatch(ctx, &V1UsageGetUsageBatchParams{Params: Params{raw: raw}})
		return err
	case "getWebhookDelivery-v1":
		_, err := c.V1Webhooks.GetWebhookDelivery(ctx, pathParams["triggered_notification_id"], &V1WebhooksGetWebhookDeliveryParams{Params: Params{raw: raw}, WebhookID: String(pathParams["webhook_id"])})
		return err
	case "ingest-v1":
		_, err := c.V1Usage.Ingest(ctx, &V1UsageIngestParams{Params: Params{raw: raw}})
		return err
	case "listAllBillableMetrics-v1":
		_, err := c.V1BillableMetrics.ListAllBillableMetrics(ctx, &V1BillableMetricsListAllBillableMetricsParams{Params: Params{raw: raw}})
		return err
	case "listApprovalRequests-v1":
		_, err := c.V1ApprovalRequests.ListApprovalRequests(ctx, &V1ApprovalRequestsListApprovalRequestsParams{Params: Params{raw: raw}})
		return err
	case "listBillableMetrics-v1":
		_, err := c.V1BillableMetrics.ListBillableMetrics(ctx, pathParams["customer_id"], &V1BillableMetricsListBillableMetricsParams{Params: Params{raw: raw}})
		return err
	case "listBreakdownInvoices-v1":
		_, err := c.V1Invoices.ListBreakdownInvoices(ctx, pathParams["customer_id"], &V1InvoicesListBreakdownInvoicesParams{Params: Params{raw: raw}})
		return err
	case "listConfiguredBillingProviders-v1":
		_, err := c.V1Settings.ListConfiguredBillingProviders(ctx, &V1SettingsListConfiguredBillingProvidersParams{Params: Params{raw: raw}})
		return err
	case "listContracts-v1":
		_, err := c.V1Contracts.ListContracts(ctx, &V1ContractsListContractsParams{Params: Params{raw: raw}})
		return err
	case "listContracts-v2":
		_, err := c.V2Contracts.ListContracts(ctx, &V2ContractsListContractsParams{Params: Params{raw: raw}})
		return err
	case "listContractsNamedSchedules-v1":
		_, err := c.V1NamedSchedules.ListContractsNamedSchedules(ctx, &V1NamedSchedulesListContractsNamedSchedulesParams{Params: Params{raw: raw}})
		return err
	case "listContractsOnPackage-v1":
		_, err := c.V1Contracts.ListContractsOnPackage(ctx, &V1ContractsListContractsOnPackageParams{Params: Params{raw: raw}})
		return err
	case "listCreditTypes-v1":
		_, err := c.V1Settings.ListCreditTypes(ctx, &V1SettingsListCreditTypesParams{Params: Params{raw: raw}})
		return err
	case "listCustomerAlerts-v1":
		_, err := c.V1Alerts.ListCustomerAlerts(ctx, &V1AlertsListCustomerAlertsParams{Params: Params{raw: raw}})
		return err
	case "listCustomerBalances-v1":
		_, err := c.V1CreditsAndCommits.ListCustomerBalances(ctx, &V1CreditsAndCommitsListCustomerBalancesParams{Params: Params{raw: raw}})
		return err
	case "listCustomerCommits-v1":
		_, err := c.V1CreditsAndCommits.ListCustomerCommits(ctx, &V1CreditsAndCommitsListCustomerCommitsParams{Params: Params{raw: raw}})
		return err
	case "listCustomerCredits-v1":
		_, err := c.V1CreditsAndCommits.ListCustomerCredits(ctx, &V1CreditsAndCommitsListCustomerCreditsParams{Params: Params{raw: raw}})
		return err
	case "listCustomers-v1":
		_, err := c.V1Customers.ListCustomers(ctx, &V1CustomersListCustomersParams{Params: Params{raw: raw}})
		return err
	case "listCustomFieldKeys-v1":
		_, err := c.V1CustomFields.ListCustomFieldKeys(ctx, &V1CustomFieldsListCustomFieldKeysParams{Params: Params{raw: raw}})
		return err
	case "listHistoricalBalances-v1":
		_, err := c.V1CreditsAndCommits.ListHistoricalBalances(ctx, &V1CreditsAndCommitsListHistoricalBalancesParams{Params: Params{raw: raw}})
		return err
	case "listInvoices-v1":
		_, err := c.V1Invoices.ListInvoices(ctx, pathParams["customer_id"], &V1InvoicesListInvoicesParams{Params: Params{raw: raw}})
		return err
	case "listOffsetNotificationConfigs-v2":
		_, err := c.V2Notifications.ListOffsetNotificationConfigs(ctx, &V2NotificationsListOffsetNotificationConfigsParams{Params: Params{raw: raw}})
		return err
	case "listPackages-v1":
		_, err := c.V1API.ListPackages(ctx, &V1APIListPackagesParams{Params: Params{raw: raw}})
		return err
	case "listProducts-v1":
		_, err := c.V1Products.ListProducts(ctx, &V1ProductsListProductsParams{Params: Params{raw: raw}})
		return err
	case "listRateCards-v1":
		_, err := c.V1RateCards.ListRateCards(ctx, &V1RateCardsListRateCardsParams{Params: Params{raw: raw}})
		return err
	case "listSeatBalances-v1":
		_, err := c.V1CreditsAndCommits.ListSeatBalances(ctx, &V1CreditsAndCommitsListSeatBalancesParams{Params: Params{raw: raw}})
		return err
	case "listSpendBreakdownInvoices-v1":
		_, err := c.V1Invoices.ListSpendBreakdownInvoices(ctx, pathParams["customer_id"], &V1InvoicesListSpendBreakdownInvoicesParams{Params: Params{raw: raw}})
		return err
	case "listSystemNotificationConfigs-v2":
		_, err := c.V2Notifications.ListSystemNotificationConfigs(ctx, &V2NotificationsListSystemNotificationConfigsParams{Params: Params{raw: raw}})
		return err
	case "listWebhookDeliveries-v1":
		_, err := c.V1Webhooks.ListWebhookDeliveries(ctx, pathParams["webhook_id"], &V1WebhooksListWebhookDeliveriesParams{Params: Params{raw: raw}})
		return err
	case "moveRateCardProducts-v1":
		_, err := c.V1RateCards.MoveRateCardProducts(ctx, &V1RateCardsMoveRateCardProductsParams{Params: Params{raw: raw}})
		return err
	case "moveToHistory-v2":
		_, err := c.V2CreditsAndCommits.MoveToHistory(ctx, &V2CreditsAndCommitsMoveToHistoryParams{Params: Params{raw: raw}})
		return err
	case "previewCustomerEvents-v1":
		_, err := c.V1Invoices.PreviewCustomerEvents(ctx, pathParams["customer_id"], &V1InvoicesPreviewCustomerEventsParams{Params: Params{raw: raw}})
		return err
	case "regenerateInvoice-v1":
		_, err := c.V1Invoices.RegenerateInvoice(ctx, &V1InvoicesRegenerateInvoiceParams{Params: Params{raw: raw}})
		return err
	case "releaseExternalPaymentGateThresholdCommit-v1":
		_, err := c.V1CreditsAndCommits.ReleaseExternalPaymentGateThresholdCommit(ctx, &V1CreditsAndCommitsReleaseExternalPaymentGateThresholdCommitParams{Params: Params{raw: raw}})
		return err
	case "resetCustomerAlerts-v1":
		_, err := c.V1Alerts.ResetCustomerAlerts(ctx, &V1AlertsResetCustomerAlertsParams{Params: Params{raw: raw}})
		return err
	case "scheduleProServicesInvoice-v1":
		_, err := c.V1Contracts.ScheduleProServicesInvoice(ctx, &V1ContractsScheduleProServicesInvoiceParams{Params: Params{raw: raw}})
		return err
	case "searchEvents-v1":
		_, err := c.V1Usage.SearchEvents(ctx, &V1UsageSearchEventsParams{Params: Params{raw: raw}})
		return err
	case "setCustomerBillingProviderConfigurations-v1":
		_, err := c.V1Customers.SetCustomerBillingProviderConfigurations(ctx, &V1CustomersSetCustomerBillingProviderConfigurationsParams{Params: Params{raw: raw}})
		return err
	case "setCustomerName-v1":
		_, err := c.V1Customers.SetCustomerName(ctx, pathParams["customer_id"], &V1CustomersSetCustomerNameParams{Params: Params{raw: raw}})
		return err
	case "setCustomerRevenueSystemConfigurations-v1":
		_, err := c.V1Customers.SetCustomerRevenueSystemConfigurations(ctx, &V1CustomersSetCustomerRevenueSystemConfigurationsParams{Params: Params{raw: raw}})
		return err
	case "setCustomFields-v1":
		_, err := c.V1CustomFields.SetCustomFields(ctx, &V1CustomFieldsSetCustomFieldsParams{Params: Params{raw: raw}})
		return err
	case "setIngestAliases-v1":
		_, err := c.V1Customers.SetIngestAliases(ctx, pathParams["customer_id"], &V1CustomersSetIngestAliasesParams{Params: Params{raw: raw}})
		return err
	case "setRateCardProductsOrder-v1":
		_, err := c.V1RateCards.SetRateCardProductsOrder(ctx, &V1RateCardsSetRateCardProductsOrderParams{Params: Params{raw: raw}})
		return err
	case "setUpBillingProvider-v1":
		_, err := c.V1Settings.SetUpBillingProvider(ctx, &V1SettingsSetUpBillingProviderParams{Params: Params{raw: raw}})
		return err
	case "setUsageFilter-v1":
		_, err := c.V1Contracts.SetUsageFilter(ctx, &V1ContractsSetUsageFilterParams{Params: Params{raw: raw}})
		return err
	case "submitApprovalRequests-v1":
		_, err := c.V1ApprovalRequests.SubmitApprovalRequests(ctx, &V1ApprovalRequestsSubmitApprovalRequestsParams{Params: Params{raw: raw}})
		return err
	case "updateBillableMetric-v1":
		_, err := c.V1BillableMetrics.UpdateBillableMetric(ctx, pathParams["billable_metric_id"], &V1BillableMetricsUpdateBillableMetricParams{Params: Params{raw: raw}})
		return err
	case "updateCommitEndDate-v1":
		_, err := c.V1CreditsAndCommits.UpdateCommitEndDate(ctx, &V1CreditsAndCommitsUpdateCommitEndDateParams{Params: Params{raw: raw}})
		return err
	case "updateContractEndDate-v1":
		_, err := c.V1Contracts.UpdateContractEndDate(ctx, &V1ContractsUpdateContractEndDateParams{Params: Params{raw: raw}})
		return err
	case "updateContractNamedSchedule-v1":
		_, err := c.V1NamedSchedules.UpdateContractNamedSchedule(ctx, &V1NamedSchedulesUpdateContractNamedScheduleParams{Params: Params{raw: raw}})
		return err
	case "updateCreditEndDate-v1":
		_, err := c.V1CreditsAndCommits.UpdateCreditEndDate(ctx, &V1CreditsAndCommitsUpdateCreditEndDateParams{Params: Params{raw: raw}})
		return err
	case "updateCustomerConfig-v1":
		_, err := c.V1Customers.UpdateCustomerConfig(ctx, pathParams["customer_id"], &V1CustomersUpdateCustomerConfigParams{Params: Params{raw: raw}})
		return err
	case "updateCustomerNamedSchedule-v1":
		_, err := c.V1NamedSchedules.UpdateCustomerNamedSchedule(ctx, &V1NamedSchedulesUpdateCustomerNamedScheduleParams{Params: Params{raw: raw}})
		return err
	case "updateInvoiceIssueDate-v1":
		_, err := c.V1Contracts.UpdateInvoiceIssueDate(ctx, &V1ContractsUpdateInvoiceIssueDateParams{Params: Params{raw: raw}})
		return err
	case "updateProduct-v1":
		_, err := c.V1Products.UpdateProduct(ctx, &V1ProductsUpdateProductParams{Params: Params{raw: raw}})
		return err
	case "updateRateCard-v1":
		_, err := c.V1RateCards.UpdateRateCard(ctx, &V1RateCardsUpdateRateCardParams{Params: Params{raw: raw}})
		return err
	case "updateRateCardNamedSchedule-v1":
		_, err := c.V1NamedSchedules.UpdateRateCardNamedSchedule(ctx, &V1NamedSchedulesUpdateRateCardNamedScheduleParams{Params: Params{raw: raw}})
		return err
	case "upsertAnrokApiToken-v1":
		_, err := c.V1Settings.UpsertAnrokAPIToken(ctx, &V1SettingsUpsertAnrokAPITokenParams{Params: Params{raw: raw}})
		return err
	case "upsertAvalaraCredentials-v1":
		_, err := c.V1Settings.UpsertAvalaraCredentials(ctx, &V1SettingsUpsertAvalaraCredentialsParams{Params: Params{raw: raw}})
		return err
	case "voidInvoice-v1":
		_, err := c.V1Invoices.VoidInvoice(ctx, &V1InvoicesVoidInvoiceParams{Params: Params{raw: raw}})
		return err
	default:
		return errUnknownOperation
	}
}
