package util

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func IDStringAttribute() schema.StringAttribute {
	return schema.StringAttribute{
		Computed:    true,
		Description: "Constant value for uniformity and efficiency. This is not actually used in the terraform plugin framework, but is still required for backwards compatibility.",
	}
}
