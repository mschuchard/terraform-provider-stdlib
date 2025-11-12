package numberfunc

import (
	"context"
	"math"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &modFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewModFunction() function.Function {
	return &modFunction{}
}

// function implementation
type modFunction struct{}

// function metadata
func (*modFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "mod"
}

// define the provider-level definition for the function
func (*modFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Determine modulus of a number",
		MarkdownDescription: "Return the remainder of the dividend number divided by the divisor number.",
		Parameters: []function.Parameter{
			function.Float64Parameter{
				Name:        "dividend",
				Description: "The dividend number from which to divide.",
			},
			function.Float64Parameter{
				Name:        "divisor",
				Description: "The divisor number by which to divide.",
			},
		},
		Return: function.Float64Return{},
	}
}

func (*modFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize input parameters
	var dividend, divisor float64

	resp.Error = req.Arguments.Get(ctx, &dividend, &divisor)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "mod: dividend", dividend)
	ctx = tflog.SetField(ctx, "mod: divisor", divisor)

	// validate input parameters
	if divisor == 0 {
		resp.Error = function.NewArgumentFuncError(1, "mod: divisor cannot be zero")
		return
	}

	// determine the modulus
	modulus := math.Mod(dividend, divisor)
	ctx = tflog.SetField(ctx, "mod: modulus", modulus)

	// store the result as a float64
	resp.Error = resp.Result.Set(ctx, &modulus)
	if resp.Error != nil {
		return
	}

	tflog.Debug(ctx, "mod: successful return", map[string]any{"success": true})
}
