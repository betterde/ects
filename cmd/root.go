package cmd

import (
	"fmt"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/build"
	"github.com/betterde/ects/internal/journal"
	"github.com/betterde/ects/internal/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var (
	cfgFile string
	verbose bool
	prefix  = strings.ToUpper(build.Name)
	rootCmd = &cobra.Command{
		Use:     build.Name,
		Short:   build.Desc,
		Version: fmt.Sprintf("Version: %s\nBuild at: %s\nCommit hash: %s", build.Version, build.Build, build.Commit),
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Init HTTP server
	server.InitHttpServer(build.Name, rootCmd.Version)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ects.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Verbose flag
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose mode")

	// Print app version
	rootCmd.SetVersionTemplate("{{printf \"%s\" .Version}}\n")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Initialize the logger using development environment.
	journal.InitLogger()

	// Parse config from file and env variables
	config.Parse(cfgFile, prefix)

	level := viper.GetString("logging.level")
	if verbose {
		level = "DEBUG"
	}

	err := journal.SetLevel(level)
	if err != nil {
		journal.Logger.Error("Unable to set logger level", err)
		os.Exit(1)
	}

	journal.Logger.Debugf("Configuration file currently in use: %s", viper.ConfigFileUsed())
}
