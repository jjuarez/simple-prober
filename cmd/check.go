// Package cmd This is only the root command, a placeholder.
package cmd

import (
	"fmt"
	"time"

	"github.com/jjuarez/simple-prober/internal/codes"
	"github.com/jjuarez/simple-prober/internal/config"
	"github.com/jjuarez/simple-prober/internal/logger"
	"github.com/jjuarez/simple-prober/internal/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	defaultConfigFileName string = "./config/endpoints.yaml"
	defaultTimeout        int    = 5
	defaultLogLevel       string = "info"
)

var (
	configFileName string
	timeout        time.Duration
	logLevel       string
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Checks for all TCP endpoints",
	Long: `This command will check for all the endpoints located in the configuration file, here's an usage example:
    simple-prober check --config config/endpoints.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		// Logger setup
		if err := logger.Setup(logLevel); err != nil {
			utils.ExitCommand(codes.ConfigurationError, err)
		}
		log.Debug(fmt.Sprintf("Logger setup as: %q", logLevel))

		// Configuration load
		eList, err := config.Load(configFileName)
		if err != nil {
			utils.ExitCommand(codes.ReadError, err)
		}

		// Connection tests
		for _, e := range *eList {
			r, err := e.TestConnection(timeout)
			if err != nil {
				log.Error(err)
			}
			log.Info(fmt.Sprintf("%v, %v\n", e, r))
		}
	},
}

func init() {
	var parameterTimeout int

	rootCmd.AddCommand(checkCmd)

	// The log level
	rootCmd.PersistentFlags().StringVar(&logLevel, "loglevel", defaultLogLevel, "The log level")

	// The confiuration file name
	rootCmd.PersistentFlags().StringVar(&configFileName, "config", defaultConfigFileName, "Config file (default is: config/endpoints.yaml)")

	// The connection timeout
	rootCmd.PersistentFlags().IntVar(&parameterTimeout, "timeout", defaultTimeout, "Connection timeout")
	timeout = time.Duration(parameterTimeout) * time.Second
}
