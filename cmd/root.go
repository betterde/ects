package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "ects",
	Short:   "Elastic Crontab System",
	Long:    "Elastic Crontab System",
	Version: "0.1.5",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
