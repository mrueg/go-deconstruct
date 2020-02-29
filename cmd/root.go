package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// VERSION is set during build
	VERSION string
)

var rootCmd = &cobra.Command{
	Use:   "go-deconstruct",
	Short: "go-deconstruct is a tool to generate go.mod and go.sum from a binary",
	Long:  `A tool to reconstruct go.mod and go.sum from a go binary built with modules`,
}

func Execute(version string) {
	VERSION = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
