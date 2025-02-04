package numberfunc

import (
	"context"
	"math"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &roundFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewRoundFunction() function.Function {
	return &roundFunction{}
}

// function implementation
type roundFunction struct{}

// function metadata
func (*roundFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "round"
}

// define the provider-level definition for the function
func (*roundFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Determine rounding of a number",
		MarkdownDescription: "Return the nearest integer of an input parameter; rounding half away from zero.",
		Parameters: []function.Parameter{
			function.Float64Parameter{
				Name:        "number",
				Description: "Input number parameter for determining the rounding.",
			},
		},
		Return: function.Int64Return{},
	}
}

func (*roundFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize input parameters
	var inputNum float64

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &inputNum))
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "round: number", inputNum)

	// determine the rounded integer
	round := int64(math.Round(inputNum))
	ctx = tflog.SetField(ctx, "round: round", round)

	// store the result as an int64
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &round))
	if resp.Error != nil {
		return
	}

	tflog.Debug(ctx, "round: successful return", map[string]any{"success": true})
}
