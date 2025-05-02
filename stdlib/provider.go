package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	mapfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/map"
	"github.com/mschuchard/terraform-provider-stdlib/stdlib/multiple"
	numberfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/number"
	slicefunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/slice"
	stringfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/string"
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
func (*stdlibProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "stdlib"
}

// define the provider-level schema for configuration data
func (*stdlibProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The stdlib provider provides additional functions for use within Terraform's HCL2 configuration language.",
	}
}

// prepare an API client for data sources and resources
func (*stdlibProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
}

// define the data sources implemented in the provider
func (*stdlibProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		mapfunc.NewEqualMapDataSource,
		mapfunc.NewFlattenMapDataSource,
		mapfunc.NewHasKeyDataSource,
		mapfunc.NewHasKeysDataSource,
		mapfunc.NewHasValueDataSource,
		mapfunc.NewHasValuesDataSource,
		mapfunc.NewKeyDeleteDataSource,
		mapfunc.NewKeysDeleteDataSource,
		multiple.NewEmptyDataSource,
		numberfunc.NewExpDataSource,
		numberfunc.NewModDataSource,
		numberfunc.NewRoundDataSource,
		numberfunc.NewSqrtDataSource,
		slicefunc.NewCompareListDataSource,
		slicefunc.NewInsertDataSource,
		slicefunc.NewLastElementDataSource,
		slicefunc.NewListIndexDataSource,
		slicefunc.NewMaxNumberDataSource,
		slicefunc.NewMaxStringDataSource,
		slicefunc.NewMinNumberDataSource,
		slicefunc.NewMinStringDataSource,
		slicefunc.NewProductDataSource,
		slicefunc.NewReplaceDataSource,
		slicefunc.NewSortListDataSource,
		stringfunc.NewCutDataSource,
		stringfunc.NewLastCharDataSource,
	}
}

// define the functions implemented in the provider
func (*stdlibProvider) Functions(_ context.Context) []func() function.Function {
	return []func() function.Function{
		mapfunc.NewEqualMapFunction,
		mapfunc.NewFlattenMapFunction,
		mapfunc.NewHasKeyFunction,
		mapfunc.NewHasKeysFunction,
		mapfunc.NewHasValueFunction,
		numberfunc.NewExpFunction,
		numberfunc.NewModFunction,
		numberfunc.NewRoundFunction,
		numberfunc.NewSqrtFunction,
		slicefunc.NewCompareListFunction,
		slicefunc.NewInsertFunction,
		slicefunc.NewLastElementFunction,
		slicefunc.NewListIndexFunction,
		slicefunc.NewMaxNumberFunction,
		slicefunc.NewMaxStringFunction,
		slicefunc.NewMinNumberFunction,
		slicefunc.NewMinStringFunction,
		slicefunc.NewProductFunction,
		slicefunc.NewReplaceFunction,
		slicefunc.NewSortListFunction,
		stringfunc.NewCutFunction,
		stringfunc.NewLastCharFunction,
	}
}

// define the resources implemented in the provider
func (*stdlibProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}
