package numberfunc_test

import (
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
	}

	util.UnitTests(testCases, resultData, numberfunc.NewModFunction(), test)
}
