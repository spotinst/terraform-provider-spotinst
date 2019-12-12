package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/terraform-providers/terraform-provider-spotinst/spotinst"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: spotinst.Provider})
}
