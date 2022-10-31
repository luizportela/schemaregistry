package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/luizportela/kafka_schema_registry_terraform_provider/restapi"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: restapi.Provider})
}
