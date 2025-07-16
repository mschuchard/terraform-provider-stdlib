package mapfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	mapfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/map"
)

func TestHasKeysFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.BoolUnknown())

	testCases := util.TestCases{
		"any-present": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar"), "baz": types.StringValue("bat")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("bar"), types.StringValue("foo")}),
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
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("bar"), types.StringValue("pizza")}),
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
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("bar"), types.StringValue("foo")}),
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
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("hello"), types.StringValue("foo"), types.StringValue("baz")}),
					types.TupleValueMust([]attr.Type{types.BoolType}, []attr.Value{types.BoolValue(true)}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(true)),
			},
		},
		"keys-too-short": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar"), "baz": types.StringValue("bat")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("bar")}),
					types.TupleValueMust([]attr.Type{}, []attr.Value{}),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "has_keys: keys parameter must be at least length 2"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, mapfunc.NewHasKeysFunction(), test)
}
