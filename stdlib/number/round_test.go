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
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(1.2), types.TupleValueMust([]attr.Type{}, []attr.Value{})}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int64Value(1)),
			},
		},
		"round-up": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(1.8), types.TupleValueMust([]attr.Type{}, []attr.Value{})}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int64Value(2)),
			},
		},
		"round-half": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(1.5), types.TupleValueMust([]attr.Type{}, []attr.Value{})}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int64Value(2)),
			},
		},
		"round-half-even-true": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(2.5), types.TupleValueMust([]attr.Type{types.BoolType}, []attr.Value{types.BoolValue(true)})}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int64Value(2)),
			},
		},
		"round-half-even-false": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(2.5), types.TupleValueMust([]attr.Type{types.BoolType}, []attr.Value{types.BoolValue(false)})}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int64Value(3)),
			},
		},
		"beyond-upper-limit": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.NumberValue(func() *big.Float { f := new(big.Float); f.SetString("1e+310"); return f }()), types.TupleValueMust([]attr.Type{}, []attr.Value{})}),
			},
			Expected: function.RunResponse{
				Result: resultData,
				Error:  function.NewArgumentFuncError(0, "round: input number is beyond the limits of float64"),
			},
		},
		"overflow": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(9.223372036854777e+18), types.TupleValueMust([]attr.Type{}, []attr.Value{})}), // math.MaxInt64 + 1
			},
			Expected: function.RunResponse{
				Result: resultData,
				Error:  function.NewArgumentFuncError(0, "round: rounded input number is beyond the limits of int64"),
			},
		},
	}

	util.UnitTests(testCases, resultData, numberfunc.NewRoundFunction(), test)
}
