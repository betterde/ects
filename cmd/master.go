package cmd

import (
	"context"
	"fmt"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/discover"
	"github.com/betterde/ects/internal/system"
	"github.com/betterde/ects/models"
	"github.com/betterde/ects/routes"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// masterCmd represents the master command
var (
	masterCmd = &cobra.Command{
		Use:   "master",
		Short: "Run a master node service",
		Long:  "Run a master node service on this server",
		Run: func(cmd *cobra.Command, args []string) {
			bootstrap()
			go register()
			discover.ServiceCluster = discover.NewCluster(config.Conf.Etcd.EndPoints)
			ctx, cancelFunc := context.WithCancel(context.Background())
			go discover.ServiceCluster.WatchNodes(ctx)
			if false {
				cancelFunc()
			}

			start(fmt.Sprintf("%s:%d", config.Conf.Service.Host, config.Conf.Service.Port))
		},
	}
	master = &models.Node{
		Mode: models.MODE_MASTER,
		Status: models.ONLINE,
	}
)

func init() {
	rootCmd.AddCommand(masterCmd)
	config.Conf = config.Init()
	masterCmd.PersistentFlags().StringVarP(&config.Path, "config", "c", "/etc/ects/ects.yaml", "Set configuration file")
	masterCmd.PersistentFlags().StringVar(&config.Conf.Service.Host, "host", "0.0.0.0", "Set listen on IP")
	masterCmd.PersistentFlags().IntVar(&config.Conf.Service.Port, "port", 9701, "Set listen on port")
	masterCmd.PersistentFlags().StringSliceVar(&config.Conf.Etcd.EndPoints, "etcd", []string{"127.0.0.1:2379"}, "Set Etcd endpoints")
	masterCmd.PersistentFlags().StringVarP(&master.Id, "node", "n", "6d037444-8667-4364-848b-0b3b79e9044a", "Set master node id")
	masterCmd.PersistentFlags().StringVar(&master.Name, "name", "", "Set master node name")
	masterCmd.PersistentFlags().StringVar(&master.Description, "desc", "", "Set master node description")
}

func bootstrap() {
	var err error
	// 判断是否已经安装
	system.Info = &system.Information{
		Version: rootCmd.Version,
	}
	system.Info.Installed, err = config.CheckConfigFile(config.Path)

	if system.Info.Installed {
		file, err := ioutil.ReadFile(config.Path)
		if err != nil {
			log.Println(err)
		}
		err = yaml.Unmarshal(file, &config.Conf)
		if err != nil {
			log.Println(err)
		}
		models.Engine, err = models.Connection()
		if err := models.Migrate(); err != nil {
			log.Println(err)
		}
	} else {
		if err != nil {
			// TODO
		}
		dir := path.Dir(config.Path)
		config.CreateConfigDir(dir)
		system.Info.Permission = config.CheckConfigDirPermisson(dir)
	}
}

func register()  {
	master.Host = config.Conf.Service.Host
	master.Port = config.Conf.Service.Port

	if master.Id == "" {
		master.Id = uuid.NewV4().String()
	}

	if master.Name == "" {
		master.Name = "master-" + master.Id
	}

	service, err := discover.NewService(master, EndPoints)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if err := service.Register(5); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func start(addr string) {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	// 注册路由
	routes.Register(app)

	if err := app.Run(iris.Addr(addr), iris.WithoutInterruptHandler, iris.WithCharset("UTF-8")); err != nil {
		log.Println(err)
	}
}
