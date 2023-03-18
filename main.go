package main

import (
  "context"

  "github.com/hashicorp/terraform-plugin-framework/providerserver"

  "github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

// provider documentation generation.
//go:generate tfplugindocs generate --provider-name stdlib

func main() {
  providerserver.Serve(context.Background(), stdlib.New, providerserver.ServeOpts{
    Address: "registry.terraform.io/mschuchard/stdlib",
  })
}
