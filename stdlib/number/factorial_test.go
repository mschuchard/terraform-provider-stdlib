// test negative, 0, and some reasonable num
package numberfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	numberfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/number"
)

func TestFactorialFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.Int64Unknown())

	testCases := util.TestCases{
		"zero": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Int64Value(0)}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int64Value(1)),
			},
		},
		"five": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Int64Value(5)}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int64Value(120)),
			},
		},
		"negative": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Int64Value(-1)}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "factorial: the input number cannot be negative"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, numberfunc.NewFactorialFunction(), test)
}
