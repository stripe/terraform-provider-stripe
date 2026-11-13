//
// File generated from our OpenAPI spec
//

package bindingruntime

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type transportFunc func(context.Context, Operation, map[string]string, url.Values, json.RawMessage) (json.RawMessage, error)

func (f transportFunc) Execute(ctx context.Context, op Operation, p map[string]string, q url.Values, b json.RawMessage) (json.RawMessage, error) {
	return f(ctx, op, p, q, b)
}
func (transportFunc) HTTPStatus(err error) int {
	var status statusError
	if errors.As(err, &status) {
		return int(status)
	}
	return 0
}

type statusError int

func (e statusError) Error() string { return "test API failure" }
func testValues(t *testing.T, source string) *Values {
	t.Helper()
	kind := types.ObjectType{AttrTypes: map[string]attr.Type{"id": types.StringType, "name": types.StringType, "aliases": types.ListType{ElemType: types.StringType}, "readiness_timeout": types.StringType}}
	v, err := FromJSON(context.Background(), kind, Literal(source), nil)
	if err != nil {
		t.Fatal(err)
	}
	return NewValues(v.(types.Object))
}

func TestPartialUpdatePersistsOnlyCompletedFields(t *testing.T) {
	ctx := context.Background()
	prior := testValues(t, `{"id":"customer","name":"old","aliases":["old"]}`)
	next := testValues(t, `{"id":"customer","name":"new","aliases":["new"]}`)
	var saved *Values
	calls := 0
	transport := transportFunc(func(_ context.Context, op Operation, _ map[string]string, _ url.Values, _ json.RawMessage) (json.RawMessage, error) {
		calls++
		if op.ID == "aliases" {
			return nil, statusError(500)
		}
		return []byte(`{}`), nil
	})
	steps := []Step{{Operation: Operation{ID: "name"}, IgnoreResponse: true, Persist: true, PersistAttributes: []string{"name"}}, {Operation: Operation{ID: "aliases"}, IgnoreResponse: true}}
	err := Run(ctx, transport, steps, next, prior, func(_ context.Context, v *Values) error { saved = NewValues(v.Object); return nil })
	if err == nil || calls != 2 || saved == nil {
		t.Fatalf("expected partial failure with saved state: %v, calls=%d", err, calls)
	}
	if saved.String("name") != "new" || !saved.Get("aliases").Equal(prior.Get("aliases")) {
		t.Fatal("unperformed alias update was recorded")
	}
	if prior.String("name") != "old" {
		t.Fatal("prior state was mutated")
	}
}

func TestCreateIdentityIsPersistedBeforeReadinessFailure(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Millisecond)
	defer cancel()
	v := testValues(t, `{"name":"bridge","readiness_timeout":"7m"}`)
	calls := 0
	saved := ""
	transport := transportFunc(func(_ context.Context, op Operation, _ map[string]string, _ url.Values, _ json.RawMessage) (json.RawMessage, error) {
		calls++
		if op.ID == "activate" {
			return []byte(`{"id":"acct_test/test"}`), nil
		}
		return nil, statusError(403)
	})
	steps := []Step{{Operation: Operation{ID: "activate"}, Response: []ResponseField{{Attribute: "id", Path: []string{"id"}, Required: true}}, Persist: true}, {Operation: Operation{ID: "probe", Read: true}, WaitForReady: true, MissingStatuses: []int{401, 403, 404}}}
	err := Run(ctx, transport, steps, v, nil, func(_ context.Context, state *Values) error { saved = state.String("id"); return nil })
	if err == nil || saved != "acct_test/test" || calls != 2 {
		t.Fatalf("activation state was lost on canceled readiness: saved=%q calls=%d err=%v", saved, calls, err)
	}
}

func TestReadErrorClassificationPreservesNonMissingFailures(t *testing.T) {
	for _, status := range []int{200, 400, 401, 403, 404, 409, 429, 500, 503} {
		t.Run(string(rune(status)), func(t *testing.T) {
			transport := transportFunc(func(context.Context, Operation, map[string]string, url.Values, json.RawMessage) (json.RawMessage, error) {
				if status == 200 {
					return []byte(`invalid`), nil
				}
				return nil, statusError(status)
			})
			err := Run(context.Background(), transport, []Step{{Operation: Operation{Read: true}, MissingStatuses: []int{401, 403, 404}}}, testValues(t, `{}`), nil, nil)
			if err == nil || IsMissing(err) != (status == 401 || status == 403 || status == 404) {
				t.Fatalf("status %d had unsafe classification: %v", status, err)
			}
		})
	}
}

func TestExplicitPaginationAndRepeatedCursorGuard(t *testing.T) {
	for _, bodyPaging := range []bool{false, true} {
		t.Run(map[bool]string{true: "body", false: "query"}[bodyPaging], func(t *testing.T) {
			calls := 0
			total := 0
			transport := transportFunc(func(_ context.Context, _ Operation, _ map[string]string, q url.Values, b json.RawMessage) (json.RawMessage, error) {
				calls++
				if calls == 2 {
					cursor := q.Get("next_page")
					if bodyPaging {
						value, _ := DecodeJSON(b)
						cursor = value.(map[string]any)["cursor"].(string)
					}
					if cursor != "next" {
						t.Fatalf("cursor not forwarded: query=%v body=%s", q, b)
					}
				}
				if calls == 1 {
					if bodyPaging {
						return []byte(`{"data":[{"id":"a"}],"cursor":"next"}`), nil
					}
					return []byte(`{"data":[{"id":"a"}],"next_page":"next"}`), nil
				}
				return []byte(`{"data":[{"id":"b"}]}`), nil
			})
			step := Step{Operation: Operation{Method: "POST", Read: true}, Root: []string{"data"}, Paginate: true, ResponseHook: func(_ context.Context, _ *Values, v any) error { total = len(v.([]any)); return nil }}
			if bodyPaging {
				step.PaginationLocation = "body"
				step.PaginationCursor = "cursor"
			}
			if err := Run(context.Background(), transport, []Step{step}, testValues(t, `{}`), nil, nil); err != nil {
				t.Fatal(err)
			}
			if calls != 2 || total != 2 {
				t.Fatalf("pagination incomplete: calls=%d rows=%d", calls, total)
			}
		})
	}
	calls := 0
	transport := transportFunc(func(context.Context, Operation, map[string]string, url.Values, json.RawMessage) (json.RawMessage, error) {
		calls++
		return []byte(`{"data":[],"next_page":"repeat"}`), nil
	})
	err := Run(context.Background(), transport, []Step{{Operation: Operation{Read: true}, Root: []string{"data"}, Paginate: true}}, testValues(t, `{}`), nil, nil)
	if err == nil || !strings.Contains(err.Error(), "repeated") || calls != 2 {
		t.Fatalf("repeated cursor was not bounded: %v, %d calls", err, calls)
	}
}

func TestMalformedSuccessIsNotRetriedDuringReadiness(t *testing.T) {
	calls := 0
	transport := transportFunc(func(context.Context, Operation, map[string]string, url.Values, json.RawMessage) (json.RawMessage, error) {
		calls++
		return []byte(`not JSON`), nil
	})
	err := Run(context.Background(), transport, []Step{{Operation: Operation{Read: true}, WaitForReady: true, MissingStatuses: []int{401, 403, 404}}}, testValues(t, `{}`), nil, nil)
	if err == nil || calls != 1 {
		t.Fatalf("malformed response retried: %v %d", err, calls)
	}
}
