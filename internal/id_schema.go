package util

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
