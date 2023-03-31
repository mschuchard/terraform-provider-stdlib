package utils

import (
  "github.com/hashicorp/terraform-plugin-framework/resource/schema"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
  "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func IDStringAttribute() schema.StringAttribute {
  return schema.StringAttribute{
    Computed: true,
    Description: "Aliased to string input parameter for efficiency.",
    PlanModifiers: []planmodifier.String{
      stringplanmodifier.UseStateForUnknown(),
    },
  }
}
