// Package cmd This is only the root command, a placeholder.
package cmd

/*
Copyright © 2023 JJ <javier.juarez@gmail.com>
*/

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "simple-prober",
	Short: "A very basic TPC prober",
	Long:  `This utility will allow you to test several TCP endpoints provided by a configuration file.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
