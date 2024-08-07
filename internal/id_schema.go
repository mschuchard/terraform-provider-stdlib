package util

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
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

func IDInt64Attribute() schema.Int64Attribute {
	return schema.Int64Attribute{
		Computed:    true,
		Description: "Aliased to number input parameter(s) for efficiency and proper plan diff detection.",
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
	}
}

func IDFloat64Attribute() schema.Float64Attribute {
	return schema.Float64Attribute{
		Computed:    true,
		Description: "Aliased to number input parameter(s) for efficiency and proper plan diff detection.",
		PlanModifiers: []planmodifier.Float64{
			float64planmodifier.UseStateForUnknown(),
		},
	}
}
