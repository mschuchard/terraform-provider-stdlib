package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// factory function for provider instantiation
var TestAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"stdlib": providerserver.NewProtocol6WithError(&stdlibProvider{
		version: "test",
	}),
}
