package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/yzhangyext/terraform-provider-bazel/bazel"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: bazel.Provider,
	})
}
