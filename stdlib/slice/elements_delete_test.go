package slicefunc_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	slicefunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/slice"
)

func TestElementsDeleteFunction(test *testing.T) {
	test.Parallel()

	standardTestCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"begin": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three")}),
					types.Int32Value(0),
					types.Int32Value(1),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("two"), types.StringValue("three")})),
			},
		},
		"end": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three")}),
					types.Int32Value(2),
					types.Int32Value(3),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one")})),
			},
		},
		// this returns two errors, and as far as I can tell this cannot be compared with func (*FuncError) Equal as it involves two structs
		/*"empty-list": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{}),
					types.Int32Value(0),
					types.Int32Value(0),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "elements_delete: list parameter must not be empty"),
				Result: function.NewResultData(types.ListUnknown(types.StringType)),
			},
		},*/
		"negative-index": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three")}),
					types.Int32Value(-1),
					types.Int32Value(2),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "elements_delete: index parameter must not be a negative number"),
				Result: function.NewResultData(types.ListUnknown(types.StringType)),
			},
		},
		"end-index-fewer-than-index": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three")}),
					types.Int32Value(2),
					types.Int32Value(1),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(2, "elements_delete: end_index parameter must be greater than or equal to the index parameter"),
				Result: function.NewResultData(types.ListUnknown(types.StringType)),
			},
		},
		"end-index-out-of-bounds": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("zero"), types.StringValue("one"), types.StringValue("two"), types.StringValue("three")}),
					types.Int32Value(2),
					types.Int32Value(4),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(2, "elements_delete: end_index parameter must not be greater than or equal to the length of the list"),
				Result: function.NewResultData(types.ListUnknown(types.StringType)),
			},
		},
	}

	for name, testCase := range standardTestCases {
		test.Run(name, func(test *testing.T) {
			// initialize result
			result := function.RunResponse{Result: function.NewResultData(types.ListUnknown(types.StringType))}

			// execute function and store result
			slicefunc.NewElementsDeleteFunction().Run(context.Background(), testCase.request, &result)

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
