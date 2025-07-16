package slicefunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	slicefunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/slice"
)

func TestCompareListFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.Int32Unknown())

	testCases := util.TestCases{
		"lesser": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("bar"), types.StringValue("b")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("foo"), types.StringValue("bar"), types.StringValue("baz")}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(-1)),
			},
		},
		"equal": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("pizza"), types.StringValue("cake")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("pizza"), types.StringValue("cake")}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(0)),
			},
		},
		"greater": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("super"), types.StringValue("hyper"), types.StringValue("turbo")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("pizza"), types.StringValue("cake"), types.StringValue("punch")}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(1)),
			},
		},
		"length": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("pizza"), types.StringValue("cake"), types.StringValue("punch")}),
					types.ListValueMust(types.StringType, []attr.Value{types.StringValue("pizza"), types.StringValue("cake")}),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.Int32Value(1)),
			},
		},
	}

	util.UnitTests(testCases, resultData, slicefunc.NewCompareListFunction(), test)
}
