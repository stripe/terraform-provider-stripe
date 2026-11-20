//
// File generated from our OpenAPI spec
//

package bindingruntime

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func tftNull(t attr.Type) tftypes.Value {
	return tftypes.NewValue(t.TerraformType(context.Background()), nil)
}
