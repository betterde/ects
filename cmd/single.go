package cmd

import (
	"fmt"
	"github.com/betterde/ects/internal/discover"
	"github.com/betterde/ects/internal/pipeline"

	"github.com/spf13/cobra"
)

var singleCmd = &cobra.Command{
	Use:   "single",
	Short: "Run a single node service",
	Long: "Run a single node service on this server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("single called")
		discover.NewClient()
		discover.GetConf("/ects/config")
		pipeline.WatchPipelines("7df52971-4894-4f01-9171-7452c4ddceca")
	},
}

func init() {
	rootCmd.AddCommand(singleCmd)
}
