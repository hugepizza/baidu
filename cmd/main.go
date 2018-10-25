package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	debug   bool
	rootCmd = cobra.Command{
		Use: "baidu-cli",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if debug {
				logrus.SetLevel(logrus.DebugLevel)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func main() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "D", false, "log debug message.")

	rootCmd.Execute()
}
