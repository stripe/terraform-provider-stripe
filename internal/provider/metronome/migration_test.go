//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func TestMigrationRejectsUnreviewedStateVersions(t *testing.T) {
	factories := []func() resource.Resource{
		NewMetronomeBillableMetricResource, NewMetronomeProductResource,
		NewMetronomeRateCardResource, NewMetronomePackageResource,
		NewMetronomeCustomerResource, NewMetronomeContractResource,
		NewMetronomeCustomerCommitResource, NewMetronomeCustomerCreditResource,
		NewMetronomeCustomerAlertResource, NewMetronomeCustomFieldKeyResource,
		NewMetronomeCreditTypeResource, NewMetronomeCustomerBillingConfigurationResource,
		NewMetronomeCustomerRevenueConfigurationResource, NewMetronomeOffsetNotificationResource,
	}
	for _, factory := range factories {
		r := factory()
		var metadata resource.MetadataResponse
		r.Metadata(context.Background(), resource.MetadataRequest{ProviderTypeName: "stripe"}, &metadata)
		t.Run(metadata.TypeName, func(t *testing.T) {
			mover := r.(resource.ResourceWithMoveState).MoveState(context.Background())[0]
			sourceType := metadata.TypeName[len("stripe_"):]
			for _, version := range []int64{-1, 1, 99} {
				var response resource.MoveStateResponse
				mover.StateMover(context.Background(), resource.MoveStateRequest{SourceProviderAddress: "registry.terraform.io/tonyxiao/metronome", SourceTypeName: sourceType, SourceSchemaVersion: version}, &response)
				if !response.Diagnostics.HasError() {
					t.Fatalf("unreviewed schema version %d accepted", version)
				}
			}
			var missing resource.MoveStateResponse
			mover.StateMover(context.Background(), resource.MoveStateRequest{SourceProviderAddress: "registry.terraform.io/tonyxiao/metronome", SourceTypeName: sourceType, SourceSchemaVersion: 0}, &missing)
			if !missing.Diagnostics.HasError() {
				t.Fatal("missing source state accepted")
			}
			for _, request := range []resource.MoveStateRequest{
				{SourceProviderAddress: "registry.terraform.io/unrelated/metronome", SourceTypeName: sourceType},
				{SourceProviderAddress: "registry.terraform.io/tonyxiao/metronome", SourceTypeName: "metronome_unreviewed"},
			} {
				var response resource.MoveStateResponse
				mover.StateMover(context.Background(), request, &response)
				if response.Diagnostics.HasError() {
					t.Fatal("mover claimed an unrelated provider or type")
				}
			}
		})
	}
}
