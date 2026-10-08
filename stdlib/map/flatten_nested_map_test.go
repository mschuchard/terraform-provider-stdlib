package mapfunc_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"

	util "github.com/mschuchard/terraform-provider-stdlib/internal"
	mapfunc "github.com/mschuchard/terraform-provider-stdlib/stdlib/map"
)

func TestFlattenNestedMapFunction(test *testing.T) {
	// initialize initial result data
	resultData := function.NewResultData(types.DynamicUnknown())

	// initialize the types of the nested structure from the terraform documentation example
	subnetType := types.ObjectType{AttrTypes: map[string]attr.Type{"cidr_block": types.StringType}}
	networkType := types.ObjectType{AttrTypes: map[string]attr.Type{
		"cidr_block": types.StringType,
		"subnets":    types.MapType{ElemType: subnetType},
	}}
	returnElemType := types.ObjectType{AttrTypes: map[string]attr.Type{
		"parent_key": types.StringType,
		"child_key":  types.StringType,
		"child":      subnetType,
	}}

	// helpers to initialize input and expected values
	subnet := func(cidrBlock string) attr.Value {
		return types.ObjectValueMust(subnetType.AttrTypes, map[string]attr.Value{"cidr_block": types.StringValue(cidrBlock)})
	}
	network := func(cidrBlock string, subnets map[string]attr.Value) attr.Value {
		return types.ObjectValueMust(networkType.AttrTypes, map[string]attr.Value{
			"cidr_block": types.StringValue(cidrBlock),
			"subnets":    types.MapValueMust(subnetType, subnets),
		})
	}
	element := func(parentKey, childKey, childCidrBlock string) attr.Value {
		return types.ObjectValueMust(returnElemType.AttrTypes, map[string]attr.Value{
			"parent_key": types.StringValue(parentKey),
			"child_key":  types.StringValue(childKey),
			"child":      subnet(childCidrBlock),
		})
	}

	networks := map[string]attr.Value{
		"private": network("10.1.0.0/16", map[string]attr.Value{"db1": subnet("10.1.0.0/24"), "db2": subnet("10.1.1.0/24")}),
		"dmz":     network("10.3.0.0/16", map[string]attr.Value{"firewall": subnet("10.3.1.0/24")}),
	}
	expectedNetworks := types.ObjectValueMust(
		map[string]attr.Type{"private.db1": returnElemType, "private.db2": returnElemType, "dmz.firewall": returnElemType},
		map[string]attr.Value{
			"private.db1":  element("private", "db1", "10.1.0.0/24"),
			"private.db2":  element("private", "db2", "10.1.1.0/24"),
			"dmz.firewall": element("dmz", "firewall", "10.3.1.0/24"),
		},
	)

	testCases := util.TestCases{
		"map-input": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.DynamicValue(types.MapValueMust(networkType, networks)),
					types.StringValue("subnets"),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.DynamicValue(expectedNetworks)),
			},
		},
		"object-input": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.DynamicValue(types.ObjectValueMust(
						map[string]attr.Type{"private": networkType, "dmz": networkType},
						networks,
					)),
					types.StringValue("subnets"),
				}),
			},
			Expected: function.RunResponse{
				Result: function.NewResultData(types.DynamicValue(expectedNetworks)),
			},
		},
		"unknown-input": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.DynamicUnknown(),
					types.StringValue("subnets"),
				}),
			},
			Expected: function.RunResponse{
				Result: resultData,
			},
		},
		"input-type": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.DynamicValue(types.StringValue("not a map")),
					types.StringValue("subnets"),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "flatten_nested_map: input parameter must be a map or object"),
				Result: resultData,
			},
		},
		"missing-children-attribute": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.DynamicValue(types.MapValueMust(networkType, networks)),
					types.StringValue("hosts"),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(1, "flatten_nested_map: input element 'dmz' does not contain the attribute 'hosts'"),
				Result: resultData,
			},
		},
		"duplicate-flattened-key": {
			Request: function.RunRequest{
				Arguments: function.NewArgumentsData([]attr.Value{
					types.DynamicValue(types.MapValueMust(networkType, map[string]attr.Value{
						"a":   network("10.1.0.0/16", map[string]attr.Value{"b.c": subnet("10.1.0.0/24")}),
						"a.b": network("10.2.0.0/16", map[string]attr.Value{"c": subnet("10.2.0.0/24")}),
					})),
					types.StringValue("subnets"),
				}),
			},
			Expected: function.RunResponse{
				Error:  function.NewArgumentFuncError(0, "flatten_nested_map: flattened key 'a.b.c' is not unique because of a '.' in a parent or nested key"),
				Result: resultData,
			},
		},
	}

	util.UnitTests(testCases, resultData, mapfunc.NewFlattenNestedMapFunction(), test)
}
