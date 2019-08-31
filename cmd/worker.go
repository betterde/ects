package cmd

import (
	"context"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/discover"
	"github.com/betterde/ects/internal/pipeline"
	"github.com/betterde/ects/internal/scheduler"
	"github.com/betterde/ects/internal/service"
	"github.com/betterde/ects/internal/utils"
	"github.com/betterde/ects/models"
	"github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

// workerCmd represents the worker command
var (
	workerCmd = &cobra.Command{
		Use:   "worker",
		Short: "Run a worker node service",
		Long:  "Run a worker node service on this server",
		Run: func(cmd *cobra.Command, args []string) {
			listen()
		},
	}

	worker = &service.Instance{
		Mode:    models.WORKER,
		Status:  models.ONLINE,
		Version: rootCmd.Version,
	}
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rootCmd.AddCommand(workerCmd)
	service.Runtime = &service.Instance{
		Mode:   models.WORKER,
		Version: rootCmd.Version,
	}
	workerCmd.Flags().StringVar(&worker.Name, "name", "", "Set worker node name")
	workerCmd.Flags().StringSliceVar(&service.EndPoints, "etcd", []string{"127.0.0.1:2379"}, "Set Etcd endpoints")
	workerCmd.Flags().StringVarP(&worker.Id, "node", "n", "", "Set node id")
	workerCmd.Flags().StringVar(&worker.Description, "desc", "worker node", "Set worker node description")
	workerCmd.Flags().StringVar(&service.ConfigKey, "config", "/ects/config", "Set the key used to get configuration information")
}

func listen() {
	if worker.Id == "" {
		worker.Id = uuid.NewV4().String()
	}

	var err error
	if worker.Name == "" {
		worker.Name, err = os.Hostname()
		if err != nil {
			log.Fatal(err)
		}
	}

	if worker.Host == "0.0.0.0" {
		ips := utils.GetIPs()
		if len(ips) > 0 {
			worker.Host = ips[0]
		}
	}

	config.Conf.Etcd.EndPoints = service.EndPoints
	discover.NewClient()
	discover.GetConf(service.ConfigKey)
	models.Engine, err = models.Connection()
	if err != nil {
		log.Fatal(err)
	}

	ips := utils.GetIPs()
	if len(ips) > 0 {
		service.Runtime.Host = ips[0]
	}

	service.Runtime = worker

	ser, err := discover.NewService(service.Runtime)
	if err != nil {
		log.Fatal(err)
	}

	go func(ser *discover.Service) {
		if err := ser.Register(5); err != nil {
			log.Fatal(err)
		}
	}(ser)

	scheduler.New()
	ctx, cancelFunc := context.WithCancel(context.Background())
	go scheduler.Instance.Run(ctx)
	go pipeline.WatchPipelines(service.Runtime.Id)

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	for {
		receiver := <-sign
		log.Printf("get a signal %s", receiver.String())
		switch receiver {
		case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL:
			ser.Stop()
			cancelFunc()
			return
		}
	}
}
