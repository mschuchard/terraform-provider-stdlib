---
page_title: "Provider: stdlib"
description: |-
  The stdlib provider provides additional functions for use within Terraform's HCL2 configuration language.
---

# STDLIB Provider

PLACEHOLDER

Use the navigation to the left to read about the available data sources which are each equivalent to functions.

## Example Usage

```terraform
terraform {
  required_providers {
    stdlib = {
      source  = "mschuchard/stdlib"
      version = "~> 1.0"
    }
  }
}

provider "stdlib" {}
```
