package numberfunc

import (
	"context"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &factorialFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewFactorialFunction() function.Function {
	return &factorialFunction{}
}

// function implementation
type factorialFunction struct{}

// function metadata
func (*factorialFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "factorial"
}

// define the provider-level definition for the function
func (*factorialFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Determine the factorial",
		MarkdownDescription: "Return the factorial of an input number parameter.",
		Parameters: []function.Parameter{
			function.Int64Parameter{
				Name:        "integer",
				Description: "Input number parameter for determining the factorial. This value cannot be negative since the resultant 'NaN' is not allowed in Terraform types.",
			},
		},
		Return: function.Int64Return{},
	}
}

func (*factorialFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize input parameters
	var inputNumber int64

	resp.Error = req.Arguments.Get(ctx, &inputNumber)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "factorial: number", inputNumber)

	// validate input parameters
	if inputNumber < 0 {
		resp.Error = function.NewArgumentFuncError(0, "factorial: the input number cannot be negative")
		return
	}

	// determine factorial
	result := (&big.Int{}).MulRange(1, inputNumber)
	if !result.IsInt64() {
		resp.Error = function.NewFuncError("factorial: result exceeds maximum int64 value")
		return
	}
	factorial := result.Int64()

	ctx = tflog.SetField(ctx, "factorial: factorial", factorial)

	// store the result as an int64
	resp.Error = resp.Result.Set(ctx, &factorial)
	if resp.Error != nil {
		return
	}

	tflog.Debug(ctx, "factorial: successful return", map[string]any{"success": true})
}
