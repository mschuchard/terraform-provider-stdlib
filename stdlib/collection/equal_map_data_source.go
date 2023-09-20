package collection

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"golang.org/x/exp/maps" // TODO: 1.21 migrate
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &equalMapDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewEqualMapDataSource() datasource.DataSource {
	return &equalMapDataSource{}
}

// data source implementation
type equalMapDataSource struct{}

// maps the data source schema data to the model
type equalMapDataSourceModel struct {
	MapOne types.Map  `tfsdk:"map_one"`
	MapTwo types.Map  `tfsdk:"map_two"`
	Result types.Bool `tfsdk:"result"`
}

// data source metadata
func (_ *equalMapDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_equal_map"
}

// define the provider-level schema for configuration data
func (_ *equalMapDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"map_one": schema.MapAttribute{
				Description: "First input map parameter to check for equality with the second.",
				ElementType: types.StringType,
				Required:    true,
			},
			"map_two": schema.MapAttribute{
				Description: "Second input map parameter to check for equality with the first.",
				ElementType: types.StringType,
				Required:    true,
			},
			"result": schema.BoolAttribute{
				Computed:    true,
				Description: "Function result storing whether the two maps are equal.",
			},
		},
		MarkdownDescription: "Return whether the two input map parameters contain the same key-value pairs (equality check). The input maps must be single-level",
	}
}

// read executes the actual function
func (_ *equalMapDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state equalMapDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// convert tf maps to go maps
	var mapOne, mapTwo map[string]string
	resp.Diagnostics.Append(state.MapOne.ElementsAs(ctx, &mapOne, false)...)
	resp.Diagnostics.Append(state.MapTwo.ElementsAs(ctx, &mapTwo, false)...)

	// check equality of maps and assign to model field member
	result := maps.Equal(mapOne, mapTwo)
	state.Result = types.BoolValue(result)

	// provide debug logging
	ctx = tflog.SetField(ctx, "stdlib_equal_map_map_one", mapOne)
	ctx = tflog.SetField(ctx, "stdlib_equal_map_map_two", mapTwo)

	// provide more debug logging
	ctx = tflog.SetField(ctx, "stdlib_equal_map_result", result)
	tflog.Debug(ctx, fmt.Sprintf("Result of whether map '%v' equals map '%v' is: %t", mapOne, mapTwo, result))

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined whether two maps are equal", map[string]any{"success": true})
}