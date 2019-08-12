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
	"time"
)

func (cluster *Cluster) WatchNodes(id string, ctx context.Context) {
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
			leaseCtx, cancelFunc := context.WithTimeout(ctx, 5 * time.Second)
			leaseRes, err := Client.Grant(ctx, 5)
			if err != nil {
				log.Println(err)
			}

			_, err = Client.KeepAlive(leaseCtx, leaseRes.ID)
			if err != nil {
				log.Println(err)
			}

			switch event.Type {
			case mvccpb.PUT:
				if err := json.Unmarshal(event.Kv.Value, &node); err != nil {
					log.Println(err)
				}

				lookKey := fmt.Sprintf("%s/%s/", config.Conf.Locker, node.Id)

				// 执行事物，开始抢锁
				txn := Client.Txn(context.TODO())
				txn.If(clientv3.Compare(clientv3.CreateRevision(lookKey), "=", 0)).
					Then(clientv3.OpPut(lookKey, id, clientv3.WithLease(leaseRes.ID))).
					Else(clientv3.OpGet(lookKey))
				txnResp, err := txn.Commit()
				if err != nil {
					log.Println(err)
				}

				// 如果没有抢到锁，则获取占用锁的节点ID
				if !txnResp.Succeeded {
					log.Printf("锁已经被节点 %s 占用", txnResp.Responses[0].GetResponseRange().Kvs[0].Value)
				} else {
					// 如果抢到锁，则创建或更新节点信息
					if err := node.CreateOrUpdate(); err != nil {
						log.Println(err)
					}

					log.Printf("节点：%s 注册成功", node.Id)
				}

				cancelFunc()
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
