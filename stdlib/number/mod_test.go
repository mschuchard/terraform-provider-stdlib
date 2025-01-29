package numberfunc_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	numberfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/number"
)

func TestModFunction(test *testing.T) {
	test.Parallel()

	standardTestCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"zero": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(4), types.Float64Value(2)}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(0)),
			},
		},
		"integer": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(5), types.Float64Value(3)}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(2)),
			},
		},
		"float": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.Float64Value(10), types.Float64Value(4.75)}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.Float64Value(0.5)),
			},
		},
	}

	for name, testCase := range standardTestCases {
		test.Run(name, func(test *testing.T) {
			// initialize result
			result := function.RunResponse{Result: function.NewResultData(types.Float64Unknown())}

			// execute function and store result
			numberfunc.NewModFunction().Run(context.Background(), testCase.request, &result)

			// compare results
			if !result.Error.Equal(testCase.expected.Error) {
				test.Errorf("expected value: %s", testCase.expected.Error)
				test.Errorf("actual value: %s", result.Error)
			}
			if !result.Result.Equal(testCase.expected.Result) {
				test.Errorf("expected value: %+q", testCase.expected.Result.Value())
				test.Errorf("actual value: %+q", result.Result.Value())
			}
		})
	}
}
