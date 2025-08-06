package slicefunc

import (
	"context"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &listIndexFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewListIndexFunction() function.Function {
	return &listIndexFunction{}
}

// function implementation
type listIndexFunction struct{}

// function metadata
func (*listIndexFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "list_index"
}

// define the provider-level definition for the function
func (*listIndexFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Return index of element in list",
		MarkdownDescription: "Return the index of the first occurrence of the element parameter in the list parameter, or return `-1` if the element parameter is not present in the input list parameter.",
		Parameters: []function.Parameter{
			function.ListParameter{
				ElementType: types.StringType,
				Name:        "list",
				Description: "Input list parameter for determining the element's index.",
			},
			function.StringParameter{
				Name:        "element",
				Description: "Element in the list to determine its index.",
			},
		},
		VariadicParameter: function.BoolParameter{
			Name:        "sorted",
			Description: "Optional: Whether the list is sorted in ascending order or not (note: see `stdlib::sort_list`). If the list is sorted then the efficient binary search algorithm will be utilized, but the combination of sorting and searching may also be less efficient overall in some situations.",
		},
		Return: function.Int32Return{},
	}
}

func (*listIndexFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize list, elem, and sorted from input parameters
	var list []string
	var elem string
	var sortedVar []bool
	sorted := false

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &list, &elem, &sortedVar))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "list_index: list", list)
	ctx = tflog.SetField(ctx, "list_index: elem", elem)
	ctx = tflog.SetField(ctx, "list_index: sorted variadic", sortedVar)

	// validation
	if len(list) == 0 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(0, "list_index: list parameter length must be at least 1"))
		return
	}

	// optional sorted argument
	if len(sortedVar) == 1 {
		sorted = sortedVar[0]
	}

	ctx = tflog.SetField(ctx, "list_index: sorted", sorted)

	// determine element index within slice
	var listIndex int

	// use efficient binary search algorithm
	if sorted {
		var found bool
		listIndex, found = slices.BinarySearch(list, elem)

		// mimic slices.Index behavior for consistency
		if !found {
			listIndex = -1
		}
	} else { // use standard search algorithm
		listIndex = slices.Index(list, elem)
	}

	// store the result as an integer
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &listIndex))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "list_index: list_index", listIndex)
	tflog.Debug(ctx, "list_index: successful return", map[string]any{"success": true})
}
