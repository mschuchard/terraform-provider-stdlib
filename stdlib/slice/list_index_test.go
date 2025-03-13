package slicefunc_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	slicefunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/slice"
)

func TestListIndexFunction(test *testing.T) {
	test.Parallel()

	standardTestCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"one": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two")}),
					types.StringValue("one"),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(1)),
			},
		},
		"sorted": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("a"), types.StringValue("b"), types.StringValue("c"), types.StringValue("d")}),
					types.StringValue("c"),
					types.TupleValueMust([]attr.Type{types.BoolType}, []attr.Value{types.BoolValue(true)}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(2)),
			},
		},
		"repeated": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three"), types.StringValue("two"), types.StringValue("one"), types.StringValue("zero")}),
					types.StringValue("two"),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(2)),
			},
		},
		"absent": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("hundred"), types.StringValue("thousand"), types.StringValue("million"), types.StringValue("billion")}),
					types.StringValue("infinity"),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(-1)),
			},
		},
		"list-length": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{}),
					types.StringValue("foo"),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "list_index: list parameter length must be at least 1"),
				Result: function.NewResultData(types.Int32Unknown()),
			},
		},
	}

	for name, testCase := range standardTestCases {
		test.Run(name, func(test *testing.T) {
			// initialize result
			result := function.RunResponse{Result: function.NewResultData(types.Int32Unknown())}

			// execute function and store result
			slicefunc.NewListIndexFunction().Run(context.Background(), testCase.request, &result)

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
