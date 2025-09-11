package numberfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	numberfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/number"
)

func TestCombinationsFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.Int64Unknown())

	testCases := util.TestCases{
		"three-choose-two": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Int64Value(3), types.Int64Value(2)}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int64Value(3)),
			},
		},
		"four-choose-five": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Int64Value(4), types.Int64Value(5)}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int64Value(0)),
			},
		},
		"negative": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Int64Value(-1), types.Int64Value(-1)}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "combinations: the input number(s) cannot be negative"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, numberfunc.NewCombinationsFunction(), test)
}
