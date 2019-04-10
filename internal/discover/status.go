package discover

import (
	"context"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/models"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"log"
	"os"
	"time"
)

func (cluster *Cluster) WatchStatus(ctx context.Context) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: config.Conf.Etcd.EndPoints,
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer func() {
		if err := client.Close(); err != nil {
			log.Println(err)
		}
	}()

	var curRevision int64 = 0

	//for {
	//	rangeResp, err := client.Get(context.TODO(), SERVICE_STATUS, clientv3.WithPrefix())
	//	if err != nil {
	//		continue
	//	}
	//
	//	cluster.Mutex.Lock()
	//
	//	for _, kv := range rangeResp.Kvs {
	//		log.Println(kv)
	//		node := models.Node{}
	//
	//		if err := json.Unmarshal(kv.Value, &node); err != nil {
	//			log.Println(err)
	//		}
	//
	//		cluster.Nodes[string(kv.Key)] = node
	//	}
	//
	//	cluster.Mutex.Unlock()
	//
	//	// 从当前版本开始订阅
	//	curRevision = rangeResp.Header.Revision + 1
	//	break
	//}

	// 监听后续的PUT与DELETE事件
	watcher := clientv3.NewWatcher(client)
	defer func() {
		if err := watcher.Close(); err != nil {
			log.Println(err)
		}
	}()

	watchNodes := watcher.Watch(ctx, SERVICE_STATUS, clientv3.WithPrefix(), clientv3.WithRev(curRevision))
	for watchResp := range watchNodes { // if ctx is Done, for loop will break
		for _, event := range watchResp.Events {
			switch (event.Type) {
			case mvccpb.PUT:
				log.Printf("节点：%s 上线", event.Kv.Value)
			case mvccpb.DELETE:
				log.Printf("节点：%s 离线", event.Kv.Key)
			}
		}
	}

	log.Println("停止监听")
}

func (cluster *Cluster) AllStatus() (nodes []models.Node) {
	nodes = make([]models.Node, 0)
	cluster.Mutex.Lock()
	for _, node := range cluster.Nodes {
		nodes = append(nodes, node)
	}
	cluster.Mutex.Unlock()
	return
}