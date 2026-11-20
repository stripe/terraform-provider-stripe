//
// File generated from our OpenAPI spec
//

// Package metronome contains the SDK bundled with the Terraform provider.
package metronome

import (
	"bytes"
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
)

const DefaultBaseURL = "https://api.metronome.com"
const responseLimit = 32 << 20

// Number and JSON preserve exact numbers and untagged or extensible API values.
type Number = json.Number
type JSON = json.RawMessage

type Params struct {
	Context context.Context `json:"-"`
	raw     *rawRequest
}

func (p *Params) GetParams() *Params { return p }

type ParamsContainer interface{ GetParams() *Params }
type APIResource struct {
	RawJSON json.RawMessage `json:"-"`
}

func (r *APIResource) setRawJSON(body json.RawMessage) { r.RawJSON = body }

type responseSetter interface{ setRawJSON(json.RawMessage) }
type Backend interface {
	Call(method, path, key string, params ParamsContainer, result responseSetter) error
}
type backend struct {
	http   *http.Client
	origin string
}
type operationInfo struct {
	Operation
	Query     []string
	HasBody   bool
	BodyField string
}
type rawRequest struct {
	body   json.RawMessage
	query  url.Values
	result json.RawMessage
}

var errUnknownOperation = errors.New("Metronome operation has no generated SDK binding")

func String(value string) *string { return &value }
func StringValue(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

// FormatURLPath encodes each argument as one path segment. The backend rejects
// dot segments and decoded separators before sending the request.
func FormatURLPath(format string, values ...string) string {
	escaped := make([]interface{}, len(values))
	for i, value := range values {
		escaped[i] = url.PathEscape(value)
	}
	return fmt.Sprintf(format, escaped...)
}

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

// New makes no network requests and uses only explicit provider configuration.
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
	return newClient(&backend{http: httpClient, origin: strings.TrimRight(origin.String(), "/")}, apiKey), nil
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

// Execute dispatches to a generated service method, preserving the provider's
// exact JSON. Page traversal and lifecycle decisions belong to the caller.
func (c *Client) Execute(ctx context.Context, operation Operation, pathParams map[string]string, query url.Values, body json.RawMessage) (json.RawMessage, error) {
	var found bool
	for _, known := range sdkOperations {
		if known.Operation == operation {
			found = true
			break
		}
	}
	if !found {
		return nil, errUnknownOperation
	}
	path := operation.Path
	for name, value := range pathParams {
		placeholder := "{" + name + "}"
		if !validSegment(value) || strings.Count(path, placeholder) != 1 {
			return nil, errors.New("invalid Metronome path parameter")
		}
		path = strings.Replace(path, placeholder, url.PathEscape(value), 1)
	}
	if strings.ContainsAny(path, "{}") {
		return nil, errors.New("missing Metronome path parameter")
	}
	if body != nil && !json.Valid(body) {
		return nil, errors.New("invalid Metronome request JSON")
	}
	raw := &rawRequest{body: body, query: query}
	if err := c.invoke(ctx, operation, pathParams, raw); err != nil {
		return nil, err
	}
	return raw.result, nil
}

func validSegment(value string) bool {
	return value != "" && value != "." && value != ".." && !strings.ContainsAny(value, "/?#")
}

func requestOperation(method, path string) (operationInfo, error) {
	for _, known := range sdkOperations {
		if known.Method == method && known.Path == path {
			return known, nil
		}
	}
	actual := strings.Split(path, "/")
	for _, known := range sdkOperations {
		expected := strings.Split(known.Path, "/")
		if known.Method != method || len(expected) != len(actual) {
			continue
		}
		match := true
		for i, segment := range expected {
			if strings.HasPrefix(segment, "{") && strings.HasSuffix(segment, "}") {
				decoded, err := url.PathUnescape(actual[i])
				if err != nil || !validSegment(decoded) {
					match = false
					break
				}
			} else if segment != actual[i] {
				match = false
				break
			}
		}
		if match {
			return known, nil
		}
	}
	return operationInfo{}, errUnknownOperation
}

// Call implements the backend contract used by the shared Go service emitter.
// Metronome encodes request bodies as JSON for every verb, including /v1 paths.
func (b *backend) Call(method, path, key string, params ParamsContainer, result responseSetter) error {
	operation, err := requestOperation(method, path)
	if err != nil {
		return err
	}
	common := params.GetParams()
	ctx := common.Context
	if ctx == nil {
		ctx = context.Background()
	}
	var body []byte
	query := url.Values{}
	if common.raw != nil {
		body, query = common.raw.body, common.raw.query
	} else {
		body, err = json.Marshal(params)
		if err != nil {
			return errors.New("invalid Metronome request JSON")
		}
		var object map[string]json.RawMessage
		if err = json.Unmarshal(body, &object); err != nil {
			return errors.New("invalid Metronome request object")
		}
		for _, name := range operation.Query {
			value, present := object[name]
			if !present {
				continue
			}
			values := []json.RawMessage{value}
			if len(value) > 0 && value[0] == '[' {
				if err := json.Unmarshal(value, &values); err != nil {
					return errors.New("invalid Metronome query array")
				}
			}
			for _, value := range values {
				var text string
				if err := json.Unmarshal(value, &text); err != nil {
					text = string(value)
				}
				query.Add(name, text)
			}
			delete(object, name)
		}
		if !operation.HasBody {
			body = nil
		} else if operation.BodyField != "" {
			body = object[operation.BodyField]
		} else if body, err = json.Marshal(object); err != nil {
			return errors.New("invalid Metronome request object")
		}
	}
	if encoded := query.Encode(); encoded != "" {
		path += "?" + encoded
	}
	for attempt := 0; ; attempt++ {
		req, err := http.NewRequestWithContext(ctx, method, b.origin+path, bytes.NewReader(body))
		if err != nil {
			return &requestError{cause: err}
		}
		req.Header.Set("Authorization", "Bearer "+key)
		if body != nil {
			req.Header.Set("Content-Type", "application/json")
		}
		resp, err := b.http.Do(req)
		if err == nil {
			raw, readErr := io.ReadAll(resp.Body)
			resp.Body.Close()
			if readErr != nil {
				return &requestError{cause: readErr}
			}
			result.setRawJSON(raw)
			if common.raw != nil {
				common.raw.result = raw
				return nil
			}
			if len(raw) == 0 {
				return nil
			}
			decoder := json.NewDecoder(bytes.NewReader(raw))
			decoder.UseNumber()
			if err := decoder.Decode(result); err != nil {
				return errors.New("invalid Metronome response JSON")
			}
			return nil
		}
		if !operation.Read || attempt == 2 || ctx.Err() != nil || !retryableRead(err) {
			return &requestError{cause: err}
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
			return ctx.Err()
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
