package numberfunc

import (
	"context"
	"math"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &sqrtFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewSqrtFunction() function.Function {
	return &sqrtFunction{}
}

// function implementation
type sqrtFunction struct{}

// function metadata
func (*sqrtFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "sqrt"
}

// define the provider-level definition for the function
func (*sqrtFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Determine square root of a number",
		MarkdownDescription: "Return the square root of an input parameter.",
		Parameters: []function.Parameter{
			function.Float64Parameter{
				Name:        "number",
				Description: "Input number parameter for determining the square root.",
			},
		},
		Return: function.Float64Return{},
	}
}

func (*sqrtFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize input parameters
	var inputNum float64

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &inputNum))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "sqrt: number", inputNum)

	// determine the square root
	sqrt := math.Sqrt(inputNum)
	/*if math.IsNaN(sqrt) {
		resp.Diagnostics.AddAttributeError(
			path.Root("param"),
			"Invalid Value",
			"The square root of the input parameter must return a valid number, but instead returned 'NaN'",
		)
		return
	}*/

	ctx = tflog.SetField(ctx, "sqrt: square root", sqrt)

	// store the result as a float64
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &sqrt))
	if resp.Error != nil {
		return
	}

	tflog.Debug(ctx, "sqrt: successful return", map[string]any{"success": true})
}
