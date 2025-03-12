package cmd

import (
	"fmt"
	"github.com/betterde/ects/internal/build"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     build.Name,
	Short:   build.Desc,
	Version: fmt.Sprintf("Version: %s\nBuild at: %s\nCommit hash: %s", build.Version, build.Build, build.Commit),
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Print app version
	rootCmd.SetVersionTemplate("{{printf \"%s\" .Version}}\n")
}
