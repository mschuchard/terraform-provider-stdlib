package util

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// validates and returns dynamic type underlying value
func GetDynamicUnderlyingValue(dynamicType types.Dynamic, ctx context.Context) (attr.Value, *function.FuncError) {
	// ascertain parameter was not refined to a specific value type
	if dynamicType.IsUnderlyingValueNull() || dynamicType.IsUnderlyingValueUnknown() {
		tflog.Error(ctx, fmt.Sprintf("GetDynamicUnderlyingValue: input parameter '%s' was refined by terraform to a specific underlying value type, and this prevents further usage", dynamicType.String()))
		return nil, function.NewArgumentFuncError(0, "GetDynamicUnderlyingValue (helper): underlying value type refined")
	}

	// access underlying value of dynamic type parameter
	return dynamicType.UnderlyingValue(), nil
}

// checks if a tf dynamic type underlying value is empty
func IsDynamicEmpty(value attr.Value, ctx context.Context) (bool, *function.FuncError) {
	// convert to one of four acceptable types
	// string
	if stringType, ok := value.(types.String); ok {
		// emptiness check
		return len(stringType.ValueString()) == 0, nil
	} else if set, ok := value.Type(ctx).(types.SetType); ok { // set
		// emptiness check
		return value.Equal(types.SetValueMust(set.ElementType(), []attr.Value{})), nil
	} else if list, ok := value.Type(ctx).(types.ListType); ok { // list
		// emptiness check
		return value.Equal(types.ListValueMust(list.ElementType(), []attr.Value{})), nil
	} else if mapType, ok := value.Type(ctx).(types.MapType); ok { // map
		// emptiness check
		return value.Equal(types.MapValueMust(mapType.ElementType(), map[string]attr.Value{})), nil
	} else {
		tflog.Error(ctx, fmt.Sprintf("IsDynamicEmpty (helper): could not convert input parameter '%s' to an acceptable terraform type", value.String()))
		return false, function.NewArgumentFuncError(0, "IsDynamicEmpty (helper): invalid input parameter type")
	}
}
