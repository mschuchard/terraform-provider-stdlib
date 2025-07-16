package mapfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	mapfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/map"
)

func TestHasValuesFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.BoolUnknown())

	testCases := util.TestCases{
		"any-present": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar"), "baz": types.StringValue("bat")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("bar")}),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(true)),
			},
		},
		"any-absent": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar"), "baz": types.StringValue("bat")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("pizza")}),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(false)),
			},
		},
		"all-absent": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar"), "baz": types.StringValue("bat")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("bar")}),
					types.TupleValueMust([]attr.Type{types.BoolType}, []attr.Value{types.BoolValue(true)}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(false)),
			},
		},
		"all-present": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar"), "baz": types.StringValue("bat")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("world"), types.StringValue("bar"), types.StringValue("bat")}),
					types.TupleValueMust([]attr.Type{types.BoolType}, []attr.Value{types.BoolValue(true)}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(true)),
			},
		},
		"values-too-short": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar"), "baz": types.StringValue("bat")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("bar")}),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "has_values: values parameter must be at least length 2"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, mapfunc.NewHasValuesFunction(), test)
}
