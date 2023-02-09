// Package cmd This is only the root command, a placeholder.
package cmd

import (
	"fmt"
	"sync"
	"time"

	"github.com/jjuarez/simple-prober/internal/codes"
	"github.com/jjuarez/simple-prober/internal/config"
	"github.com/jjuarez/simple-prober/internal/logger"
	"github.com/jjuarez/simple-prober/internal/model"
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

func doTests(endpoints []model.Endpoint) map[string]bool {
	start := time.Now()
	defer func() {
		log.Info(fmt.Sprintf("Execution time: %s", time.Since(start)))
	}()

	connectionResults := make(map[string]bool, len(endpoints))
	wg := sync.WaitGroup{}
	// Connection tests
	for _, e := range endpoints {
		wg.Add(1)
		go func(e model.Endpoint) {
			r, err := e.Connect(timeout)
			if err != nil {
				return
			}
			connectionResults[e.Name] = r
			log.Debug(fmt.Sprintf("Tested: %s, with result: %v", e.Name, r))
			wg.Done()
		}(e)
	}
	wg.Wait()
	return connectionResults
}

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
		results := doTests(*eList)
		log.Info(results)
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
