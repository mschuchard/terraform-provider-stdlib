package numberfunc_test

import (
	"math"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	numberfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/number"
)

func TestSqrtFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.Float64Unknown())

	testCases := util.TestCases{
		"four": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(4)}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(2)),
			},
		},
		"zero": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(0)}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(0)),
			},
		},
		"two": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(2)}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(1.4142135623730951)),
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
		/*"nan": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(math.NaN())}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "sqrt: the input number cannot be 'not a number'"),
				Result: resultData,
			},
		},*/
	}

	util.UnitTests(testCases, resultData, numberfunc.NewSqrtFunction(), test)
}
