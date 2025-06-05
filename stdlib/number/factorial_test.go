// test negative, 0, and some reasonable num
package numberfunc_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	numberfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/number"
)

func TestFactorialFunction(test *testing.T) {
	test.Parallel()

	standardTestCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"zero": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Int64Value(0)}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.Int64Value(1)),
			},
		},
		"five": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Int64Value(5)}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.Int64Value(120)),
			},
		},
		"negative": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Int64Value(-1)}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "factorial: the input number cannot be negative"),
				Result: function.NewResultData(types.Int64Unknown()),
			},
		},
	}

	for name, testCase := range standardTestCases {
		test.Run(name, func(test *testing.T) {
			// initialize result
			result := function.RunResponse{Result: function.NewResultData(types.Int64Unknown())}

			// execute function and store result
			numberfunc.NewFactorialFunction().Run(context.Background(), testCase.request, &result)

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
