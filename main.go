package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/MaxAnderson95/terraform-provider-windns/windns"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: windns.Provider,
	})
}
