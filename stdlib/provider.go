package stdlib

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// ensure the implementation satisfies the expected interfaces
var _ provider.Provider = &stdlibProvider{}

// helper pseudo-constructors to simplify provider server and testing implementation (second needed due to nuance in TF testing framework)
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return NewStruct(version)
	}
}

func NewStruct(version string) provider.Provider {
	return &stdlibProvider{
		version: version,
	}
}

// provider implementation
type stdlibProvider struct {
	version string
}

// provider metadata
func (_ *stdlibProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "stdlib"
}

// define the provider-level schema for configuration data
func (_ *stdlibProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

// prepare an API client for data sources and resources
func (_ *stdlibProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

// define the data sources implemented in the provider
func (_ *stdlibProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewFlattenMapDataSource,
		NewHasKeyDataSource,
		NewHasValueDataSource,
		NewKeyDeleteDataSource,
		NewLastCharDataSource,
	}
}

// define the resources implemented in the provider
func (_ *stdlibProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
