package discover

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/models"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"log"
	"os"
	"sync"
	"time"
)

type (
	Cluster struct {
		mutex sync.Mutex
		nodes map[string]models.Node
	}
)

func NewCluster() *Cluster {
	return &Cluster{
		nodes: make(map[string]models.Node),
	}
}

func (cluster *Cluster) Watch(ctx context.Context) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: config.Conf.Etcd.EndPoints,
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		os.Exit(1)
	}

	defer func() {
		if err := client.Close(); err != nil {
			log.Println(err)
		}
	}()

	var curRevision int64 = 0

	kv := clientv3.NewKV(client)

	for {
		rangeResp, err := kv.Get(context.TODO(), SERVICE_NODE, clientv3.WithPrefix())
		if err != nil {
			continue
		}

		cluster.mutex.Lock()

		for _, kv := range rangeResp.Kvs {
			node := models.Node{}

			if err := json.Unmarshal(kv.Value, &node); err != nil {
				log.Println(err)
			}

			cluster.nodes[string(kv.Key)] = node
		}

		cluster.mutex.Unlock()

		// 从当前版本开始订阅
		curRevision = rangeResp.Header.Revision + 1
		break
	}

	// 监听后续的PUT与DELETE事件
	watcher := clientv3.NewWatcher(client)
	defer watcher.Close()

	watchChan := watcher.Watch(ctx, SERVICE_NODE, clientv3.WithPrefix(), clientv3.WithRev(curRevision))
	for watchResp := range watchChan { // if ctx is Done, for loop will break
		for _, event := range watchResp.Events {
			cluster.mutex.Lock()
			switch (event.Type) {
			case mvccpb.PUT:
				node := models.Node{}
				if err := json.Unmarshal(event.Kv.Value, node); err != nil {
					log.Println(err)
				}
				fmt.Println("PUT事件")
				cluster.nodes[string(event.Kv.Key)] = node

			case mvccpb.DELETE:
				delete(cluster.nodes, string(event.Kv.Key))
				fmt.Println("DELETE事件")
			}
			cluster.mutex.Unlock()
		}
	}

	fmt.Println("停止监听")
}

func (cluster *Cluster) Nodes() (nodes []models.Node) {
	nodes = make([]models.Node, 0)

	// 按endpoint去重
	cluster.mutex.Lock()
	log.Println(cluster.nodes)
	for _, node := range cluster.nodes {
		nodes = append(nodes, node)
	}
	cluster.mutex.Unlock()
	return
}
