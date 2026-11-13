//
// File generated from our OpenAPI spec
//

package bindingruntime

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"net/url"
	"regexp"
	"strings"
)

func Literal(text string) any {
	v, err := DecodeJSON([]byte(text))
	if err != nil {
		panic("invalid generated JSON literal")
	}
	return v
}
func EmptyValues(t types.ObjectType) *Values {
	attrs := map[string]attr.Value{}
	for key, kind := range t.AttrTypes {
		attrs[key] = Null(kind)
	}
	return NewValues(types.ObjectValueMust(t.AttrTypes, attrs))
}
func ComposeIdentity(ctx context.Context, v *Values, attribute string, fields []string, delimiter string) error {
	parts := make([]string, len(fields))
	for i, key := range fields {
		parts[i] = v.String(key)
		if parts[i] == "" {
			return errors.New("identity component is missing")
		}
	}
	return v.SetJSON(ctx, attribute, strings.Join(parts, delimiter))
}
func ImportIdentity(ctx context.Context, v *Values, id string, fields []string, delimiter string) error {
	parts := strings.SplitN(id, delimiter, len(fields))
	if len(parts) != len(fields) {
		return errors.New("import ID does not match the documented composite identity")
	}
	for i, field := range fields {
		if strings.TrimSpace(parts[i]) == "" {
			return errors.New("import identity must not be empty")
		}
		if err := v.SetJSON(ctx, field, parts[i]); err != nil {
			return err
		}
	}
	return nil
}

var parameter = regexp.MustCompile(`\{([^{}]+)\}`)

func InvokeOperation(ctx context.Context, t Transport, op Operation, v *Values) (json.RawMessage, string, error) {
	if t == nil {
		return nil, "", errors.New("provider is not configured")
	}
	paths := map[string]string{}
	query := url.Values{}
	known := map[string]bool{}
	for _, match := range parameter.FindAllStringSubmatch(op.Path, -1) {
		known[match[1]] = true
	}
	for _, field := range []string{"path_params", "query_params"} {
		value, err := JSONValue(v.Get(field), false)
		if err != nil {
			return nil, "", err
		}
		if value == nil {
			continue
		}
		object, ok := value.(map[string]any)
		if !ok {
			return nil, "", errors.New("operation parameters must be string maps")
		}
		for key, x := range object {
			text, ok := x.(string)
			if !ok {
				return nil, "", errors.New("operation parameter must be a string")
			}
			if field == "path_params" {
				if !known[key] {
					return nil, "", errors.New("unknown API path parameter")
				}
				paths[key] = text
			} else {
				query.Set(key, text)
			}
		}
	}
	renderedPath := op.Path
	for key := range known {
		value := paths[key]
		if value == "" || value == "." || value == ".." || strings.ContainsAny(value, "/?#") {
			return nil, "", errors.New("path parameter must be a nonempty path segment")
		}
		renderedPath = strings.ReplaceAll(renderedPath, "{"+key+"}", url.PathEscape(value))
	}
	var body json.RawMessage
	request := v.Get("request")
	raw := v.Get("request_json")
	if request != nil && !request.IsNull() {
		if raw != nil && !raw.IsNull() && v.String("request_json") != "" {
			return nil, "", errors.New("use request or request_json, not both")
		}
		value, err := JSONValue(request, false)
		if err != nil {
			return nil, "", err
		}
		body, err = json.Marshal(value)
		if err != nil {
			return nil, "", errors.New("invalid JSON request")
		}
	}
	if raw != nil && !raw.IsNull() {
		if raw.IsUnknown() {
			return nil, "", errors.New("request_json must be known")
		}
		if v.String("request_json") != "" {
			body = json.RawMessage(v.String("request_json"))
			if !json.Valid(body) {
				return nil, "", errors.New("request_json is not valid JSON")
			}
		}
	}
	result, err := t.Execute(ctx, op, paths, query, body)
	if err != nil {
		return nil, "", err
	}
	canonical := []byte("null")
	if body != nil {
		buffer := new(bytes.Buffer)
		if err := json.Compact(buffer, body); err != nil {
			return nil, "", err
		}
		canonical = buffer.Bytes()
	}
	sum := sha256.Sum256([]byte(op.ID + "\x00" + renderedPath + "\x00" + query.Encode() + "\x00" + string(canonical)))
	return result, hex.EncodeToString(sum[:]), nil
}
func ReadOperation(ctx context.Context, t Transport, op Operation, v *Values, binary bool) error {
	if !op.Read {
		return errors.New("operation is not a reviewed read")
	}
	result, id, err := InvokeOperation(ctx, t, op, v)
	if err != nil {
		return err
	}
	if binary {
		if err = v.SetJSON(ctx, "response_base64", base64.StdEncoding.EncodeToString(result)); err != nil {
			return err
		}
		if err = v.SetJSON(ctx, "response_json", nil); err != nil {
			return err
		}
	} else {
		if !json.Valid(result) {
			return errors.New("API returned invalid JSON")
		}
		if err = v.SetJSON(ctx, "response_json", string(result)); err != nil {
			return err
		}
		if err = v.SetJSON(ctx, "response_base64", nil); err != nil {
			return err
		}
	}
	return v.SetJSON(ctx, "id", id)
}
