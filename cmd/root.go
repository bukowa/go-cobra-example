package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/bukowa/go-cobra-example/cmd/backend"
	"github.com/bukowa/go-cobra-example/cmd/frontend"
	"strings"
	"os"
)


var LogLevel string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use: "cli",
	Long: ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		switch strings.ToLower(LogLevel) {
		case "info":
		case "debug":
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&LogLevel, "loglevel", "", "info", "log level: info,debug")
	RootCmd.AddCommand(backend.Cmd)
	RootCmd.AddCommand(frontend.Cmd)
}
