// File generated from our OpenAPI spec and handwritten Terraform acceptance templates.
// Handwritten Terraform acceptance source emitted by sdk-codegen.
package runner

import "os"

func ShouldRunCase(def CaseDefinition) bool {
	return matchesSelector("STRIPE_TF_ACCTEST_GROUP", []string{def.Group}) &&
		matchesSelector("STRIPE_TF_ACCTEST_SURFACE", []string{def.Surface}) &&
		matchesSelector("STRIPE_TF_ACCTEST_CASE", []string{def.Name})
}

func matchesSelector(envKey string, values []string) bool {
	selector := os.Getenv(envKey)
	if selector == "" {
		return true
	}

	for _, value := range values {
		if value == selector {
			return true
		}
	}

	return false
}
