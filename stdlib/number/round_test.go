package numberfunc_test

import (
	"math/big"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	numberfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/number"
)

func TestRoundFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.Int64Unknown())

	testCases := util.TestCases{
		"round-down": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(1.2)}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int64Value(1)),
			},
		},
		"round-up": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(1.8)}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int64Value(2)),
			},
		},
		"round-half": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(1.5)}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int64Value(2)),
			},
		},
		"beyond-upper-limit": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.NumberValue(func() *big.Float { f := new(big.Float); f.SetString("1e+310"); return f }())}),
			},
			Expected: function.RunResponse{
				Result: resultData,
				Error:  function.NewFuncError("round: input number is beyond the limits of float64 for rounding"),
			},
		},
	}

	util.UnitTests(testCases, resultData, numberfunc.NewRoundFunction(), test)
}
