// Package cmd This is only the root command, a placeholder.
package cmd

import (
	"github.com/jjuarez/simple-prober/internal/codes"
	"github.com/jjuarez/simple-prober/internal/utils"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "simple-prober",
	Short: "simple-prober - a really simple TPC tester",
	Long:  `This utility will allow you to test several TCP endpoints provided by a configuration file.`,
	Run: func(cmd *cobra.Command, arg []string) {
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		utils.ExitCommand(codes.UnknownError, err)
	}
}

func init() {
}
