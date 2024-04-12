package slicefunc

import (
	"context"
	"fmt"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &insertDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewInsertDataSource() datasource.DataSource {
	return &insertDataSource{}
}

// data source implementation
type insertDataSource struct{}

// maps the data source schema data to the model
type insertDataSourceModel struct {
	ID           types.String `tfsdk:"id"`
	Index        types.Int64  `tfsdk:"index"`
	InsertValues types.List   `tfsdk:"insert_values"`
	ListParam    types.List   `tfsdk:"list_param"`
	Result       types.List   `tfsdk:"result"`
}

// data source metadata
func (_ *insertDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_insert"
}

// define the provider-level schema for configuration data
func (_ *insertDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDStringAttribute(),
			"list_param": schema.ListAttribute{
				Description: "Input list parameter into which the values will be inserted.",
				ElementType: types.StringType,
				Required:    true,
			},
			"insert_values": schema.ListAttribute{
				Description: "Input list of values which will be inserted into the list.",
				ElementType: types.StringType,
				Required:    true,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"index": schema.Int64Attribute{
				Description: "Index in the list at which to insert the values.",
				Required:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListAttribute{
				Computed:    true,
				Description: "The resulting list with the inserted values.",
				ElementType: types.StringType,
				Required:    true,
			},
		},
		MarkdownDescription: "Return the list where values are inserted into a list at a specific index. The elments at the index in the original list are shifted up to make room. This function errors if the specified index is out of range for the list.",
	}
}

// read executes the actual function
func (_ *insertDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state insertDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// convert tf list to go slice and tf string to go string
	var listParam, insertValues []string
	resp.Diagnostics.Append(state.ListParam.ElementsAs(ctx, &listParam, false)...)
	resp.Diagnostics.Append(state.InsertValues.ElementsAs(ctx, &insertValues, false)...)
	index := state.Index.ValueInt64()

	// determine if index is out of bounds for slice
	if int(index) >= len(listParam) {
		resp.Diagnostics.AddAttributeError(
			path.Root("index"),
			"Invalid Value",
			"The index at which to insert the values cannot be greater than or equal to the length of the list into which the values will be inserted as that would be out of range.",
		)
		return
	}

	// insert values into list at index
	result := slices.Insert(listParam, int(index), insertValues...)

	// provide debug logging
	ctx = tflog.SetField(ctx, "stdlib_insert_result", result)
	tflog.Debug(ctx, fmt.Sprintf("Values \"%v\" inserted into \"%v\" at index \"%d\"", insertValues, listParam, index))
	tflog.Debug(ctx, fmt.Sprintf("Resulting list is \"%s\"", result))

	// store number of entries of output map as id
	state.ID = types.StringValue(listParam[0])
	// store list with values inserted at index in state
	var listConvertDiags diag.Diagnostics
	state.Result, listConvertDiags = types.ListValueFrom(ctx, types.StringType, result)
	resp.Diagnostics.Append(listConvertDiags...)

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined list with values inserted at index", map[string]any{"success": true})
}
