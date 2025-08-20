package slicefunc

import (
	"context"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &elementsDeleteFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewElementsDeleteFunction() function.Function {
	return &elementsDeleteFunction{}
}

// function implementation
type elementsDeleteFunction struct{}

// function metadata
func (*elementsDeleteFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "elements_delete"
}

// define the provider-level definition for the function
func (*elementsDeleteFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Delete elements from a list",
		MarkdownDescription: "Return the list where values are deleted beginning and ending inclusively at two specific element indices. This function errors if the `end_index` is out of range for the original list (greater than or equal to the length of the `list` parameter).",
		Parameters: []function.Parameter{
			function.ListParameter{
				ElementType: types.StringType,
				Name:        "list",
				Description: "Input list parameter for which the values will be deleted.",
			},
			function.Int32Parameter{
				Name:        "index",
				Description: "Index in the list at which to begin deleting the values (inclusive).",
			},
			function.Int32Parameter{
				Name:        "end_index",
				Description: "Index in the list at which to finish deleting the values (inclusive).",
			},
		},
		Return: function.ListReturn{ElementType: types.StringType},
	}
}

func (*elementsDeleteFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize list and indices from input parameters
	var list []string
	var index, endIndex int

	resp.Error = req.Arguments.Get(ctx, &list, &index, &endIndex)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "elements_delete: list", list)
	ctx = tflog.SetField(ctx, "elements_delete: index", index)
	ctx = tflog.SetField(ctx, "elements_delete: end_index", endIndex)

	// validation
	if len(list) == 0 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(0, "elements_delete: list parameter must not be empty"))
	}
	if index < 0 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(1, "elements_delete: index parameter must not be a negative number"))
	}
	if endIndex < index {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(2, "elements_delete: end_index parameter must be greater than or equal to the index parameter"))
	}
	if endIndex >= len(list) {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(2, "elements_delete: end_index parameter must not be greater than or equal to the length of the list"))
	}
	if resp.Error != nil {
		return
	}

	// add one to the endIndex since TF list begins at 0 and not 1 (go code uses slicing), and endIndex is inclusive
	endIndex += 1

	// delete elements from list between index and endIndex
	result := slices.Delete(list, index, endIndex)

	// store the result as a list of strings
	resp.Error = resp.Result.Set(ctx, &result)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "elements_delete: result", result)
	tflog.Debug(ctx, "elements_delete: successful return", map[string]any{"success": true})
}
