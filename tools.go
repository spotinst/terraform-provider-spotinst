//go:build tools
// +build tools

package main

//go:generate go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

import (
	_ "github.com/bflad/tfproviderlint/cmd/tfproviderlint"
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
	_ "golang.org/x/lint/golint"
)
