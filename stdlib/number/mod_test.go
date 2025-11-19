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

func TestModFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.Float64Unknown())

	testCases := util.TestCases{
		"zero": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(4), types.Float64Value(2)}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(0)),
			},
		},
		"integer": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(5), types.Float64Value(3)}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(2)),
			},
		},
		"float": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(10), types.Float64Value(4.75)}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(0.5)),
			},
		},
		"zero-divisor": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(10), types.Float64Value(0)}),
			},
			Expected: function.RunResponse{
				Result: resultData,
				Error:  function.NewArgumentFuncError(1, "mod: divisor cannot be zero"),
			},
		},
		"dividend-beyond-upper-limit": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.NumberValue(func() *big.Float { f := new(big.Float); f.SetString("1e+310"); return f }()), types.Float64Value(10)}),
			},
			Expected: function.RunResponse{
				Result: resultData,
				Error:  function.NewArgumentFuncError(0, "mod: dividend is beyond the limits of float64"),
			},
		},
		"divisor-beyond-upper-limit": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(10), types.NumberValue(func() *big.Float { f := new(big.Float); f.SetString("1e+310"); return f }())}),
			},
			Expected: function.RunResponse{
				Result: resultData,
				Error:  function.NewArgumentFuncError(1, "mod: divisor is beyond the limits of float64"),
			},
		},
	}

	util.UnitTests(testCases, resultData, numberfunc.NewModFunction(), test)
}
