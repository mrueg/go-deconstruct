package main

import (
	"github.com/mrueg/go-deconstruct/cmd"
)

var (
	// VERSION is set during build
	VERSION = "unknown"
)

func main() {
	cmd.Execute(VERSION)
}
