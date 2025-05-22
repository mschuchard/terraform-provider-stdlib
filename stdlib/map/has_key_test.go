package mapfunc_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mapfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/map"
)

func TestHasKeyFunction(test *testing.T) {
	test.Parallel()

	standardTestCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"present": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar")}),
					types.StringValue("foo"),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(true)),
			},
		},
		"absent": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar")}),
					types.StringValue("bar"),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(false)),
			},
		},
		"empty-key": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar")}),
					types.StringValue(""),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "has_key: input key parameter must not be empty"),
				Result: function.NewResultData(types.BoolUnknown()),
			},
		},
	}

	for name, testCase := range standardTestCases {
		test.Run(name, func(test *testing.T) {
			// initialize result
			result := function.RunResponse{Result: function.NewResultData(types.BoolUnknown())}

			// execute function and store result
			mapfunc.NewHasKeyFunction().Run(context.Background(), testCase.request, &result)

			// compare results
			if !result.Error.Equal(testCase.expected.Error) {
				test.Errorf("expected error: %s", testCase.expected.Error)
				test.Errorf("actual error: %s", result.Error)
			}
			if !result.Result.Equal(testCase.expected.Result) {
				test.Errorf("expected value: %t", testCase.expected.Result.Value())
				test.Errorf("actual value: %t", result.Result.Value())
			}
		})
	}
}
