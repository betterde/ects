package cmd

import (
	"context"
	"github.com/betterde/ects/internal/build"
	"github.com/betterde/ects/internal/global"
	"github.com/betterde/ects/internal/journal"
	"github.com/betterde/ects/internal/server"
	"github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/discover"
	"github.com/betterde/ects/internal/service"
	"github.com/betterde/ects/internal/utils"
	"github.com/betterde/ects/models"
)

// masterCmd represents the master command
var (
	masterCmd = &cobra.Command{
		Use:   "server",
		Short: "Run a server node",
		Run: func(cmd *cobra.Command, args []string) {
			bootstrap()
			watch()
			start()
		},
	}

	master = &service.Instance{
		Mode:    models.MASTER,
		Status:  models.ONLINE,
		Version: build.Version,
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

	if err := ser.Register(300); err != nil {
		log.Fatal(err)
	}
}

func start() {
	service.Runtime = master
	server.Instance.Run(verbose)
	go register()
	if err := signalHandler(global.CancelFunc); err != nil {
		journal.Logger.Errorw("Failed to shutdown CDNS server:", err)
	}
}

func signalHandler(cancel context.CancelFunc) error {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	select {
	case <-shutdown:
		cancel()
		return server.Instance.Engine.Shutdown()
	}
}
