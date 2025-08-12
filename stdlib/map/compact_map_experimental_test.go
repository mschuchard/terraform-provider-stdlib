package mapfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	mapfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/map"
)

func TestCompactMapExperimentalFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.MapUnknown(types.DynamicType))

	testCases := util.TestCases{
		"string": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(
						types.DynamicType,
						map[string]attr.Value{
							"hello": types.DynamicValue(types.StringValue("world")),
							"foo":   types.DynamicValue(types.StringValue("")),
							"bar":   types.DynamicValue(types.StringNull()),
						},
					),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.MapValueMust(types.DynamicType, map[string]attr.Value{"hello": types.DynamicValue(types.StringValue("world"))})),
			},
		},
		"set": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(
						types.DynamicType,
						map[string]attr.Value{
							"hello": types.DynamicValue(types.SetValueMust(types.StringType, []attr.Value{types.StringValue("world")})),
							"bar":   types.DynamicValue(types.SetNull(types.StringType)),
							"baz":   types.DynamicValue(types.SetValueMust(types.StringType, []attr.Value{})),
						},
					),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.MapValueMust(types.DynamicType, map[string]attr.Value{"hello": types.DynamicValue(types.SetValueMust(types.StringType, []attr.Value{types.StringValue("world")}))})),
			},
		},
		"list": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(
						types.DynamicType,
						map[string]attr.Value{
							"hello": types.DynamicValue(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("world")})),
							"bar":   types.DynamicValue(types.ListNull(types.StringType)),
							"baz":   types.DynamicValue(types.ListValueMust(types.StringType, []attr.Value{})),
						},
					),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.MapValueMust(types.DynamicType, map[string]attr.Value{"hello": types.DynamicValue(types.ListValueMust(types.StringType, []attr.Value{types.StringValue("world")}))})),
			},
		},
		"map": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(
						types.DynamicType,
						map[string]attr.Value{
							"hello": types.DynamicValue(types.MapValueMust(types.StringType, map[string]attr.Value{"world": types.StringValue("!")})),
							"bar":   types.DynamicValue(types.MapNull(types.StringType)),
							"baz":   types.DynamicValue(types.MapValueMust(types.StringType, map[string]attr.Value{})),
						},
					),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.MapValueMust(types.DynamicType, map[string]attr.Value{"hello": types.DynamicValue(types.MapValueMust(types.StringType, map[string]attr.Value{"world": types.StringValue("!")}))})),
			},
		},
		"empty": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.MapValueMust(types.DynamicType, map[string]attr.Value{}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.MapValueMust(types.DynamicType, map[string]attr.Value{})),
			},
		},
	}

	util.UnitTests(testCases, resultData, mapfunc.NewCompactMapExperimentalFunction(), test)
}
