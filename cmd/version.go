// Package cmd This is only the root command, a placeholder.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version ...
	Version string = "v0.0.0+unknown"
)

// versionCmd shows the tool version
var versionCmd = &cobra.Command{
	Use:     "version",
	Version: Version,
	Short:   "Shows the tool version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version: %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
