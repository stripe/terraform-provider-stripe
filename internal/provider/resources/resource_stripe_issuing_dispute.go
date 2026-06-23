//
// File generated from our OpenAPI spec
//

package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	stripe "github.com/stripe/stripe-go/v86"
)

var _ resource.Resource = &IssuingDisputeResource{}

var _ resource.ResourceWithConfigure = &IssuingDisputeResource{}

var _ resource.ResourceWithImportState = &IssuingDisputeResource{}

func NewIssuingDisputeResource() resource.Resource {
	return &IssuingDisputeResource{}
}

type IssuingDisputeResource struct {
	client *stripe.Client
}

type IssuingDisputeResourceModel struct {
	Object      types.String `tfsdk:"object"`
	Amount      types.Int64  `tfsdk:"amount"`
	Created     types.Int64  `tfsdk:"created"`
	Currency    types.String `tfsdk:"currency"`
	Evidence    types.Object `tfsdk:"evidence"`
	ID          types.String `tfsdk:"id"`
	Livemode    types.Bool   `tfsdk:"livemode"`
	LossReason  types.String `tfsdk:"loss_reason"`
	Metadata    types.Map    `tfsdk:"metadata"`
	Status      types.String `tfsdk:"status"`
	Transaction types.String `tfsdk:"transaction"`
	Treasury    types.Object `tfsdk:"treasury"`
}

func (r *IssuingDisputeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*stripe.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected *stripe.Client, got: %T", req.ProviderData))
		return
	}

	r.client = client
}

func (r *IssuingDisputeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_issuing_dispute"
}

func (r *IssuingDisputeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "As a [card issuer](https://docs.stripe.com/issuing), you can dispute transactions that the cardholder does not recognize, suspects to be fraudulent, or has other issues with.\n\nRelated guide: [Issuing disputes](https://docs.stripe.com/issuing/purchases/disputes)",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed:      true,
				Description:   "String representing the object's type. Objects of the same type share the same value.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("issuing.dispute")},
			},
			"amount": schema.Int64Attribute{
				Optional:      true,
				Computed:      true,
				Description:   "Disputed amount in the card's currency and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal). Usually the amount of the `transaction`, but can differ (usually because of currency fluctuation).",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"created": schema.Int64Attribute{
				Computed:      true,
				Description:   "Time at which the object was created. Measured in seconds since the Unix epoch.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
			},
			"currency": schema.StringAttribute{
				Computed:      true,
				Description:   "The currency the `transaction` was made in.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"evidence": schema.SingleNestedAttribute{
				Optional: true,
				Computed: true,

				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
				Attributes: map[string]schema.Attribute{
					"canceled": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"additional_documentation": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "(ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"canceled_at": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Date when order was canceled.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"cancellation_policy_provided": schema.BoolAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Whether the cardholder was provided with a cancellation policy.",
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"cancellation_reason": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Reason for canceling the order.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"expected_at": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Date when the cardholder expected to receive the product.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"explanation": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Explanation of why the cardholder is disputing this transaction.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"product_description": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Description of the merchandise or service that was purchased.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"product_type": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Whether the product was a merchandise or service.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("merchandise", "service")},
							},
							"return_status": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Result of cardholder's attempt to return the product.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("merchant_rejected", "successful")},
							},
							"returned_at": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Date when the product was returned or attempted to be returned.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"duplicate": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"additional_documentation": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "(ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"card_statement": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "(ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Copy of the card statement showing that the product had already been paid for.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"cash_receipt": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "(ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Copy of the receipt showing that the product had been paid for in cash.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"check_image": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "(ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Image of the front and back of the check that was used to pay for the product.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"explanation": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Explanation of why the cardholder is disputing this transaction.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"original_transaction": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Transaction (e.g., ipi_...) that the disputed transaction is a duplicate of. Of the two or more transactions that are copies of each other, this is original undisputed one.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"fraudulent": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"additional_documentation": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "(ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"explanation": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Explanation of why the cardholder is disputing this transaction.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"merchandise_not_as_described": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"additional_documentation": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "(ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"explanation": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Explanation of why the cardholder is disputing this transaction.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"received_at": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Date when the product was received.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"return_description": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Description of the cardholder's attempt to return the product.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"return_status": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Result of cardholder's attempt to return the product.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("merchant_rejected", "successful")},
							},
							"returned_at": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Date when the product was returned or attempted to be returned.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
					"no_valid_authorization": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"additional_documentation": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "(ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"explanation": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Explanation of why the cardholder is disputing this transaction.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
						},
					},
					"not_received": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"additional_documentation": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "(ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"expected_at": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Date when the cardholder expected to receive the product.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"explanation": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Explanation of why the cardholder is disputing this transaction.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"product_description": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Description of the merchandise or service that was purchased.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"product_type": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Whether the product was a merchandise or service.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("merchandise", "service")},
							},
						},
					},
					"other": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"additional_documentation": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "(ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"explanation": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Explanation of why the cardholder is disputing this transaction.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"product_description": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Description of the merchandise or service that was purchased.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"product_type": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Whether the product was a merchandise or service.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
								Validators:    []validator.String{stringvalidator.OneOf("merchandise", "service")},
							},
						},
					},
					"reason": schema.StringAttribute{
						Optional:      true,
						Computed:      true,
						Description:   "The reason for filing the dispute. Its value will match the field containing the evidence.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
						Validators:    []validator.String{stringvalidator.OneOf("canceled", "duplicate", "fraudulent", "merchandise_not_as_described", "no_valid_authorization", "not_received", "other", "service_not_as_described")},
					},
					"service_not_as_described": schema.SingleNestedAttribute{
						Optional: true,
						Computed: true,

						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown()},
						Attributes: map[string]schema.Attribute{
							"additional_documentation": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "(ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"canceled_at": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Date when order was canceled.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"cancellation_reason": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Reason for canceling the order.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"explanation": schema.StringAttribute{
								Optional:      true,
								Computed:      true,
								Description:   "Explanation of why the cardholder is disputing this transaction.",
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"received_at": schema.Int64Attribute{
								Optional:      true,
								Computed:      true,
								Description:   "Date when the product was received.",
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
						},
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Unique identifier for the object.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"livemode": schema.BoolAttribute{
				Computed:      true,
				Description:   "If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"loss_reason": schema.StringAttribute{
				Computed:      true,
				Description:   "The enum that describes the dispute loss outcome. If the dispute is not lost, this field will be absent. New enum values may be added in the future, so be sure to handle unknown values.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("cardholder_authentication_issuer_liability", "eci5_token_transaction_with_tavv", "excess_disputes_in_timeframe", "has_not_met_the_minimum_dispute_amount_requirements", "invalid_duplicate_dispute", "invalid_incorrect_amount_dispute", "invalid_no_authorization", "invalid_use_of_disputes", "merchandise_delivered_or_shipped", "merchandise_or_service_as_described", "not_cancelled", "other", "refund_issued", "submitted_beyond_allowable_time_limit", "transaction_3ds_required", "transaction_approved_after_prior_fraud_dispute", "transaction_authorized", "transaction_electronically_read", "transaction_qualifies_for_visa_easy_payment_service", "transaction_unattended")},
			},
			"metadata": schema.MapAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.",
				PlanModifiers: []planmodifier.Map{mapplanmodifier.UseStateForUnknown()},
				ElementType:   types.StringType,
			},
			"status": schema.StringAttribute{
				Computed:      true,
				Description:   "Current status of the dispute.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
				Validators:    []validator.String{stringvalidator.OneOf("expired", "lost", "submitted", "unsubmitted", "won")},
			},
			"transaction": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "The transaction being disputed.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"treasury": schema.SingleNestedAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "[Treasury](https://docs.stripe.com/api/treasury) details related to this dispute if it was created on a [FinancialAccount](https://docs.stripe.com/api/treasury/financial_accounts)",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseStateForUnknown(), objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"debit_reversal": schema.StringAttribute{
						Computed:      true,
						Description:   "The Treasury [DebitReversal](https://docs.stripe.com/api/treasury/debit_reversals) representing this Issuing dispute",
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
					},
					"received_debit": schema.StringAttribute{
						Required:      true,
						Description:   "The Treasury [ReceivedDebit](https://docs.stripe.com/api/treasury/received_debits) that is being disputed.",
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
				},
			},
		},
	}
}

func (r *IssuingDisputeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IssuingDisputeResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	params, err := expandIssuingDisputeCreate(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error building IssuingDispute create params", err.Error())
		return
	}

	obj, err := r.client.V1IssuingDisputes.Create(ctx, params)
	if err != nil {
		resp.Diagnostics.AddError("Error creating IssuingDispute", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1IssuingDisputes.B, r.client.V1IssuingDisputes.Key, stripe.FormatURLPath("/v1/issuing/disputes/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating IssuingDispute create raw response", err.Error())
		return
	}

	if err := flattenIssuingDispute(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening IssuingDispute create response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *IssuingDisputeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorState IssuingDisputeResourceModel
	diags := req.State.Get(ctx, &priorState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state IssuingDisputeResourceModel
	state = priorState

	obj, err := r.client.V1IssuingDisputes.Retrieve(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading IssuingDispute", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1IssuingDisputes.B, r.client.V1IssuingDisputes.Key, stripe.FormatURLPath("/v1/issuing/disputes/%s", state.ID.ValueString()), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating IssuingDispute raw response", err.Error())
		return
	}

	if err := flattenIssuingDispute(obj, &state); err != nil {
		resp.Diagnostics.AddError("Error flattening IssuingDispute read response", err.Error())
		return
	}
	normalizeUnknownValues(&state)
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *IssuingDisputeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IssuingDisputeResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state IssuingDisputeResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diffPlan := plan
	diffState := state

	params, err := expandIssuingDisputeUpdate(diffPlan, diffState)
	if err != nil {
		resp.Diagnostics.AddError("Error building IssuingDispute update params", err.Error())
		return
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() &&
		!state.Metadata.IsNull() && !state.Metadata.IsUnknown() {
		if !assignMetadataDiffToNamedField(params, "Metadata", plan.Metadata, state.Metadata) {
			resp.Diagnostics.AddError("Error building IssuingDispute update params", fmt.Sprintf("failed to assign metadata diff %q on %T", "metadata", params))
			return
		}
	}
	obj, err := r.client.V1IssuingDisputes.Update(ctx, state.ID.ValueString(), params)
	if err != nil {
		resp.Diagnostics.AddError("Error updating IssuingDispute", err.Error())
		return
	}

	if err := ensureRawResponse(obj, r.client.V1IssuingDisputes.B, r.client.V1IssuingDisputes.Key, stripe.FormatURLPath("/v1/issuing/disputes/%s", obj.ID), nil); err != nil {
		resp.Diagnostics.AddError("Error hydrating IssuingDispute update raw response", err.Error())
		return
	}

	if err := flattenIssuingDispute(obj, &plan); err != nil {
		resp.Diagnostics.AddError("Error flattening IssuingDispute update response", err.Error())
		return
	}
	normalizeUnknownValues(&plan)
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *IssuingDisputeResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// This resource does not support deletion from Stripe.
	// Removing it from Terraform state only.
}

func (r *IssuingDisputeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func expandIssuingDisputeCreate(plan IssuingDisputeResourceModel) (*stripe.IssuingDisputeCreateParams, error) {
	params := &stripe.IssuingDisputeCreateParams{}

	if !plan.Amount.IsNull() && !plan.Amount.IsUnknown() {
		params.Amount = stripe.Int64(plan.Amount.ValueInt64())
	}
	if !plan.Evidence.IsNull() && !plan.Evidence.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Evidence", plan.Evidence) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "evidence", params)
		}
	}
	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "metadata", params)
		}
	}
	if !plan.Transaction.IsNull() && !plan.Transaction.IsUnknown() {
		if !assignStringToNamedFieldOrMethod(params, "TransactionID", "Transaction", plan.Transaction.ValueString()) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "transaction", params)
		}
	}
	if !plan.Treasury.IsNull() && !plan.Treasury.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Treasury", plan.Treasury) {
			return nil, fmt.Errorf("failed to assign attribute %q on %T", "treasury", params)
		}
	}

	return params, nil
}

func expandIssuingDisputeUpdate(plan IssuingDisputeResourceModel, state IssuingDisputeResourceModel) (*stripe.IssuingDisputeUpdateParams, error) {
	params := &stripe.IssuingDisputeUpdateParams{}

	if !plan.Amount.Equal(state.Amount) && !plan.Amount.IsNull() && !plan.Amount.IsUnknown() {
		params.Amount = stripe.Int64(plan.Amount.ValueInt64())
	}
	if !plan.Evidence.Equal(state.Evidence) && !plan.Evidence.IsNull() && !plan.Evidence.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Evidence", plan.Evidence) {
			if !plan.Evidence.Equal(state.Evidence) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "evidence", params)
			}
		}
	}
	if !plan.Metadata.Equal(state.Metadata) && !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		if !assignAttrValueToNamedField(params, "Metadata", plan.Metadata) {
			if !plan.Metadata.Equal(state.Metadata) {
				return nil, fmt.Errorf("failed to assign changed attribute %q on %T", "metadata", params)
			}
		}
	}

	return params, nil
}

func flattenIssuingDispute(obj *stripe.IssuingDispute, state *IssuingDisputeResourceModel) error {
	raw, hasRaw := rawResponseToPlain(obj)
	{
		if rawValueObject, rawOk := plainValueAtPath(raw, "object"); rawOk {
			if valueObject, err := flattenPlainValue(rawValueObject, types.StringType, "object", "raw response"); err != nil {
				return err
			} else {
				if typedObject, ok := valueObject.(types.String); ok {
					state.Object = typedObject
				}
			}
		} else if !hasRaw {
			if responseValueObject, ok := plainFromResponseField(obj, "Object"); ok {
				if valueObject, err := flattenPlainValue(responseValueObject, types.StringType, "object", "response struct"); err != nil {
					return err
				} else {
					if typedObject, ok := valueObject.(types.String); ok {
						state.Object = typedObject
					}
				}
			}
		}
	}
	{
		if rawValueAmount, rawOk := plainValueAtPath(raw, "amount"); rawOk {
			if valueAmount, err := flattenPlainValue(rawValueAmount, types.Int64Type, "amount", "raw response"); err != nil {
				return err
			} else {
				if typedAmount, ok := valueAmount.(types.Int64); ok {
					state.Amount = typedAmount
				}
			}
		} else if !hasRaw {
			if responseValueAmount, ok := plainFromResponseField(obj, "Amount"); ok {
				if valueAmount, err := flattenPlainValue(responseValueAmount, types.Int64Type, "amount", "response struct"); err != nil {
					return err
				} else {
					if typedAmount, ok := valueAmount.(types.Int64); ok {
						state.Amount = typedAmount
					}
				}
			}
		}
	}
	{
		if rawValueCreated, rawOk := plainValueAtPath(raw, "created"); rawOk {
			if valueCreated, err := flattenPlainValue(rawValueCreated, types.Int64Type, "created", "raw response"); err != nil {
				return err
			} else {
				if typedCreated, ok := valueCreated.(types.Int64); ok {
					state.Created = typedCreated
				}
			}
		} else if !hasRaw {
			if responseValueCreated, ok := plainFromResponseField(obj, "Created"); ok {
				if valueCreated, err := flattenPlainValue(responseValueCreated, types.Int64Type, "created", "response struct"); err != nil {
					return err
				} else {
					if typedCreated, ok := valueCreated.(types.Int64); ok {
						state.Created = typedCreated
					}
				}
			}
		}
	}
	{
		if rawValueCurrency, rawOk := plainValueAtPath(raw, "currency"); rawOk {
			if valueCurrency, err := flattenPlainValue(rawValueCurrency, types.StringType, "currency", "raw response"); err != nil {
				return err
			} else {
				if typedCurrency, ok := valueCurrency.(types.String); ok {
					state.Currency = typedCurrency
				}
			}
		} else if !hasRaw {
			if responseValueCurrency, ok := plainFromResponseField(obj, "Currency"); ok {
				if valueCurrency, err := flattenPlainValue(responseValueCurrency, types.StringType, "currency", "response struct"); err != nil {
					return err
				} else {
					if typedCurrency, ok := valueCurrency.(types.String); ok {
						state.Currency = typedCurrency
					}
				}
			}
		}
	}
	{
		assignedEvidence := false
		hadRawEvidence := false
		if rawValueEvidence, rawOk := plainValueAtPath(raw, "evidence"); rawOk {
			hadRawEvidence = true
			if rawValueEvidence != nil {
				sourceEvidence := applyConfiguredKeyedListShapes(rawValueEvidence, attrValueToPlain(state.Evidence))
				if valueEvidence, err := flattenPlainValue(sourceEvidence, types.ObjectType{AttrTypes: map[string]attr.Type{"canceled": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "canceled_at": types.Int64Type, "cancellation_policy_provided": types.BoolType, "cancellation_reason": types.StringType, "expected_at": types.Int64Type, "explanation": types.StringType, "product_description": types.StringType, "product_type": types.StringType, "return_status": types.StringType, "returned_at": types.Int64Type}}, "duplicate": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "card_statement": types.StringType, "cash_receipt": types.StringType, "check_image": types.StringType, "explanation": types.StringType, "original_transaction": types.StringType}}, "fraudulent": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "explanation": types.StringType}}, "merchandise_not_as_described": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "explanation": types.StringType, "received_at": types.Int64Type, "return_description": types.StringType, "return_status": types.StringType, "returned_at": types.Int64Type}}, "no_valid_authorization": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "explanation": types.StringType}}, "not_received": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "expected_at": types.Int64Type, "explanation": types.StringType, "product_description": types.StringType, "product_type": types.StringType}}, "other": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "explanation": types.StringType, "product_description": types.StringType, "product_type": types.StringType}}, "reason": types.StringType, "service_not_as_described": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "canceled_at": types.Int64Type, "cancellation_reason": types.StringType, "explanation": types.StringType, "received_at": types.Int64Type}}}}, "evidence", "raw response"); err != nil {
					return err
				} else {
					if typedEvidence, ok := valueEvidence.(types.Object); ok {
						state.Evidence = typedEvidence
						assignedEvidence = true
					}
				}
			}
		}
		if !assignedEvidence {
			if !hasRaw {
				if responseValueEvidence, ok := plainFromResponseField(obj, "Evidence"); ok {
					sourceEvidence := applyConfiguredKeyedListShapes(responseValueEvidence, attrValueToPlain(state.Evidence))
					if valueEvidence, err := flattenPlainValue(
						sourceEvidence,
						types.ObjectType{AttrTypes: map[string]attr.Type{"canceled": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "canceled_at": types.Int64Type, "cancellation_policy_provided": types.BoolType, "cancellation_reason": types.StringType, "expected_at": types.Int64Type, "explanation": types.StringType, "product_description": types.StringType, "product_type": types.StringType, "return_status": types.StringType, "returned_at": types.Int64Type}}, "duplicate": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "card_statement": types.StringType, "cash_receipt": types.StringType, "check_image": types.StringType, "explanation": types.StringType, "original_transaction": types.StringType}}, "fraudulent": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "explanation": types.StringType}}, "merchandise_not_as_described": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "explanation": types.StringType, "received_at": types.Int64Type, "return_description": types.StringType, "return_status": types.StringType, "returned_at": types.Int64Type}}, "no_valid_authorization": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "explanation": types.StringType}}, "not_received": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "expected_at": types.Int64Type, "explanation": types.StringType, "product_description": types.StringType, "product_type": types.StringType}}, "other": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "explanation": types.StringType, "product_description": types.StringType, "product_type": types.StringType}}, "reason": types.StringType, "service_not_as_described": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "canceled_at": types.Int64Type, "cancellation_reason": types.StringType, "explanation": types.StringType, "received_at": types.Int64Type}}}},
						"evidence",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedEvidence, ok := valueEvidence.(types.Object); ok {
							state.Evidence = typedEvidence
							assignedEvidence = true
						}
					}
				}
			}
		}
		if !assignedEvidence && hadRawEvidence {
			if nullEvidence, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"canceled": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "canceled_at": types.Int64Type, "cancellation_policy_provided": types.BoolType, "cancellation_reason": types.StringType, "expected_at": types.Int64Type, "explanation": types.StringType, "product_description": types.StringType, "product_type": types.StringType, "return_status": types.StringType, "returned_at": types.Int64Type}}, "duplicate": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "card_statement": types.StringType, "cash_receipt": types.StringType, "check_image": types.StringType, "explanation": types.StringType, "original_transaction": types.StringType}}, "fraudulent": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "explanation": types.StringType}}, "merchandise_not_as_described": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "explanation": types.StringType, "received_at": types.Int64Type, "return_description": types.StringType, "return_status": types.StringType, "returned_at": types.Int64Type}}, "no_valid_authorization": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "explanation": types.StringType}}, "not_received": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "expected_at": types.Int64Type, "explanation": types.StringType, "product_description": types.StringType, "product_type": types.StringType}}, "other": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "explanation": types.StringType, "product_description": types.StringType, "product_type": types.StringType}}, "reason": types.StringType, "service_not_as_described": types.ObjectType{AttrTypes: map[string]attr.Type{"additional_documentation": types.StringType, "canceled_at": types.Int64Type, "cancellation_reason": types.StringType, "explanation": types.StringType, "received_at": types.Int64Type}}}}); ok {
				if typedEvidence, ok := nullEvidence.(types.Object); ok {
					state.Evidence = typedEvidence
				}
			}
		}
	}
	{
		if rawValueID, rawOk := plainValueAtPath(raw, "id"); rawOk {
			if valueID, err := flattenPlainValue(rawValueID, types.StringType, "id", "raw response"); err != nil {
				return err
			} else {
				if typedID, ok := valueID.(types.String); ok {
					state.ID = typedID
				}
			}
		} else if !hasRaw {
			if responseValueID, ok := plainFromResponseField(obj, "ID"); ok {
				if valueID, err := flattenPlainValue(responseValueID, types.StringType, "id", "response struct"); err != nil {
					return err
				} else {
					if typedID, ok := valueID.(types.String); ok {
						state.ID = typedID
					}
				}
			}
		}
	}
	{
		if rawValueLivemode, rawOk := plainValueAtPath(raw, "livemode"); rawOk {
			if valueLivemode, err := flattenPlainValue(rawValueLivemode, types.BoolType, "livemode", "raw response"); err != nil {
				return err
			} else {
				if typedLivemode, ok := valueLivemode.(types.Bool); ok {
					state.Livemode = typedLivemode
				}
			}
		} else if !hasRaw {
			if responseValueLivemode, ok := plainFromResponseField(obj, "Livemode"); ok {
				if valueLivemode, err := flattenPlainValue(responseValueLivemode, types.BoolType, "livemode", "response struct"); err != nil {
					return err
				} else {
					if typedLivemode, ok := valueLivemode.(types.Bool); ok {
						state.Livemode = typedLivemode
					}
				}
			}
		}
	}
	{
		if rawValueLossReason, rawOk := plainValueAtPath(raw, "loss_reason"); rawOk {
			if valueLossReason, err := flattenPlainValue(rawValueLossReason, types.StringType, "loss_reason", "raw response"); err != nil {
				return err
			} else {
				if typedLossReason, ok := valueLossReason.(types.String); ok {
					state.LossReason = typedLossReason
				}
			}
		} else if !hasRaw {
			if responseValueLossReason, ok := plainFromResponseField(obj, "LossReason"); ok {
				if valueLossReason, err := flattenPlainValue(responseValueLossReason, types.StringType, "loss_reason", "response struct"); err != nil {
					return err
				} else {
					if typedLossReason, ok := valueLossReason.(types.String); ok {
						state.LossReason = typedLossReason
					}
				}
			}
		}
	}
	{
		if rawValueMetadata, rawOk := plainValueAtPath(raw, "metadata"); rawOk {
			if valueMetadata, err := flattenPlainValue(applyConfiguredKeyedListShapes(rawValueMetadata, attrValueToPlain(state.Metadata)), types.MapType{ElemType: types.StringType}, "metadata", "raw response"); err != nil {
				return err
			} else {
				if typedMetadata, ok := valueMetadata.(types.Map); ok {
					state.Metadata = typedMetadata
				}
			}
		} else if !hasRaw {
			if responseValueMetadata, ok := plainFromResponseField(obj, "Metadata"); ok {
				if valueMetadata, err := flattenPlainValue(
					applyConfiguredKeyedListShapes(responseValueMetadata, attrValueToPlain(state.Metadata)),
					types.MapType{ElemType: types.StringType},
					"metadata",
					"response struct",
				); err != nil {
					return err
				} else {
					if typedMetadata, ok := valueMetadata.(types.Map); ok {
						state.Metadata = typedMetadata
					}
				}
			}
		}
	}
	{
		if rawValueStatus, rawOk := plainValueAtPath(raw, "status"); rawOk {
			if valueStatus, err := flattenPlainValue(rawValueStatus, types.StringType, "status", "raw response"); err != nil {
				return err
			} else {
				if typedStatus, ok := valueStatus.(types.String); ok {
					state.Status = typedStatus
				}
			}
		} else if !hasRaw {
			if responseValueStatus, ok := plainFromResponseField(obj, "Status"); ok {
				if valueStatus, err := flattenPlainValue(responseValueStatus, types.StringType, "status", "response struct"); err != nil {
					return err
				} else {
					if typedStatus, ok := valueStatus.(types.String); ok {
						state.Status = typedStatus
					}
				}
			}
		}
	}
	{
		if state.Transaction.IsNull() || state.Transaction.IsUnknown() {
			if rawValueTransaction, rawOk := plainValueAtPath(raw, "transaction"); rawOk {
				if typedTransaction, ok := plainToStringIDValue(rawValueTransaction); ok {
					state.Transaction = typedTransaction
				}
			} else if !hasRaw {
				if responseValueTransaction, ok := plainFromResponseField(obj, "Transaction"); ok {
					if typedTransaction, ok := plainToStringIDValue(responseValueTransaction); ok {
						state.Transaction = typedTransaction
					}
				}
			}
		}
	}
	{
		assignedTreasury := false
		hadRawTreasury := false
		if rawValueTreasury, rawOk := plainValueAtPath(raw, "treasury"); rawOk {
			hadRawTreasury = true
			if rawValueTreasury != nil {
				sourceTreasury := applyConfiguredKeyedListShapes(rawValueTreasury, attrValueToPlain(state.Treasury))
				if valueTreasury, err := flattenPlainValue(sourceTreasury, types.ObjectType{AttrTypes: map[string]attr.Type{"debit_reversal": types.StringType, "received_debit": types.StringType}}, "treasury", "raw response"); err != nil {
					return err
				} else {
					if typedTreasury, ok := valueTreasury.(types.Object); ok {
						state.Treasury = typedTreasury
						assignedTreasury = true
					}
				}
			}
		}
		if !assignedTreasury {
			if !hasRaw {
				if responseValueTreasury, ok := plainFromResponseField(obj, "Treasury"); ok {
					sourceTreasury := applyConfiguredKeyedListShapes(responseValueTreasury, attrValueToPlain(state.Treasury))
					if valueTreasury, err := flattenPlainValue(
						sourceTreasury,
						types.ObjectType{AttrTypes: map[string]attr.Type{"debit_reversal": types.StringType, "received_debit": types.StringType}},
						"treasury",
						"response struct",
					); err != nil {
						return err
					} else {
						if typedTreasury, ok := valueTreasury.(types.Object); ok {
							state.Treasury = typedTreasury
							assignedTreasury = true
						}
					}
				}
			}
		}
		if !assignedTreasury && hadRawTreasury {
			if nullTreasury, ok := nullTerraformValue(types.ObjectType{AttrTypes: map[string]attr.Type{"debit_reversal": types.StringType, "received_debit": types.StringType}}); ok {
				if typedTreasury, ok := nullTreasury.(types.Object); ok {
					state.Treasury = typedTreasury
				}
			}
		}
	}
	return nil
}
