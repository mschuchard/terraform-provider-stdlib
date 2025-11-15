package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	provider "github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

// provider documentation generation.
//go:generate tfplugindocs generate --provider-name stdlib

func main() {
	// start provider server
	if err := providerserver.Serve(context.Background(), provider.New("2.2.1"), providerserver.ServeOpts{
		Address: "registry.terraform.io/mschuchard/stdlib",
	}); err != nil {
		log.Fatal(err)
	}
}
