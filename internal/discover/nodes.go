package discover

import (
	"context"
	"encoding/json"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/models"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"log"
)

func (cluster *Cluster) WatchNodes(ctx context.Context) {
	var curRevision int64 = 0

	for {
		rangeResp, err := Client.Get(context.TODO(), config.Conf.Etcd.Service, clientv3.WithPrefix())

		if err != nil {
			continue
		}
		// 从当前版本开始订阅
		curRevision = rangeResp.Header.Revision + 1
		break
	}

	watchChan := Client.Watch(ctx, config.Conf.Etcd.Service, clientv3.WithPrefix(), clientv3.WithRev(curRevision), clientv3.WithPrevKV())
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			var node models.Node
			switch event.Type {
			case mvccpb.PUT:
				if err := json.Unmarshal(event.Kv.Value, &node); err != nil {
					log.Println(err)
				}

				if err := node.CreateOrUpdate(); err != nil {
					log.Println(err)
				}

				log.Printf("节点：%s 注册成功", node.Id)
			case mvccpb.DELETE:
				if err := json.Unmarshal(event.PrevKv.Value, &node); err != nil {
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
