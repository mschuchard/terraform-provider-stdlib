package stringfunc

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &cutFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewCutFunction() function.Function {
	return &cutFunction{}
}

// function implementation
type cutFunction struct{}

// function metadata
func (*cutFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "cut"
}

// define the provider-level definition for the function
func (*cutFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Cut a string in two",
		MarkdownDescription: "Returns the strings before and after the first instance of the separator in the input string. Also returns whether or not the separator was found in the input string. The return is a tuple: `before`, `after`, `found`. If the separator is not found in the input string, then `found` will be false, `before` will be equal to `param`, and `after` will be an empty string.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "param",
				Description: "Input string parameter for cutting around a separator.",
			},
			function.StringParameter{
				Name:        "separator",
				Description: "The separator for cutting the input string.",
			},
		},
		Return: function.DynamicReturn{},
	}
}

func (*cutFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize input string param and separator from input parameters
	var inputString, separator string

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &inputString, &separator))
	if resp.Error != nil {
		return
	}

	// validate input parameters
	if len(inputString) < 1 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(0, "cut: input string parameter must be at least length 1"))
	}
	if len(separator) < 1 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(1, "cut: separator parameter must be at least length 1"))
	}

	// determine string cut
	before, after, found := strings.Cut(inputString, separator)

	// provide debug logging TODO

	// return the result as a tuple of string, string, bool
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, []any{before, after, found}))
}
