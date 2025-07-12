# terraform-provider-stdlib

The Terraform Provider Plugin "stdlib" provides additional functions to enhance the capabilities of the Terraform DSL through HCL2. These additional functions will not overlap with the core functionality of Terraform. If that were ever to occur, then the corresponding function in this plugin will be gradually deprecated.

This README is purposefully short as all documentation is generated with `terraform-plugin-docs` and hosted at the [Terraform Registry](https://registry.terraform.io/providers/mschuchard/stdlib/latest/docs) as per usual.

This repository additionally does accept feature requests for e.g. additional functions or enhancements to current functions in the Github issue tracker.

If your version of Terraform is >= 1.8 then you can additionally invoke the provider custom functions instead of the data sources. Otherwise you must declare the data sources to utilize this plugin's functions.
