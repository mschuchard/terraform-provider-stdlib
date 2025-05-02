package mapfunc_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mapfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/map"
)

func TestHasKeysFunction(test *testing.T) {
	test.Parallel()

	standardTestCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"any-present": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar"), "baz": types.StringValue("bat")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("bar"), types.StringValue("foo")}),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(true)),
			},
		},
		"any-absent": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar"), "baz": types.StringValue("bat")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("bar"), types.StringValue("pizza")}),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(false)),
			},
		},
		"all-absent": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar"), "baz": types.StringValue("bat")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("bar"), types.StringValue("foo")}),
					types.TupleValueMust([]attr.Type{types.BoolType}, []attr.Value{types.BoolValue(true)}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(false)),
			},
		},
		"all-present": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar"), "baz": types.StringValue("bat")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("hello"), types.StringValue("foo"), types.StringValue("baz")}),
					types.TupleValueMust([]attr.Type{types.BoolType}, []attr.Value{types.BoolValue(true)}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(true)),
			},
		},
		"keys-too-short": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar"), "baz": types.StringValue("bat")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("bar")}),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "has_keys: keys parameter must be at least length 2"),
				Result: function.NewResultData(types.BoolUnknown()),
			},
		},
	}

	for name, testCase := range standardTestCases {
		test.Run(name, func(test *testing.T) {
			// initialize result
			result := function.RunResponse{Result: function.NewResultData(types.BoolUnknown())}

			// execute function and store result
			mapfunc.NewHasKeysFunction().Run(context.Background(), testCase.request, &result)

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
