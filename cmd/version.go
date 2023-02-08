/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ToolName The external name for the tool.
const ToolName string = "simple-prober"

// Version The version of the tool.
var Version string = "v0.0.0*unknown"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version",
	Long:  `This command shows the version of the tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%v version %v", ToolName, Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
