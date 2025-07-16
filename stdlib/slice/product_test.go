package slicefunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	slicefunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/slice"
)

func TestProductFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.Float64Unknown())

	testCases := util.TestCases{
		"zero": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.SetValueMust(types.Float64Type, []attr.Value{types.Float64Value(0), types.Float64Value(1), types.Float64Value(1), types.Float64Value(2)}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(0)),
			},
		},
		"single": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.SetValueMust(types.Float64Type, []attr.Value{types.Float64Value(5)}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(5)),
			},
		},
		"permutation": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.SetValueMust(types.Float64Type, []attr.Value{types.Float64Value(1), types.Float64Value(2), types.Float64Value(3), types.Float64Value(4), types.Float64Value(5)}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(120)),
			},
		},
		"set-length": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.SetValueMust(types.Float64Type, []attr.Value{}),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "product: set parameter length must be at least 1"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, slicefunc.NewProductFunction(), test)
}
