package discover

import (
	"github.com/betterde/ects/models"
	"github.com/coreos/etcd/clientv3"
	"log"
	"sync"
	"time"
)

type (
	Cluster struct {
		Mutex sync.Mutex
		Nodes map[string]models.Node
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
		Nodes: make(map[string]models.Node),
		client: client,
	}
}
