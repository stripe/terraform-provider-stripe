//
// File generated from our OpenAPI spec
//

// Package providerdata routes each provider instance to its explicit service clients.
package providerdata

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	stripe "github.com/stripe/stripe-go/v87"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
	metronomeclient "github.com/stripe/terraform-provider-stripe/metronome-go"
)

type Status struct{ State, Detail string }

func (s Status) Err(service string) error {
	if s.State == "ready" {
		return nil
	}
	return fmt.Errorf("%s service %s: %s", service, s.State, s.Detail)
}

type Data struct {
	Stripe                        *stripe.Client
	Metronome                     *metronomeclient.Client
	StripeStatus, MetronomeStatus Status
	StripeAccount                 string
}

func (d *Data) HTTPStatus(err error) int { return metronomeclient.HTTPStatus(err) }

func (d *Data) Execute(ctx context.Context, operation bindingruntime.Operation, paths map[string]string, query url.Values, body json.RawMessage) (json.RawMessage, error) {
	if operation.Service != "metronome" {
		return nil, errors.New("unknown Terraform service binding")
	}
	if err := d.MetronomeStatus.Err("Metronome"); err != nil {
		return nil, err
	}
	if d.Metronome == nil {
		return nil, errors.New("Metronome client missing")
	}
	if d.StripeAccount != "" {
		return nil, errors.New("Metronome uses the API key's account and does not support stripe_account; use a separate provider alias with that account's key")
	}
	return d.Metronome.Execute(ctx, metronomeclient.Operation{ID: operation.ID, Method: operation.Method, Path: operation.Path, Read: operation.Read}, paths, query, body)
}
