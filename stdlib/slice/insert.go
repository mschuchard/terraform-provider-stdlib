package slicefunc

import (
	"context"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &insertFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewInsertFunction() function.Function {
	return &insertFunction{}
}

// function implementation
type insertFunction struct{}

// function metadata
func (*insertFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "insert"
}

// define the provider-level definition for the function
func (*insertFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Insert elements into a list",
		MarkdownDescription: "Return the list where values are inserted into a list at a specific index. The elements at the index in the original list are shifted up to make room. This function errors if the specified index is out of range for the list (length + 1).",
		Parameters: []function.Parameter{
			function.ListParameter{
				ElementType: types.StringType,
				Name:        "list",
				Description: "Input list parameter into which the values will be inserted.",
			},
			function.ListParameter{
				ElementType: types.StringType,
				Name:        "insert_values",
				Description: "Input list of values which will be inserted into the list.",
			},
			function.Int32Parameter{
				Name:        "index",
				Description: "Index in the list at which to insert the values.",
			},
		},
		Return: function.ListReturn{ElementType: types.StringType},
	}
}

func (*insertFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize list and insert from input parameters
	var list, insertValues []string
	var index int

	resp.Error = req.Arguments.Get(ctx, &list, &insertValues, &index)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "insert: list", list)
	ctx = tflog.SetField(ctx, "insert: insert values", insertValues)
	ctx = tflog.SetField(ctx, "insert: index", index)

	// validation
	if len(insertValues) < 1 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(1, "insert: insert values parameter must be at least length 1"))
	}
	if index < 0 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(2, "insert: index parameter must not be a negative number"))
	}
	// determine if index is out of bounds for slice
	if int(index) > len(list) {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(2, "insert: index parameter must not be out of range for list"))
	}
	if resp.Error != nil {
		return
	}

	// insert values into list at index
	result := slices.Insert(list, index, insertValues...)

	// store the result as a list of strings
	resp.Error = resp.Result.Set(ctx, &result)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "insert: result", result)
	tflog.Debug(ctx, "insert: successful return", map[string]any{"success": true})
}
