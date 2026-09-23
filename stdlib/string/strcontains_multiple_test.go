package stringfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	stringfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/string"
)

func TestStrcontainsMultipleFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.BoolUnknown())

	testCases := util.TestCases{
		"empty-substrings": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("foobarbazbot"), types.ListValueMust(types.StringType, []attr.Value{}), types.TupleValueMust([]attr.Type{}, []attr.Value{})}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(false)),
			},
		},
		"any-true": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("foobarbazbot"), types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("pizza")}), types.TupleValueMust([]attr.Type{}, []attr.Value{})}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(true)),
			},
		},
		"any-false": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("foobarbazbot"), types.ListValueMust(types.StringType, []attr.Value{types.StringValue("pizza"), types.StringValue("cake")}), types.TupleValueMust([]attr.Type{}, []attr.Value{})}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(false)),
			},
		},
		"all-false": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("foobarbazbot"), types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("pizza")}), types.TupleValueMust([]attr.Type{types.BoolType}, []attr.Value{types.BoolValue(true)})}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(false)),
			},
		},
		"all-true": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{types.StringValue("foobarbazbot"), types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("baz")}), types.TupleValueMust([]attr.Type{types.BoolType}, []attr.Value{types.BoolValue(true)})}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.BoolValue(true)),
			},
		},
	}

	util.UnitTests(testCases, resultData, stringfunc.NewStrcontainsMultipleFunction(), test)
}
