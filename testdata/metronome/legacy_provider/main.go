//
// File generated from our OpenAPI spec
//

// This test-only server replays the pinned standalone schema. It cannot read or
// mutate resources; all moved-state handling must go through the destination.
package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"log"
	"sort"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	providerschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

//go:embed schema.json
var schemaJSON []byte

type capturedAttribute struct {
	Shape, Element, Description             string
	Required, Optional, Computed, Sensitive bool
	Attributes                              map[string]capturedAttribute
}
type capturedResource struct {
	Description string
	Attributes  map[string]capturedAttribute
}
type emptyProvider struct{}

func (*emptyProvider) Metadata(_ context.Context, _ provider.MetadataRequest, r *provider.MetadataResponse) {
	r.TypeName = "metronome"
	r.Version = "fixture"
}
func (*emptyProvider) Schema(_ context.Context, _ provider.SchemaRequest, r *provider.SchemaResponse) {
	r.Schema = providerschema.Schema{}
}
func (*emptyProvider) Configure(context.Context, provider.ConfigureRequest, *provider.ConfigureResponse) {
}
func (*emptyProvider) Resources(context.Context) []func() resource.Resource       { return nil }
func (*emptyProvider) DataSources(context.Context) []func() datasource.DataSource { return nil }

type schemaReplay struct {
	tfprotov6.ProviderServer
	schemas map[string]*tfprotov6.Schema
}

func (s *schemaReplay) GetProviderSchema(ctx context.Context, req *tfprotov6.GetProviderSchemaRequest) (*tfprotov6.GetProviderSchemaResponse, error) {
	response, err := s.ProviderServer.GetProviderSchema(ctx, req)
	if err != nil {
		return nil, err
	}
	response.ResourceSchemas = s.schemas
	return response, nil
}
func (s *schemaReplay) GetMetadata(ctx context.Context, req *tfprotov6.GetMetadataRequest) (*tfprotov6.GetMetadataResponse, error) {
	response, err := s.ProviderServer.GetMetadata(ctx, req)
	if err != nil {
		return nil, err
	}
	for name := range s.schemas {
		response.Resources = append(response.Resources, tfprotov6.ResourceMetadata{TypeName: name})
	}
	return response, nil
}
func objectType(attributes map[string]capturedAttribute) tftypes.Type {
	fields := map[string]tftypes.Type{}
	for name, attribute := range attributes {
		fields[name] = attributeType(attribute)
	}
	return tftypes.Object{AttributeTypes: fields}
}
func attributeType(attribute capturedAttribute) tftypes.Type {
	switch attribute.Shape {
	case "string":
		return tftypes.String
	case "bool":
		return tftypes.Bool
	case "int64", "float64", "number":
		return tftypes.Number
	case "dynamic":
		return tftypes.DynamicPseudoType
	case "singlenested":
		return objectType(attribute.Attributes)
	case "listnested":
		return tftypes.List{ElementType: objectType(attribute.Attributes)}
	case "setnested":
		return tftypes.Set{ElementType: objectType(attribute.Attributes)}
	case "map", "list", "set":
		if attribute.Element != "basetypes.StringType" {
			panic("unreviewed legacy collection type")
		}
		switch attribute.Shape {
		case "map":
			return tftypes.Map{ElementType: tftypes.String}
		case "list":
			return tftypes.List{ElementType: tftypes.String}
		default:
			return tftypes.Set{ElementType: tftypes.String}
		}
	default:
		panic("unreviewed legacy schema type: " + attribute.Shape)
	}
}
func main() {
	var captured map[string]capturedResource
	if err := json.Unmarshal(schemaJSON, &captured); err != nil {
		log.Fatal(err)
	}
	schemas := map[string]*tfprotov6.Schema{}
	for name, resource := range captured {
		block := &tfprotov6.SchemaBlock{Description: resource.Description}
		names := make([]string, 0, len(resource.Attributes))
		for field := range resource.Attributes {
			names = append(names, field)
		}
		sort.Strings(names)
		for _, field := range names {
			a := resource.Attributes[field]
			block.Attributes = append(block.Attributes, &tfprotov6.SchemaAttribute{Name: field, Type: attributeType(a), Required: a.Required, Optional: a.Optional, Computed: a.Computed, Sensitive: a.Sensitive, Description: a.Description})
		}
		schemas[name] = &tfprotov6.Schema{Version: 0, Block: block}
	}
	if err := tf6server.Serve("registry.terraform.io/tonyxiao/metronome", func() tfprotov6.ProviderServer {
		return &schemaReplay{ProviderServer: providerserver.NewProtocol6(&emptyProvider{})(), schemas: schemas}
	}); err != nil {
		log.Fatal(err)
	}
}
