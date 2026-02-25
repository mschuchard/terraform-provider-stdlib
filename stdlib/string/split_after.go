package stringfunc

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &splitAfterFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewSplitAfterFunction() function.Function {
	return &splitAfterFunction{}
}

// function implementation
type splitAfterFunction struct{}

// function metadata
func (*splitAfterFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "split_after"
}

// define the provider-level definition for the function
func (*splitAfterFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Produces a list by splitting a given string after each occurrence of a separator.",
		MarkdownDescription: "Return a list of strings by splitting an input string parameter after each occurrence of a separator. If the separator is not empty nor present within the string, then the return value will be a list containing the original string as the only element instead. This function is roughly equivalent to the `split` function in Terraform core, but additionally retains the separator as part of the returned list elements.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "string",
				Description: "Input string parameter for searching for the substring.",
			},
			function.StringParameter{
				Name:        "separator",
				Description: "Input separator parameter for searching within the string.",
			},
		},
		Return: function.ListReturn{ElementType: types.StringType},
	}
}

func (*splitAfterFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize input parameters
	var inputString, separator string

	resp.Error = req.Arguments.Get(ctx, &inputString, &separator)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "split_after: string", inputString)
	ctx = tflog.SetField(ctx, "split_after: separator", separator)

	// validate input parameters
	if len(inputString) < 1 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(0, "split_after: input string parameter must be at least length 1"))
		return
	}

	// determine last char
	splitAfter := strings.SplitAfter(inputString, separator)
	ctx = tflog.SetField(ctx, "split_after: result", splitAfter)

	// store the result as a list of strings
	resp.Error = resp.Result.Set(ctx, splitAfter)
	if resp.Error != nil {
		return
	}

	tflog.Debug(ctx, "split_after: successful return", map[string]any{"success": true})
}
