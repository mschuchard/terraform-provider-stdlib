package mapfunc_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	mapfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/map"
)

func TestInsertFunction(test *testing.T) {
	test.Parallel()

	standardTestCases := map[string]struct {
		request  function.RunRequest
		expected function.RunResponse
	}{
		"prepend": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(
						types.MapType{ElemType: types.StringType},
						[]attr.Value{types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world")}), types.MapValueMust(types.StringType, map[string]attr.Value{"foo": types.StringValue("bar")})},
					),
				}),
			},
			expected: function.RunResponse{
				Result: function.NewResultData(types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar")})),
			},
		},
		"list-maps-length": {
			request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.MapType{ElemType: types.StringType}, []attr.Value{}),
				}),
			},
			expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "flatten_map: list of maps parameter must be at least length 1"),
				Result: function.NewResultData(types.MapUnknown(types.StringType)),
			},
		},
	}

	for name, testCase := range standardTestCases {
		test.Run(name, func(test *testing.T) {
			// initialize result
			result := function.RunResponse{Result: function.NewResultData(types.MapUnknown(types.StringType))}

			// execute function and store result
			mapfunc.NewFlattenMapFunction().Run(context.Background(), testCase.request, &result)

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
