package slicefunc_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	slicefunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/slice"
)

func TestCutFunction(test *testing.T) {
	test.Parallel()

	standardTestCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"lesser": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("bar"), types.StringValue("b")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("bar"), types.StringValue("baz")}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(-1)),
			},
		},
		"equal": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("pizza"), types.StringValue("cake")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("pizza"), types.StringValue("cake")}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(0)),
			},
		},
		"greater": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("super"), types.StringValue("hyper"), types.StringValue("turbo")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("pizza"), types.StringValue("cake"), types.StringValue("punch")}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(1)),
			},
		},
		"length": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("pizza"), types.StringValue("cake"), types.StringValue("punch")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("pizza"), types.StringValue("cake")}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(1)),
			},
		},
	}

	for name, testCase := range standardTestCases {
		test.Run(name, func(test *testing.T) {
			// initialize result
			result := function.RunResponse{Result: function.NewResultData(types.Int32Unknown())}

			// execute function and store result
			slicefunc.NewCompareListFunction().Run(context.Background(), testCase.request, &result)

			// compare results
			if !result.Error.Equal(testCase.expected.Error) {
				test.Errorf("expected value: %s", testCase.expected.Error)
				test.Errorf("actual value: %s", result.Error)
			}
			if !result.Result.Equal(testCase.expected.Result) {
				test.Errorf("expected value: %d", testCase.expected.Result.Value())
				test.Errorf("actual value: %d", result.Result.Value())
			}
		})
	}
}
