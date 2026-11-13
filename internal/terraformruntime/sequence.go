//
// File generated from our OpenAPI spec
//

package bindingruntime

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/url"
	"strings"
	"time"
)

type Operation struct {
	ID, Service, Method, Path string
	Read                      bool
}
type Transport interface {
	Execute(context.Context, Operation, map[string]string, url.Values, json.RawMessage) (json.RawMessage, error)
	HTTPStatus(error) int
}
type Mapping struct {
	Attribute         string
	Constant          json.RawMessage
	Path              []string
	Location          string
	Wrap, OmitUnknown bool
}
type ResponseField struct {
	Attribute                                                        string
	Path                                                             []string
	Required, PreserveMissing, PreserveConfigured, EmptyStringIsNull bool
}
type RequestHook func(context.Context, *Values, *Values, map[string]any) (any, error)
type ResponseHook func(context.Context, *Values, any) error
type ItemsHook func(context.Context, *Values) ([]any, error)
type Step struct {
	Operation          Operation
	Request            []Mapping
	Response           []ResponseField
	Root               []string
	Selector           map[string]string
	Paginate           bool
	PaginationLocation string
	PaginationCursor   string
	RequestHook        RequestHook
	ResponseHook       ResponseHook
	ForEach            ItemsHook
	ForEachQuery       string
	WhenChanged        []string
	WhenMissing        string
	Persist            bool
	PersistAttributes  []string
	ArchivedPaths      [][]string
	MissingStatuses    []int
	WaitForReady       bool
	IgnoreResponse     bool
	OnError            []Step
}

var ErrMissing = errors.New("remote object is absent or archived")

func IsMissing(err error) bool { return errors.Is(err, ErrMissing) }

type Persist func(context.Context, *Values) error
type SkipRequest struct{}

// Run executes only explicit operation bindings. Writes have no retry loop; the
// SDK transport owns the reviewed read retry policy.
func Run(ctx context.Context, transport Transport, steps []Step, values, prior *Values, persist Persist) error {
	if transport == nil && len(steps) > 0 {
		return errors.New("provider is not configured")
	}
	var committed *Values
	if prior != nil {
		committed = NewValues(prior.Object)
	}
	for _, step := range steps {
		if len(step.WhenChanged) > 0 && !Changed(values, prior, step.WhenChanged...) {
			continue
		}
		if step.WhenMissing != "" && values.Known(step.WhenMissing) {
			continue
		}
		items := []any{nil}
		if step.ForEach != nil {
			var err error
			items, err = step.ForEach(ctx, values)
			if err != nil {
				return err
			}
		}
		for _, item := range items {
			if err := runStep(ctx, transport, step, values, prior, item); err != nil {
				if len(step.OnError) > 0 && !IsMissing(err) {
					recovery := Run(ctx, transport, step.OnError, values, prior, nil)
					if IsMissing(recovery) {
						return nil
					}
					if recovery != nil {
						return fmt.Errorf("%s failed and recovery also failed: %w", step.Operation.ID, recovery)
					}
				}
				return fmt.Errorf("%s: %w", step.Operation.ID, err)
			}
			if step.Persist && persist != nil {
				state := values
				if committed != nil && len(step.PersistAttributes) > 0 {
					for _, key := range step.PersistAttributes {
						if err := committed.Set(key, values.Get(key)); err != nil {
							return err
						}
					}
					state = committed
				}
				if err := persist(ctx, state); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func runStep(ctx context.Context, t Transport, s Step, v, prior *Values, item any) error {
	body := map[string]any{}
	paths := map[string]string{}
	query := url.Values{}
	for _, mapping := range s.Request {
		var value any
		if mapping.Constant != nil {
			var err error
			value, err = DecodeJSON(mapping.Constant)
			if err != nil {
				return err
			}
		} else {
			source := v.Get(mapping.Attribute)
			if source == nil {
				return fmt.Errorf("unknown mapping source %s", mapping.Attribute)
			}
			if source.IsNull() || (source.IsUnknown() && mapping.OmitUnknown) {
				continue
			}
			var err error
			value, err = JSONValue(source, false)
			if err != nil {
				return err
			}
		}
		if mapping.Wrap {
			value = []any{value}
		}
		switch mapping.Location {
		case "path", "query":
			if len(mapping.Path) != 1 {
				return errors.New("path/query mapping requires one key")
			}
			text, ok := value.(string)
			if !ok {
				return errors.New("path/query value must be a string")
			}
			if mapping.Location == "path" {
				paths[mapping.Path[0]] = text
			} else {
				query.Set(mapping.Path[0], text)
			}
		case "", "body":
			if err := Put(body, mapping.Path, value); err != nil {
				return err
			}
		default:
			return errors.New("unreviewed request location")
		}
	}
	var payload any = body
	if s.RequestHook != nil {
		var err error
		payload, err = s.RequestHook(ctx, v, prior, body)
		if err != nil {
			return err
		}
	}
	if _, skip := payload.(SkipRequest); skip {
		return nil
	}
	if item != nil {
		if s.ForEachQuery != "" {
			text, ok := item.(string)
			if !ok {
				return errors.New("query iterator requires strings")
			}
			query.Set(s.ForEachQuery, text)
			payload = nil
		} else {
			payload = item
		}
	}
	if s.Operation.Method == "GET" && len(body) == 0 && s.RequestHook == nil && item == nil {
		payload = nil
	}
	ctx, cancel := pollContext(ctx, s, v)
	defer cancel()
	for {
		raw, err := callPages(ctx, t, s, paths, query, payload, v)
		if err == nil {
			if s.IgnoreResponse {
				return nil
			}
			for _, path := range s.ArchivedPaths {
				value, _ := at(raw, path)
				if value != nil && value != "" && value != false {
					return ErrMissing
				}
			}
			if s.ResponseHook != nil {
				if err = s.ResponseHook(ctx, v, raw); err != nil {
					return err
				}
			}
			for _, mapping := range s.Response {
				previous := v.Get(mapping.Attribute)
				if mapping.PreserveConfigured && previous != nil && !previous.IsNull() && !previous.IsUnknown() {
					continue
				}
				value, found := at(raw, mapping.Path)
				if !found && mapping.Required {
					return fmt.Errorf("response omitted %s", strings.Join(mapping.Path, "."))
				}
				if !found && mapping.PreserveMissing {
					continue
				}
				if mapping.EmptyStringIsNull && value == "" {
					value = nil
				}
				if mapping.Required && (value == nil || value == "") {
					return errors.New("response omitted a required identity or field")
				}
				if err = v.SetJSON(ctx, mapping.Attribute, value); err != nil {
					return err
				}
			}
			return nil
		}
		status := t.HTTPStatus(err)
		var network net.Error
		if !s.WaitForReady || (!missingStatus(s, status) && status != 429 && (status < 500 || status >= 600) && !errors.As(err, &network)) {
			return classify(s, status, err)
		}
		timer := time.NewTimer(2 * time.Second)
		select {
		case <-ctx.Done():
			timer.Stop()
			return errors.New("activation completed but Metronome readiness timed out; keep this resource in state and retry after authorization caches expire")
		case <-timer.C:
		}
	}
}
func pollContext(ctx context.Context, s Step, v *Values) (context.Context, context.CancelFunc) {
	if !s.WaitForReady {
		return context.WithCancel(ctx)
	}
	timeout := 7 * time.Minute
	if text := v.String("readiness_timeout"); text != "" {
		if duration, err := time.ParseDuration(text); err == nil {
			timeout = duration
		}
	}
	return context.WithTimeout(ctx, timeout)
}
func missingStatus(s Step, status int) bool {
	statuses := s.MissingStatuses
	if len(statuses) == 0 {
		statuses = []int{404}
	}
	for _, missing := range statuses {
		if status == missing {
			return true
		}
	}
	return false
}
func classify(s Step, status int, err error) error {
	if missingStatus(s, status) {
		return ErrMissing
	}
	return err
}
func callPages(ctx context.Context, t Transport, s Step, paths map[string]string, query url.Values, payload any, v *Values) (any, error) {
	seen := map[string]bool{}
	var combined []any
	for page := 0; page < 10000; page++ {
		var body json.RawMessage
		if payload != nil {
			encoded, err := json.Marshal(payload)
			if err != nil {
				return nil, errors.New("invalid exact JSON request")
			}
			body = encoded
		}
		raw, err := t.Execute(ctx, s.Operation, paths, query, body)
		if err != nil {
			return nil, err
		}
		if s.IgnoreResponse {
			return nil, nil
		}
		value, err := DecodeJSON(raw)
		if err != nil {
			return nil, err
		}
		selected, exists := at(value, s.Root)
		if !exists {
			return nil, errors.New("API response envelope is missing")
		}
		if len(s.Selector) > 0 || s.Paginate {
			items, ok := selected.([]any)
			if !ok {
				return nil, errors.New("API response must contain a list")
			}
			if len(s.Selector) > 0 {
				var match any
				for _, item := range items {
					matches := true
					for field, attribute := range s.Selector {
						x, _ := at(item, strings.Split(field, "."))
						if x != v.String(attribute) {
							matches = false
							break
						}
					}
					if matches {
						if match != nil {
							return nil, errors.New("ambiguous API identity")
						}
						match = item
					}
				}
				if match != nil {
					return match, nil
				}
			} else {
				combined = append(combined, items...)
			}
			cursorKey := s.PaginationCursor
			if cursorKey == "" {
				cursorKey = "next_page"
			}
			next, exists := at(value, []string{cursorKey})
			if exists && next != nil && next != "" {
				cursor, ok := next.(string)
				if !ok {
					return nil, errors.New("invalid pagination cursor")
				}
				if !s.Paginate {
					return nil, errors.New("API returned additional pages for a non-paginated managed lookup")
				}
				if seen[cursor] {
					return nil, errors.New("API repeated a pagination cursor")
				}
				seen[cursor] = true
				if s.PaginationLocation == "body" {
					object, ok := payload.(map[string]any)
					if !ok {
						return nil, errors.New("pagination requires object body")
					}
					object[cursorKey] = cursor
				} else {
					query.Set(cursorKey, cursor)
				}
				continue
			}
			if len(s.Selector) > 0 {
				return nil, ErrMissing
			}
			if combined == nil {
				combined = []any{}
			}
			return combined, nil
		}
		return selected, nil
	}
	return nil, errors.New("API pagination exceeded the safety bound")
}
