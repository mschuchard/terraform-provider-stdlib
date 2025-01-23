package stringfunc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &lastCharFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewLastCharFunction() function.Function {
	return &lastCharFunction{}
}

// function implementation
type lastCharFunction struct{}

// function metadata
func (*lastCharFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "last_char"
}

// define the provider-level definition for the function
func (*lastCharFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Determine last character(s) of a string",
		MarkdownDescription: "Return the last character(s) of an input string parameter. Only the terminating character is returned by default unless a value for `num_chars` is defined.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "string",
				Description: "Input string parameter for determining the last character.",
			},
		},
		VariadicParameter: function.Int32Parameter{
			Name:        "number_of_characters",
			Description: "Optional: The number of terminating characters at the end of the string to return (default: 1).",
		},
		Return: function.StringReturn{},
	}
}

func (*lastCharFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize input parameters
	var inputString string
	var numCharsVar []int
	var numChars int

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &inputString, &numCharsVar))
	if resp.Error != nil {
		return
	}

	// validate input parameters
	if len(inputString) < 1 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(0, "lastChar: input string parameter must be at least length 1"))
	}
	if len(numCharsVar) == 0 {
		// assign default numChars value of 1
		numChars = 1
	} else {
		// assign numChars from variadic
		numChars = numCharsVar[0]

		// and then continue validation
		if numChars < 1 {
			resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(1, "lastChar: number_of_characters parameter must be at least 1"))
		} else if numChars >= len(inputString) {
			resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(1, "lastChar: number_of_characters parameter must be fewer than the length of the input string parameter"))
		}
	}
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "lastChar: string", inputString)
	ctx = tflog.SetField(ctx, "lastChar: number_of_characters", numChars)

	// determine last char
	lastCharacter := inputString[len(inputString)-numChars:]
	ctx = tflog.SetField(ctx, "lastChar: last_character(s)", lastCharacter)

	// store the result as a tuple of string, string, bool
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &lastCharacter))
	if resp.Error != nil {
		return
	}

	tflog.Debug(ctx, "lastChar: successful return", map[string]any{"success": true})
}
