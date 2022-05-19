package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/spotinst/terraform-provider-spotinst/spotinst"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: spotinst.Provider})
}
