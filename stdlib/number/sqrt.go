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
				Description: "Input number parameter for determining the square root. This number cannot be negative, infinite (positive or negative), or NaN.",
			},
		},
		Return: function.Float64Return{},
	}
}

func (*sqrtFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize input parameters
	var inputNum float64

	resp.Error = req.Arguments.Get(ctx, &inputNum)

	// validate input parameters
	if inputNum < 0 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(0, "sqrt: the input number cannot be negative"))
	} else if math.IsInf(inputNum, 0) {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(0, "sqrt: the input number cannot be 'positive or negative infinity'"))
	} else if math.IsNaN(inputNum) {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(0, "sqrt: the input number cannot be 'not a number'"))
	}
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "sqrt: number", inputNum)

	// determine the square root
	sqrt := math.Sqrt(inputNum)
	ctx = tflog.SetField(ctx, "sqrt: square root", sqrt)

	// store the result as a float64
	resp.Error = resp.Result.Set(ctx, &sqrt)
	if resp.Error != nil {
		return
	}

	tflog.Debug(ctx, "sqrt: successful return", map[string]any{"success": true})
}
