package slicefunc

import (
	"context"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &minStringFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewMinStringFunction() function.Function {
	return &minStringFunction{}
}

// function implementation
type minStringFunction struct{}

// function metadata
func (*minStringFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "min_string"
}

// define the provider-level definition for the function
func (*minStringFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Return minimum string in list",
		MarkdownDescription: "Return the minimum string (first by lexical ordering) from the elements of an input list parameter.",
		Parameters: []function.Parameter{
			function.ListParameter{
				ElementType: types.StringType,
				Name:        "list",
				Description: "Input list parameter for determining the minimum string.",
			},
		},
		Return: function.StringReturn{},
	}
}

func (*minStringFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize list
	var list []string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &list))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "min_string: list", list)

	// validation
	if len(list) == 0 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(0, "min_string: list parameter length must be at least 1"))
		return
	}

	// determine maximum string element of slice
	minString := slices.Min(list)

	// store the result as a string
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &minString))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "min_string: minimum string", minString)
	tflog.Debug(ctx, "min_string: successful return", map[string]any{"success": true})
}
