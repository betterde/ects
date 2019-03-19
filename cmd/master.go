package cmd

import (
	"fmt"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/routes"
	"github.com/betterde/ects/internal/utils/system"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/spf13/cobra"
	"log"
	"path"
)

// masterCmd represents the master command
var masterCmd = &cobra.Command{
	Use:   "master",
	Short: "Run a master node service",
	Long:  "Run a master node service on this server",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap()
		start(fmt.Sprintf("%s:%d", config.Conf.Service.Host, config.Conf.Service.Port))
	},
}

func init() {
	rootCmd.AddCommand(masterCmd)
	config.Conf = config.Init()
	masterCmd.PersistentFlags().StringVarP(&config.Path, "config", "c", "/etc/ects/ects.yaml", "Set configuration file")
	masterCmd.PersistentFlags().StringVar(&config.Conf.Service.Host, "host", "0.0.0.0", "Set listen on IP")
	masterCmd.PersistentFlags().IntVar(&config.Conf.Service.Port, "port", 9701, "Set listen on port")
}

func bootstrap() {
	var err error
	// 判断是否已经安装
	system.Info = &system.Information{
		Version: rootCmd.Version,
	}
	system.Info.Installed, err = config.CheckConfigFile(config.Path)
	if err != nil {
		// TODO
	}
	dir := path.Dir(config.Path)
	config.CreateConfigDir(dir)
	system.Info.Permission = config.CheckConfigDirPermisson(dir)
}

func start(addr string) {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	// 注册路由
	routes.Register(app)

	if err := app.Run(iris.Addr(addr), iris.WithoutInterruptHandler); err != nil {
		log.Println(err)
	}
}
