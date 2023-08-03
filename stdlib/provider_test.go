package stdlib

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// reusable provider config
const providerConfig = `provider "stdlib" {}` + "\n"

// factory function for provider instantiation
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"stdlib": providerserver.NewProtocol6WithError(NewStruct("test")),
}
