package numberfunc_test

import (
	"math"
	"math/big"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	numberfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/number"
)

func TestSqrtFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.NumberUnknown())

	testCases := util.TestCases{
		"four": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(4)}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.NumberValue(big.NewFloat(2.0))),
			},
		},
		"zero": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(0)}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.NumberValue(big.NewFloat(0.0))),
			},
		},
		"two": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(2)}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.NumberValue(big.NewFloat(1.4142135623730951))),
			},
		},
		"negative": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(-1)}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "sqrt: the input number cannot be negative"),
				Result: resultData,
			},
		},
		"infinite": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(math.Inf(1))}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "sqrt: the input number cannot be 'positive or negative infinity'"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, numberfunc.NewSqrtFunction(), test)
}
