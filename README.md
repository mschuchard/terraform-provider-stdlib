# terraform-provider-stdlib

The Terraform Provider Plugin "stdlib" provides additional functions to enhance the capabilities of the Terraform DSL through HCL2. These additional functions will not overlap with the core functionality of Terraform. If that were ever to occur, then the corresponding function in this plugin will be gradually deprecated.

This README is purposefully short as all documentation is generated with `terraform-plugin-docs` and hosted at the [Terraform Registry](https://registry.terraform.io/providers/mschuchard/stdlib/latest/docs) as per usual.

This repository additionally does accept feature requests for e.g. additional functions or enhancements to current functions in the Github issue tracker.

### Upcoming 2.0.0 Release Announcement
- All functions as of version 1.6.0 will be re-implemented as custom provider functions.
- All new functions implemented after the release of version 2.0.0 will be custom provider functions only, and not data sources.
- All data source functions that exist at the time of the release of version 1.6.0 will be maintained afterwards for any necessary bug fixes.
- Please upgrade to Terraform version >= 1.8 by the release of version 2.1.0 to ensure support for any new functions supported by this provider.
