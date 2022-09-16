package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/satori/go.uuid"
	"github.com/spf13/cobra"

	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/discover"
	"github.com/betterde/ects/internal/service"
	"github.com/betterde/ects/internal/utils"
	"github.com/betterde/ects/models"
	"github.com/betterde/ects/routes"
)

// masterCmd represents the master command
var (
	masterCmd = &cobra.Command{
		Use:   "master",
		Short: "Run a master node service",
		Long:  "Run a master node service on this server",
		Run: func(cmd *cobra.Command, args []string) {
			bootstrap()
			watch()
			start()
		},
	}

	master = &service.Instance{
		Mode:    models.MASTER,
		Status:  models.ONLINE,
		Version: rootCmd.Version,
	}

	ctx, cancelFunc = context.WithCancel(context.Background())
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rootCmd.AddCommand(masterCmd)
	config.Conf = config.Init()
	service.Initialize()
	masterCmd.Flags().StringVar(&master.Host, "host", "0.0.0.0", "Set listen on IP")
	masterCmd.Flags().IntVar(&master.Port, "port", 9701, "Set listen on port")
	masterCmd.Flags().StringSliceVar(&service.EndPoints, "etcd", []string{"127.0.0.1:2379"}, "Set Etcd endpoints")
	masterCmd.Flags().StringVarP(&master.Id, "node", "n", uuid.NewV4().String(), "Set master node id")
	masterCmd.Flags().StringVar(&master.Name, "name", "", "Set master node name")
	masterCmd.Flags().StringVar(&master.Description, "desc", "master node", "Set master node description")
	masterCmd.Flags().StringVar(&service.ConfigKey, "config", "/ects/config", "Set the key used to get configuration information")
}

func bootstrap() {
	var err error
	config.Conf.Etcd.EndPoints = service.EndPoints
	discover.NewClient()
	discover.GetConf(service.ConfigKey)
	models.Engine, err = models.Connection()
	if err != nil {
		log.Fatal(err)
	}
}

func watch() {
	go discover.ServiceCluster.WatchNodes(master.Id, ctx)
}

// Service registry
func register() {
	if master.Host == "0.0.0.0" {
		ips := utils.GetIPs()
		if len(ips) > 0 {
			master.Host = ips[0]
		}
	}

	var err error

	if master.Name == "" {
		master.Name, err = os.Hostname()
		if err != nil {
			log.Fatal(err)
		}
	}

	ser, err := discover.NewService(master)
	if err != nil {
		log.Fatal(err)
	}

	if err := ser.Register(5); err != nil {
		log.Fatal(err)
	}
}

func start() {
	service.Runtime = master
	addr := fmt.Sprintf("%s:%d", service.Runtime.Host, service.Runtime.Port)
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Logger().SetLevel("disable")

	routes.Register(app)

	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// 当主节点关闭时，先检查是否有其他在线主节点
		if count, err := models.Engine.Where("mode = ? AND status = ? AND id != ?", "master", "online", service.Runtime.Id).Count(&models.Node{}); err != nil {
			log.Println(err)
		} else {
			if count == 0 {
				node := &models.Node{
					Id: service.Runtime.Id,
				}
				node.Offline()
			}
		}
		cancelFunc()
		if err := app.Shutdown(ctx); err != nil {
			log.Println(err)
		}
	})

	go register()

	if err := app.Run(iris.Addr(addr), iris.WithoutInterruptHandler, iris.WithOptimizations, iris.WithCharset("UTF-8")); err != nil {
		log.Fatal(err)
	}
}
