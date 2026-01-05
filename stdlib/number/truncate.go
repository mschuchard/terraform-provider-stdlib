package numberfunc

import (
	"context"
	"math"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &truncateFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewTruncateFunction() function.Function {
	return &truncateFunction{}
}

// function implementation
type truncateFunction struct{}

// function metadata
func (*truncateFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "truncate"
}

// define the provider-level definition for the function
func (*truncateFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Determine integer value of a number",
		MarkdownDescription: "Return the integer value of an input parameter.",
		Parameters: []function.Parameter{
			function.NumberParameter{
				Name:        "number",
				Description: "Input number parameter for determining the truncated integer value.",
			},
		},
		Return: function.Int64Return{},
	}
}

func (*truncateFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize input parameters
	var inputNum *big.Float

	resp.Error = req.Arguments.Get(ctx, &inputNum)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "truncate: number", inputNum)

	// convert to float64
	float, _ := inputNum.Float64()
	if math.IsNaN(float) || math.IsInf(float, 0) {
		resp.Error = function.NewArgumentFuncError(0, "truncate: input number is beyond the limits of float64")
		return
	}

	// determine the truncated integer
	truncate := math.Trunc(float)
	if truncate > float64(math.MaxInt64) || truncate < float64(math.MinInt64) {
		resp.Error = function.NewArgumentFuncError(0, "truncate: truncated input number is beyond the limits of int64")
		return
	}
	result := int64(truncate)

	ctx = tflog.SetField(ctx, "truncate: truncated", result)

	// store the result as an int64
	resp.Error = resp.Result.Set(ctx, &result)
	if resp.Error != nil {
		return
	}

	tflog.Debug(ctx, "truncate: successful return", map[string]any{"success": true})
}
