package stringfunc

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &compareStringFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewCompareStringFunction() function.Function {
	return &compareStringFunction{}
}

// function implementation
type compareStringFunction struct{}

// function metadata
func (*compareStringFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "compare_string"
}

// define the provider-level definition for the function
func (*compareStringFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Compare two strings similar to the spaceship operator.",
		MarkdownDescription: "Lexicographically compares two strings and returns an integer indicating their relationship. A `-1` indicates the first string is lexicographically less than the second string, `0` indicates they are equal, and a `1` indicates the first string is lexicographically greater than the second string.",
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
		Return: function.Int32Return{},
	}
}

func (*compareStringFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize input parameters
	var stringOne, stringTwo string

	resp.Error = req.Arguments.Get(ctx, &stringOne, &stringTwo)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "compare_string: string_one", stringOne)
	ctx = tflog.SetField(ctx, "compare_string: string_two", stringTwo)

	// determine lexicographical relationship between the two strings
	result := strings.Compare(stringOne, stringTwo)
	ctx = tflog.SetField(ctx, "compare_string: result", result)

	// store the result as a list of strings
	resp.Error = resp.Result.Set(ctx, &result)
	if resp.Error != nil {
		return
	}

	tflog.Debug(ctx, "compare_string: successful return", map[string]any{"success": true})
}
