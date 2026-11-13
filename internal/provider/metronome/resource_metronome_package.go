//
// File generated from our OpenAPI spec
//

package metronome

import (
	"context"
	"errors"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
)

type MetronomePackageResource struct{ transport bindingruntime.Transport }

type MetronomePackageResourceModel struct {
	ContractName    types.String  `tfsdk:"contract_name"`
	ID              types.String  `tfsdk:"id"`
	Name            types.String  `tfsdk:"name"`
	RateCardID      types.String  `tfsdk:"rate_card_id"`
	RecurringCommit types.Object  `tfsdk:"recurring_commit"`
	Request         types.Dynamic `tfsdk:"request"`
	RequestJson     types.String  `tfsdk:"request_json"`
	StripeBilling   types.Bool    `tfsdk:"stripe_billing"`
	UniquenessKey   types.String  `tfsdk:"uniqueness_key"`
	Package         types.Object  `tfsdk:"package"`
}

func expandMetronomePackageResourceModel(ctx context.Context, model *MetronomePackageResourceModel) *bindingruntime.Values {
	return bindingruntime.NewValues(types.ObjectValueMust(map[string]attr.Type{"contract_name": model.ContractName.Type(ctx), "id": model.ID.Type(ctx), "name": model.Name.Type(ctx), "rate_card_id": model.RateCardID.Type(ctx), "recurring_commit": model.RecurringCommit.Type(ctx), "request": model.Request.Type(ctx), "request_json": model.RequestJson.Type(ctx), "stripe_billing": model.StripeBilling.Type(ctx), "uniqueness_key": model.UniquenessKey.Type(ctx), "package": model.Package.Type(ctx)}, map[string]attr.Value{"contract_name": model.ContractName, "id": model.ID, "name": model.Name, "rate_card_id": model.RateCardID, "recurring_commit": model.RecurringCommit, "request": model.Request, "request_json": model.RequestJson, "stripe_billing": model.StripeBilling, "uniqueness_key": model.UniquenessKey, "package": model.Package}))
}

func NewMetronomePackageResource() resource.Resource { return &MetronomePackageResource{} }

func (r *MetronomePackageResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metronome_package"
}

func (r *MetronomePackageResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Metronome package. Prefer the typed package attribute, generated from CreatePackagePayload in the pinned OpenAPI spec. request and request_json remain for compatibility and API escape hatches. Destroy archives the package.",
		Attributes: map[string]schema.Attribute{
			"contract_name": schema.StringAttribute{
				Optional: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Optional: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"rate_card_id": schema.StringAttribute{
				Optional: true,

				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"recurring_commit": schema.SingleNestedAttribute{
				Optional:      true,
				Description:   "One monthly prepaid commit; omit for a usage-only package.",
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
				Attributes: map[string]schema.Attribute{
					"applicable_product_ids": schema.ListAttribute{
						Required: true,

						ElementType: types.StringType,
					},
					"credit_type_id": schema.StringAttribute{
						Required: true,
					},
					"fixed_product_id": schema.StringAttribute{
						Required: true,
					},
					"grant_duration_periods": schema.Int64Attribute{
						Required:    true,
						Description: "Number of monthly periods before each grant expires.",
					},
					"monthly_grant_cents": schema.Int64Attribute{
						Required: true,
					},
					"monthly_price_cents": schema.Int64Attribute{
						Required: true,
					},
					"name": schema.StringAttribute{
						Required: true,
					},
					"priority": schema.Int64Attribute{
						Required:    true,
						Description: "Use 10 so monthly grants spend before runtime top-ups at 100.",
					},
					"rollover_fraction": schema.NumberAttribute{
						Required:    true,
						Description: "Fraction of unused grant that rolls over on contract transition, from 0 to 1.",
					},
				},
			},
			"request": schema.DynamicAttribute{
				Optional:    true,
				Description: "API-shaped HCL escape hatch. Use package for typed fields.",
			},
			"request_json": schema.StringAttribute{
				Optional:    true,
				Description: "JSON createPackage request escape hatch.",
			},
			"stripe_billing": schema.BoolAttribute{
				Optional:      true,
				Description:   "Legacy monthly configuration: use Stripe direct billing.",
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"uniqueness_key": schema.StringAttribute{
				Optional:      true,
				Description:   "Optional duplicate-prevention key, including the package: prefix. Omit it for repeatable create/destroy; Metronome generates the package ID stored in state. Explicit keys remain reserved after archive and replacements require a new key. Raw modes put the key inside their payload.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"package": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Typed CreatePackagePayload from the pinned Metronome OpenAPI specification. Use this, request, request_json, or legacy fields exclusively.",
				Attributes: map[string]schema.Attribute{
					"aliases": schema.ListNestedAttribute{
						Optional:    true,
						Description: "Reference this alias when creating a contract. If the same alias is assigned to multiple packages, it will reference the package to which it was most recently assigned. It is not exposed to end customers.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"ending_before": schema.StringAttribute{
									Optional: true,
								},
								"name": schema.StringAttribute{
									Required: true,
								},
								"starting_at": schema.StringAttribute{
									Optional: true,
								},
							},
						},
					},
					"billing_provider": schema.StringAttribute{
						Optional: true,
					},
					"commits": schema.ListNestedAttribute{
						Optional: true,

						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"access_schedule": schema.SingleNestedAttribute{
									Required:    true,
									Description: "Required: Schedule for distributing the commit to the customer. For \"POSTPAID\" commits only one schedule item is allowed and amount must match invoice_schedule total.",
									Attributes: map[string]schema.Attribute{
										"access_type": schema.StringAttribute{
											Optional:    true,
											Description: "Determines how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used. Defaults to `SPEND` if omitted.",
										},
										"credit_type_id": schema.StringAttribute{
											Optional:    true,
											Description: "Defaults to USD (cents) if not passed",
										},
										"schedule_items": schema.ListNestedAttribute{
											Required: true,

											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"amount": schema.NumberAttribute{
														Required: true,
													},
													"duration": schema.SingleNestedAttribute{
														Required:    true,
														Description: "Offset relative to the start of this segment indicating when it should end.",
														Attributes: map[string]schema.Attribute{
															"unit": schema.StringAttribute{
																Required: true,
															},
															"value": schema.NumberAttribute{
																Required: true,
															},
														},
													},
													"starting_at_offset": schema.SingleNestedAttribute{
														Required:    true,
														Description: "Date relative to the contract start date indicating the start of this schedule segment.",
														Attributes: map[string]schema.Attribute{
															"unit": schema.StringAttribute{
																Required: true,
															},
															"value": schema.NumberAttribute{
																Required: true,
															},
														},
													},
												},
											},
										},
									},
								},
								"applicable_product_ids": schema.ListAttribute{
									Optional:    true,
									Description: "Which products the commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.",
									ElementType: types.StringType,
								},
								"applicable_product_tags": schema.ListAttribute{
									Optional:    true,
									Description: "Which tags the commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.",
									ElementType: types.StringType,
								},
								"custom_fields": schema.MapAttribute{
									Optional: true,

									ElementType: types.StringType,
								},
								"description": schema.StringAttribute{
									Optional:    true,
									Description: "Used only in UI/API. It is not exposed to end customers.",
								},
								"invoice_schedule": schema.SingleNestedAttribute{
									Optional:    true,
									Description: "Required for \"POSTPAID\" commits: the true up invoice will be generated at this time and only one schedule item is allowed; the total must match access_schedule amount. Optional for \"PREPAID\" commits: if not provided, this will be a \"complimentary\" commit with no invoice.",
									Attributes: map[string]schema.Attribute{
										"credit_type_id": schema.StringAttribute{
											Optional:    true,
											Description: "Defaults to USD (cents) if not passed.",
										},
										"do_not_invoice": schema.BoolAttribute{
											Optional:    true,
											Description: "If true, this schedule will not generate an invoice.",
										},
										"schedule_items": schema.ListNestedAttribute{
											Required:    true,
											Description: "Either provide amount or provide both unit_price and quantity.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"date_offset": schema.SingleNestedAttribute{
														Required:    true,
														Description: "Date relative to the contract start date.",
														Attributes: map[string]schema.Attribute{
															"unit": schema.StringAttribute{
																Required: true,
															},
															"value": schema.NumberAttribute{
																Required: true,
															},
														},
													},
													"quantity": schema.NumberAttribute{
														Required:    true,
														Description: "Quantity for the charge. Will be multiplied by unit_price to determine the amount.",
													},
													"unit_price": schema.NumberAttribute{
														Required:    true,
														Description: "Unit price for the charge. Will be multiplied by quantity to determine the amount.",
													},
												},
											},
										},
									},
								},
								"name": schema.StringAttribute{
									Optional:    true,
									Description: "displayed on invoices",
								},
								"priority": schema.NumberAttribute{
									Optional:    true,
									Description: "If multiple commits are applicable, the one with the lower priority will apply first.",
								},
								"product_id": schema.StringAttribute{
									Required: true,
								},
								"rate_type": schema.StringAttribute{
									Optional: true,
								},
								"rollover_fraction": schema.NumberAttribute{
									Optional:    true,
									Description: "Fraction of unused segments that will be rolled over. Must be between 0 and 1.",
								},
								"specifiers": schema.ListNestedAttribute{
									Optional:    true,
									Description: "List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"exclude": schema.ListNestedAttribute{
												Optional:    true,
												Description: "If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"product_tags": schema.ListAttribute{
															Optional:    true,
															Description: "If provided, the specifier will not apply to products with all the specified tags.",
															ElementType: types.StringType,
														},
													},
												},
											},
											"presentation_group_values": schema.MapAttribute{
												Optional:    true,
												Description: "If provided, the specifier will apply to product usage with these set of presentation group values.",
												ElementType: types.StringType,
											},
											"pricing_group_values": schema.MapAttribute{
												Optional:    true,
												Description: "If provided, the specifier will apply to product usage with these set of pricing group values.",
												ElementType: types.StringType,
											},
											"product_id": schema.StringAttribute{
												Optional:    true,
												Description: "If provided, the specifier will only apply to the product with the specified ID.",
											},
											"product_tags": schema.ListAttribute{
												Optional:    true,
												Description: "If provided, the specifier will only apply to products with all the specified tags.",
												ElementType: types.StringType,
											},
										},
									},
								},
								"temporary_id": schema.StringAttribute{
									Optional:    true,
									Description: "A temporary ID for the commit that can be used to reference the commit for commit specific overrides.",
								},
								"type": schema.StringAttribute{
									Required: true,
								},
							},
						},
					},
					"contract_name": schema.StringAttribute{
						Optional: true,
					},
					"credits": schema.ListNestedAttribute{
						Optional: true,

						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"access_schedule": schema.SingleNestedAttribute{
									Required:    true,
									Description: "Schedule for distributing the credit to the customer.",
									Attributes: map[string]schema.Attribute{
										"access_type": schema.StringAttribute{
											Optional:    true,
											Description: "Determines how the balance is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used. Defaults to `SPEND` if omitted.",
										},
										"credit_type_id": schema.StringAttribute{
											Optional:    true,
											Description: "Defaults to USD (cents) if not passed",
										},
										"schedule_items": schema.ListNestedAttribute{
											Required: true,

											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"amount": schema.NumberAttribute{
														Required: true,
													},
													"duration": schema.SingleNestedAttribute{
														Required:    true,
														Description: "Offset relative to the start of this segment indicating when it should end.",
														Attributes: map[string]schema.Attribute{
															"unit": schema.StringAttribute{
																Required: true,
															},
															"value": schema.NumberAttribute{
																Required: true,
															},
														},
													},
													"starting_at_offset": schema.SingleNestedAttribute{
														Required:    true,
														Description: "Date relative to the contract start date indicating the start of this schedule segment.",
														Attributes: map[string]schema.Attribute{
															"unit": schema.StringAttribute{
																Required: true,
															},
															"value": schema.NumberAttribute{
																Required: true,
															},
														},
													},
												},
											},
										},
									},
								},
								"applicable_product_ids": schema.ListAttribute{
									Optional:    true,
									Description: "Which products the credit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the credit applies to all products.",
									ElementType: types.StringType,
								},
								"applicable_product_tags": schema.ListAttribute{
									Optional:    true,
									Description: "Which tags the credit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the credit applies to all products.",
									ElementType: types.StringType,
								},
								"custom_fields": schema.MapAttribute{
									Optional: true,

									ElementType: types.StringType,
								},
								"description": schema.StringAttribute{
									Optional:    true,
									Description: "Used only in UI/API. It is not exposed to end customers.",
								},
								"name": schema.StringAttribute{
									Optional:    true,
									Description: "displayed on invoices",
								},
								"priority": schema.NumberAttribute{
									Optional:    true,
									Description: "If multiple credits are applicable, the one with the lower priority will apply first.",
								},
								"product_id": schema.StringAttribute{
									Required: true,
								},
								"rate_type": schema.StringAttribute{
									Optional: true,
								},
								"specifiers": schema.ListNestedAttribute{
									Optional:    true,
									Description: "List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"exclude": schema.ListNestedAttribute{
												Optional:    true,
												Description: "If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"product_tags": schema.ListAttribute{
															Optional:    true,
															Description: "If provided, the specifier will not apply to products with all the specified tags.",
															ElementType: types.StringType,
														},
													},
												},
											},
											"presentation_group_values": schema.MapAttribute{
												Optional:    true,
												Description: "If provided, the specifier will apply to product usage with these set of presentation group values.",
												ElementType: types.StringType,
											},
											"pricing_group_values": schema.MapAttribute{
												Optional:    true,
												Description: "If provided, the specifier will apply to product usage with these set of pricing group values.",
												ElementType: types.StringType,
											},
											"product_id": schema.StringAttribute{
												Optional:    true,
												Description: "If provided, the specifier will only apply to the product with the specified ID.",
											},
											"product_tags": schema.ListAttribute{
												Optional:    true,
												Description: "If provided, the specifier will only apply to products with all the specified tags.",
												ElementType: types.StringType,
											},
										},
									},
								},
							},
						},
					},
					"delivery_method": schema.StringAttribute{
						Optional: true,
					},
					"duration": schema.SingleNestedAttribute{
						Optional: true,

						Attributes: map[string]schema.Attribute{
							"unit": schema.StringAttribute{
								Required: true,
							},
							"value": schema.NumberAttribute{
								Required: true,
							},
						},
					},
					"multiplier_override_prioritization": schema.StringAttribute{
						Optional:    true,
						Description: "Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list prices automatically. EXPLICIT prioritization requires specifying priorities for each multiplier; the one with the lowest priority value will be prioritized first. If tiered overrides are used, prioritization must be explicit.",
					},
					"name": schema.StringAttribute{
						Required: true,
					},
					"net_payment_terms_days": schema.NumberAttribute{
						Optional: true,
					},
					"overrides": schema.DynamicAttribute{
						Optional:    true,
						Description: " (contains an opaque API field)",
					},
					"prepaid_balance_threshold_configuration": schema.SingleNestedAttribute{
						Optional: true,

						Attributes: map[string]schema.Attribute{
							"commit": schema.SingleNestedAttribute{
								Required: true,

								Attributes: map[string]schema.Attribute{
									"applicable_product_ids": schema.ListAttribute{
										Optional:    true,
										Description: "Which products the threshold commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.",
										ElementType: types.StringType,
									},
									"applicable_product_tags": schema.ListAttribute{
										Optional:    true,
										Description: "Which tags the threshold commit applies to. If applicable_product_ids, applicable_product_tags or specifiers are not provided, the commit applies to all products.",
										ElementType: types.StringType,
									},
									"description": schema.StringAttribute{
										Optional: true,
									},
									"duration": schema.SingleNestedAttribute{
										Optional:    true,
										Description: "The length of time the created commit will be valid, starting from the end of the invoice's service period. If not provided, defaults to one year.",
										Attributes: map[string]schema.Attribute{
											"unit": schema.StringAttribute{
												Required: true,
											},
											"value": schema.NumberAttribute{
												Required: true,
											},
										},
									},
									"name": schema.StringAttribute{
										Optional:    true,
										Description: "Specify the name of the line item for the threshold charge. If left blank, it will default to the commit product name.",
									},
									"priority": schema.NumberAttribute{
										Optional:    true,
										Description: "The priority of the commit, used to determine drawdown order. Lower priority commits are consumed first. Defaults to 100 if not specified.",
									},
									"product_id": schema.StringAttribute{
										Required:    true,
										Description: "The commit product that will be used to generate the line item for commit payment.",
									},
									"rate_type": schema.StringAttribute{
										Optional:    true,
										Description: "Whether the created commits will be charged at commit rate or list rate.",
									},
									"rollover_fraction": schema.NumberAttribute{
										Optional:    true,
										Description: "Fraction of the created commit's unused balance that will roll over. Must be between 0 and 1.",
									},
									"specifiers": schema.ListNestedAttribute{
										Optional:    true,
										Description: "List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"exclude": schema.ListNestedAttribute{
													Optional:    true,
													Description: "If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"product_tags": schema.ListAttribute{
																Optional:    true,
																Description: "If provided, the specifier will not apply to products with all the specified tags.",
																ElementType: types.StringType,
															},
														},
													},
												},
												"presentation_group_values": schema.MapAttribute{
													Optional:    true,
													Description: "If provided, the specifier will apply to product usage with these set of presentation group values.",
													ElementType: types.StringType,
												},
												"pricing_group_values": schema.MapAttribute{
													Optional:    true,
													Description: "If provided, the specifier will apply to product usage with these set of pricing group values.",
													ElementType: types.StringType,
												},
												"product_id": schema.StringAttribute{
													Optional:    true,
													Description: "If provided, the specifier will only apply to the product with the specified ID.",
												},
												"product_tags": schema.ListAttribute{
													Optional:    true,
													Description: "If provided, the specifier will only apply to products with all the specified tags.",
													ElementType: types.StringType,
												},
											},
										},
									},
								},
							},
							"custom_credit_type_id": schema.StringAttribute{
								Optional:    true,
								Description: "If provided, the threshold, recharge-to amount, and the resulting threshold commit amount will be in terms of this credit type instead of the fiat currency.",
							},
							"discount_configuration": schema.SingleNestedAttribute{
								Optional: true,

								Attributes: map[string]schema.Attribute{
									"cap": schema.SingleNestedAttribute{
										Optional:    true,
										Description: "If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.",
										Attributes: map[string]schema.Attribute{
											"amount": schema.NumberAttribute{
												Required:    true,
												Description: "Accumulated spend ceiling above which the discount stops applying.",
											},
											"spend_tracker_alias": schema.StringAttribute{
												Required:    true,
												Description: "Alias of the spend tracker this cap is measured against.",
											},
										},
									},
									"payment_fraction": schema.NumberAttribute{
										Required:    true,
										Description: "The fraction of the original amount that the customer pays after applying the discount. For example, 0.85 means the customer pays 85% of the original amount (a 15% discount).",
									},
								},
							},
							"is_enabled": schema.BoolAttribute{
								Required:    true,
								Description: "When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.",
							},
							"payment_gate_config": schema.SingleNestedAttribute{
								Required: true,

								Attributes: map[string]schema.Attribute{
									"payment_gate_type": schema.StringAttribute{
										Required:    true,
										Description: "Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.",
									},
									"precalculated_tax_config": schema.SingleNestedAttribute{
										Optional:    true,
										Description: "Only applicable if using PRECALCULATED as your tax type.",
										Attributes: map[string]schema.Attribute{
											"tax_amount": schema.NumberAttribute{
												Required:    true,
												Description: "Amount of tax to be applied. This should be in the same currency and denomination  as the commit's invoice schedule",
											},
											"tax_name": schema.StringAttribute{
												Optional:    true,
												Description: "Name of the tax to be applied. This may be used in an invoice line item description.",
											},
										},
									},
									"stripe_config": schema.SingleNestedAttribute{
										Optional:    true,
										Description: "Only applicable if using STRIPE as your payment gate type.",
										Attributes: map[string]schema.Attribute{
											"invoice_metadata": schema.MapAttribute{
												Optional:    true,
												Description: "Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as your payment type.",
												ElementType: types.StringType,
											},
											"payment_type": schema.StringAttribute{
												Required:    true,
												Description: "If left blank, will default to INVOICE",
											},
										},
									},
									"tax_type": schema.StringAttribute{
										Optional:    true,
										Description: "Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.",
									},
								},
							},
							"recharge_to_amount": schema.NumberAttribute{
								Required:    true,
								Description: "Specify the amount the balance should be recharged to.",
							},
							"threshold_amount": schema.NumberAttribute{
								Required:    true,
								Description: "Specify the threshold amount for the contract. Each time the contract's prepaid balance lowers to this amount, a threshold charge will be initiated.",
							},
							"threshold_balance_specifiers": schema.ListNestedAttribute{
								Optional:    true,
								Description: "Determines which balances are excluded from remaining balance calculation for threshold billing.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"exclude": schema.ListNestedAttribute{
											Required:    true,
											Description: "If any of the exclude specifier is met, the balance is not considered when evaluating threshold billing",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"custom_field_filters": schema.ListNestedAttribute{
														Required:    true,
														Description: "If provided, balances with all the custom fields will not be considered when evaluating threshold billing",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"entity": schema.StringAttribute{
																	Required: true,
																},
																"key": schema.StringAttribute{
																	Required: true,
																},
																"value": schema.StringAttribute{
																	Required: true,
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
					"rate_card_alias": schema.StringAttribute{
						Optional:    true,
						Description: "Selects the rate card linked to the specified alias as of the contract's start date.",
					},
					"rate_card_id": schema.StringAttribute{
						Optional: true,
					},
					"recurring_commits": schema.ListNestedAttribute{
						Optional: true,

						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"access_amount": schema.SingleNestedAttribute{
									Required:    true,
									Description: "The amount of commit to grant.",
									Attributes: map[string]schema.Attribute{
										"access_type": schema.StringAttribute{
											Optional:    true,
											Description: "Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.",
										},
										"credit_type_id": schema.StringAttribute{
											Optional:    true,
											Description: "Defaults to USD (cents) if not passed",
										},
										"quantity": schema.NumberAttribute{
											Optional:    true,
											Description: "This field is required unless a subscription is attached via `subscription_config`.",
										},
										"unit_price": schema.NumberAttribute{
											Required: true,
										},
									},
								},
								"applicable_product_ids": schema.ListAttribute{
									Optional:    true,
									Description: "Will be passed down to the individual commits",
									ElementType: types.StringType,
								},
								"applicable_product_tags": schema.ListAttribute{
									Optional:    true,
									Description: "Will be passed down to the individual commits",
									ElementType: types.StringType,
								},
								"commit_duration": schema.SingleNestedAttribute{
									Required:    true,
									Description: "Defines the length of the access schedule for each created commit/credit. The value represents the number of units. Unit defaults to \"PERIODS\", where the length of a period is determined by the recurrence_frequency.",
									Attributes: map[string]schema.Attribute{
										"unit": schema.StringAttribute{
											Optional: true,
										},
										"value": schema.NumberAttribute{
											Required: true,
										},
									},
								},
								"description": schema.StringAttribute{
									Optional:    true,
									Description: "Will be passed down to the individual commits",
								},
								"duration": schema.SingleNestedAttribute{
									Optional:    true,
									Description: "Offset relative to the recurring credit start that determines when the contract will stop creating recurring commits. optional",
									Attributes: map[string]schema.Attribute{
										"unit": schema.StringAttribute{
											Required: true,
										},
										"value": schema.NumberAttribute{
											Required: true,
										},
									},
								},
								"invoice_amount": schema.SingleNestedAttribute{
									Optional:    true,
									Description: "The amount the customer should be billed for the commit. Not required.",
									Attributes: map[string]schema.Attribute{
										"credit_type_id": schema.StringAttribute{
											Required: true,
										},
										"quantity": schema.NumberAttribute{
											Required: true,
										},
										"unit_price": schema.NumberAttribute{
											Required: true,
										},
									},
								},
								"name": schema.StringAttribute{
									Optional:    true,
									Description: "displayed on invoices. will be passed through to the individual commits",
								},
								"priority": schema.NumberAttribute{
									Required:    true,
									Description: "Will be passed down to the individual commits",
								},
								"product_id": schema.StringAttribute{
									Required: true,
								},
								"proration": schema.StringAttribute{
									Optional:    true,
									Description: "Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).",
								},
								"proration_rounding": schema.SingleNestedAttribute{
									Optional:    true,
									Description: "Optional rounding configuration for prorated recurring commit amounts.",
									Attributes: map[string]schema.Attribute{
										"access": schema.SingleNestedAttribute{
											Optional: true,

											Attributes: map[string]schema.Attribute{
												"decimal_places": schema.NumberAttribute{
													Required:    true,
													Description: "Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit. For USD, this means rounding to the nearest dollar).",
												},
												"rounding_method": schema.StringAttribute{
													Required: true,
												},
											},
										},
										"invoice": schema.SingleNestedAttribute{
											Optional: true,

											Attributes: map[string]schema.Attribute{
												"decimal_places": schema.NumberAttribute{
													Required:    true,
													Description: "Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit. For USD, this means rounding to the nearest dollar).",
												},
												"rounding_method": schema.StringAttribute{
													Required: true,
												},
											},
										},
									},
								},
								"rate_type": schema.StringAttribute{
									Optional:    true,
									Description: "Whether the created commits will use the commit rate or list rate",
								},
								"recurrence_frequency": schema.StringAttribute{
									Optional:    true,
									Description: "The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions",
								},
								"rollover_fraction": schema.NumberAttribute{
									Optional:    true,
									Description: "Will be passed down to the individual commits. This controls how much of an individual unexpired commit will roll over upon contract transition. Must be between 0 and 1.",
								},
								"specifiers": schema.ListNestedAttribute{
									Optional:    true,
									Description: "List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"exclude": schema.ListNestedAttribute{
												Optional:    true,
												Description: "If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"product_tags": schema.ListAttribute{
															Optional:    true,
															Description: "If provided, the specifier will not apply to products with all the specified tags.",
															ElementType: types.StringType,
														},
													},
												},
											},
											"presentation_group_values": schema.MapAttribute{
												Optional:    true,
												Description: "If provided, the specifier will apply to product usage with these set of presentation group values.",
												ElementType: types.StringType,
											},
											"pricing_group_values": schema.MapAttribute{
												Optional:    true,
												Description: "If provided, the specifier will apply to product usage with these set of pricing group values.",
												ElementType: types.StringType,
											},
											"product_id": schema.StringAttribute{
												Optional:    true,
												Description: "If provided, the specifier will only apply to the product with the specified ID.",
											},
											"product_tags": schema.ListAttribute{
												Optional:    true,
												Description: "If provided, the specifier will only apply to products with all the specified tags.",
												ElementType: types.StringType,
											},
										},
									},
								},
								"starting_at_offset": schema.SingleNestedAttribute{
									Required:    true,
									Description: "Offset relative to the contract start date that determines the start time for the first commit",
									Attributes: map[string]schema.Attribute{
										"unit": schema.StringAttribute{
											Required: true,
										},
										"value": schema.NumberAttribute{
											Required: true,
										},
									},
								},
								"subscription_config": schema.SingleNestedAttribute{
									Optional:    true,
									Description: "Attach a subscription to the recurring commit/credit.",
									Attributes: map[string]schema.Attribute{
										"access_policy": schema.StringAttribute{
											Optional:    true,
											Description: "Controls when a customer can access recurring commit or credit child balances when the attached subscription is payment-gated. `BILLING_PERIOD_PAID` releases each balance after payment for its billing period. `INITIAL_BILLING_PERIOD_PAID_ONLY` releases all balances after the first payment.",
										},
										"allocation": schema.StringAttribute{
											Optional: true,
										},
										"apply_seat_increase_config": schema.SingleNestedAttribute{
											Required: true,

											Attributes: map[string]schema.Attribute{
												"is_prorated": schema.BoolAttribute{
													Required:    true,
													Description: "Indicates whether a mid-period seat increase should be prorated.",
												},
											},
										},
										"subscription_id": schema.StringAttribute{
											Required:    true,
											Description: "ID of the subscription to configure on the recurring commit/credit.",
										},
									},
								},
								"temporary_id": schema.StringAttribute{
									Optional:    true,
									Description: "A temporary ID that can be used to reference the recurring commit for commit specific overrides.",
								},
							},
						},
					},
					"recurring_credits": schema.ListNestedAttribute{
						Optional: true,

						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"access_amount": schema.SingleNestedAttribute{
									Required:    true,
									Description: "The amount of commit to grant.",
									Attributes: map[string]schema.Attribute{
										"access_type": schema.StringAttribute{
											Optional:    true,
											Description: "Indicates how the balance of child commits is drawn down. `SPEND` deducts the dollar cost of usage. `QUANTITY` deducts the number of units used.",
										},
										"credit_type_id": schema.StringAttribute{
											Optional:    true,
											Description: "Defaults to USD (cents) if not passed",
										},
										"quantity": schema.NumberAttribute{
											Optional:    true,
											Description: "This field is required unless a subscription is attached via `subscription_config`.",
										},
										"unit_price": schema.NumberAttribute{
											Required: true,
										},
									},
								},
								"applicable_product_ids": schema.ListAttribute{
									Optional:    true,
									Description: "Will be passed down to the individual commits",
									ElementType: types.StringType,
								},
								"applicable_product_tags": schema.ListAttribute{
									Optional:    true,
									Description: "Will be passed down to the individual commits",
									ElementType: types.StringType,
								},
								"commit_duration": schema.SingleNestedAttribute{
									Required:    true,
									Description: "Defines the length of the access schedule for each created commit/credit. The value represents the number of units. Unit defaults to \"PERIODS\", where the length of a period is determined by the recurrence_frequency.",
									Attributes: map[string]schema.Attribute{
										"unit": schema.StringAttribute{
											Optional: true,
										},
										"value": schema.NumberAttribute{
											Required: true,
										},
									},
								},
								"description": schema.StringAttribute{
									Optional:    true,
									Description: "Will be passed down to the individual commits",
								},
								"duration": schema.SingleNestedAttribute{
									Optional:    true,
									Description: "Offset relative to the recurring credit start that determines when the contract will stop creating recurring commits. optional",
									Attributes: map[string]schema.Attribute{
										"unit": schema.StringAttribute{
											Required: true,
										},
										"value": schema.NumberAttribute{
											Required: true,
										},
									},
								},
								"name": schema.StringAttribute{
									Optional:    true,
									Description: "displayed on invoices. will be passed through to the individual commits",
								},
								"priority": schema.NumberAttribute{
									Required:    true,
									Description: "Will be passed down to the individual commits",
								},
								"product_id": schema.StringAttribute{
									Required: true,
								},
								"proration": schema.StringAttribute{
									Optional:    true,
									Description: "Determines whether the first and last commit will be prorated.  If not provided, the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).",
								},
								"proration_rounding": schema.SingleNestedAttribute{
									Optional:    true,
									Description: "Optional rounding configuration for prorated recurring credit amounts.",
									Attributes: map[string]schema.Attribute{
										"access": schema.SingleNestedAttribute{
											Optional: true,

											Attributes: map[string]schema.Attribute{
												"decimal_places": schema.NumberAttribute{
													Required:    true,
													Description: "Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit. For USD, this means rounding to the nearest dollar).",
												},
												"rounding_method": schema.StringAttribute{
													Required: true,
												},
											},
										},
									},
								},
								"rate_type": schema.StringAttribute{
									Optional:    true,
									Description: "Whether the created commits will use the commit rate or list rate",
								},
								"recurrence_frequency": schema.StringAttribute{
									Optional:    true,
									Description: "The frequency at which the recurring commits will be created. If not provided: - The commits will be created on the usage invoice frequency. If provided: - The period defined in the duration will correspond to this frequency. - Commits will be created aligned with the recurring commit's starting_at rather than the usage invoice dates. - Daily recurring commits have a limit of one per contract, and are unable to be created with seat-based subscriptions",
								},
								"rollover_fraction": schema.NumberAttribute{
									Optional:    true,
									Description: "Will be passed down to the individual commits. This controls how much of an individual unexpired commit will roll over upon contract transition. Must be between 0 and 1.",
								},
								"specifiers": schema.ListNestedAttribute{
									Optional:    true,
									Description: "List of filters that determine what kind of customer usage draws down a commit or credit. A customer's usage needs to meet the condition of at least one of the specifiers to contribute to a commit's or credit's drawdown. This field cannot be used together with `applicable_product_ids` or `applicable_product_tags`.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"exclude": schema.ListNestedAttribute{
												Optional:    true,
												Description: "If provided, the specifier will not apply to product usage that matches the inclusion criteria and any of the excluding values.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"product_tags": schema.ListAttribute{
															Optional:    true,
															Description: "If provided, the specifier will not apply to products with all the specified tags.",
															ElementType: types.StringType,
														},
													},
												},
											},
											"presentation_group_values": schema.MapAttribute{
												Optional:    true,
												Description: "If provided, the specifier will apply to product usage with these set of presentation group values.",
												ElementType: types.StringType,
											},
											"pricing_group_values": schema.MapAttribute{
												Optional:    true,
												Description: "If provided, the specifier will apply to product usage with these set of pricing group values.",
												ElementType: types.StringType,
											},
											"product_id": schema.StringAttribute{
												Optional:    true,
												Description: "If provided, the specifier will only apply to the product with the specified ID.",
											},
											"product_tags": schema.ListAttribute{
												Optional:    true,
												Description: "If provided, the specifier will only apply to products with all the specified tags.",
												ElementType: types.StringType,
											},
										},
									},
								},
								"starting_at_offset": schema.SingleNestedAttribute{
									Required:    true,
									Description: "Offset relative to the contract start date that determines the start time for the first commit",
									Attributes: map[string]schema.Attribute{
										"unit": schema.StringAttribute{
											Required: true,
										},
										"value": schema.NumberAttribute{
											Required: true,
										},
									},
								},
								"subscription_config": schema.SingleNestedAttribute{
									Optional:    true,
									Description: "Attach a subscription to the recurring commit/credit.",
									Attributes: map[string]schema.Attribute{
										"access_policy": schema.StringAttribute{
											Optional:    true,
											Description: "Controls when a customer can access recurring commit or credit child balances when the attached subscription is payment-gated. `BILLING_PERIOD_PAID` releases each balance after payment for its billing period. `INITIAL_BILLING_PERIOD_PAID_ONLY` releases all balances after the first payment.",
										},
										"allocation": schema.StringAttribute{
											Optional: true,
										},
										"apply_seat_increase_config": schema.SingleNestedAttribute{
											Required: true,

											Attributes: map[string]schema.Attribute{
												"is_prorated": schema.BoolAttribute{
													Required:    true,
													Description: "Indicates whether a mid-period seat increase should be prorated.",
												},
											},
										},
										"subscription_id": schema.StringAttribute{
											Required:    true,
											Description: "ID of the subscription to configure on the recurring commit/credit.",
										},
									},
								},
								"temporary_id": schema.StringAttribute{
									Optional:    true,
									Description: "A temporary ID that can be used to reference the recurring commit for commit specific overrides.",
								},
							},
						},
					},
					"scheduled_charges": schema.ListNestedAttribute{
						Optional: true,

						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"custom_fields": schema.MapAttribute{
									Optional: true,

									ElementType: types.StringType,
								},
								"name": schema.StringAttribute{
									Optional:    true,
									Description: "displayed on invoices",
								},
								"product_id": schema.StringAttribute{
									Required: true,
								},
								"schedule": schema.SingleNestedAttribute{
									Required: true,

									Attributes: map[string]schema.Attribute{
										"credit_type_id": schema.StringAttribute{
											Optional:    true,
											Description: "Defaults to USD (cents) if not passed.",
										},
										"schedule_items": schema.ListNestedAttribute{
											Required:    true,
											Description: "Either provide amount or provide both unit_price and quantity.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"date_offset": schema.SingleNestedAttribute{
														Required:    true,
														Description: "Date relative to the contract start date.",
														Attributes: map[string]schema.Attribute{
															"unit": schema.StringAttribute{
																Required: true,
															},
															"value": schema.NumberAttribute{
																Required: true,
															},
														},
													},
													"quantity": schema.NumberAttribute{
														Required:    true,
														Description: "Quantity for the charge. Will be multiplied by unit_price to determine the amount.",
													},
													"unit_price": schema.NumberAttribute{
														Required:    true,
														Description: "Unit price for the charge. Will be multiplied by quantity to determine the amount.",
													},
												},
											},
										},
									},
								},
							},
						},
					},
					"scheduled_charges_on_usage_invoices": schema.StringAttribute{
						Optional: true,
					},
					"spend_threshold_configuration": schema.SingleNestedAttribute{
						Optional: true,

						Attributes: map[string]schema.Attribute{
							"commit": schema.SingleNestedAttribute{
								Required: true,

								Attributes: map[string]schema.Attribute{
									"description": schema.StringAttribute{
										Optional: true,
									},
									"name": schema.StringAttribute{
										Optional:    true,
										Description: "Specify the name of the line item for the threshold charge. If left blank, it will default to the commit product name.",
									},
									"priority": schema.NumberAttribute{
										Optional:    true,
										Description: "The priority of the commit, used to determine drawdown order. Lower priority commits are consumed first. Defaults to 100 if not specified.",
									},
									"product_id": schema.StringAttribute{
										Required:    true,
										Description: "The commit product that will be used to generate the line item for commit payment.",
									},
								},
							},
							"discount_configuration": schema.SingleNestedAttribute{
								Optional: true,

								Attributes: map[string]schema.Attribute{
									"cap": schema.SingleNestedAttribute{
										Optional:    true,
										Description: "If provided, the discount stops applying once the spend tracker has accumulated this much spend in the billing period.",
										Attributes: map[string]schema.Attribute{
											"amount": schema.NumberAttribute{
												Required:    true,
												Description: "Accumulated spend ceiling above which the discount stops applying.",
											},
											"spend_tracker_alias": schema.StringAttribute{
												Required:    true,
												Description: "Alias of the spend tracker this cap is measured against.",
											},
										},
									},
									"payment_fraction": schema.NumberAttribute{
										Required:    true,
										Description: "The fraction of the original amount that the customer pays after applying the discount. For example, 0.85 means the customer pays 85% of the original amount (a 15% discount).",
									},
								},
							},
							"is_enabled": schema.BoolAttribute{
								Required:    true,
								Description: "When set to false, the contract will not be evaluated against the threshold_amount. Toggling to true will result an immediate evaluation, regardless of prior state.",
							},
							"payment_gate_config": schema.SingleNestedAttribute{
								Required: true,

								Attributes: map[string]schema.Attribute{
									"payment_gate_type": schema.StringAttribute{
										Required:    true,
										Description: "Gate access to the commit balance based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe. Select EXTERNAL to facilitate payment using your own payment integration. Select NONE if you do not wish to payment gate the commit balance.",
									},
									"precalculated_tax_config": schema.SingleNestedAttribute{
										Optional:    true,
										Description: "Only applicable if using PRECALCULATED as your tax type.",
										Attributes: map[string]schema.Attribute{
											"tax_amount": schema.NumberAttribute{
												Required:    true,
												Description: "Amount of tax to be applied. This should be in the same currency and denomination  as the commit's invoice schedule",
											},
											"tax_name": schema.StringAttribute{
												Optional:    true,
												Description: "Name of the tax to be applied. This may be used in an invoice line item description.",
											},
										},
									},
									"stripe_config": schema.SingleNestedAttribute{
										Optional:    true,
										Description: "Only applicable if using STRIPE as your payment gate type.",
										Attributes: map[string]schema.Attribute{
											"invoice_metadata": schema.MapAttribute{
												Optional:    true,
												Description: "Metadata to be added to the Stripe invoice. Only applicable if using INVOICE as your payment type.",
												ElementType: types.StringType,
											},
											"payment_type": schema.StringAttribute{
												Required:    true,
												Description: "If left blank, will default to INVOICE",
											},
										},
									},
									"tax_type": schema.StringAttribute{
										Optional:    true,
										Description: "Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.",
									},
								},
							},
							"threshold_amount": schema.NumberAttribute{
								Required:    true,
								Description: "Specify the threshold amount for the contract. Each time the contract's usage hits this amount, a threshold charge will be initiated.",
							},
						},
					},
					"spend_trackers": schema.ListNestedAttribute{
						Optional: true,

						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"alias": schema.StringAttribute{
									Required:    true,
									Description: "Human-readable identifier, unique per contract.",
								},
								"applicable_spend_specifiers": schema.ListNestedAttribute{
									Required: true,

									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"discounted": schema.StringAttribute{
												Optional:    true,
												Description: "Filter by whether the spend was discounted. Defaults to ANY if omitted.",
											},
											"sources": schema.ListAttribute{
												Required: true,

												ElementType: types.StringType,
											},
											"spend_type": schema.StringAttribute{
												Required: true,
											},
										},
									},
								},
								"credit_type_id": schema.StringAttribute{
									Required: true,
								},
								"reset_frequency": schema.StringAttribute{
									Required: true,
								},
							},
						},
					},
					"subscriptions": schema.ListNestedAttribute{
						Optional: true,

						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"billing_cycle_config": schema.SingleNestedAttribute{
									Optional: true,

									Attributes: map[string]schema.Attribute{
										"invoice_placement": schema.StringAttribute{
											Optional:    true,
											Description: "Controls whether subscriptions consolidate onto usage invoices. Defaults to ON_USAGE_INVOICE if omitted.",
										},
									},
								},
								"collection_schedule": schema.StringAttribute{
									Required: true,
								},
								"custom_fields": schema.MapAttribute{
									Optional: true,

									ElementType: types.StringType,
								},
								"description": schema.StringAttribute{
									Optional: true,
								},
								"duration": schema.SingleNestedAttribute{
									Optional:    true,
									Description: "Lifetime of the subscription from its start. If not provided, subscription inherits contract end date.",
									Attributes: map[string]schema.Attribute{
										"unit": schema.StringAttribute{
											Required: true,
										},
										"value": schema.NumberAttribute{
											Required: true,
										},
									},
								},
								"initial_quantity": schema.NumberAttribute{
									Optional:    true,
									Description: "The initial quantity for the subscription. It must be non-negative value. Required if quantity_management_mode is QUANTITY_ONLY.",
								},
								"name": schema.StringAttribute{
									Optional: true,
								},
								"payment_gate_config": schema.SingleNestedAttribute{
									Optional: true,

									Attributes: map[string]schema.Attribute{
										"payment_gate_type": schema.StringAttribute{
											Required:    true,
											Description: "Gate access to the subscription based on successful collection of payment. Select STRIPE for Metronome to facilitate payment via Stripe.",
										},
										"tax_type": schema.StringAttribute{
											Optional:    true,
											Description: "Stripe tax is only supported for Stripe payment gateway. Select NONE if you do not wish Metronome to calculate tax on your behalf. Leaving this field blank will default to NONE.",
										},
									},
								},
								"proration": schema.SingleNestedAttribute{
									Required: true,

									Attributes: map[string]schema.Attribute{
										"invoice_behavior": schema.StringAttribute{
											Optional:    true,
											Description: "Indicates how mid-period quantity adjustments are invoiced.\n**BILL_IMMEDIATELY**: Only available when collection schedule is `ADVANCE`. The quantity increase will be billed immediately on the scheduled date.\n**BILL_ON_NEXT_COLLECTION_DATE**: The quantity increase will be billed for in-arrears at the end of the period.\n",
										},
										"is_prorated": schema.BoolAttribute{
											Optional:    true,
											Description: "Indicates if the partial period will be prorated or charged a full amount.",
										},
										"rounding": schema.SingleNestedAttribute{
											Optional: true,

											Attributes: map[string]schema.Attribute{
												"decimal_places": schema.NumberAttribute{
													Required:    true,
													Description: "Number of decimal places to round to. Applied directly to the stored monetary representation. Negative values round to powers of 10 (e.g., -2 rounds to nearest 100 in the stored unit. For USD, this means rounding to the nearest dollar).",
												},
												"rounding_method": schema.StringAttribute{
													Required: true,
												},
											},
										},
									},
								},
								"quantity_management_mode": schema.StringAttribute{
									Optional:    true,
									Description: "Determines how the subscription's quantity is controlled. Defaults to QUANTITY_ONLY. **QUANTITY_ONLY**: The subscription quantity is specified directly on the subscription. `initial_quantity` must be provided with this option. Compatible with recurring commits/credits that use POOLED allocation. **SEAT_BASED**: Use when you want to pass specific seat identifiers (e.g. add user_123) to increment and decrement a subscription quantity, rather than directly providing the quantity. You must use a **SEAT_BASED** subscription to use a linked recurring credit with an allocation per seat. `seat_config` must be provided with this option.",
								},
								"seat_config": schema.SingleNestedAttribute{
									Optional: true,

									Attributes: map[string]schema.Attribute{
										"initial_unassigned_seats": schema.NumberAttribute{
											Optional:    true,
											Description: "The initial amount of unassigned seats on this subscription.",
										},
										"seat_group_key": schema.StringAttribute{
											Required:    true,
											Description: "The property name, sent on usage events, that identifies the seat ID associated with the usage event. For example, the property name might be seat_id or user_id. The property must be set as a group key on billable metrics and a presentation/pricing group key on contract products. This allows linked recurring credits with an allocation per seat to be consumed by only one seat's usage.",
										},
									},
								},
								"starting_at_offset": schema.SingleNestedAttribute{
									Optional:    true,
									Description: "Relative date from contract start date corresponding to the inclusive start time for the subscription. If not provided, defaults to contract start date",
									Attributes: map[string]schema.Attribute{
										"unit": schema.StringAttribute{
											Required: true,
										},
										"value": schema.NumberAttribute{
											Required: true,
										},
									},
								},
								"subscription_alias": schema.StringAttribute{
									Optional:    true,
									Description: "Optional human-readable, per-package-unique alias for this subscription template. Lets a contract customization reference the subscription by alias instead of its template id.",
								},
								"subscription_rate": schema.SingleNestedAttribute{
									Required: true,

									Attributes: map[string]schema.Attribute{
										"billing_frequency": schema.StringAttribute{
											Required:    true,
											Description: "Frequency to bill subscription with. Together with product_id, must match existing rate on the rate card.",
										},
										"product_id": schema.StringAttribute{
											Required:    true,
											Description: "Must be subscription type product",
										},
									},
								},
								"temporary_id": schema.StringAttribute{
									Optional:    true,
									Description: "A temporary ID used to reference the subscription in recurring commit/credit subscription configs created within the same payload.",
								},
							},
						},
					},
					"uniqueness_key": schema.StringAttribute{
						Optional: true,
					},
					"usage_statement_schedule": schema.SingleNestedAttribute{
						Optional: true,

						Attributes: map[string]schema.Attribute{
							"day": schema.StringAttribute{
								Optional:    true,
								Description: "If not provided, defaults to the first day of the month.",
							},
							"frequency": schema.StringAttribute{
								Required: true,
							},
							"invoice_generation_starting_at_offset": schema.SingleNestedAttribute{
								Optional:    true,
								Description: "The offset at which Metronome should start generating usage invoices, relative to the contract start date.  If unspecified, contract start date will be used. This is useful to set if you want to import historical invoices via our 'Create Historical Invoices' API rather than having Metronome automatically generate them.",
								Attributes: map[string]schema.Attribute{
									"unit": schema.StringAttribute{
										Required: true,
									},
									"value": schema.NumberAttribute{
										Required: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *MetronomePackageResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	data, ok := req.ProviderData.(bindingruntime.Transport)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", "Expected the combined provider service clients.")
		return
	}
	r.transport = data
}

func (r *MetronomePackageResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var valuesModel MetronomePackageResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomePackageResourceModel(ctx, &valuesModel)
	if err := validatePackage(ctx, values); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome configuration", err.Error())
		return
	}
}

func (r *MetronomePackageResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() || req.State.Raw.IsNull() {
		return
	}
	var valuesModel MetronomePackageResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomePackageResourceModel(ctx, &valuesModel)
	var priorModel MetronomePackageResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &priorModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	prior := expandMetronomePackageResourceModel(ctx, &priorModel)
	replacement, err := planPackage(ctx, values, prior)
	if err != nil {
		resp.Diagnostics.AddError("Unsafe Metronome change", err.Error())
		return
	}
	for _, attribute := range replacement {
		resp.RequiresReplace = append(resp.RequiresReplace, path.Root(attribute))
	}
	resp.Diagnostics.Append(resp.Plan.Set(ctx, values.Object)...)
}

func (r *MetronomePackageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var valuesModel MetronomePackageResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomePackageResourceModel(ctx, &valuesModel)
	if err := validatePackage(ctx, values); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome configuration", err.Error())
		return
	}

	persist := func(ctx context.Context, state *bindingruntime.Values) error {
		if err := state.ResolveUnknowns(); err != nil {
			return err
		}
		resp.Diagnostics.Append(resp.State.Set(ctx, state.Object)...)
		if resp.Diagnostics.HasError() {
			return errors.New("could not persist created identity")
		}
		return nil
	}
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "createPackage-v1", Service: "metronome", Method: "POST", Path: "/v1/packages/create", Read: false},
		Request:   []bindingruntime.Mapping{},
		Response:  []bindingruntime.ResponseField{{Attribute: "id", Path: []string{"id"}, Required: true, PreserveMissing: false, PreserveConfigured: false, EmptyStringIsNull: false}},
		Root:      []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: packageRequest, ResponseHook: nil, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{}, WhenMissing: "", Persist: true, PersistAttributes: []string{},
		ArchivedPaths: [][]string{}, MissingStatuses: []int{},
		WaitForReady: false, IgnoreResponse: false, OnError: nil,
	}}, values, nil, persist); err != nil {
		resp.Diagnostics.AddError("Could not create Metronome resource", err.Error())
		return
	}
	if err := values.ResolveUnknowns(); err != nil {
		resp.Diagnostics.AddError("Invalid Terraform state", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, values.Object)...)
}

func (r *MetronomePackageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var valuesModel MetronomePackageResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomePackageResourceModel(ctx, &valuesModel)
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "getPackage-v1", Service: "metronome", Method: "POST", Path: "/v1/packages/get", Read: true},
		Request:   []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"package_id"}, Location: "body", Wrap: false, OmitUnknown: false}},
		Response:  []bindingruntime.ResponseField{},
		Root:      []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: readPackage, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
		ArchivedPaths: [][]string{[]string{"archived_at"}}, MissingStatuses: []int{},
		WaitForReady: false, IgnoreResponse: false, OnError: nil,
	}}, values, nil, nil); err != nil {
		if bindingruntime.IsMissing(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Could not refresh Metronome resource", err.Error())
		return
	}

	if err := values.ResolveUnknowns(); err != nil {
		resp.Diagnostics.AddError("Invalid Terraform state", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, values.Object)...)
}

func (r *MetronomePackageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var valuesModel MetronomePackageResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomePackageResourceModel(ctx, &valuesModel)
	var priorModel MetronomePackageResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &priorModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	prior := expandMetronomePackageResourceModel(ctx, &priorModel)
	if err := values.Set("id", prior.Get("id")); err != nil {
		resp.Diagnostics.AddError("Invalid resource identity", err.Error())
		return
	}
	if err := updatePackage(ctx, values, prior); err != nil {
		resp.Diagnostics.AddError("Unsupported Metronome update", err.Error())
		return
	}
	persist := func(ctx context.Context, state *bindingruntime.Values) error {
		resp.Diagnostics.Append(resp.State.Set(ctx, state.Object)...)
		if resp.Diagnostics.HasError() {
			return errors.New("could not persist partial update")
		}
		return nil
	}
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{}, values, prior, persist); err != nil {
		resp.Diagnostics.AddError("Could not update Metronome resource", err.Error())
		return
	}
	if err := values.ResolveUnknowns(); err != nil {
		resp.Diagnostics.AddError("Invalid Terraform state", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, values.Object)...)
}

func (r *MetronomePackageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var valuesModel MetronomePackageResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &valuesModel)...)
	if resp.Diagnostics.HasError() {
		return
	}
	values := expandMetronomePackageResourceModel(ctx, &valuesModel)
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "archivePackage-v1", Service: "metronome", Method: "POST", Path: "/v1/packages/archive", Read: false},
		Request:   []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"package_id"}, Location: "body", Wrap: false, OmitUnknown: false}},
		Response:  []bindingruntime.ResponseField{},
		Root:      []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: nil, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
		ArchivedPaths: [][]string{}, MissingStatuses: []int{},
		WaitForReady: false, IgnoreResponse: true, OnError: nil,
	}}, values, nil, nil); err != nil && !bindingruntime.IsMissing(err) {
		resp.Diagnostics.AddError("Could not archive Metronome resource", err.Error())
		return
	}
}

func (r *MetronomePackageResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var definition resource.SchemaResponse
	r.Schema(ctx, resource.SchemaRequest{}, &definition)
	values := bindingruntime.EmptyValues(definition.Schema.Type().(types.ObjectType))
	if err := importPackage(ctx, values, req.ID); err != nil {
		resp.Diagnostics.AddError("Invalid Metronome import", err.Error())
		return
	}
	if err := bindingruntime.Run(ctx, r.transport, []bindingruntime.Step{{
		Operation: bindingruntime.Operation{ID: "getPackage-v1", Service: "metronome", Method: "POST", Path: "/v1/packages/get", Read: true},
		Request:   []bindingruntime.Mapping{{Attribute: "id", Constant: nil, Path: []string{"package_id"}, Location: "body", Wrap: false, OmitUnknown: false}},
		Response:  []bindingruntime.ResponseField{},
		Root:      []string{"data"}, Selector: map[string]string{},
		Paginate: false, PaginationLocation: "query", PaginationCursor: "next_page",
		RequestHook: nil, ResponseHook: readPackage, ForEach: nil, ForEachQuery: "",
		WhenChanged: []string{}, WhenMissing: "", Persist: false, PersistAttributes: []string{},
		ArchivedPaths: [][]string{[]string{"archived_at"}}, MissingStatuses: []int{},
		WaitForReady: false, IgnoreResponse: false, OnError: nil,
	}}, values, nil, nil); err != nil {
		resp.Diagnostics.AddError("Could not verify import", err.Error())
		return
	}

	if err := values.ResolveUnknowns(); err != nil {
		resp.Diagnostics.AddError("Invalid Terraform state", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, values.Object)...)
}

func (r *MetronomePackageResource) MoveState(_ context.Context) []resource.StateMover {
	return []resource.StateMover{{StateMover: func(ctx context.Context, req resource.MoveStateRequest, resp *resource.MoveStateResponse) {
		if req.SourceProviderAddress != "registry.terraform.io/tonyxiao/metronome" || req.SourceTypeName != "metronome_package" {
			return
		}
		switch req.SourceSchemaVersion {
		case 0:
		default:
			resp.Diagnostics.AddError("Unsupported Metronome source state", "The standalone schema version is not in the reviewed migration contract.")
			return
		}
		if req.SourceRawState == nil {
			resp.Diagnostics.AddError("Missing source state", "No legacy resource state was supplied.")
			return
		}
		var definition resource.SchemaResponse
		r.Schema(ctx, resource.SchemaRequest{}, &definition)
		raw, err := req.SourceRawState.Unmarshal(definition.Schema.Type().TerraformType(ctx))
		if err != nil {
			resp.Diagnostics.AddError("Invalid legacy Metronome state", err.Error())
			return
		}
		resp.TargetState = tfsdk.State{Schema: definition.Schema, Raw: raw}
		resp.TargetPrivate = req.SourcePrivate
	}}}
}
