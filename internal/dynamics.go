package util

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// checks if a tf dynamic type underlying value is null or empty
func IsDynamicEmpty(dynamicType types.Dynamic, ctx context.Context) (bool, *function.FuncError) {
	// access underlying value of dynamic type parameter
	value := dynamicType.UnderlyingValue()

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
		tflog.Error(ctx, fmt.Sprintf("IsDynamicEmpty (helper): could not convert input parameter '%s' to an acceptable terraform type", dynamicType.String()))
		return false, function.NewArgumentFuncError(0, "IsDynamicEntry (helper): invalid input parameter type")
	}
}
