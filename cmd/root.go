package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "ects",
	Short:   "Elastic Crontab System",
	Long:    "Elastic Crontab System",
	Version: "0.6.1",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
