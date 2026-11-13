//
// File generated from our OpenAPI spec
//

// Package metronomeclient supplies transport compatibility for the official SDK.
package metronomeclient

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	metronome "github.com/Metronome-Industries/metronome-go/v3"
	"github.com/Metronome-Industries/metronome-go/v3/option"
)

const DefaultBaseURL = "https://api.metronome.com"
const responseLimit = 32 << 20

// Client keeps the SDK and its explicit transport options isolated per provider.
type Client struct{ SDK *metronome.Client }

// Operation's Read flag comes from the reviewed read policy, never HTTP method.
type Operation struct {
	ID, Method, Path string
	Read             bool
}

// APIError retains raw failures for explicit inspection, while diagnostics exclude
// API bodies because those bodies can echo credentials or customer fields.
type APIError struct {
	Status     int
	Body       []byte
	retryAfter time.Duration
}

func (e *APIError) Error() string { return fmt.Sprintf("Metronome API returned HTTP %d", e.Status) }
func HTTPStatus(err error) int {
	var api *APIError
	if errors.As(err, &api) {
		return api.Status
	}
	var sdk *metronome.Error
	if errors.As(err, &sdk) {
		return sdk.StatusCode
	}
	return 0
}
func IsNotFound(err error) bool { return HTTPStatus(err) == http.StatusNotFound }

type requestError struct{ cause error }

func (e *requestError) Error() string {
	if status := HTTPStatus(e.cause); status != 0 {
		return fmt.Sprintf("Metronome API returned HTTP %d", status)
	}
	if errors.Is(e.cause, context.Canceled) {
		return "Metronome API request canceled"
	}
	if errors.Is(e.cause, context.DeadlineExceeded) {
		return "Metronome API request timed out"
	}
	return "Metronome API request failed"
}
func (e *requestError) Unwrap() error { return e.cause }

// New makes no network requests and deliberately avoids SDK DefaultClientOptions,
// which read ambient bearer tokens, webhook secrets, origins and custom headers.
func New(apiKey, baseURL, caFile, version string) (*Client, error) {
	if strings.TrimSpace(apiKey) == "" {
		return nil, errors.New("Stripe API key is required for Metronome")
	}
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	origin, err := url.Parse(baseURL)
	if err != nil || origin.Scheme != "https" || origin.Hostname() == "" || origin.User != nil || origin.RawQuery != "" || origin.ForceQuery || origin.Fragment != "" || (origin.Path != "" && origin.Path != "/") || origin.Opaque != "" {
		return nil, errors.New("Metronome base URL must be an HTTPS origin")
	}
	base, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		return nil, errors.New("default HTTP transport does not support Metronome TLS configuration")
	}
	transport := base.Clone()
	transport.TLSClientConfig = &tls.Config{MinVersion: tls.VersionTLS12}
	if caFile != "" {
		pemData, err := os.ReadFile(caFile)
		if err != nil {
			return nil, fmt.Errorf("reading Metronome CA certificate: %w", err)
		}
		roots, err := x509.SystemCertPool()
		if err != nil {
			return nil, fmt.Errorf("loading system certificate pool: %w", err)
		}
		if !roots.AppendCertsFromPEM(pemData) {
			return nil, errors.New("Metronome CA file contains no valid certificates")
		}
		transport.TLSClientConfig.RootCAs = roots
	}
	httpClient := &http.Client{Timeout: 30 * time.Second, Transport: &originTransport{base: transport, origin: origin, apiKey: apiKey, version: version}, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
	opts := []option.RequestOption{option.WithBaseURL(origin.String()), option.WithBearerToken(apiKey), option.WithHTTPClient(httpClient), option.WithMaxRetries(0)}
	sdk := &metronome.Client{Options: opts, V1: metronome.NewV1Service(opts...), V2: metronome.NewV2Service(opts...)}
	return &Client{SDK: sdk}, nil
}

type originTransport struct {
	base            http.RoundTripper
	origin          *url.URL
	apiKey, version string
}

func (t *originTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.Scheme != t.origin.Scheme || req.URL.Host != t.origin.Host || req.URL.User != nil {
		return nil, errors.New("Metronome request must use configured HTTPS origin")
	}
	req = req.Clone(req.Context())
	req.Header.Set("Authorization", "Bearer "+t.apiKey)
	req.Header.Set("User-Agent", "terraform-provider-stripe/"+t.version)
	resp, err := t.base.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		defer resp.Body.Close()
		body, readErr := io.ReadAll(io.LimitReader(resp.Body, 64<<10))
		if readErr != nil {
			return nil, readErr
		}
		return nil, &APIError{Status: resp.StatusCode, Body: body, retryAfter: retryAfter(resp.Header.Get("Retry-After"))}
	}
	resp.Body = &boundedBody{ReadCloser: resp.Body, remaining: responseLimit}
	return resp, nil
}

type boundedBody struct {
	io.ReadCloser
	remaining int
}

func (b *boundedBody) Read(p []byte) (int, error) {
	if b.remaining < 0 {
		return 0, errors.New("Metronome API response exceeds 32 MiB")
	}
	if len(p) > b.remaining+1 {
		p = p[:b.remaining+1]
	}
	n, err := b.ReadCloser.Read(p)
	b.remaining -= n
	if b.remaining < 0 {
		return 0, errors.New("Metronome API response exceeds 32 MiB")
	}
	return n, err
}

// Execute returns one response page unchanged. Pagination, archived-state and
// lifecycle interpretation remain with generated callers; no SDK auto-pager runs.
func (c *Client) Execute(ctx context.Context, operation Operation, pathParams map[string]string, query url.Values, body json.RawMessage) (json.RawMessage, error) {
	if operation.Method == "" {
		return nil, errors.New("Metronome operation method is required")
	}
	path := operation.Path
	for key, value := range pathParams {
		if value == "" || value == "." || value == ".." || strings.ContainsAny(value, "/?#") {
			return nil, fmt.Errorf("invalid %s path segment", key)
		}
		placeholder := "{" + key + "}"
		if strings.Count(path, placeholder) != 1 {
			return nil, fmt.Errorf("missing or repeated %s in endpoint", placeholder)
		}
		path = strings.Replace(path, placeholder, url.PathEscape(value), 1)
	}
	parsed, err := url.Parse(path)
	if err != nil || !strings.HasPrefix(path, "/") || strings.HasPrefix(path, "//") || parsed.IsAbs() || parsed.Host != "" || parsed.RawQuery != "" || parsed.Fragment != "" || strings.ContainsAny(path, "{}") {
		return nil, errors.New("invalid Metronome operation path")
	}
	if body != nil && !json.Valid(body) {
		return nil, errors.New("invalid Metronome request JSON")
	}
	for attempt := 0; ; attempt++ {
		var result []byte
		opts := []option.RequestOption{option.WithResponseBodyInto(&result), option.WithMaxRetries(0)}
		if body != nil {
			opts = append(opts, option.WithRequestBody("application/json", []byte(body)))
		} else {
			opts = append(opts, option.WithRequestBody("", []byte(nil)))
		}
		for key, values := range query {
			opts = append(opts, option.WithQueryDel(key))
			for _, value := range values {
				opts = append(opts, option.WithQueryAdd(key, value))
			}
		}
		err = c.invoke(ctx, operation, path, pathParams, opts)
		if err == nil {
			return json.RawMessage(result), nil
		}
		if !operation.Read || attempt == 2 || ctx.Err() != nil || !retryableRead(err) {
			return nil, &requestError{cause: err}
		}
		delay := time.Duration(1<<attempt) * 200 * time.Millisecond
		var api *APIError
		if errors.As(err, &api) && api.retryAfter > delay {
			delay = api.retryAfter
		}
		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			timer.Stop()
			return nil, ctx.Err()
		case <-timer.C:
		}
	}
}
func retryableRead(err error) bool {
	if status := HTTPStatus(err); status != 0 {
		switch status {
		case 429, 500, 502, 503, 504:
			return true
		}
		return false
	}
	var network net.Error
	return errors.As(err, &network) && network.Timeout()
}
func retryAfter(value string) time.Duration {
	var delay time.Duration
	if seconds, err := strconv.ParseInt(value, 10, 64); err == nil {
		if seconds > 5 {
			return 5 * time.Second
		}
		if seconds > 0 {
			delay = time.Duration(seconds) * time.Second
		}
	} else if at, err := http.ParseTime(value); err == nil {
		delay = time.Until(at)
	}
	if delay < 0 {
		return 0
	}
	if delay > 5*time.Second {
		return 5 * time.Second
	}
	return delay
}
