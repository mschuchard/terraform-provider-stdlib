package slicefunc

import (
	"context"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &replaceFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewReplaceFunction() function.Function {
	return &replaceFunction{}
}

// function implementation
type replaceFunction struct{}

// function metadata
func (*replaceFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "replace"
}

// define the provider-level definition for the function
func (*replaceFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Replace elements in a list",
		MarkdownDescription: "Return the list where values are replaced at a specific element index (inclusive). This function errors if the `end_index`, or the specified `index` plus the length of the `replace_values` list, is out of range for the original list (greater than or equal to the length of the `list` parameter).",
		Parameters: []function.Parameter{
			function.ListParameter{
				ElementType: types.StringType,
				Name:        "list",
				Description: "Input list parameter for which the values will be replaced.",
			},
			function.ListParameter{
				ElementType: types.StringType,
				Name:        "replace_values",
				Description: "Input list of values which will replace values in the list.",
			},
			function.Int32Parameter{
				Name:        "index",
				Description: "Index in the list at which to begin replacing the values (inclusive).",
			},
		},
		VariadicParameter: function.Int32Parameter{
			Name:                "end_index",
			MarkdownDescription: "Optional: The index in the list at which to finish replacing values (inclusive). If the difference between this and the `index` is greater than or equal to the length of the list of the `replace_values`, then the additional elements in the original `list` will all be zeroed (i.e. removed; see third example of `provider::stdlib::replace`). This parameter input value is only necessary for that situation as otherwise its value will be automatically deduced by the provider function.",
		},
		Return: function.ListReturn{ElementType: types.StringType},
	}
}

func (*replaceFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize list and replace from input parameters
	var list, replaceValues []string
	var index, endIndex int
	var endIndexVar []int

	resp.Error = req.Arguments.Get(ctx, &list, &replaceValues, &index, &endIndexVar)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "replace: list", list)
	ctx = tflog.SetField(ctx, "replace: replace values", replaceValues)
	ctx = tflog.SetField(ctx, "replace: index", index)
	ctx = tflog.SetField(ctx, "replace: end_index variadic", endIndexVar)

	// validation
	if len(list) == 0 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(0, "replace: list parameter must not be empty"))
	}
	if len(replaceValues) == 0 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(1, "replace: replace values parameter must not be empty"))
	}
	if index < 0 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(2, "replace: index parameter must not be a negative number"))
	}

	// optional end index
	if len(endIndexVar) == 1 {
		// add one to the endIndex since TF list begins at 0 and not 1 (go code uses slicing)
		endIndex = endIndexVar[0] + 1

		// validation
		if endIndex <= index {
			resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(3, "replace: end index parameter must be greater than or equal to the index parameter"))
		}
	} else {
		// s[i:j] element ordering
		endIndex = index + len(replaceValues)
	}

	ctx = tflog.SetField(ctx, "replace: end_index", endIndex)

	// validate end index is within bounds of slice
	if endIndex > len(list) {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(3, "replace: The index at which to replace the values added to the length of the replacement values (i.e. 'endIndex') cannot be greater than the length of the list where the values will be replaced as that would be out of range."))
	}
	// rolls up all parameter validation
	if resp.Error != nil {
		return
	}

	// replace values into list at index
	result := slices.Replace(list, index, endIndex, replaceValues...)

	// store the result as a list of strings
	resp.Error = resp.Result.Set(ctx, &result)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "replace: result", result)
	tflog.Debug(ctx, "replace: successful return", map[string]any{"success": true})
}
