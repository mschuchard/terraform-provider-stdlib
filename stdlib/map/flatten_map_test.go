package mapfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	mapfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/map"
)

func TestFlattenMapFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.MapUnknown(types.StringType))

	testCases := util.TestCases{
		"prepend": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(
						types.MapType{ElemType: types.StringType},
						[]attr.Value{types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world")}), types.MapValueMust(types.StringType, map[string]attr.Value{"foo": types.StringValue("bar")})},
					),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar")})),
			},
		},
		"list-maps-length": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.MapType{ElemType: types.StringType}, []attr.Value{}),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "flatten_map: list of maps parameter must be at least length 1"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, mapfunc.NewFlattenMapFunction(), test)
}
