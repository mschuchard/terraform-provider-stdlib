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
func GetDynamicUnderlyingValue(dynamicType types.Dynamic, ctx context.Context) (attr.Value, bool, bool) {
	// initialize unknown and null vars
	var unknown, null bool

	// ascertain if dynamic type is unknown...
	if dynamicType.IsUnknown() || dynamicType.IsUnderlyingValueUnknown() {
		tflog.Info(ctx, fmt.Sprintf("GetDynamicUnderlyingValue (helper): input parameter '%s' is unknown, or was refined by terraform to an unknown", dynamicType.String()))
		unknown = true
	} else if dynamicType.IsNull() || dynamicType.IsUnderlyingValueNull() { // ...or null
		tflog.Info(ctx, fmt.Sprintf("GetDynamicUnderlyingValue (helper): input parameter '%s' is null, or was refined by terraform to a specific underlying type (but value is still null)", dynamicType.String()))
		null = true
	}

	// access underlying value of dynamic type parameter, and return unknown/null status
	return dynamicType.UnderlyingValue(), unknown, null
}

// checks if a tf dynamic type underlying value is empty
func IsDynamicEmpty(value attr.Value, position int64, ctx context.Context) (bool, *function.FuncError) {
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
		return false, function.NewArgumentFuncError(position, "IsDynamicEmpty (helper): invalid input parameter type")
	}
}
