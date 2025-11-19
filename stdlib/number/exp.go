package numberfunc

import (
	"context"
	"math"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &expFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewExpFunction() function.Function {
	return &expFunction{}
}

// function implementation
type expFunction struct{}

// function metadata
func (*expFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "exp"
}

// define the provider-level definition for the function
func (*expFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Determine exponential of a number",
		MarkdownDescription: "Return the base-e exponential of an input number parameter.",
		Parameters: []function.Parameter{
			function.NumberParameter{
				Name:        "number",
				Description: "Input number parameter for determining the base-e exponential.",
			},
		},
		Return: function.Float64Return{},
	}
}

func (*expFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize input parameters
	var inputNum *big.Float

	resp.Error = req.Arguments.Get(ctx, &inputNum)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "exp: number", inputNum)

	// convert to float64
	float, _ := inputNum.Float64()
	if math.IsNaN(float) || math.IsInf(float, 0) {
		resp.Error = function.NewArgumentFuncError(0, "exp: input number is beyond the limits of float64")
		return
	}

	// determine base e exponential
	exponential := math.Exp(float)
	ctx = tflog.SetField(ctx, "exp: exponential", exponential)

	// store the result as a float64
	resp.Error = resp.Result.Set(ctx, &exponential)
	if resp.Error != nil {
		return
	}

	tflog.Debug(ctx, "exp: successful return", map[string]any{"success": true})
}
