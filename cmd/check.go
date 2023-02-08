// Package cmd This is only the root command, a placeholder.
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/jjuarez/simple-prober/internal/codes"
	"github.com/jjuarez/simple-prober/internal/config"
	"github.com/jjuarez/simple-prober/internal/utils"
	"github.com/spf13/cobra"
)

const (
	defaultConfigFileName string = "./config/endpoints.yaml"
	defaultTimeout        int    = 5
)

var (
	configFileName string
	timeout        time.Duration
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Checks for all TCP endpoints",
	Long: `This command will check for all the endpoints located in the configuration file, here's an usage example:
    simple-prober check --config config/endpoints.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		eList, err := config.Load(configFileName)
		if err != nil {
			utils.ExitCommand(codes.ReadError, err)
		}
		for _, e := range *eList {
			r, err := e.TestConnection(timeout)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%v, %v\n", e, r)
		}
	},
}

func init() {
	var parameterTimeout int

	rootCmd.AddCommand(checkCmd)
	rootCmd.PersistentFlags().StringVar(&configFileName, "config", defaultConfigFileName, "Config file (default is: config/endpoints.yaml)")
	rootCmd.PersistentFlags().IntVar(&parameterTimeout, "timeout", defaultTimeout, "Connection timeout")
	timeout = time.Duration(parameterTimeout) * time.Second
}
