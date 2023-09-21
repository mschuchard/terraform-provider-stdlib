package collection

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"golang.org/x/exp/maps" // TODO: 1.21 migrate

	"github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &flattenMapDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewFlattenMapDataSource() datasource.DataSource {
	return &flattenMapDataSource{}
}

// data source implementation
type flattenMapDataSource struct{}

// maps the data source schema data to the model
type flattenMapDataSourceModel struct {
	ID     types.String `tfsdk:"id"`
	Param  types.List   `tfsdk:"param"`
	Result types.Map    `tfsdk:"result"`
}

// data source metadata
func (_ *flattenMapDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_flatten_map"
}

// define the provider-level schema for configuration data
func (_ *flattenMapDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDStringAttribute(),
			// TODO: also support set
			"param": schema.ListAttribute{
				Description: "Input list of maps to flatten.",
				ElementType: types.MapType{
					ElemType: types.StringType,
				},
				Required: true,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"result": schema.MapAttribute{
				Computed:    true,
				Description: "Function result storing the flattened map.",
				// TODO: allow non-strings with interface or generics
				ElementType: types.StringType,
			},
		},
		MarkdownDescription: "Return the flattened map of an input list of maps parameter. Note that if a key is repeated then the last entry will overwrite any previous entries in the result.",
	}
}

// TODO: need to revisit when plugin framework supports list(map) in the schema
// read executes the actual function
func (_ *flattenMapDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state flattenMapDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: allow non-strings with interface or generics
	// initialize input list of maps, nested maps, and output map
	var inputList []types.Map
	resp.Diagnostics.Append(state.Param.ElementsAs(ctx, &inputList, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var nestedMap map[string]string
	outputMap := map[string]string{}

	// iterate through list of maps and merge the maps into new map
	for _, nestedTFMap := range inputList {
		nestedTFMap.ElementsAs(ctx, &nestedMap, false)
		maps.Copy(outputMap, nestedMap)
	}
	// provide debug logging
	ctx = tflog.SetField(ctx, "stdlib_flatten_map_param", state.Param)
	ctx = tflog.SetField(ctx, "stdlib_flatten_map_result", outputMap)
	tflog.Debug(ctx, fmt.Sprintf("Flattened map is \"%v\"", outputMap))

	// store number of entries of output map as id
	state.ID = types.StringValue(fmt.Sprint(len(outputMap)))
	// TODO: allow non-strings with interface or generics
	var mapConvertDiags diag.Diagnostics
	state.Result, mapConvertDiags = types.MapValueFrom(ctx, types.StringType, outputMap)
	resp.Diagnostics.Append(mapConvertDiags...)

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined flattened map", map[string]any{"success": true})
}
