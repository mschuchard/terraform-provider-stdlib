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
var _ datasource.DataSource = &replaceDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewReplaceDataSource() datasource.DataSource {
	return &replaceDataSource{}
}

// data source implementation
type replaceDataSource struct{}

// maps the data source schema data to the model
type replaceDataSourceModel struct {
	ID            types.String `tfsdk:"id"`
	EndIndex      types.Int64  `tfsdk:"end_index"`
	Index         types.Int64  `tfsdk:"index"`
	ReplaceValues types.List   `tfsdk:"replace_values"`
	ListParam     types.List   `tfsdk:"list_param"`
	Result        types.List   `tfsdk:"result"`
}

// data source metadata
func (_ *replaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_replace"
}

// define the provider-level schema for configuration data
func (_ *replaceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDStringAttribute(),
			"end_index": schema.Int64Attribute{
				Description: "The index in the list at which to end replacing values. If the difference between this and the index is greater than or equal to the length of the list of the replace_values, then the additional elements will all be zeroed (i.e. removed). This parameter input value is only necessary for that condition as otherwise it will be automatically deduced by the provider function.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"list_param": schema.ListAttribute{
				Description: "Input list parameter for which the values will be replaced.",
				ElementType: types.StringType,
				Required:    true,
			},
			"replace_values": schema.ListAttribute{
				Description: "Input list of values which will replace values in the list_param.",
				ElementType: types.StringType,
				Required:    true,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"index": schema.Int64Attribute{
				Description: "Index in the list at which to begin replacing the values.",
				Required:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListAttribute{
				Computed:    true,
				Description: "The resulting list with the replaced values.",
				ElementType: types.StringType,
			},
		},
		MarkdownDescription: "Return the list where values are replaced at a specific element index. This function errors if the specified index plus the length of the replace_values list, or the end_index, is out of range for the list (length of list_param + 1). Note also that the terminating index is determined by generic slice s[i:j] in Go, and so it may be helpful in Terraform to consider the terminating index as beginning at element 1 (does not apply to the the end_index input value), and that the length of the resulting list will therefore be one greater than the original unless end_index is specified.",
	}
}

// read executes the actual function
func (_ *replaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state replaceDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// convert tf list to go slice, tf int64 to go int
	var listParam, replaceValues []string
	resp.Diagnostics.Append(state.ListParam.ElementsAs(ctx, &listParam, false)...)
	resp.Diagnostics.Append(state.ReplaceValues.ElementsAs(ctx, &replaceValues, false)...)
	index := int(state.Index.ValueInt64())

	// determine end_index
	var endIndex int
	if state.EndIndex.IsNull() {
		// s[i:j] element ordering and not s[i] so subtract 1
		endIndex = index + len(replaceValues) - 1
	} else {
		endIndex = int(state.EndIndex.ValueInt64()) + 1
	}

	// determine if end index is out of bounds for slice
	if endIndex > len(listParam) {
		resp.Diagnostics.AddAttributeError(
			path.Root("index"),
			"Invalid Value",
			"The index at which to replace the values added to the length of the replacement values cannot be greater than the length of the list where the values will be replaced as that would be out of range.",
		)
		return
	}

	// replace values into list at index
	result := slices.Replace(listParam, index, endIndex, replaceValues...)

	// provide debug logging
	ctx = tflog.SetField(ctx, "stdlib_replace_result", result)
	tflog.Debug(ctx, fmt.Sprintf("Values \"%v\" replaced with \"%v\" at index \"%d\" to \"%d\"", listParam, replaceValues, index, endIndex))
	tflog.Debug(ctx, fmt.Sprintf("Resulting list is \"%s\"", result))

	// store zeroth element of input as id
	state.ID = types.StringValue(listParam[0])
	// store list with values replaced at index in state
	var listConvertDiags diag.Diagnostics
	state.Result, listConvertDiags = types.ListValueFrom(ctx, types.StringType, result)
	resp.Diagnostics.Append(listConvertDiags...)

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined list with values replaced at index", map[string]any{"success": true})
}
