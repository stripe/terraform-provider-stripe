//
// File generated from our OpenAPI spec
//

package provider_test

import (
	"testing"

	"github.com/stripe/terraform-provider-stripe/internal/provider"
)

func TestProviderFactoryReturnsProvider(t *testing.T) {
	factory := provider.New("test")
	if factory == nil {
		t.Fatal("New returned nil factory")
	}
	p := factory()
	if p == nil {
		t.Fatal("factory returned nil provider")
	}
}
