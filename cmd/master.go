package cmd

import (
	"fmt"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/server"
	"github.com/kataras/iris"
	"github.com/spf13/cobra"
	"log"
)

// masterCmd represents the master command
var masterCmd = &cobra.Command{
	Use:   "master",
	Short: "Run a master node service",
	Long:  "Run a master node service on this server",
	Run: func(cmd *cobra.Command, args []string) {
		start(fmt.Sprintf("%s:%d", server.Conf.Service.Host, server.Conf.Service.Port))
	},
}

func init() {
	rootCmd.AddCommand(masterCmd)
	server.Conf = config.Init()
	masterCmd.PersistentFlags().StringVar(&server.Conf.Service.Host, "host", "0.0.0.0", "Set listen on IP")
	masterCmd.PersistentFlags().IntVar(&server.Conf.Service.Port, "port", 7368, "Set listen on port")
	masterCmd.PersistentFlags().StringVar(&server.Conf.Database.Host, "db_host", "localhost", "Set mysql service host")
	masterCmd.PersistentFlags().IntVar(&server.Conf.Database.Port, "db_port", 3306, "Set mysql service host")
	masterCmd.PersistentFlags().StringVar(&server.Conf.Database.Name, "db_name", "ects", "Set mysql service db name")
	masterCmd.PersistentFlags().StringVar(&server.Conf.Database.User, "db_user", "root", "Set mysql service user")
	masterCmd.PersistentFlags().StringVar(&server.Conf.Database.Host, "db_pass", "", "Set mysql service pass")
	masterCmd.PersistentFlags().StringVar(&server.Conf.Auth.Secret, "secret", "ects", "Set JWT Secret")
	masterCmd.PersistentFlags().Int64Var(&server.Conf.Auth.TTL, "ttl", 10e9, "Set JWT TTL")
}

func start(addr string) {
	app := server.Start()
	if err := app.Run(iris.Addr(addr), iris.WithoutInterruptHandler); err != nil {
		log.Println(err)
	}
}
