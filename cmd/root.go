package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hclq",
	Short: "A jq-like tool for HCL files",
	Long:  "hclq is a command-line tool for querying and manipulating HCL files, similar to how jq works with JSON.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing hclq: %s\n", err)
		os.Exit(1)
	}
}
