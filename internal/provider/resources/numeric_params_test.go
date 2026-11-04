// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
package resources

import (
	"net/url"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stripe/stripe-go/v87/form"
)

func TestNumericParams(t *testing.T) {
	encode := func(params interface{}, err error) url.Values {
		t.Helper()
		if err != nil {
			t.Fatal(err)
		}
		values := &form.Values{}
		form.AppendTo(values, params)
		return values.ToValues()
	}
	for _, tc := range []struct {
		name      string
		value     types.Float64
		want      string
		wantFloat string
	}{
		{"fraction", types.Float64Value(1.23456789), "1.23456789", "1.2346"},
		{"small fraction", types.Float64Value(0.000000000001), "0.000000000001", "0.0000"},
		{"zero", types.Float64Value(0), "0", "0.0000"},
		{"negative", types.Float64Value(-12.5), "-12.5", "-12.5000"},
		{"null", types.Float64Null(), "", ""},
		{"unknown", types.Float64Unknown(), "", ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			invoice := InvoiceItemResourceModel{QuantityDecimal: tc.value, UnitAmountDecimal: tc.value}
			climate := ClimateOrderResourceModel{MetricTons: tc.value}
			plan := PlanResourceModel{AmountDecimal: tc.value}
			taxRate := TaxRateResourceModel{Percentage: tc.value}
			for name, values := range map[string]url.Values{
				"climate create":  encode(expandClimateOrderCreate(climate)),
				"tax rate create": encode(expandTaxRateCreate(taxRate)),
				"plan create":     encode(expandPlanCreate(plan)),
				"invoice create":  encode(expandInvoiceItemCreate(invoice)),
				"invoice update":  encode(expandInvoiceItemUpdate(invoice, InvoiceItemResourceModel{})),
			} {
				for _, key := range map[string][]string{
					"climate create":  {"metric_tons"},
					"tax rate create": {"percentage"},
					"plan create":     {"amount_decimal"},
					"invoice create":  {"quantity_decimal", "unit_amount_decimal"},
					"invoice update":  {"quantity_decimal", "unit_amount_decimal"},
				}[name] {
					want := tc.want
					if name == "tax rate create" {
						want = tc.wantFloat
					}
					if got := values.Get(key); got != want {
						t.Errorf("%s %s = %q, want %q", name, key, got, want)
					}
				}
			}
			unchanged := encode(expandInvoiceItemUpdate(invoice, invoice))
			if unchanged.Has("quantity_decimal") || unchanged.Has("unit_amount_decimal") {
				t.Error("unchanged numeric fields were sent in update")
			}
		})
	}
}

func TestPriceDecimalParams(t *testing.T) {
	for _, value := range []string{"0", "0.0015000000", "9007199254740993.123456789012"} {
		t.Run(value, func(t *testing.T) {
			params, err := expandPriceCreate(PriceResourceModel{UnitAmountDecimal: types.StringValue(value)})
			if err != nil {
				t.Fatal(err)
			}
			if params.UnitAmountDecimal == nil {
				t.Fatal("unit_amount_decimal was omitted")
			}
			values := &form.Values{}
			form.AppendTo(values, params)
			got := values.ToValues().Get("unit_amount_decimal")
			want := params.UnitAmountDecimal.String()
			if got != want {
				t.Fatalf("encoded decimal = %q, want %q", got, want)
			}
			if value == "9007199254740993.123456789012" && got != value {
				t.Fatalf("decimal lost precision: got %q, want %q", got, value)
			}
		})
	}
	if _, err := expandPriceCreate(PriceResourceModel{UnitAmountDecimal: types.StringValue("invalid")}); err == nil {
		t.Fatal("invalid decimal was accepted")
	}
	for _, value := range []types.String{types.StringNull(), types.StringUnknown()} {
		params, err := expandPriceCreate(PriceResourceModel{UnitAmountDecimal: value})
		if err != nil {
			t.Fatal(err)
		}
		if params.UnitAmountDecimal != nil {
			t.Fatal("null or unknown decimal was sent")
		}
	}
}
