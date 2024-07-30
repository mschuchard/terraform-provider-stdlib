package mapfunc

import (
	"context"
	"golang.org/x/exp/maps"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

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
	ID     types.Int64 `tfsdk:"id"`
	Param  types.List  `tfsdk:"param"`
	Result types.Map   `tfsdk:"result"`
}

// data source metadata
func (_ *flattenMapDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_flatten_map"
}

// define the provider-level schema for configuration data
func (_ *flattenMapDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDInt64Attribute(),
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
				ElementType: types.StringType,
			},
		},
		MarkdownDescription: "Return the flattened map of an input list of maps parameter. Note that if a key is repeated then the last entry will overwrite any previous entries in the result.",
	}
}

// read executes the actual function
func (_ *flattenMapDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state flattenMapDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// initialize input list of maps, nested maps, and output map
	var inputList []types.Map
	resp.Diagnostics.Append(state.Param.ElementsAs(ctx, &inputList, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var nestedMap map[string]string
	outputMap := map[string]string{}

	// iterate through list of maps, convert each map, and merge each map into new map
	for _, nestedTFMap := range inputList {
		resp.Diagnostics.Append(nestedTFMap.ElementsAs(ctx, &nestedMap, false)...)
		if resp.Diagnostics.HasError() {
			return
		}

		maps.Copy(outputMap, nestedMap)
	}

	// store number of entries of output map as id
	state.ID = types.Int64Value(int64(len(outputMap)))
	// store flattened map in state
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
