package stringfunc

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &countFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewCountFunction() function.Function {
	return &countFunction{}
}

// function implementation
type countFunction struct{}

// function metadata
func (*countFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "count"
}

// define the provider-level definition for the function
func (*countFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Counts the number of non-overlapping instances of a substring in a string.",
		MarkdownDescription: "Return the number of non-overlapping instances of a substring in a string. The substring cannot be empty as this counts the number of Unicode code points in the string (before and after each rune), and this is awkward within Terraform.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "string",
				Description: "Input string parameter for determining the number of occurrences of the substrings within it.",
			},
			function.StringParameter{
				Name:        "substring",
				Description: "Input substring parameter for determining the number of occurrences of that substring in a string.",
			},
		},
		Return: function.Int32Return{},
	}
}

func (*countFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize input parameters
	var inputString, subString string

	resp.Error = req.Arguments.Get(ctx, &inputString, &subString)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "count: string", inputString)
	ctx = tflog.SetField(ctx, "count: substring", subString)

	// validate input parameters
	if len(subString) < 1 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(1, "count: substring parameter must be at least length 1"))
		return
	}

	// count number of occurrences of the substring within the string
	count := strings.Count(inputString, subString)
	ctx = tflog.SetField(ctx, "count: result", count)

	// store the result as a string
	resp.Error = resp.Result.Set(ctx, &count)
	if resp.Error != nil {
		return
	}

	tflog.Debug(ctx, "count: successful return", map[string]any{"success": true})
}
