package mapfunc

import (
	"context"
	"fmt"
	"maps"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	util "github.com/mschuchard/terraform-provider-stdlib/internal"
)

// ensure the implementation satisfies the expected interfaces
var _ function.Function = &flattenNestedMapFunction{}

// helper pseudo-constructor to simplify provider server and testing implementation
func NewFlattenNestedMapFunction() function.Function {
	return &flattenNestedMapFunction{}
}

// function implementation
type flattenNestedMapFunction struct{}

// function metadata
func (*flattenNestedMapFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "flatten_nested_map"
}

// define the provider-level definition for the function
func (*flattenNestedMapFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Flatten a map of objects that each contain a nested map into a single-level map.",
		MarkdownDescription: "Return a single-level map with one element for each entry of each nested map in the input. Each result key is the parent key and the nested key joined by `.`. Each result element is an object with the attributes `parent_key`, `child_key`, and `child` (the nested entry). Attributes of the parent object can be retrieved by indexing the input with `parent_key`. This function replaces the standard nested map flattening algorithm (`flatten` with nested `for` expressions and `tomap`) for use with the `for_each` meta-parameter.",
		Parameters: []function.Parameter{
			function.DynamicParameter{
				Name:        "input",
				Description: "Input map (or object) of objects to flatten. Each object must contain the nested map (or object) attribute named by the `children_attribute` parameter.",
			},
			function.StringParameter{
				Name:        "children_attribute",
				Description: "Name of the attribute in each input object that contains the nested map to flatten.",
			},
		},
		Return: function.DynamicReturn{},
	}
}

// retrieve the key-value entries from a map or object value, because terraform supplies either type for nested structures
// a variable declared with a map type arrives as a map, whereas an object literal (or a value with differing element types) arrives as an object
// a second return value of false indicates that the value is neither a map nor an object
func nestedEntries(value attr.Value) (map[string]attr.Value, bool) {
	if mapValue, ok := value.(types.Map); ok {
		return mapValue.Elements(), true
	}
	if objectValue, ok := value.(types.Object); ok {
		return objectValue.Attributes(), true
	}

	return nil, false
}

// run the function logic: validate the input structure, and then flatten each nested entry into a single-level object
func (*flattenNestedMapFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// initialize input and children attribute from input parameters
	var input types.Dynamic
	var childrenAttribute string

	resp.Error = req.Arguments.Get(ctx, &input, &childrenAttribute)
	if resp.Error != nil {
		return
	}

	// log the parameters for debugging
	ctx = tflog.SetField(ctx, "flatten_nested_map: input", input.String())
	ctx = tflog.SetField(ctx, "flatten_nested_map: children attribute", childrenAttribute)

	// retrieve underlying dynamic value; the concrete type is only known at runtime, so the structure is validated manually in the remainder of the function
	inputValue, unknown, null := util.GetDynamicUnderlyingValue(ctx, input)
	if null {
		resp.Error = function.NewArgumentFuncError(0, "flatten_nested_map: input parameter is null")
		return
	}
	// the flattened keys cannot be determined from an unknown input, so the result is also unknown
	if unknown {
		resp.Error = resp.Result.Set(ctx, types.DynamicUnknown())
		return
	}

	// retrieve the parent entries (the top-level keys of the input)
	parents, ok := nestedEntries(inputValue)
	if !ok {
		resp.Error = function.NewArgumentFuncError(0, "flatten_nested_map: input parameter must be a map or object")
		return
	}

	// the type and value of each flattened element are accumulated separately, because the result object requires both and is keyed by the flattened key
	// iterate through the parents in sorted order for deterministic error reporting, because go randomizes map iteration order
	flatTypes := map[string]attr.Type{}
	flatValues := map[string]attr.Value{}
	for _, parentKey := range slices.Sorted(maps.Keys(parents)) {
		// validate the parent
		parent, ok := parents[parentKey].(types.Object)
		if !ok {
			resp.Error = function.NewArgumentFuncError(0, fmt.Sprintf("flatten_nested_map: input element '%s' must be an object", parentKey))
			return
		}
		// the nested keys of an unknown parent cannot be determined, so the result is also unknown
		// otherwise the plan would fail on a value that is only known after apply
		if parent.IsUnknown() {
			resp.Error = resp.Result.Set(ctx, types.DynamicUnknown())
			return
		}
		if parent.IsNull() {
			resp.Error = function.NewArgumentFuncError(0, fmt.Sprintf("flatten_nested_map: input element '%s' is null", parentKey))
			return
		}

		// retrieve the nested map from the parent; the validation of the attribute name occurs for each parent, and not before the iteration begins
		nested, exists := parent.Attributes()[childrenAttribute]
		if !exists {
			resp.Error = function.NewArgumentFuncError(1, fmt.Sprintf("flatten_nested_map: input element '%s' does not contain the attribute '%s'", parentKey, childrenAttribute))
			return
		}
		// the flattened keys depend on the keys of the nested map, so an unknown nested map also results in an unknown result
		if nested.IsUnknown() {
			resp.Error = resp.Result.Set(ctx, types.DynamicUnknown())
			return
		}
		if nested.IsNull() {
			resp.Error = function.NewArgumentFuncError(0, fmt.Sprintf("flatten_nested_map: attribute '%s' of input element '%s' is null", childrenAttribute, parentKey))
			return
		}
		// the nested structure is also allowed to be either a map or an object
		children, ok := nestedEntries(nested)
		if !ok {
			resp.Error = function.NewArgumentFuncError(0, fmt.Sprintf("flatten_nested_map: attribute '%s' of input element '%s' must be a map or object", childrenAttribute, parentKey))
			return
		}

		// create one flattened element for each nested entry
		for _, childKey := range slices.Sorted(maps.Keys(children)) {
			// the separator is always a period, which matches the key format of the documentation example (network key and subnet key joined by a period)
			flatKey := parentKey + "." + childKey

			// distinct parent and nested key pairs can still join to the same flattened key
			if _, duplicate := flatValues[flatKey]; duplicate {
				resp.Error = function.NewArgumentFuncError(0, fmt.Sprintf("flatten_nested_map: flattened key '%s' is not unique because of a '.' in a parent or nested key", flatKey))
				return
			}

			// each element retains the parent key and the nested key, so that the consuming resource can look up the parent (e.g. `aws_vpc.example[each.value.parent_key]`) without a separate local value
			// the nested entry is stored under its own attribute, so that its attribute names cannot collide with the key attribute names
			elementTypes := map[string]attr.Type{
				"parent_key": types.StringType,
				"child_key":  types.StringType,
				"child":      children[childKey].Type(ctx),
			}
			element, diags := types.ObjectValue(elementTypes, map[string]attr.Value{
				"parent_key": types.StringValue(parentKey),
				"child_key":  types.StringValue(childKey),
				"child":      children[childKey],
			})
			if diags.HasError() {
				resp.Error = function.FuncErrorFromDiags(ctx, diags)
				return
			}

			flatTypes[flatKey] = element.Type(ctx)
			flatValues[flatKey] = element
		}
	}

	// the result is an object, because the element types are not required to be identical
	// terraform accepts an object for the `for_each` meta-parameter in addition to a map
	flattened, diags := types.ObjectValue(flatTypes, flatValues)
	if diags.HasError() {
		resp.Error = function.FuncErrorFromDiags(ctx, diags)
		return
	}

	// store the result as a dynamic value, because the return type is dynamic
	result := types.DynamicValue(flattened)
	resp.Error = resp.Result.Set(ctx, &result)
	if resp.Error != nil {
		return
	}

	// log the result for debugging
	ctx = tflog.SetField(ctx, "flatten_nested_map: result", result)
	tflog.Debug(ctx, "flatten_nested_map: successful return", map[string]any{"success": true})
}
