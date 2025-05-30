package mapfunc

import (
	"context"
	"fmt"
	"maps"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	util "github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &hasValueDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewHasValueDataSource() datasource.DataSource {
	return &hasValueDataSource{}
}

// data source implementation
type hasValueDataSource struct{}

// maps the data source schema data to the model
type hasValueDataSourceModel struct {
	ID     types.String `tfsdk:"id"`
	Value  types.String `tfsdk:"value"`
	Map    types.Map    `tfsdk:"map"`
	Result types.Bool   `tfsdk:"result"`
}

// data source metadata
func (*hasValueDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_has_value"
}

// define the provider-level schema for configuration data
func (*hasValueDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDStringAttribute(),
			"value": schema.StringAttribute{
				Description: "Name of the value to check for existence in the map.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"map": schema.MapAttribute{
				Description: "Input map parameter from which to check a value's existence.",
				ElementType: types.StringType,
				Required:    true,
			},
			"result": schema.BoolAttribute{
				Computed:    true,
				Description: "Function result storing whether the key exists in the map.",
			},
		},
		MarkdownDescription: "Return whether the input value parameter is present in the input map parameter. The input map must be single-level.",
	}
}

// read executes the actual function
func (*hasValueDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state hasValueDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// initialize map and key
	valueCheck := state.Value.ValueString()
	var inputMap map[string]string
	resp.Diagnostics.Append(state.Map.ElementsAs(ctx, &inputMap, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// assign values of map and check input value's existence
	mapValues := slices.Collect(maps.Values(inputMap))
	valueExists := slices.Contains(mapValues, valueCheck)

	// provide more debug logging
	ctx = tflog.SetField(ctx, "stdlib_has_value_result", valueExists)
	tflog.Debug(ctx, fmt.Sprintf("Result of whether value '%s' is in map is: %t", valueCheck, valueExists))

	// store resultant map in state
	state.ID = types.StringValue(valueCheck)
	state.Result = types.BoolValue(valueExists)

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined whether value exists in map", map[string]any{"success": true})
}
