package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var singleCmd = &cobra.Command{
	Use:   "single",
	Short: "Run a single node service",
	Long: "Run a single node service on this server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("single called")
	},
}

func init() {
	rootCmd.AddCommand(singleCmd)
}
