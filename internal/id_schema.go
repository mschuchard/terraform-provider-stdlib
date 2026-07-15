package util

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func IDStringAttribute() schema.StringAttribute {
	return schema.StringAttribute{
		Computed:    true,
		Description: "Constant value for efficiency. This is not used in plugin framework.",
	}
}

func IDInt64Attribute() schema.Int64Attribute {
	return schema.Int64Attribute{
		Computed:    true,
		Description: "Constant value for efficiency. This is not used in plugin framework.",
	}
}

func IDFloat64Attribute() schema.Float64Attribute {
	return schema.Float64Attribute{
		Computed:    true,
		Description: "Constant value for efficiency. This is not used in plugin framework.",
	}
}
