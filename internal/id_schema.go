package util

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func IDStringAttribute() schema.StringAttribute {
	return schema.StringAttribute{
		Computed:    true,
		Description: "Aliased to string input parameter(s) for efficiency and proper plan diff detection.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}
}

func IDIntAttribute() schema.Int64Attribute {
	return schema.Int64Attribute{
		Computed:    true,
		Description: "Aliased to number input parameter(s) for efficiency and proper plan diff detection.",
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
	}
}
