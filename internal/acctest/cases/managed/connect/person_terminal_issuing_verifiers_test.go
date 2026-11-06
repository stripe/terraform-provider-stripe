// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	stripe "github.com/stripe/stripe-go/v86"

	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

type personExpectations struct {
	Address                   string
	ExpectedFirstName         string
	ExpectedLastName          string
	ExpectedEmail             string
	ExpectedPhone             string
	ExpectedAddressLine1      string
	ExpectedAddressLine2      string
	ExpectedCity              string
	ExpectedState             string
	ExpectedPostalCode        string
	ExpectedCountry           string
	ExpectedNationality       string
	ExpectedPoliticalExposure string
	ExpectedTitle             string
	ExpectedRepresentative    bool
	ExpectedExecutive         bool
	ExpectedDirector          bool
	CheckOwner                bool
	ExpectedOwner             bool
	CheckPercentOwnership     bool
	ExpectedPercentOwnership  float64
	ExpectedDOBDay            int64
	ExpectedDOBMonth          int64
	ExpectedDOBYear           int64
	ExpectedMetadata          map[string]string
}

type terminalLocationExpectations struct {
	Address                              string
	ExpectedDisplayName                  string
	ExpectedPhone                        string
	ExpectedAddressLine1                 string
	ExpectedAddressLine2                 string
	ExpectedCity                         string
	ExpectedState                        string
	ExpectedPostalCode                   string
	ExpectedCountry                      string
	ExpectedConfigurationOverrideAddress string
	ExpectedMetadata                     map[string]string
}

type terminalConfigurationExpectations struct {
	Address                          string
	ExpectedName                     string
	ExpectedOfflineEnabled           bool
	ExpectedCellularEnabled          bool
	ExpectedRebootWindowStartHour    int64
	ExpectedRebootWindowEndHour      int64
	ExpectedTippingUSDFixedAmounts   []int64
	ExpectedTippingUSDPercentages    []int64
	ExpectedTippingUSDSmartThreshold int64
}

type issuingCardholderExpectations struct {
	Address                        string
	ExpectedName                   string
	ExpectedEmail                  string
	ExpectedPhoneNumber            string
	ExpectedType                   string
	ExpectedStatus                 string
	ExpectedFirstName              string
	ExpectedLastName               string
	ExpectedBillingLine1           string
	ExpectedBillingCity            string
	ExpectedBillingPostalCode      string
	ExpectedBillingCountry         string
	ExpectedPreferredLocales       []string
	ExpectedAllowedCardPresences   []string
	ExpectedSpendingLimitAmount    int64
	ExpectedSpendingLimitInterval  string
	ExpectedSpendingLimitsCurrency string
	ExpectedMetadata               map[string]string
}

type issuingCardExpectations struct {
	Address                       string
	ExpectedCardholderAddress     string
	ExpectedCurrency              string
	ExpectedType                  string
	ExpectedStatus                string
	ExpectedLifecyclePaymentCount int64
	ExpectedMetadata              map[string]string
}

type issuingDisputeExpectations struct {
	Address                            string
	ExpectedStatus                     string
	ExpectedReason                     string
	ExpectedFraudulentExplanation      string
	ExpectedOtherExplanation           string
	ExpectedCanceledExplanation        string
	ExpectedCanceledAt                 int64
	ExpectedCancellationReason         string
	ExpectedCancellationPolicyProvided bool
	ExpectedExpectedAt                 int64
	ExpectedProductDescription         string
	ExpectedProductType                string
	ExpectedReturnStatus               string
	ExpectedReturnedAt                 int64
	ExpectedMetadata                   map[string]string
}

type issuingPersonalizationDesignExpectations struct {
	Address               string
	ExpectedName          string
	CompareStateLookupKey bool
	CheckCarrierText      bool
	ExpectedHeaderTitle   string
	ExpectedHeaderBody    string
	ExpectedFooterTitle   string
	ExpectedFooterBody    string
	CheckIsDefault        bool
	ExpectedIsDefault     bool
	ExpectedMetadata      map[string]string
}

func verifyPerson(expect personExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		person, err := retrievePerson(client, state, expect.Address)
		if err != nil {
			return err
		}
		accountID, err := runner.ResourceAttribute(state, expect.Address, "account")
		if err != nil {
			return err
		}
		if person.Account != accountID {
			return fmt.Errorf(
				"remote %s.account mismatch: expected %q, got %q",
				expect.Address,
				accountID,
				person.Account,
			)
		}
		if person.FirstName != expect.ExpectedFirstName {
			return fmt.Errorf(
				"remote %s.first_name mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedFirstName,
				person.FirstName,
			)
		}
		if person.LastName != expect.ExpectedLastName {
			return fmt.Errorf(
				"remote %s.last_name mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedLastName,
				person.LastName,
			)
		}
		if person.Email != expect.ExpectedEmail {
			return fmt.Errorf(
				"remote %s.email mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedEmail,
				person.Email,
			)
		}
		if person.Phone != expect.ExpectedPhone {
			return fmt.Errorf(
				"remote %s.phone mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedPhone,
				person.Phone,
			)
		}
		if person.Address == nil {
			return fmt.Errorf("remote %s.address missing", expect.Address)
		}
		if person.Address.Line1 != expect.ExpectedAddressLine1 {
			return fmt.Errorf(
				"remote %s.address.line1 mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedAddressLine1,
				person.Address.Line1,
			)
		}
		if person.Address.Line2 != expect.ExpectedAddressLine2 {
			return fmt.Errorf(
				"remote %s.address.line2 mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedAddressLine2,
				person.Address.Line2,
			)
		}
		if person.Address.City != expect.ExpectedCity {
			return fmt.Errorf(
				"remote %s.address.city mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedCity,
				person.Address.City,
			)
		}
		if person.Address.State != expect.ExpectedState {
			return fmt.Errorf(
				"remote %s.address.state mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedState,
				person.Address.State,
			)
		}
		if person.Address.PostalCode != expect.ExpectedPostalCode {
			return fmt.Errorf(
				"remote %s.address.postal_code mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedPostalCode,
				person.Address.PostalCode,
			)
		}
		if person.Address.Country != expect.ExpectedCountry {
			return fmt.Errorf(
				"remote %s.address.country mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedCountry,
				person.Address.Country,
			)
		}
		if expect.ExpectedNationality != "" && person.Nationality != expect.ExpectedNationality {
			return fmt.Errorf(
				"remote %s.nationality mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedNationality,
				person.Nationality,
			)
		}
		if expect.ExpectedPoliticalExposure != "" &&
			string(person.PoliticalExposure) != expect.ExpectedPoliticalExposure {
			return fmt.Errorf(
				"remote %s.political_exposure mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedPoliticalExposure,
				string(person.PoliticalExposure),
			)
		}
		if person.DOB == nil {
			return fmt.Errorf("remote %s.dob missing", expect.Address)
		}
		if person.DOB.Day != expect.ExpectedDOBDay ||
			person.DOB.Month != expect.ExpectedDOBMonth ||
			person.DOB.Year != expect.ExpectedDOBYear {
			return fmt.Errorf(
				"remote %s.dob mismatch: expected %d-%d-%d, got %d-%d-%d",
				expect.Address,
				expect.ExpectedDOBYear,
				expect.ExpectedDOBMonth,
				expect.ExpectedDOBDay,
				person.DOB.Year,
				person.DOB.Month,
				person.DOB.Day,
			)
		}
		if person.Relationship == nil {
			return fmt.Errorf("remote %s.relationship missing", expect.Address)
		}
		if person.Relationship.Representative != expect.ExpectedRepresentative {
			return fmt.Errorf(
				"remote %s.relationship.representative mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedRepresentative,
				person.Relationship.Representative,
			)
		}
		if person.Relationship.Executive != expect.ExpectedExecutive {
			return fmt.Errorf(
				"remote %s.relationship.executive mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedExecutive,
				person.Relationship.Executive,
			)
		}
		if person.Relationship.Director != expect.ExpectedDirector {
			return fmt.Errorf(
				"remote %s.relationship.director mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedDirector,
				person.Relationship.Director,
			)
		}
		if expect.CheckOwner && person.Relationship.Owner != expect.ExpectedOwner {
			return fmt.Errorf(
				"remote %s.relationship.owner mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedOwner,
				person.Relationship.Owner,
			)
		}
		if expect.CheckPercentOwnership &&
			person.Relationship.PercentOwnership != expect.ExpectedPercentOwnership {
			return fmt.Errorf(
				"remote %s.relationship.percent_ownership mismatch: expected %f, got %f",
				expect.Address,
				expect.ExpectedPercentOwnership,
				person.Relationship.PercentOwnership,
			)
		}
		if person.Relationship.Title != expect.ExpectedTitle {
			return fmt.Errorf(
				"remote %s.relationship.title mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedTitle,
				person.Relationship.Title,
			)
		}
		if err := expectMetadataSubset(
			expect.Address+".metadata",
			person.Metadata,
			expect.ExpectedMetadata,
		); err != nil {
			return err
		}

		return nil
	}
}

func verifyPersonDestroyMissing(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	id, err := runner.ResourcePrimaryID(state, "stripe_person.test")
	if err != nil {
		return err
	}
	accountID, err := runner.ResourceAttribute(state, "stripe_person.test", "account")
	if err != nil {
		return err
	}

	params := &stripe.PersonRetrieveParams{
		Account: stripe.String(accountID),
	}
	person, err := client.V1Persons.Retrieve(context.Background(), id, params)
	if err == nil && person != nil && person.Deleted {
		return nil
	}
	return expectRemoteMissing("stripe_person.test", id, err)
}

func verifyTerminalLocation(expect terminalLocationExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		location, err := retrieveTerminalLocation(client, state, expect.Address)
		if err != nil {
			return err
		}
		if location.DisplayName != expect.ExpectedDisplayName {
			return fmt.Errorf(
				"remote %s.display_name mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedDisplayName,
				location.DisplayName,
			)
		}
		if location.Phone != expect.ExpectedPhone {
			return fmt.Errorf(
				"remote %s.phone mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedPhone,
				location.Phone,
			)
		}
		if location.Address == nil {
			return fmt.Errorf("remote %s.address missing", expect.Address)
		}
		if location.Address.Line1 != expect.ExpectedAddressLine1 {
			return fmt.Errorf(
				"remote %s.address.line1 mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedAddressLine1,
				location.Address.Line1,
			)
		}
		if location.Address.Line2 != expect.ExpectedAddressLine2 {
			return fmt.Errorf(
				"remote %s.address.line2 mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedAddressLine2,
				location.Address.Line2,
			)
		}
		if location.Address.City != expect.ExpectedCity {
			return fmt.Errorf(
				"remote %s.address.city mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedCity,
				location.Address.City,
			)
		}
		if location.Address.State != expect.ExpectedState {
			return fmt.Errorf(
				"remote %s.address.state mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedState,
				location.Address.State,
			)
		}
		if location.Address.PostalCode != expect.ExpectedPostalCode {
			return fmt.Errorf(
				"remote %s.address.postal_code mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedPostalCode,
				location.Address.PostalCode,
			)
		}
		if location.Address.Country != expect.ExpectedCountry {
			return fmt.Errorf(
				"remote %s.address.country mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedCountry,
				location.Address.Country,
			)
		}
		if expect.ExpectedConfigurationOverrideAddress != "" {
			expectedConfigurationOverrideID, err := runner.ResourcePrimaryID(
				state,
				expect.ExpectedConfigurationOverrideAddress,
			)
			if err != nil {
				return err
			}
			if location.ConfigurationOverrides != expectedConfigurationOverrideID {
				return fmt.Errorf(
					"remote %s.configuration_overrides mismatch: expected %q, got %q",
					expect.Address,
					expectedConfigurationOverrideID,
					location.ConfigurationOverrides,
				)
			}
		}
		if err := expectMetadataSubset(
			expect.Address+".metadata",
			location.Metadata,
			expect.ExpectedMetadata,
		); err != nil {
			return err
		}

		return nil
	}
}

func verifyTerminalLocationDestroyMissing(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	id, err := runner.ResourcePrimaryID(state, "stripe_terminal_location.test")
	if err != nil {
		return err
	}

	location, err := client.V1TerminalLocations.Retrieve(context.Background(), id, nil)
	if err == nil && location != nil && location.Deleted {
		return nil
	}
	return expectRemoteMissing("stripe_terminal_location.test", id, err)
}

func verifyTerminalConfiguration(
	expect terminalConfigurationExpectations,
) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		configuration, err := retrieveTerminalConfiguration(client, state, expect.Address)
		if err != nil {
			return err
		}
		if configuration.Name != expect.ExpectedName {
			return fmt.Errorf(
				"remote %s.name mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedName,
				configuration.Name,
			)
		}
		if configuration.Offline == nil {
			return fmt.Errorf("remote %s.offline missing", expect.Address)
		}
		if configuration.Offline.Enabled != expect.ExpectedOfflineEnabled {
			return fmt.Errorf(
				"remote %s.offline.enabled mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedOfflineEnabled,
				configuration.Offline.Enabled,
			)
		}
		if configuration.Cellular == nil {
			return fmt.Errorf("remote %s.cellular missing", expect.Address)
		}
		if configuration.Cellular.Enabled != expect.ExpectedCellularEnabled {
			return fmt.Errorf(
				"remote %s.cellular.enabled mismatch: expected %t, got %t",
				expect.Address,
				expect.ExpectedCellularEnabled,
				configuration.Cellular.Enabled,
			)
		}
		if configuration.RebootWindow == nil {
			return fmt.Errorf("remote %s.reboot_window missing", expect.Address)
		}
		if configuration.RebootWindow.StartHour != expect.ExpectedRebootWindowStartHour {
			return fmt.Errorf(
				"remote %s.reboot_window.start_hour mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedRebootWindowStartHour,
				configuration.RebootWindow.StartHour,
			)
		}
		if configuration.RebootWindow.EndHour != expect.ExpectedRebootWindowEndHour {
			return fmt.Errorf(
				"remote %s.reboot_window.end_hour mismatch: expected %d, got %d",
				expect.Address,
				expect.ExpectedRebootWindowEndHour,
				configuration.RebootWindow.EndHour,
			)
		}
		if len(expect.ExpectedTippingUSDFixedAmounts) > 0 ||
			len(expect.ExpectedTippingUSDPercentages) > 0 ||
			expect.ExpectedTippingUSDSmartThreshold != 0 {
			if configuration.Tipping == nil || configuration.Tipping.USD == nil {
				return fmt.Errorf("remote %s.tipping.usd missing", expect.Address)
			}
			if err := expectInt64Slice(
				expect.Address+".tipping.usd.fixed_amounts",
				configuration.Tipping.USD.FixedAmounts,
				expect.ExpectedTippingUSDFixedAmounts,
			); err != nil {
				return err
			}
			if err := expectInt64Slice(
				expect.Address+".tipping.usd.percentages",
				configuration.Tipping.USD.Percentages,
				expect.ExpectedTippingUSDPercentages,
			); err != nil {
				return err
			}
			if configuration.Tipping.USD.SmartTipThreshold != expect.ExpectedTippingUSDSmartThreshold {
				return fmt.Errorf(
					"remote %s.tipping.usd.smart_tip_threshold mismatch: expected %d, got %d",
					expect.Address,
					expect.ExpectedTippingUSDSmartThreshold,
					configuration.Tipping.USD.SmartTipThreshold,
				)
			}
		}

		return nil
	}
}

func verifyTerminalConfigurationDestroyMissing(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	id, err := runner.ResourcePrimaryID(state, "stripe_terminal_configuration.test")
	if err != nil {
		return err
	}

	configuration, err := client.V1TerminalConfigurations.Retrieve(
		context.Background(),
		id,
		nil,
	)
	if err == nil && configuration != nil && configuration.Deleted {
		return nil
	}
	return expectRemoteMissing("stripe_terminal_configuration.test", id, err)
}

func verifyIssuingCardholder(expect issuingCardholderExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		cardholder, err := retrieveIssuingCardholder(client, state, expect.Address)
		if err != nil {
			return err
		}
		if cardholder.Name != expect.ExpectedName {
			return fmt.Errorf(
				"remote %s.name mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedName,
				cardholder.Name,
			)
		}
		if cardholder.Email != expect.ExpectedEmail {
			return fmt.Errorf(
				"remote %s.email mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedEmail,
				cardholder.Email,
			)
		}
		if cardholder.PhoneNumber != expect.ExpectedPhoneNumber {
			return fmt.Errorf(
				"remote %s.phone_number mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedPhoneNumber,
				cardholder.PhoneNumber,
			)
		}
		if string(cardholder.Type) != expect.ExpectedType {
			return fmt.Errorf(
				"remote %s.type mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedType,
				string(cardholder.Type),
			)
		}
		if string(cardholder.Status) != expect.ExpectedStatus {
			return fmt.Errorf(
				"remote %s.status mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedStatus,
				string(cardholder.Status),
			)
		}
		if cardholder.Billing == nil || cardholder.Billing.Address == nil {
			return fmt.Errorf("remote %s.billing.address missing", expect.Address)
		}
		if cardholder.Billing.Address.Line1 != expect.ExpectedBillingLine1 {
			return fmt.Errorf(
				"remote %s.billing.address.line1 mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedBillingLine1,
				cardholder.Billing.Address.Line1,
			)
		}
		if cardholder.Billing.Address.City != expect.ExpectedBillingCity {
			return fmt.Errorf(
				"remote %s.billing.address.city mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedBillingCity,
				cardholder.Billing.Address.City,
			)
		}
		if cardholder.Billing.Address.PostalCode != expect.ExpectedBillingPostalCode {
			return fmt.Errorf(
				"remote %s.billing.address.postal_code mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedBillingPostalCode,
				cardholder.Billing.Address.PostalCode,
			)
		}
		if cardholder.Billing.Address.Country != expect.ExpectedBillingCountry {
			return fmt.Errorf(
				"remote %s.billing.address.country mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedBillingCountry,
				cardholder.Billing.Address.Country,
			)
		}
		if cardholder.Individual == nil {
			return fmt.Errorf("remote %s.individual missing", expect.Address)
		}
		if cardholder.Individual.FirstName != expect.ExpectedFirstName {
			return fmt.Errorf(
				"remote %s.individual.first_name mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedFirstName,
				cardholder.Individual.FirstName,
			)
		}
		if cardholder.Individual.LastName != expect.ExpectedLastName {
			return fmt.Errorf(
				"remote %s.individual.last_name mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedLastName,
				cardholder.Individual.LastName,
			)
		}
		if len(expect.ExpectedPreferredLocales) > 0 {
			if err := expectStringSlice(
				expect.Address+".preferred_locales",
				stringSliceFromIssuingCardholderLocales(cardholder.PreferredLocales),
				expect.ExpectedPreferredLocales,
			); err != nil {
				return err
			}
		}
		if len(expect.ExpectedAllowedCardPresences) > 0 {
			if cardholder.SpendingControls == nil {
				return fmt.Errorf("remote %s.spending_controls missing", expect.Address)
			}
			if err := expectStringSlice(
				expect.Address+".spending_controls.allowed_card_presences",
				stringSliceFromIssuingCardholderCardPresences(
					cardholder.SpendingControls.AllowedCardPresences,
				),
				expect.ExpectedAllowedCardPresences,
			); err != nil {
				return err
			}
			if string(cardholder.SpendingControls.SpendingLimitsCurrency) != expect.ExpectedSpendingLimitsCurrency {
				return fmt.Errorf(
					"remote %s.spending_controls.spending_limits_currency mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedSpendingLimitsCurrency,
					string(cardholder.SpendingControls.SpendingLimitsCurrency),
				)
			}
			if len(cardholder.SpendingControls.SpendingLimits) != 1 || cardholder.SpendingControls.SpendingLimits[0] == nil {
				return fmt.Errorf(
					"remote %s.spending_controls.spending_limits mismatch: expected exactly one limit",
					expect.Address,
				)
			}
			limit := cardholder.SpendingControls.SpendingLimits[0]
			if limit.Amount != expect.ExpectedSpendingLimitAmount {
				return fmt.Errorf(
					"remote %s.spending_controls.spending_limits[0].amount mismatch: expected %d, got %d",
					expect.Address,
					expect.ExpectedSpendingLimitAmount,
					limit.Amount,
				)
			}
			if string(limit.Interval) != expect.ExpectedSpendingLimitInterval {
				return fmt.Errorf(
					"remote %s.spending_controls.spending_limits[0].interval mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedSpendingLimitInterval,
					string(limit.Interval),
				)
			}
		}
		if err := expectMetadataSubset(
			expect.Address+".metadata",
			cardholder.Metadata,
			expect.ExpectedMetadata,
		); err != nil {
			return err
		}

		return nil
	}
}

func verifyIssuingCardholderDestroyStillExists(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	_, err := retrieveIssuingCardholder(client, state, "stripe_issuing_cardholder.test")
	return err
}

func verifyIssuingCard(expect issuingCardExpectations) runner.StateVerifier {
	return func(
		_ runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		card, err := retrieveIssuingCard(client, state, expect.Address)
		if err != nil {
			return err
		}
		expectedCardholderID, err := runner.ResourcePrimaryID(state, expect.ExpectedCardholderAddress)
		if err != nil {
			return err
		}
		if card.Cardholder == nil || card.Cardholder.ID != expectedCardholderID {
			actualCardholderID := ""
			if card.Cardholder != nil {
				actualCardholderID = card.Cardholder.ID
			}
			return fmt.Errorf(
				"remote %s.cardholder mismatch: expected %q, got %q",
				expect.Address,
				expectedCardholderID,
				actualCardholderID,
			)
		}
		if string(card.Currency) != expect.ExpectedCurrency {
			return fmt.Errorf(
				"remote %s.currency mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedCurrency,
				string(card.Currency),
			)
		}
		if string(card.Type) != expect.ExpectedType {
			return fmt.Errorf(
				"remote %s.type mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedType,
				string(card.Type),
			)
		}
		if string(card.Status) != expect.ExpectedStatus {
			return fmt.Errorf(
				"remote %s.status mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedStatus,
				string(card.Status),
			)
		}
		if card.Brand == "" {
			return fmt.Errorf("remote %s.brand missing", expect.Address)
		}
		if card.Last4 == "" {
			return fmt.Errorf("remote %s.last4 missing", expect.Address)
		}
		if err := expectMetadataSubset(
			expect.Address+".metadata",
			card.Metadata,
			expect.ExpectedMetadata,
		); err != nil {
			return err
		}
		if expect.ExpectedLifecyclePaymentCount != 0 {
			if card.LifecycleControls == nil || card.LifecycleControls.CancelAfter == nil {
				return fmt.Errorf("remote %s.lifecycle_controls.cancel_after missing", expect.Address)
			}
			if card.LifecycleControls.CancelAfter.PaymentCount != expect.ExpectedLifecyclePaymentCount {
				return fmt.Errorf(
					"remote %s.lifecycle_controls.cancel_after.payment_count mismatch: expected %d, got %d",
					expect.Address,
					expect.ExpectedLifecyclePaymentCount,
					card.LifecycleControls.CancelAfter.PaymentCount,
				)
			}
		}

		return nil
	}
}

func verifyIssuingCardDestroyStillExists(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	_, err := retrieveIssuingCard(client, state, "stripe_issuing_card.test")
	return err
}

func verifyIssuingDispute(expect issuingDisputeExpectations) runner.StateVerifier {
	return func(
		env runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		dispute, err := retrieveIssuingDispute(client, state, expect.Address)
		if err != nil {
			return err
		}
		expectedTransactionID := env.IssuingDisputeTransaction
		if expect.ExpectedReason == "other" && env.IssuingDisputeOtherTransaction != "" {
			expectedTransactionID = env.IssuingDisputeOtherTransaction
		}
		if expectedTransactionID != "" {
			if dispute.Transaction == nil || dispute.Transaction.ID != expectedTransactionID {
				actualTransactionID := ""
				if dispute.Transaction != nil {
					actualTransactionID = dispute.Transaction.ID
				}
				return fmt.Errorf(
					"remote %s.transaction mismatch: expected %q, got %q",
					expect.Address,
					expectedTransactionID,
					actualTransactionID,
				)
			}
		}
		if string(dispute.Status) != expect.ExpectedStatus {
			return fmt.Errorf(
				"remote %s.status mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedStatus,
				string(dispute.Status),
			)
		}
		if dispute.Evidence == nil {
			return fmt.Errorf("remote %s.evidence missing", expect.Address)
		}
		if string(dispute.Evidence.Reason) != expect.ExpectedReason {
			return fmt.Errorf(
				"remote %s.evidence.reason mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedReason,
				string(dispute.Evidence.Reason),
			)
		}
		if expect.ExpectedReason == "fraudulent" {
			if dispute.Evidence.Fraudulent == nil {
				return fmt.Errorf("remote %s.evidence.fraudulent missing", expect.Address)
			}
			if dispute.Evidence.Fraudulent.Explanation != expect.ExpectedFraudulentExplanation {
				return fmt.Errorf(
					"remote %s.evidence.fraudulent.explanation mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedFraudulentExplanation,
					dispute.Evidence.Fraudulent.Explanation,
				)
			}
		}
		if expect.ExpectedReason == "other" {
			if dispute.Evidence.Other == nil {
				return fmt.Errorf("remote %s.evidence.other missing", expect.Address)
			}
			if dispute.Evidence.Other.Explanation != expect.ExpectedOtherExplanation {
				return fmt.Errorf(
					"remote %s.evidence.other.explanation mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedOtherExplanation,
					dispute.Evidence.Other.Explanation,
				)
			}
			if dispute.Evidence.Other.ProductDescription != expect.ExpectedProductDescription {
				return fmt.Errorf(
					"remote %s.evidence.other.product_description mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedProductDescription,
					dispute.Evidence.Other.ProductDescription,
				)
			}
			if string(dispute.Evidence.Other.ProductType) != expect.ExpectedProductType {
				return fmt.Errorf(
					"remote %s.evidence.other.product_type mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedProductType,
					string(dispute.Evidence.Other.ProductType),
				)
			}
		}
		if expect.ExpectedReason == "canceled" {
			if dispute.Evidence.Canceled == nil {
				return fmt.Errorf("remote %s.evidence.canceled missing", expect.Address)
			}
			if dispute.Evidence.Canceled.Explanation != expect.ExpectedCanceledExplanation {
				return fmt.Errorf(
					"remote %s.evidence.canceled.explanation mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedCanceledExplanation,
					dispute.Evidence.Canceled.Explanation,
				)
			}
			if dispute.Evidence.Canceled.CanceledAt != expect.ExpectedCanceledAt {
				return fmt.Errorf(
					"remote %s.evidence.canceled.canceled_at mismatch: expected %d, got %d",
					expect.Address,
					expect.ExpectedCanceledAt,
					dispute.Evidence.Canceled.CanceledAt,
				)
			}
			if dispute.Evidence.Canceled.CancellationReason != expect.ExpectedCancellationReason {
				return fmt.Errorf(
					"remote %s.evidence.canceled.cancellation_reason mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedCancellationReason,
					dispute.Evidence.Canceled.CancellationReason,
				)
			}
			if dispute.Evidence.Canceled.CancellationPolicyProvided != expect.ExpectedCancellationPolicyProvided {
				return fmt.Errorf(
					"remote %s.evidence.canceled.cancellation_policy_provided mismatch: expected %t, got %t",
					expect.Address,
					expect.ExpectedCancellationPolicyProvided,
					dispute.Evidence.Canceled.CancellationPolicyProvided,
				)
			}
			if dispute.Evidence.Canceled.ExpectedAt != expect.ExpectedExpectedAt {
				return fmt.Errorf(
					"remote %s.evidence.canceled.expected_at mismatch: expected %d, got %d",
					expect.Address,
					expect.ExpectedExpectedAt,
					dispute.Evidence.Canceled.ExpectedAt,
				)
			}
			if dispute.Evidence.Canceled.ProductDescription != expect.ExpectedProductDescription {
				return fmt.Errorf(
					"remote %s.evidence.canceled.product_description mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedProductDescription,
					dispute.Evidence.Canceled.ProductDescription,
				)
			}
			if string(dispute.Evidence.Canceled.ProductType) != expect.ExpectedProductType {
				return fmt.Errorf(
					"remote %s.evidence.canceled.product_type mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedProductType,
					string(dispute.Evidence.Canceled.ProductType),
				)
			}
			if string(dispute.Evidence.Canceled.ReturnStatus) != expect.ExpectedReturnStatus {
				return fmt.Errorf(
					"remote %s.evidence.canceled.return_status mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedReturnStatus,
					string(dispute.Evidence.Canceled.ReturnStatus),
				)
			}
			if dispute.Evidence.Canceled.ReturnedAt != expect.ExpectedReturnedAt {
				return fmt.Errorf(
					"remote %s.evidence.canceled.returned_at mismatch: expected %d, got %d",
					expect.Address,
					expect.ExpectedReturnedAt,
					dispute.Evidence.Canceled.ReturnedAt,
				)
			}
		}
		if err := expectMetadataSubset(
			expect.Address+".metadata",
			dispute.Metadata,
			expect.ExpectedMetadata,
		); err != nil {
			return err
		}

		return nil
	}
}

func stringSliceFromIssuingCardholderLocales(
	locales []stripe.IssuingCardholderPreferredLocale,
) []string {
	values := make([]string, 0, len(locales))
	for _, locale := range locales {
		values = append(values, string(locale))
	}
	return values
}

func stringSliceFromIssuingCardholderCardPresences(
	presences []stripe.IssuingCardholderSpendingControlsAllowedCardPresence,
) []string {
	values := make([]string, 0, len(presences))
	for _, presence := range presences {
		values = append(values, string(presence))
	}
	return values
}

func expectStringSlice(name string, actual []string, expected []string) error {
	if len(actual) != len(expected) {
		return fmt.Errorf(
			"%s length mismatch: expected %d, got %d",
			name,
			len(expected),
			len(actual),
		)
	}
	for i := range expected {
		if actual[i] != expected[i] {
			return fmt.Errorf(
				"%s[%d] mismatch: expected %q, got %q",
				name,
				i,
				expected[i],
				actual[i],
			)
		}
	}
	return nil
}

func expectInt64Slice(name string, actual []int64, expected []int64) error {
	if len(actual) != len(expected) {
		return fmt.Errorf(
			"%s length mismatch: expected %d, got %d",
			name,
			len(expected),
			len(actual),
		)
	}
	for i := range expected {
		if actual[i] != expected[i] {
			return fmt.Errorf(
				"%s[%d] mismatch: expected %d, got %d",
				name,
				i,
				expected[i],
				actual[i],
			)
		}
	}
	return nil
}

func verifyIssuingDisputeDestroyStillExists(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	_, err := retrieveIssuingDispute(client, state, "stripe_issuing_dispute.test")
	return err
}

func verifyIssuingPersonalizationDesign(
	expect issuingPersonalizationDesignExpectations,
) runner.StateVerifier {
	return func(
		env runner.TestEnv,
		client *stripe.Client,
		state *terraform.State,
	) error {
		design, err := retrieveIssuingPersonalizationDesign(client, state, expect.Address)
		if err != nil {
			return err
		}
		if env.IssuingPhysicalBundle != "" {
			if design.PhysicalBundle == nil || design.PhysicalBundle.ID != env.IssuingPhysicalBundle {
				actualBundleID := ""
				if design.PhysicalBundle != nil {
					actualBundleID = design.PhysicalBundle.ID
				}
				return fmt.Errorf(
					"remote %s.physical_bundle mismatch: expected %q, got %q",
					expect.Address,
					env.IssuingPhysicalBundle,
					actualBundleID,
				)
			}
		}
		if design.Name != expect.ExpectedName {
			return fmt.Errorf(
				"remote %s.name mismatch: expected %q, got %q",
				expect.Address,
				expect.ExpectedName,
				design.Name,
			)
		}
		if expect.CompareStateLookupKey {
			expectedLookupKey, err := runner.ResourceAttribute(state, expect.Address, "lookup_key")
			if err != nil {
				return err
			}
			if design.LookupKey != expectedLookupKey {
				return fmt.Errorf(
					"remote %s.lookup_key mismatch: expected %q, got %q",
					expect.Address,
					expectedLookupKey,
					design.LookupKey,
				)
			}
		}
		if expect.CheckCarrierText {
			if design.CarrierText == nil {
				return fmt.Errorf("remote %s.carrier_text missing", expect.Address)
			}
			if design.CarrierText.HeaderTitle != expect.ExpectedHeaderTitle {
				return fmt.Errorf(
					"remote %s.carrier_text.header_title mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedHeaderTitle,
					design.CarrierText.HeaderTitle,
				)
			}
			if design.CarrierText.HeaderBody != expect.ExpectedHeaderBody {
				return fmt.Errorf(
					"remote %s.carrier_text.header_body mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedHeaderBody,
					design.CarrierText.HeaderBody,
				)
			}
			if design.CarrierText.FooterTitle != expect.ExpectedFooterTitle {
				return fmt.Errorf(
					"remote %s.carrier_text.footer_title mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedFooterTitle,
					design.CarrierText.FooterTitle,
				)
			}
			if design.CarrierText.FooterBody != expect.ExpectedFooterBody {
				return fmt.Errorf(
					"remote %s.carrier_text.footer_body mismatch: expected %q, got %q",
					expect.Address,
					expect.ExpectedFooterBody,
					design.CarrierText.FooterBody,
				)
			}
		}
		if expect.CheckIsDefault {
			if design.Preferences == nil {
				return fmt.Errorf("remote %s.preferences missing", expect.Address)
			}
			if design.Preferences.IsDefault != expect.ExpectedIsDefault {
				return fmt.Errorf(
					"remote %s.preferences.is_default mismatch: expected %t, got %t",
					expect.Address,
					expect.ExpectedIsDefault,
					design.Preferences.IsDefault,
				)
			}
		}
		if err := expectMetadataSubset(
			expect.Address+".metadata",
			design.Metadata,
			expect.ExpectedMetadata,
		); err != nil {
			return err
		}

		return nil
	}
}

func verifyIssuingPersonalizationDesignDestroyStillExists(
	_ runner.TestEnv,
	client *stripe.Client,
	state *terraform.State,
) error {
	_, err := retrieveIssuingPersonalizationDesign(
		client,
		state,
		"stripe_issuing_personalization_design.test",
	)
	return err
}

func retrievePerson(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.Person, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}
	accountID, err := runner.ResourceAttribute(state, address, "account")
	if err != nil {
		return nil, err
	}

	params := &stripe.PersonRetrieveParams{
		Account: stripe.String(accountID),
	}
	person, err := client.V1Persons.Retrieve(context.Background(), id, params)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return person, nil
}

func retrieveTerminalLocation(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.TerminalLocation, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	location, err := client.V1TerminalLocations.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return location, nil
}

func retrieveTerminalConfiguration(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.TerminalConfiguration, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	configuration, err := client.V1TerminalConfigurations.Retrieve(
		context.Background(),
		id,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return configuration, nil
}

func retrieveIssuingCardholder(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.IssuingCardholder, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	cardholder, err := client.V1IssuingCardholders.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return cardholder, nil
}

func retrieveIssuingCard(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.IssuingCard, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	card, err := client.V1IssuingCards.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return card, nil
}

func retrieveIssuingDispute(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.IssuingDispute, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	dispute, err := client.V1IssuingDisputes.Retrieve(context.Background(), id, nil)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return dispute, nil
}

func retrieveIssuingPersonalizationDesign(
	client *stripe.Client,
	state *terraform.State,
	address string,
) (*stripe.IssuingPersonalizationDesign, error) {
	id, err := runner.ResourcePrimaryID(state, address)
	if err != nil {
		return nil, err
	}

	design, err := client.V1IssuingPersonalizationDesigns.Retrieve(
		context.Background(),
		id,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("retrieve %s (%s): %w", address, id, err)
	}

	return design, nil
}
