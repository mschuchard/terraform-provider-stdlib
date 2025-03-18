package slicefunc_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	slicefunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/slice"
)

func TestProductFunction(test *testing.T) {
	test.Parallel()

	standardTestCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"zero": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.SetValueMust(types.Float64Type, []attr.Value{types.Float64Value(0), types.Float64Value(1), types.Float64Value(1), types.Float64Value(2)}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(0)),
			},
		},
		"single": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.SetValueMust(types.Float64Type, []attr.Value{types.Float64Value(5)}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(5)),
			},
		},
		"permutation": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.SetValueMust(types.Float64Type, []attr.Value{types.Float64Value(1), types.Float64Value(2), types.Float64Value(3), types.Float64Value(4), types.Float64Value(5)}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(120)),
			},
		},
		"set-length": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.SetValueMust(types.Float64Type, []attr.Value{}),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "product: set parameter length must be at least 1"),
				Result: function.NewResultData(types.Float64Unknown()),
			},
		},
	}

	for name, testCase := range standardTestCases {
		test.Run(name, func(test *testing.T) {
			// initialize result
			result := function.RunResponse{Result: function.NewResultData(types.Float64Unknown())}

			// execute function and store result
			slicefunc.NewProductFunction().Run(context.Background(), testCase.request, &result)

			// compare results
			if !result.Error.Equal(testCase.expected.Error) {
				test.Errorf("expected error: %s", testCase.expected.Error)
				test.Errorf("actual error: %s", result.Error)
			}
			if !result.Result.Equal(testCase.expected.Result) {
				test.Errorf("expected value: %f", testCase.expected.Result.Value())
				test.Errorf("actual value: %f", result.Result.Value())
			}
		})
	}
}
