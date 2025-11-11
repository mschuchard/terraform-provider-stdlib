package numberfunc

import (
	"context"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &combinationsFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewCombinationsFunction() function.Function {
	return &combinationsFunction{}
}

// function implementation
type combinationsFunction struct{}

// function metadata
func (*combinationsFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "combinations"
}

// define the provider-level definition for the function
func (*combinationsFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Determine the number of combinations",
		MarkdownDescription: "Return the number of combinations given a set with a number of elements, and a selection size.",
		Parameters: []function.Parameter{
			function.Int64Parameter{
				Name:        "num_elements",
				Description: "Input number parameter for the number of elements in a set from which to make a selection. This value cannot be negative since the resultant 'NaN' is not allowed in Terraform types.",
			},
			function.Int64Parameter{
				Name:        "selection_size",
				Description: "Input number parameter for the number of elements to select from the set. This value cannot be negative since the resultant 'NaN' is not allowed in Terraform types.",
			},
		},
		Return: function.Int64Return{},
	}
}

func (*combinationsFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize input parameters
	var numElements, selectionSize int64

	resp.Error = req.Arguments.Get(ctx, &numElements, &selectionSize)
	if resp.Error != nil {
		return
	}

	ctx = tflog.SetField(ctx, "combinations: num_elements", numElements)
	ctx = tflog.SetField(ctx, "combinations: selection_size", selectionSize)

	// validate input parameters
	if numElements < 0 || selectionSize < 0 {
		resp.Error = function.NewArgumentFuncError(0, "combinations: the input number(s) cannot be negative")
		return
	}

	// initalize result
	var combinations int64

	// determine number of combinations
	if selectionSize > numElements {
		// in this situation the result is zero and should not be calculated
		combinations = 0
	} else {
		// n!
		numerator := (&big.Int{}).MulRange(1, numElements)
		// k!
		kFact := (&big.Int{}).MulRange(1, selectionSize)
		// (n-k)!
		nMinuskFact := (&big.Int{}).MulRange(1, numElements-selectionSize)
		// k! * (n-k)!
		denominator := (&big.Int{}).Mul(kFact, nMinuskFact)
		// n! / (k! * (n-k)!)
		result := big.NewInt(0).Div(numerator, denominator)
		if !result.IsInt64() {
			resp.Error = function.NewArgumentFuncError(0, "combinations: result exceeds maximum int64 value")
			return
		}
		combinations = result.Int64()
	}

	ctx = tflog.SetField(ctx, "combinations: combinations", combinations)

	// store the result as an int64
	resp.Error = resp.Result.Set(ctx, &combinations)
	if resp.Error != nil {
		return
	}

	tflog.Debug(ctx, "combinations: successful return", map[string]any{"success": true})
}
