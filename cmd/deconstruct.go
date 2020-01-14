package cmd

import (
	"fmt"

	"github.com/mrueg/go-deconstruct/pkg"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deconstructCmd)
	deconstructCmd.Flags().StringP("output", "o", "", "Output path for go.sum, StdOut if unused")
}

var deconstructCmd = &cobra.Command{
	Use:   "generate binary",
	Short: "Generate go.mod file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		modFile, err := pkg.GetInfoFromBinary(args[0])
		if err != nil {
			fmt.Printf("%s", err)
		}
		outputPath, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Printf("%s", err)
		}
		pkg.WriteMod(modFile, outputPath)

	},
}
