package util

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func IDStringAttribute() schema.StringAttribute {
	return schema.StringAttribute{
		Computed:    true,
		Description: "Aliased to string input parameter(s) for efficiency and proper plan diff detection.",
	}
}

func IDInt64Attribute() schema.Int64Attribute {
	return schema.Int64Attribute{
		Computed:    true,
		Description: "Aliased to number input parameter(s) for efficiency and proper plan diff detection.",
	}
}

func IDFloat64Attribute() schema.Float64Attribute {
	return schema.Float64Attribute{
		Computed:    true,
		Description: "Aliased to number input parameter(s) for efficiency and proper plan diff detection.",
	}
}
