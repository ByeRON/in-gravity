package main

import (
	"in-gravity/cmd"
	"os"
)

func main() {
	switch os.Args[1] {
	case "serve":
		cmd.Serve(os.Args[2:])
	default:
		os.Exit(1)
	}
}
