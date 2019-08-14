package cmd

import (
	"github.com/betterde/ects/internal/discover"
	"github.com/betterde/ects/internal/pipeline"
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

	worker = &models.Node{
		Mode:   models.WORKER,
		Status: models.ONLINE,
		Version: rootCmd.Version,
	}

	EndPoints []string
)

type (
	Server struct {
	}
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rootCmd.AddCommand(workerCmd)
	workerCmd.PersistentFlags().StringVar(&worker.Name, "name", "", "Set worker node name")
	workerCmd.PersistentFlags().StringSliceVar(&EndPoints, "etcd", []string{"127.0.0.1:2379"}, "Set Etcd endpoints")
	workerCmd.PersistentFlags().StringVarP(&worker.Id, "node", "n", "", "Set node id")
	workerCmd.PersistentFlags().StringVar(&worker.Description, "desc", "worker node", "Set worker node description")
	workerCmd.PersistentFlags().StringVar(&confKey, "config", "/ects/config", "Set the key used to get configuration information")
}

func listen() {
	if worker.Id == "" {
		worker.Id = uuid.NewV4().String()
	}

	if worker.Name == "" {
		worker.Name = "worker-" + worker.Id
	}

	discover.NewClient()
	discover.GetConf(confKey)

	ips := utils.GetIPs()
	if len(ips) > 0 {
		worker.Host = ips[0]
	}

	service, err := discover.NewService(worker)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	go func(service *discover.Service) {
		if err := service.Register(5); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}(service)

	go pipeline.WatchPipelines(worker.Id)

	sign := make(chan os.Signal, 1)

	signal.Notify(sign, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	for {
		receiver := <-sign
		log.Printf("get a signal %s", receiver.String())
		switch receiver {
		case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL:
			service.Stop()
			return
		}
	}
}
