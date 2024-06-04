package slicefunc

import (
	"context"
	"fmt"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ datasource.DataSource = &listIndexDataSource{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewListIndexDataSource() datasource.DataSource {
	return &listIndexDataSource{}
}

// data source implementation
type listIndexDataSource struct{}

// maps the data source schema data to the model
type listIndexDataSourceModel struct {
	ID        types.String `tfsdk:"id"`
	ElemParam types.String `tfsdk:"elem_param"`
	ListParam types.List   `tfsdk:"list_param"`
	Sorted    types.Bool   `tfsdk:"sorted"`
	Result    types.Int64  `tfsdk:"result"`
}

// data source metadata
func (_ *listIndexDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_list_index"
}

// define the provider-level schema for configuration data
func (_ *listIndexDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": util.IDStringAttribute(),
			"list_param": schema.ListAttribute{
				Description: "Input list parameter for determining the element's index.",
				ElementType: types.StringType,
				Required:    true,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"elem_param": schema.StringAttribute{
				Description: "Element in the list to determine its index.",
				Required:    true,
			},
			"sorted": schema.BoolAttribute{
				Description: "Whether the list is sorted in ascending order or not (note: see `stdlib_sort_list`). If the list is sorted then the efficient binary search algorithm will be utilized, but the combination of sorting and searching may be less efficient overall in some situations.",
				Optional:    true,
			},
			"result": schema.Int64Attribute{
				Computed:    true,
				Description: "Function result storing the index of the element.",
			},
		},
		MarkdownDescription: "Return the index of the first occurrence of the element parameter in the list parameter, or return '-1' if the element parameter is not present in the input list parameter.",
	}
}

// read executes the actual function
func (_ *listIndexDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// determine input values
	var state listIndexDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// convert tf list to go slice and tf string to go string
	var listParam []string
	resp.Diagnostics.Append(state.ListParam.ElementsAs(ctx, &listParam, false)...)
	elemParam := state.ElemParam.ValueString()

	// determine element index within slice
	var listIndex int
	var found bool

	if state.Sorted.ValueBool() {
		// use efficient binary search algorithm
		listIndex, found = slices.BinarySearch(listParam, elemParam)
		// mimic slices.Index behavior for consistency
		if !found {
			listIndex= -1
		}
	} else {
		// use standard search algorithm
		listIndex = slices.Index(listParam, elemParam)
	}

	// provide debug logging
	ctx = tflog.SetField(ctx, "stdlib_list_index_result", listIndex)
	tflog.Debug(ctx, fmt.Sprintf("Input element parameter \"%s\" index in list parameter \"%v\" is \"%d\"", elemParam, listParam, listIndex))

	// store listIndex element(s) of list in state
	state.ID = types.StringValue(listParam[0])
	state.Result = types.Int64Value(int64(listIndex))

	// set state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Determined element index within list", map[string]any{"success": true})
}
