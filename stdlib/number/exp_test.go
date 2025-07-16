package numberfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	numberfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/number"
)

func TestExpFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.Float64Unknown())

	testCases := util.TestCases{
		"zero-one": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(0)}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(1)),
			},
		},
		"float-one-three": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(1.0986122)}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(2.9999997339956828)),
			},
		},
	}

	util.UnitTests(testCases, resultData, numberfunc.NewExpFunction(), test)
}
