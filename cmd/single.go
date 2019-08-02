package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var singleCmd = &cobra.Command{
	Use:   "single",
	Short: "Run a single node service",
	Long:  "Run a single node service on this server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("single called")
		//discover.NewClient()
		//config.Conf = config.Init()
		//discover.GetConf("/ects/config")
		//pipeline.WatchPipelines("7df52971-4894-4f01-9171-7452c4ddceca")
	},
}

func init() {
	rootCmd.AddCommand(singleCmd)
}
