package slicefunc

import (
	"context"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &repeatFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewRepeatFunction() function.Function {
	return &repeatFunction{}
}

// function implementation
type repeatFunction struct{}

// function metadata
func (*repeatFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "repeat"
}

// define the provider-level definition for the function
func (*repeatFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Return the repeated list.",
		MarkdownDescription: "Return a list that repeats the input list the input number of times.",
		Parameters: []function.Parameter{
			function.ListParameter{
				ElementType: types.StringType,
				Name:        "list",
				Description: "Input list parameter for repeating. The list must contain at least one element.",
			},
			function.Int32Parameter{
				Name:        "count",
				Description: "Number of times to repeat the input list. This cannot be a negative number.",
			},
		},
		Return: function.ListReturn{ElementType: types.StringType},
	}
}

func (*repeatFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize list and count
	var list []string
	var count int

	resp.Error = req.Arguments.Get(ctx, &list, &count)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "repeat: list", list)
	ctx = tflog.SetField(ctx, "repeat: count", count)

	// validation
	if len(list) == 0 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(0, "repeat: list parameter length must be at least 1"))
	}
	if count < 0 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(1, "repeat: count parameter value cannot be negative"))
	}
	if resp.Error != nil {
		return
	}

	// determine the repeated list
	repeated := slices.Repeat(list, count)

	// store the result as a slice of strings
	resp.Error = resp.Result.Set(ctx, &repeated)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "repeat: repeated", repeated)
	tflog.Debug(ctx, "repeat: successful return", map[string]any{"success": true})
}
