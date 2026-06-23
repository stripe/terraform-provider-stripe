// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package checks

import (
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func ResourceIDIsSet(address string) statecheck.StateCheck {
	return statecheck.ExpectKnownValue(address, tfjsonpath.New("id"), knownvalue.NotNull())
}

func OutputIsNotNull(name string) statecheck.StateCheck {
	return statecheck.ExpectKnownOutputValue(name, knownvalue.NotNull())
}

func OutputStringExact(name string, value string) statecheck.StateCheck {
	return statecheck.ExpectKnownOutputValue(name, knownvalue.StringExact(value))
}

func OutputIsNull(name string) statecheck.StateCheck {
	return statecheck.ExpectKnownOutputValue(name, knownvalue.Null())
}
