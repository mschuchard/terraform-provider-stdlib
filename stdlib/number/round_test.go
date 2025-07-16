package numberfunc_test

import (
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
	}

	util.UnitTests(testCases, resultData, numberfunc.NewRoundFunction(), test)
}
