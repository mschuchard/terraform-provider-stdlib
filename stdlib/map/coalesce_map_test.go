package mapfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	mapfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/map"
)

func TestCoalesceMapFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.MapUnknown(types.StringType))

	testCases := util.TestCases{
		"standard": {
			Request: function.RunRequest{
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
			Expected: function.RunResponse{
				Result: function.NewResultData(types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world")})),
			},
		},
		"no-input-args": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.TupleValueMust([]attr.Type{}, []attr.Value{})}),
			},
			Expected: function.RunResponse{
				Error:  function.NewFuncError("coalesce_map: at least one argument is required"),
				Result: resultData,
			},
		},
		"all-args-empty-maps": {
			Request: function.RunRequest{
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
			Expected: function.RunResponse{
				Error:  function.NewFuncError("coalesce_map: all arguments are empty maps"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, mapfunc.NewCoalesceMapFunction(), test)
}
