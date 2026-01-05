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

func TestTruncateFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.Int64Unknown())

	testCases := util.TestCases{
		"truncate-one": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(1.2345)}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int64Value(1)),
			},
		},
		"truncate-negative-sixty-seven": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(-67.89)}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int64Value(-67)),
			},
		},
		"beyond-upper-limit": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.NumberValue(func() *big.Float { f := new(big.Float); f.SetString("1e+310"); return f }())}),
			},
			Expected: function.RunResponse{
				Result: resultData,
				Error:  function.NewArgumentFuncError(0, "truncate: input number is beyond the limits of float64"),
			},
		},
		"overflow": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(9.223372036854777e+18)}), // math.MaxInt64 + 1
			},
			Expected: function.RunResponse{
				Result: resultData,
				Error:  function.NewArgumentFuncError(0, "truncate: truncated input number is beyond the limits of int64"),
			},
		},
	}

	util.UnitTests(testCases, resultData, numberfunc.NewTruncateFunction(), test)
}
