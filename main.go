package main

import (
	"v2dat/cmd"
	_ "v2dat/cmd/unpack"
)

func main() {
	_ = cmd.RootCmd.Execute()
}
