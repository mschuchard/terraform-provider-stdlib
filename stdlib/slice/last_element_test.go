package slicefunc_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	slicefunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/slice"
)

func TestLastElementFunction(test *testing.T) {
	test.Parallel()

	standardTestCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"single": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("h"), types.StringValue("e"), types.StringValue("l"), types.StringValue("l"), types.StringValue("o")}),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("o")})),
			},
		},
		"three": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("h"), types.StringValue("e"), types.StringValue("l"), types.StringValue("l"), types.StringValue("o")}),
					types.TupleValueMust([]attr.Type{types.Int32Type}, []attr.Value{types.Int32Value(3)}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("l"), types.StringValue("l"), types.StringValue("o")})),
			},
		},
		"num-elem-too-great": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("h"), types.StringValue("e"), types.StringValue("l"), types.StringValue("l"), types.StringValue("o")}),
					types.TupleValueMust([]attr.Type{types.Int32Type}, []attr.Value{types.Int32Value(10)}),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "last_element: the number of terminating elements to return must be fewer than the length of the input list parameter"),
				Result: function.NewResultData(types.ListUnknown(types.StringType)),
			},
		},
		"num-elem-minimum": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("h"), types.StringValue("e"), types.StringValue("l"), types.StringValue("l"), types.StringValue("o")}),
					types.TupleValueMust([]attr.Type{types.Int32Type}, []attr.Value{types.Int32Value(0)}),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "last_element: number of elements parameter value must be at least 1"),
				Result: function.NewResultData(types.ListUnknown(types.StringType)),
			},
		},
	}

	for name, testCase := range standardTestCases {
		test.Run(name, func(test *testing.T) {
			// initialize result
			result := function.RunResponse{Result: function.NewResultData(types.ListUnknown(types.StringType))}

			// execute function and store result
			slicefunc.NewLastElementFunction().Run(context.Background(), testCase.request, &result)

			// compare results
			if !result.Error.Equal(testCase.expected.Error) {
				test.Errorf("expected error: %s", testCase.expected.Error)
				test.Errorf("actual error: %s", result.Error)
			}
			if !result.Result.Equal(testCase.expected.Result) {
				test.Errorf("expected value: %+q", testCase.expected.Result.Value())
				test.Errorf("actual value: %+q", result.Result.Value())
			}
		})
	}
}
