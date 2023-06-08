package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/mschuchard/terraform-provider-stdlib/stdlib"
)

// provider documentation generation.
//go:generate tfplugindocs generate --provider-name stdlib

func main() {
	const version string = "1.0.0"

	// start provider server
	err := providerserver.Serve(context.Background(), stdlib.New(version), providerserver.ServeOpts{
		Address: "registry.terraform.io/mschuchard/stdlib",
	})

	if err != nil {
		log.Fatal(err.Error())
	}
}
