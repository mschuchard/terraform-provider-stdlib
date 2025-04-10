---
page_title: "stdlib Provider"
subcategory: ""
description: |-
  The stdlib provider provides additional functions for use within Terraform's HCL2 configuration language.
---

# stdlib Provider

The stdlib provider provides additional functions for use within Terraform's HCL2 configuration language.

The Terraform provider plugin "stdlib" provides additional functions for Terraform available as data sources and custom functions. These data sources and custom functions enable functionality either not intrinsically available to Terraform, or instead streamlined within a single invocation. However, data sources are not as robustly invoked with inputs or returns compared to true functions. Without the true support for custom functions in Terraform >= 1.8 then this becomes the next best available option. If you are using Terraform >= 1.8 then it is advised to use the custom functions instead of the data sources, but otherwise you will need to declare the data sources.

Use the navigation to the left to read about the available custom functions, and the alternative data sources which are each equivalent to Terraform functions.

## Example Usage

```terraform
terraform {
  required_providers {
    stdlib = {
      source  = "mschuchard/stdlib"
      version = "~> 2.0"
    }
  }
}

provider "stdlib" {} # can be omitted
```

<!-- schema generated by tfplugindocs -->
## Schema
