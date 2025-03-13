package slicefunc

import (
	"context"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &maxStringFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewMaxStringFunction() function.Function {
	return &maxStringFunction{}
}

// function implementation
type maxStringFunction struct{}

// function metadata
func (*maxStringFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "max_string"
}

// define the provider-level definition for the function
func (*maxStringFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Return maximum string in list",
		MarkdownDescription: "Return the maximum string (last by lexical ordering) from the elements of an input list parameter.",
		Parameters: []function.Parameter{
			function.ListParameter{
				ElementType: types.StringType,
				Name:        "list",
				Description: "Input list parameter for determining the maximum string.",
			},
		},
		Return: function.StringReturn{},
	}
}

func (*maxStringFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize list
	var list []string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &list))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "max_string: list", list)

	// validation
	if len(list) == 0 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(0, "max_string: list parameter length must be at least 1"))
	}
	if resp.Error != nil {
		return
	}

	// determine maximum string element of slice
	maxString := slices.Max(list)

	// store the result as a string
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &maxString))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "max_string: maximum string", maxString)
	tflog.Debug(ctx, "max_string: successful return", map[string]any{"success": true})
}
