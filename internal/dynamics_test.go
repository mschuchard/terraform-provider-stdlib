package util_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	util "github.com/mschuchard/terraform-provider-stdlib/internal"
)

func TestGetDynamicUnderlyingValue(test *testing.T) {
	ctx := context.Background()

	dynamicValue := types.DynamicValue(types.StringValue("test"))
	value, unknown, null := util.GetDynamicUnderlyingValue(dynamicValue, ctx)
	if unknown {
		test.Error("expected unknown to be false")
	}
	if null {
		test.Error("expected null to be false")
	}
	if value == nil || !value.Equal(types.StringValue("test")) {
		test.Error("expected value to not be nil")
	}

	dynamicValue = types.DynamicUnknown()
	_, unknown, null = util.GetDynamicUnderlyingValue(dynamicValue, ctx)
	if !unknown {
		test.Error("expected unknown to be true")
	}
	if null {
		test.Error("expected null to be false")
	}

	dynamicValue = types.DynamicNull()
	_, unknown, null = util.GetDynamicUnderlyingValue(dynamicValue, ctx)
	if unknown {
		test.Error("expected unknown to be false")
	}
	if !null {
		test.Error("expected null to be true")
	}
}

func TestIsDynamicEmpty(test *testing.T) {
	ctx := context.Background()

	isEmpty, err := util.IsDynamicEmpty(types.StringValue(""), ctx)
	if err != nil {
		test.Fatalf("unexpected error: %s", err)
	}
	if !isEmpty {
		test.Error("expected empty string to return true")
	}

	isEmpty, err = util.IsDynamicEmpty(types.StringValue("test"), ctx)
	if err != nil {
		test.Fatalf("unexpected error: %s", err)
	}
	if isEmpty {
		test.Error("expected non-empty string to return false")
	}

	emptyList := types.ListValueMust(types.StringType, []attr.Value{})
	isEmpty, err = util.IsDynamicEmpty(emptyList, ctx)
	if err != nil {
		test.Fatalf("unexpected error: %s", err)
	}
	if !isEmpty {
		test.Error("expected empty list to return true")
	}

	nonEmptyList := types.ListValueMust(types.StringType, []attr.Value{types.StringValue("item")})
	isEmpty, err = util.IsDynamicEmpty(nonEmptyList, ctx)
	if err != nil {
		test.Fatalf("unexpected error: %s", err)
	}
	if isEmpty {
		test.Error("expected non-empty list to return false")
	}

	emptySet := types.SetValueMust(types.StringType, []attr.Value{})
	isEmpty, err = util.IsDynamicEmpty(emptySet, ctx)
	if err != nil {
		test.Fatalf("unexpected error: %s", err)
	}
	if !isEmpty {
		test.Error("expected empty set to return true")
	}

	nonEmptySet := types.SetValueMust(types.StringType, []attr.Value{types.StringValue("item")})
	isEmpty, err = util.IsDynamicEmpty(nonEmptySet, ctx)
	if err != nil {
		test.Fatalf("unexpected error: %s", err)
	}
	if isEmpty {
		test.Error("expected non-empty set to return false")
	}

	emptyMap := types.MapValueMust(types.StringType, map[string]attr.Value{})
	isEmpty, err = util.IsDynamicEmpty(emptyMap, ctx)
	if err != nil {
		test.Fatalf("unexpected error: %s", err)
	}
	if !isEmpty {
		test.Error("expected empty map to return true")
	}

	nonEmptyMap := types.MapValueMust(types.StringType, map[string]attr.Value{"key": types.StringValue("value")})
	isEmpty, err = util.IsDynamicEmpty(nonEmptyMap, ctx)
	if err != nil {
		test.Fatalf("unexpected error: %s", err)
	}
	if isEmpty {
		test.Error("expected non-empty map to return false")
	}

	_, err = util.IsDynamicEmpty(types.BoolValue(true), ctx)
	if err == nil || err.Error() != "IsDynamicEmpty (helper): invalid input parameter type" {
		if err == nil {
			test.Error("expected error for invalid type")
		} else {
			test.Errorf("unexpected error message: %s", err)
		}
	}
}
