package stringfunc

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &equalFoldFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewEqualFoldFunction() function.Function {
	return &equalFoldFunction{}
}

// function implementation
type equalFoldFunction struct{}

// function metadata
func (*equalFoldFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "equal_fold"
}

// define the provider-level definition for the function
func (*equalFoldFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Determine case-insensitive equality of two strings.",
		MarkdownDescription: "Determine if two strings, interpreted as UTF-8 strings, are equal under simple Unicode case-folding. This is a more general form of case-insensitivity.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "string_one",
				Description: "First input string parameter to compare with the second.",
			},
			function.StringParameter{
				Name:        "string_two",
				Description: "Second input string parameter to compare with the first.",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (*equalFoldFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize input parameters
	var stringOne, stringTwo string

	resp.Error = req.Arguments.Get(ctx, &stringOne, &stringTwo)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "equal_fold: string_one", stringOne)
	ctx = tflog.SetField(ctx, "equal_fold: string_two", stringTwo)

	// determine case insensitive equality between the two strings
	result := strings.EqualFold(stringOne, stringTwo)
	ctx = tflog.SetField(ctx, "equal_fold: result", result)

	// store the result as a boolean
	resp.Error = resp.Result.Set(ctx, &result)
	if resp.Error != nil {
		return
	}

	tflog.Debug(ctx, "equal_fold: successful return", map[string]any{"success": true})
}
