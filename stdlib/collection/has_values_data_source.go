package collection

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"golang.org/x/exp/maps"   // TODO: 1.21 migrate
	"golang.org/x/exp/slices" // TODO: 1.21 migrate

	"github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &hasValuesDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewHasValuesDataSource() datasource.DataSource {
	return &hasValuesDataSource{}
}

// data source implementation
type hasValuesDataSource struct{}

// maps the data source schema data to the model
type hasValuesDataSourceModel struct {
	ID     types.String `tfsdk:"id"`
	All    types.Bool   `tfsdk:"all"`
	Values types.List   `tfsdk:"values"`
	Map    types.Map    `tfsdk:"map"`
	Result types.Bool   `tfsdk:"result"`
}

// data source metadata
func (_ *hasValuesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_has_values"
}

// define the provider-level schema for configuration data
func (_ *hasValuesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDStringAttribute(),
			"all": schema.BoolAttribute{
				Description: "Whether to check for all of the values instead of the default any of the values.",
				Optional:    true,
			},
			"values": schema.ListAttribute{
				Description: "Names of the values to check for existence in the map.",
				Required:    true,
				ElementType: types.StringType,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(2),
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
		MarkdownDescription: "Return whether any or all of the input value parameters are present in the input map parameter. The input map must be single-level.",
	}
}

// read executes the actual function
func (_ *hasValuesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state hasValuesDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// initialize values, map, and return
	var valuesCheck []string
	resp.Diagnostics.Append(state.Values.ElementsAs(ctx, &valuesCheck, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var inputMap map[string]string
	resp.Diagnostics.Append(state.Map.ElementsAs(ctx, &inputMap, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// declare value existence and all vs. any, and then determine all value
	var valueExists, all bool
	if !state.All.IsNull() {
		all = state.All.ValueBool()
	}

	// provide debug logging
	ctx = tflog.SetField(ctx, "stdlib_has_values_values", valuesCheck)
	ctx = tflog.SetField(ctx, "stdlib_has_values_map", inputMap)
	ctx = tflog.SetField(ctx, "stdlib_has_values_all", all)

	// assign values of map
	mapValues := maps.Values(inputMap)
	// switch between any of the values or all of the values
	if all {
		// assume all of the values exist until single check proves otherwise
		valueExists = true

		// iterate through values to check
		for _, value := range valuesCheck {
			// check input values' existence
			if !slices.Contains(mapValues, value) {
				valueExists = false
				break
			}
		}
	} else {
		// assume none of the values exist until single check proves otherwise
		valueExists = false

		// iterate through values to check
		for _, value := range valuesCheck {
			// check input values' existence
			if slices.Contains(mapValues, value) {
				valueExists = true
				break
			}
		}
	}

	// provide more debug logging
	ctx = tflog.SetField(ctx, "stdlib_has_values_result", valueExists)
	tflog.Debug(ctx, fmt.Sprintf("Result of whether values '%q' are in map is: %t", valuesCheck, valueExists))

	// store resultant map in state
	state.ID = types.StringValue(valuesCheck[0])
	state.Result = types.BoolValue(valueExists)

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined whether value exists in map", map[string]any{"success": true})
}
