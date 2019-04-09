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
	ID string
	Host string
	Port int
	EndPoints []string
)

type (
	Server struct {
	}
)

func init() {
	rootCmd.AddCommand(workerCmd)
	workerCmd.PersistentFlags().StringVar(&Host, "host", "0.0.0.0", "Set listen on IP")
	workerCmd.PersistentFlags().IntVar(&Port, "port", 9412, "Set listen on port")
	workerCmd.PersistentFlags().StringSliceVar(&EndPoints, "etcd", nil, "Set Etcd endpoints")
	workerCmd.PersistentFlags().StringVarP(&ID, "node", "n", "", "Set node id")
}

func (server *Server) Run(ctx context.Context, request *rpc.Request) (*rpc.Response, error) {
	return &rpc.Response{
		Output: "Hello",
	}, nil
}

func listen() {
	addr := fmt.Sprintf("%s:%d", Host, Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if ID == "" {
		ID = uuid.NewV4().String()
	}

	node := &models.Node{
		Id: ID,
		Name: "Worker",
		Host: Host,
		Port: Port,
		Mode: models.MODE_WORKER,
	}

	service, err := discover.NewService(node, EndPoints)
	if err := service.Register(600); err != nil {
		log.Println(err)
	}

	server := grpc.NewServer()
	rpc.RegisterTaskServer(server, &Server{})
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
