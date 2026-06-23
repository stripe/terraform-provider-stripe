// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package runner

import (
	"testing"

	stripe "github.com/stripe/stripe-go/v86"

	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

type StateVerifier func(TestEnv, *stripe.Client, *terraform.State) error
type StateCheck = statecheck.StateCheck
type ActionConfigPreparer func(*testing.T, TestEnv, *stripe.Client) map[string]string
type ActionRenderEnv func(TestEnv) TestEnv

type CaseDefinition struct {
	Name        string
	Surface     string
	Group       string
	Kind        string
	RequiredEnv []string
}

type ManagedCase struct {
	Definition              CaseDefinition
	ResourceAddress         string
	CreateTemplate          string
	UpdateTemplate          string
	ImportStable            bool
	ImportStateVerifyIgnore []string
	VerifyCreate            StateVerifier
	VerifyUpdate            StateVerifier
	VerifyDestroy           StateVerifier
}

type EphemeralCase struct {
	Definition           CaseDefinition
	ConfigTemplate       string
	ConfigStateChecks    []statecheck.StateCheck
	VerifyPersistedState StateVerifier
}

type ActionCase struct {
	Definition                 CaseDefinition
	PrerequisiteConfigTemplate string
	ConfigTemplate             string
	ConfigStateChecks          []statecheck.StateCheck
	ConfigPreparer             ActionConfigPreparer
	RenderEnv                  ActionRenderEnv
	VerifyInvoke               StateVerifier
}
