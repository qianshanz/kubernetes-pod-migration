package main

import (
	"kubernetes-pod-migration/pkg/plugin"
)

//kubectl plugin checkpoint POD_ID
func main() {
	cmd := plugin.NewPluginCmd()
	cmd.Execute()
}
