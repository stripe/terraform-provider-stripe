// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package cases

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
	stripe "github.com/stripe/stripe-go/v86"

	"github.com/stripe/terraform-provider-stripe/internal/acctest/runner"
)

func runBaseManagedCase(
	t *testing.T,
	name string,
	surface string,
	resourceAddress string,
	createTemplate string,
	updateTemplate string,
	importStable bool,
	verifyCreate runner.StateVerifier,
	verifyUpdate runner.StateVerifier,
	verifyDestroy runner.StateVerifier,
) {
	t.Helper()
	runner.RunManagedCase(t, runner.ManagedCase{
		Definition:      runner.CaseDefinition{Name: name, Surface: surface, Group: "base", Kind: "resource"},
		ResourceAddress: resourceAddress,
		CreateTemplate:  createTemplate,
		UpdateTemplate:  updateTemplate,
		ImportStable:    importStable,
		VerifyCreate:    verifyCreate,
		VerifyUpdate:    verifyUpdate,
		VerifyDestroy:   verifyDestroy,
	})
}

func runBaseManagedCaseWithImportIgnore(
	t *testing.T,
	name string,
	surface string,
	resourceAddress string,
	createTemplate string,
	updateTemplate string,
	importStable bool,
	importStateVerifyIgnore []string,
	verifyCreate runner.StateVerifier,
	verifyUpdate runner.StateVerifier,
	verifyDestroy runner.StateVerifier,
) {
	t.Helper()
	runner.RunManagedCase(t, runner.ManagedCase{
		Definition:              runner.CaseDefinition{Name: name, Surface: surface, Group: "base", Kind: "resource"},
		ResourceAddress:         resourceAddress,
		CreateTemplate:          createTemplate,
		UpdateTemplate:          updateTemplate,
		ImportStable:            importStable,
		ImportStateVerifyIgnore: importStateVerifyIgnore,
		VerifyCreate:            verifyCreate,
		VerifyUpdate:            verifyUpdate,
		VerifyDestroy:           verifyDestroy,
	})
}

func expectRemoteStateString(state *terraform.State, address string, attribute string, actual string) error {
	expected, err := runner.ResourceAttribute(state, address, attribute)
	if err != nil {
		return err
	}
	if expected != actual {
		return fmt.Errorf("remote %s.%s mismatch: expected %q, got %q", address, attribute, expected, actual)
	}
	return nil
}

func expectFloat(name string, actual float64, expected float64) error {
	if math.Abs(actual-expected) > 0.000001 {
		return fmt.Errorf("remote %s mismatch: expected %f, got %f", name, expected, actual)
	}
	return nil
}

func expectRemoteStringList(name string, actual []string, expected []string) error {
	if !reflect.DeepEqual(actual, expected) {
		return fmt.Errorf("remote %s mismatch: expected %v, got %v", name, expected, actual)
	}
	return nil
}

func expectRemoteMissing(address string, id string, err error) error {
	if err == nil {
		return fmt.Errorf("expected %s (%s) to be missing remotely", address, id)
	}
	var stripeErr *stripe.Error
	if errors.As(err, &stripeErr) && stripeErr.HTTPStatusCode == http.StatusNotFound {
		return nil
	}
	return fmt.Errorf("expected %s (%s) missing error, got: %w", address, id, err)
}
