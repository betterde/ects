package discover

import (
	"github.com/coreos/etcd/clientv3"
	"log"
	"time"
)

type (
	Cluster struct {
		client *clientv3.Client
	}
)

var (
	ServiceCluster *Cluster
)

func NewCluster(endpoints []string) *Cluster {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 10 * time.Second,
	})

	if err != nil {
		log.Println(err)
	}

	return &Cluster{
		client: client,
	}
}
