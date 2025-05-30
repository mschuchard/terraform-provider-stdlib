package mapfunc

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	util "github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &hasKeyDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewHasKeyDataSource() datasource.DataSource {
	return &hasKeyDataSource{}
}

// data source implementation
type hasKeyDataSource struct{}

// maps the data source schema data to the model
type hasKeyDataSourceModel struct {
	ID     types.String `tfsdk:"id"`
	Key    types.String `tfsdk:"key"`
	Map    types.Map    `tfsdk:"map"`
	Result types.Bool   `tfsdk:"result"`
}

// data source metadata
func (*hasKeyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_has_key"
}

// define the provider-level schema for configuration data
func (*hasKeyDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDStringAttribute(),
			"key": schema.StringAttribute{
				Description: "Name of the key to check for existence in the map.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"map": schema.MapAttribute{
				Description: "Input map parameter from which to check a key's existence.",
				ElementType: types.StringType,
				Required:    true,
			},
			"result": schema.BoolAttribute{
				Computed:    true,
				Description: "Function result storing whether the key exists in the map.",
			},
		},
		MarkdownDescription: "Return whether the input key parameter is present in the input map parameter. The input map must be single-level.",
	}
}

// read executes the actual function
func (*hasKeyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state hasKeyDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// initialize map and key
	keyCheck := state.Key.ValueString()
	var inputMap map[string]string
	resp.Diagnostics.Append(state.Map.ElementsAs(ctx, &inputMap, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// check key's existence
	keyExists := false
	if _, ok := inputMap[keyCheck]; ok {
		keyExists = true
	}

	// provide more debug logging
	ctx = tflog.SetField(ctx, "stdlib_has_key_result", keyExists)
	tflog.Debug(ctx, fmt.Sprintf("Result of whether key '%s' is in map is: %t", keyCheck, keyExists))

	// store resultant map in state
	state.ID = types.StringValue(keyCheck)
	state.Result = types.BoolValue(keyExists)

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined whether key exists in map", map[string]any{"success": true})
}
