package stdlib

import (
  "github.com/hashicorp/terraform-plugin-framework/providerserver"
  "github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const (
  // reusable provider config
  providerConfig = `
provider "stdlib" {}
`
)

var (
  // factory function for provider instantiation
  testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
    "stdlib": providerserver.NewProtocol6WithError(New()),
  }
)
