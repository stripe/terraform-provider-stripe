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
	stripe "github.com/stripe/stripe-go/v86"
	metronomeclient "github.com/stripe/terraform-provider-stripe/internal/metronomeclient"
	bindingruntime "github.com/stripe/terraform-provider-stripe/internal/terraformruntime"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
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
	StripeKey, StripeAccount      string
	ActivationBackend             stripe.RawRequestBackend
}

func NewActivationBackend() stripe.RawRequestBackend {
	client := &http.Client{Timeout: 30 * time.Second, Transport: stripeOriginTransport{base: http.DefaultTransport}, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
	backend := stripe.GetBackendWithConfig(stripe.APIBackend, &stripe.BackendConfig{URL: stripe.String("https://api.stripe.com"), HTTPClient: client, MaxNetworkRetries: stripe.Int64(0), LeveledLogger: &stripe.LeveledLogger{Level: stripe.LevelNull}})
	return backend.(stripe.RawRequestBackend)
}

type stripeOriginTransport struct{ base http.RoundTripper }

func (t stripeOriginTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.Scheme != "https" || req.URL.Host != "api.stripe.com" || req.URL.User != nil {
		return nil, errors.New("activation must use the Stripe API HTTPS origin")
	}
	response, err := t.base.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	response.Body = &stripeResponseBody{Reader: io.LimitReader(response.Body, (32<<20)+1), Closer: response.Body}
	return response, nil
}

type stripeResponseBody struct {
	io.Reader
	io.Closer
}

func (d *Data) HTTPStatus(err error) int {
	var stripeErr *stripe.Error
	if errors.As(err, &stripeErr) {
		return stripeErr.HTTPStatusCode
	}
	return metronomeclient.HTTPStatus(err)
}
func (d *Data) Execute(ctx context.Context, operation bindingruntime.Operation, paths map[string]string, query url.Values, body json.RawMessage) (json.RawMessage, error) {
	switch operation.Service {
	case "metronome":
		if err := d.MetronomeStatus.Err("Metronome"); err != nil {
			return nil, err
		}
		if d.Metronome == nil {
			return nil, errors.New("Metronome client missing")
		}
		if d.StripeAccount != "" {
			if _, err := d.accountIdentity(ctx); err != nil {
				return nil, err
			}
		}
		return d.Metronome.Execute(ctx, metronomeclient.Operation{ID: operation.ID, Method: operation.Method, Path: operation.Path, Read: operation.Read}, paths, query, body)
	case "stripe":
		if err := d.StripeStatus.Err("Stripe"); err != nil {
			return nil, err
		}
		if len(paths) != 0 || len(query) != 0 {
			return nil, errors.New("unreviewed Stripe activation parameters")
		}
		if operation.ID == "stripeAccountIdentity" && operation.Method == "GET" && operation.Path == "/v1/account" && operation.Read {
			return d.accountIdentity(ctx)
		}
		if operation.ID == "stripeMetronomeActivate" && operation.Method == "POST" && operation.Path == "/v1/experimental-met-activate" && !operation.Read {
			if len(body) > 0 && string(body) != "{}" {
				return nil, errors.New("activation takes an empty payload")
			}
			if _, err := d.accountIdentity(ctx); err != nil {
				return nil, err
			}
			return d.stripeRequest(ctx, "POST", operation.Path)
		}
		return nil, errors.New("Stripe operation has no reviewed activation binding")
	default:
		return nil, errors.New("unknown Terraform service binding")
	}
}
func (d *Data) stripeRequest(ctx context.Context, method, path string) (json.RawMessage, error) {
	if d.ActivationBackend == nil {
		return nil, errors.New("Stripe activation backend is unavailable")
	}
	response, err := d.ActivationBackend.RawRequest(method, path, d.StripeKey, "", &stripe.RawParams{Params: stripe.Params{Context: ctx}})
	if err != nil {
		return nil, &stripeRequestError{cause: err}
	}
	if response == nil || len(response.RawJSON) > 32<<20 {
		return nil, errors.New("invalid Stripe activation response")
	}
	return json.RawMessage(response.RawJSON), nil
}

type stripeRequestError struct{ cause error }

func (e *stripeRequestError) Error() string {
	var api *stripe.Error
	if errors.As(e.cause, &api) {
		return fmt.Sprintf("Stripe API returned HTTP %d", api.HTTPStatusCode)
	}
	return "Stripe API request failed"
}
func (e *stripeRequestError) Unwrap() error { return e.cause }
func (d *Data) accountIdentity(ctx context.Context) (json.RawMessage, error) {
	mode := ""
	switch {
	case strings.HasPrefix(d.StripeKey, "sk_test_"), strings.HasPrefix(d.StripeKey, "rkcs_test_"):
		mode = "test"
	case strings.HasPrefix(d.StripeKey, "sk_live_"):
		mode = "live"
	default:
		return nil, errors.New("Metronome activation requires a supported Stripe account key: sk_test, sk_live, or rkcs_test")
	}
	raw, err := d.stripeRequest(ctx, "GET", "/v1/account")
	if err != nil {
		return nil, err
	}
	var account struct {
		ID string `json:"id"`
	}
	if err = json.Unmarshal(raw, &account); err != nil || !strings.HasPrefix(account.ID, "acct_") {
		return nil, errors.New("Stripe did not return an authenticated account identity")
	}
	if d.StripeAccount != "" && d.StripeAccount != account.ID {
		return nil, errors.New("Metronome uses the API key's account; Stripe Connect context does not select another Metronome account")
	}
	encoded, err := json.Marshal(map[string]string{"id": account.ID, "metronome_mode": mode})
	return json.RawMessage(encoded), err
}
