package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/discover"
	"github.com/betterde/ects/internal/system"
	"github.com/betterde/ects/models"
	"github.com/betterde/ects/routes"
	"github.com/coreos/etcd/clientv3"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

// masterCmd represents the master command
var (
	masterCmd = &cobra.Command{
		Use:   "master",
		Short: "Run a master node service",
		Long:  "Run a master node service on this server",
		Run: func(cmd *cobra.Command, args []string) {
			bootstrap()
			go watch()
			start()
		},
	}
	master = &models.Node{
		Mode:   models.MODE_MASTER,
		Status: models.ONLINE,
	}

	confKey string

	ctx, cancelFunc = context.WithCancel(context.Background())
)

func init() {
	rootCmd.AddCommand(masterCmd)
	config.Conf = config.Init()
	masterCmd.PersistentFlags().StringVar(&config.Conf.Service.Host, "host", "0.0.0.0", "Set listen on IP")
	masterCmd.PersistentFlags().IntVar(&config.Conf.Service.Port, "port", 9701, "Set listen on port")
	masterCmd.PersistentFlags().StringSliceVar(&config.Conf.Etcd.EndPoints, "etcd", []string{"127.0.0.1:2379"}, "Set Etcd endpoints")
	masterCmd.PersistentFlags().StringVarP(&master.Id, "node", "n", "", "Set master node id")
	masterCmd.PersistentFlags().StringVar(&master.Name, "name", "", "Set master node name")
	masterCmd.PersistentFlags().StringVar(&master.Description, "desc", "", "Set master node description")
	masterCmd.PersistentFlags().StringVar(&confKey, "config", "ects_config", "Set the key used to get configuration information")
}

func bootstrap() {
	var err error
	system.Info = &system.Information{
		Version: rootCmd.Version,
	}

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   config.Conf.EndPoints,
		DialTimeout: 10 * time.Second,
	})

	defer func() {
		if err := client.Close(); err != nil {
			log.Println(err)
		}
	}()

	if err != nil {
		log.Println(err)
	}

	requestctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := client.Get(requestctx, confKey)
	cancel()

	if res != nil {
		if err := json.Unmarshal(res.Kvs[0].Value, &config.Conf); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	models.Engine, err = models.Connection()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func watch() {
	discover.ServiceCluster = discover.NewCluster(config.Conf.Etcd.EndPoints)
	discover.ServiceCluster.WatchNodes(ctx)
}

// Service registry
func register() {
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
	if err := service.Register(2); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func start() {
	addr := fmt.Sprintf("%s:%d", config.Conf.Service.Host, config.Conf.Service.Port)
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	routes.Register(app)

	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		node := &models.Node{
			Id: master.Id,
		}
		node.Offline()
		cancelFunc()
		if err := app.Shutdown(ctx); err != nil {
			log.Println(err)
		}
	})

	go register()

	if err := app.Run(iris.Addr(addr), iris.WithoutInterruptHandler, iris.WithOptimizations, iris.WithCharset("UTF-8")); err != nil {
		log.Println(err)
	}
}
