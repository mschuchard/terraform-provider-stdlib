package mapfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	mapfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/map"
)

func TestKeysDeleteFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.MapUnknown(types.StringType))

	testCases := util.TestCases{
		"present": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar"), "baz": types.StringValue("bat")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("baz")}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world")})),
			},
		},
		"absent": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar"), "baz": types.StringValue("bat")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("bar")}),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "keys_delete: the key to be deleted 'bar' does not exist in the input map"),
				Result: resultData,
			},
		},
		"keys-too-short": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.StringType, map[string]attr.Value{"hello": types.StringValue("world"), "foo": types.StringValue("bar"), "baz": types.StringValue("bat")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("bar")}),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "delete_keys: keys parameter must be at least length 2"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, mapfunc.NewKeysDeleteFunction(), test)
}
