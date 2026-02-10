package stringfunc

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &lastIndexFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewLastIndexFunction() function.Function {
	return &lastIndexFunction{}
}

// function implementation
type lastIndexFunction struct{}

// function metadata
func (*lastIndexFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "last_index"
}

// define the provider-level definition for the function
func (*lastIndexFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Determine index of the last instance of a substring within a string.",
		MarkdownDescription: "Return the index of the last instance of a substring within an input string parameter. If the substring is not present within the string then the value '-1' will be returned instead.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "string",
				Description: "Input string parameter for searching for the substring.",
			},
			function.StringParameter{
				Name:        "substring",
				Description: "Input substring parameter for searching within the string.",
			},
		},
		Return: function.Int32Return{},
	}
}

func (*lastIndexFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize input parameters
	var inputString, substring string

	resp.Error = req.Arguments.Get(ctx, &inputString, &substring)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "last_index: string", inputString)
	ctx = tflog.SetField(ctx, "last_index: substring", substring)

	// validate input parameters
	if len(inputString) < 1 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(0, "last_index: input string parameter must be at least length 1"))
	}

	// determine last char
	lastIndex := strings.LastIndex(inputString, substring)
	ctx = tflog.SetField(ctx, "last_index: result", lastIndex)

	// store the result as an int32
	resp.Error = resp.Result.Set(ctx, int32(lastIndex))
	if resp.Error != nil {
		return
	}

	tflog.Debug(ctx, "last_index: successful return", map[string]any{"success": true})
}
