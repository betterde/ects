package discover

import (
	"context"
	"encoding/json"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/models"
	"github.com/betterde/ects/services"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"log"
	"os"
	"time"
)

func (cluster *Cluster) WatchNodes(ctx context.Context) {
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

	for {
		rangeResp, err := client.Get(context.TODO(), SERVICE_NODE, clientv3.WithPrefix())
		if err != nil {
			continue
		}
		// 从当前版本开始订阅
		curRevision = rangeResp.Header.Revision + 1
		break
	}

	// 监听后续的PUT与DELETE事件
	watcher := clientv3.NewWatcher(client)
	defer func() {
		if err := watcher.Close(); err != nil {
			log.Println(err)
		}
	}()

	nodeService := services.NewNodeService()

	watchNodes := watcher.Watch(ctx, SERVICE_NODE, clientv3.WithPrefix(), clientv3.WithRev(curRevision))
	for watchResp := range watchNodes {
		for _, event := range watchResp.Events {
			switch (event.Type) {
			case mvccpb.PUT:
				var node models.Node
				if err := json.Unmarshal(event.Kv.Value, &node); err != nil {
					log.Println(err)
				}

				if err := node.CreateOrUpdate(); err != nil {
					log.Println(err)
				}

				log.Printf("节点：%s 注册成功", node.Id)
			case mvccpb.DELETE:
				id := string(event.Kv.Key)[12:]
				node, err := nodeService.FindByID(id)

				if err != nil {
					log.Println(err)
				}

				node.Status = models.OFFLINE
				if err := node.Update(); err != nil {
					log.Println(err)
				}
				log.Printf("节点：%s 离线", node.Id)
			}
		}
	}
}
