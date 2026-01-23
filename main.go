package main

import (
	"golanglearn/cli"
)

func main() {
	rc := cli.NewRootCmd()
	err := rc.Cmd.Execute()
	if err != nil {
		// Do Stuff Here
	}
}
