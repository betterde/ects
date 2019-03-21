package cmd

import (
	"context"
	"fmt"
	"github.com/betterde/ects/internal/rpc"
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
		Short: "Run a master node service",
		Long:  "Run a master node service on this server",
		Run: func(cmd *cobra.Command, args []string) {
			listen()
		},
	}
	host string
	port int
)

type (
	Server struct {
	}
)

func init() {
	rootCmd.AddCommand(workerCmd)
	workerCmd.PersistentFlags().StringVar(&host, "host", "0.0.0.0", "Set listen on IP")
	workerCmd.PersistentFlags().IntVar(&port, "port", 9412, "Set listen on port")
}

func (server *Server) Run(ctx context.Context, request *rpc.Request) (*rpc.Response, error) {
	return &rpc.Response{
		Output: "Hello",
	}, nil
}

func listen() {
	addr := fmt.Sprintf("%s:%d", host, port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	rpc.RegisterTaskServer(server, &Server{})
	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
