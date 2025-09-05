package slicefunc

import (
	"context"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &sortListNumberFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewSortListNumberFunction() function.Function {
	return &sortListNumberFunction{}
}

// function implementation
type sortListNumberFunction struct{}

// function metadata
func (*sortListNumberFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "sort_list_number"
}

// define the provider-level definition for the function
func (*sortListNumberFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Sort the numeric elements in a list",
		MarkdownDescription: "Return the list where numeric values are sorted in ascending order.",
		Parameters: []function.Parameter{
			function.ListParameter{
				ElementType: types.Float32Type,
				Name:        "list",
				Description: "Input list parameter for sorting. This must be at least size 2.",
			},
		},
		Return: function.ListReturn{ElementType: types.Float32Type},
	}
}

func (*sortListNumberFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize list
	var list []float32

	resp.Error = req.Arguments.Get(ctx, &list)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "sort_list_number: list", list)

	// validation
	if len(list) < 2 {
		resp.Error = function.NewArgumentFuncError(0, "sort_list_number: list parameter length must be at least 2")
		return
	}

	// sort the list
	slices.Sort(list)

	// store the result as a list of floats
	resp.Error = resp.Result.Set(ctx, &list)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "sort_list_number: sorted list", list)
	tflog.Debug(ctx, "sort_list_number: successful return", map[string]any{"success": true})
}
