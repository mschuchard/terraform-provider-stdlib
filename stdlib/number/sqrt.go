package numberfunc

import (
	"context"
	"math/big"

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
			function.NumberParameter{
				Name:        "number",
				Description: "Input number parameter for determining the square root. This number cannot be negative, infinite (positive or negative), or NaN.",
			},
		},
		Return: function.NumberReturn{},
	}
}

func (*sqrtFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize input parameters
	var inputNum *big.Float

	resp.Error = req.Arguments.Get(ctx, &inputNum)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "sqrt: number", inputNum)

	// validate input parameters
	if inputNum.Cmp(big.NewFloat(0.0)) < 0 {
		resp.Error = function.NewArgumentFuncError(0, "sqrt: the input number cannot be negative")
		return
	}
	if inputNum.IsInf() {
		resp.Error = function.NewArgumentFuncError(0, "sqrt: the input number cannot be 'positive or negative infinity'")
		return
	}

	// determine the square root
	sqrt := (&big.Float{}).Sqrt(inputNum)
	ctx = tflog.SetField(ctx, "sqrt: square root", sqrt)

	// store the result as a number
	resp.Error = resp.Result.Set(ctx, &sqrt)
	if resp.Error != nil {
		return
	}

	tflog.Debug(ctx, "sqrt: successful return", map[string]any{"success": true})
}
