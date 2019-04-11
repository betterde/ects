package cmd

import (
	"context"
	"fmt"
	"github.com/betterde/ects/internal/discover"
	"github.com/betterde/ects/internal/rpc"
	"github.com/betterde/ects/models"
	"github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
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
		Mode: models.MODE_WORKER,
		Status: models.ONLINE,
	}

	EndPoints []string
)

type (
	Server struct {
	}
)

func init() {
	rootCmd.AddCommand(workerCmd)
	workerCmd.PersistentFlags().StringVar(&worker.Name, "name", "", "Set worker node name")
	workerCmd.PersistentFlags().StringVar(&worker.Host, "host", "0.0.0.0", "Set listen on IP")
	workerCmd.PersistentFlags().IntVar(&worker.Port, "port", 9412, "Set listen on port")
	workerCmd.PersistentFlags().StringSliceVar(&EndPoints, "etcd", []string{"127.0.0.1:2379"}, "Set Etcd endpoints")
	workerCmd.PersistentFlags().StringVarP(&worker.Id, "node", "n", "6d037444-8667-4364-848b-0b3b79e9044b", "Set node id")
	workerCmd.PersistentFlags().StringVar(&worker.Description, "desc", "", "Set worker node description")
}

func (server *Server) Run(ctx context.Context, request *rpc.Request) (*rpc.Response, error) {
	return &rpc.Response{
		Output: "Hello",
	}, nil
}

func listen() {
	addr := fmt.Sprintf("%s:%d", worker.Host, worker.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if worker.Id == "" {
		worker.Id = uuid.NewV4().String()
	}

	if worker.Name == "" {
		worker.Name = "worker-" + worker.Id
	}

	service, err := discover.NewService(worker, EndPoints)
	if err := service.Register(5); err != nil {
		log.Println(err)
	}

	server := grpc.NewServer()
	rpc.RegisterTaskServer(server, &Server{})
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
