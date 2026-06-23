// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package runner

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func RootModuleState(state *terraform.State) (*terraform.ModuleState, error) {
	if state == nil {
		return nil, fmt.Errorf("terraform state is nil")
	}
	if len(state.Modules) == 0 || state.Modules[0] == nil {
		return nil, fmt.Errorf("terraform state has no root module")
	}

	return state.Modules[0], nil
}

func ResourcePrimaryID(state *terraform.State, address string) (string, error) {
	resourceState, err := ResourceState(state, address)
	if err != nil {
		return "", err
	}
	if resourceState.Primary == nil || resourceState.Primary.ID == "" {
		return "", fmt.Errorf("resource %s has no primary id in state", address)
	}

	return resourceState.Primary.ID, nil
}

func ResourceAttribute(state *terraform.State, address string, attribute string) (string, error) {
	resourceState, err := ResourceState(state, address)
	if err != nil {
		return "", err
	}
	if resourceState.Primary == nil {
		return "", fmt.Errorf("resource %s has no primary state", address)
	}

	value, ok := resourceState.Primary.Attributes[attribute]
	if !ok {
		return "", fmt.Errorf("resource %s is missing attribute %s in state", address, attribute)
	}

	return value, nil
}

func ResourceState(state *terraform.State, address string) (*terraform.ResourceState, error) {
	rootModule, err := RootModuleState(state)
	if err != nil {
		return nil, err
	}

	resourceState, ok := rootModule.Resources[address]
	if !ok {
		return nil, fmt.Errorf("resource %s not found in terraform state", address)
	}

	return resourceState, nil
}

func ExpectResourceAbsent(state *terraform.State, address string) error {
	rootModule, err := RootModuleState(state)
	if err != nil {
		return err
	}

	if _, ok := rootModule.Resources[address]; ok {
		return fmt.Errorf("resource %s unexpectedly persisted in terraform state", address)
	}

	return nil
}

func OutputValue(state *terraform.State, name string) (string, error) {
	rootModule, err := RootModuleState(state)
	if err != nil {
		return "", err
	}

	outputState, ok := rootModule.Outputs[name]
	if !ok {
		return "", fmt.Errorf("output %s not found in terraform state", name)
	}
	if outputState == nil {
		return "", fmt.Errorf("output %s has nil state", name)
	}
	value, ok := outputState.Value.(string)
	if !ok {
		return "", fmt.Errorf("output %s is %T, expected string", name, outputState.Value)
	}

	return value, nil
}

func ExpectOutputAbsent(state *terraform.State, name string) error {
	rootModule, err := RootModuleState(state)
	if err != nil {
		return err
	}

	if _, ok := rootModule.Outputs[name]; ok {
		return fmt.Errorf("output %s unexpectedly persisted in terraform state", name)
	}

	return nil
}
