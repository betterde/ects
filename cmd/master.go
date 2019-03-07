package cmd

import (
	"fmt"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/models"
	"github.com/betterde/ects/internal/routes"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/spf13/cobra"
	"log"
)

// masterCmd represents the master command
var masterCmd = &cobra.Command{
	Use:   "master",
	Short: "Run a master node service",
	Long:  "Run a master node service on this server",
	Run: func(cmd *cobra.Command, args []string) {
		start(fmt.Sprintf("%s:%d", config.Conf.Service.Host, config.Conf.Service.Port))
	},
}

func init() {
	rootCmd.AddCommand(masterCmd)
	config.Conf = config.Init()
	masterCmd.PersistentFlags().StringVar(&config.Conf.Service.Host, "host", "0.0.0.0", "Set listen on IP")
	masterCmd.PersistentFlags().IntVar(&config.Conf.Service.Port, "port", 7368, "Set listen on port")
	masterCmd.PersistentFlags().StringVar(&config.Conf.Database.Host, "db_host", "localhost", "Set mysql service host")
	masterCmd.PersistentFlags().IntVar(&config.Conf.Database.Port, "db_port", 3306, "Set mysql service host")
	masterCmd.PersistentFlags().StringVar(&config.Conf.Database.Name, "db_name", "ects", "Set mysql service db name")
	masterCmd.PersistentFlags().StringVar(&config.Conf.Database.User, "db_user", "root", "Set mysql service user")
	masterCmd.PersistentFlags().StringVar(&config.Conf.Database.Pass, "db_pass", "George@1994", "Set mysql service pass")
	masterCmd.PersistentFlags().StringVar(&config.Conf.Database.Char, "db_char", "utf8mb4", "Set mysql character")
	masterCmd.PersistentFlags().StringVar(&config.Conf.Auth.Secret, "secret", "ects", "Set JWT Secret")
	masterCmd.PersistentFlags().DurationVar(&config.Conf.Auth.TTL, "ttl", 86400, "Set JWT TTL")
}

func start(addr string) {
	bootstrap()
	app := iris.New()
	app.Use(iris.Gzip)
	app.Use(recover.New())
	app.Use(logger.New())

	// 注册路由
	routes.Register(app)

	if err := app.Run(iris.Addr(addr), iris.WithoutInterruptHandler); err != nil {
		log.Println(err)
	}
}

func bootstrap() {
	models.Engine = models.Connection()
	if err := models.Migrate(); err != nil {
		log.Println(err)
	}
}
