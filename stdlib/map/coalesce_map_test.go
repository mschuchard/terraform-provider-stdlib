package mapfunc_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mapfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/map"
)

func TestCoalesceMapFunction(test *testing.T) {
	test.Parallel()

	standardTestCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"standard": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.TupleValueMust(
						[]attr.Type{
							types.MapType{ElemType: types.StringType}, types.MapType{ElemType: types.StringType}, types.MapType{ElemType: types.StringType}, types.MapType{ElemType: types.StringType},
						},
						[]attr.Value{
							types.MapValueMust(types.StringType, map[string]attr.Value{}),
							types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world")}),
							types.MapValueMust(types.StringType, map[string]attr.Value{"foo": types.StringValue("bar")}),
							types.MapValueMust(types.StringType, map[string]attr.Value{}),
						}),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world")})),
			},
		},
		"no-input-args": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.TupleValueMust([]attr.Type{}, []attr.Value{})}),
			},
			expected: function.RunResponse{
				Error:  function.NewFuncError("coalesce_map: at least one argument is required"),
				Result: function.NewResultData(types.MapUnknown(types.StringType)),
			},
		},
		"all-args-empty-maps": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.TupleValueMust(
						[]attr.Type{
							types.MapType{ElemType: types.StringType}, types.MapType{ElemType: types.StringType},
						},
						[]attr.Value{
							types.MapValueMust(types.StringType, map[string]attr.Value{}),
							types.MapValueMust(types.StringType, map[string]attr.Value{}),
						}),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewFuncError("coalesce_map: all arguments are empty maps"),
				Result: function.NewResultData(types.MapUnknown(types.StringType)),
			},
		},
	}

	for name, testCase := range standardTestCases {
		test.Run(name, func(test *testing.T) {
			// initialize result
			result := function.RunResponse{Result: function.NewResultData(types.MapUnknown(types.StringType))}

			// execute function and store result
			mapfunc.NewCoalesceMapFunction().Run(context.Background(), testCase.request, &result)

			// compare results
			if !result.Error.Equal(testCase.expected.Error) {
				test.Errorf("expected error: %s", testCase.expected.Error)
				test.Errorf("actual error: %s", result.Error)
			}
			if !result.Result.Equal(testCase.expected.Result) {
				test.Errorf("expected value: %+v", testCase.expected.Result.Value())
				test.Errorf("actual value: %+v", result.Result.Value())
			}
		})
	}
}
