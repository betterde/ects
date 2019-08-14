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

			switch event.Type {
			case mvccpb.PUT:
				if err := json.Unmarshal(event.Kv.Value, &node); err != nil {
					log.Println(err)
				}
				result := seize(leaseCtx, node.Id, id)
				if result {
					if err := node.CreateOrUpdate(); err != nil {
						log.Println(err)
					}
					log.Printf("节点：%s 注册成功", node.Id)
				}
				break
			case mvccpb.DELETE:
				if err := json.Unmarshal(event.PrevKv.Value, &node); err != nil {
					log.Println(err)
				}

				result := seize(leaseCtx, node.Id, id)
				if result {
					node.Status = models.OFFLINE
					if err := node.Update(); err != nil {
						log.Println(err)
					}
					log.Printf("节点：%s 离线", node.Id)
				}
				break
			}
			cancelFunc()
		}
	}
}

// 抢用于更新节点信息的锁
func seize(ctx context.Context, key, val string) bool {
	res, err := Client.Grant(ctx, 5)
	if err != nil {
		log.Println(err)
		return false
	}
	resc, err := Client.KeepAlive(ctx, res.ID)
	if err != nil {
		log.Println(err)
		return false
	}

	go func(ctx context.Context) {
		for {
			select {
			case kresp := <-resc:
				if kresp != nil {
					log.Println("续租成功，LeaseID: ", kresp.ID)
				} else if resc == nil {
					log.Println("续租失败")
					return
				}
			case <-ctx.Done():
				if _, err := Client.Revoke(context.TODO(), res.ID); err != nil {
					log.Println(err)
				}
				return
			}
			time.Sleep(1 * time.Second)
		}
	}(ctx)

	lockey := fmt.Sprintf("%s/%s/", config.Conf.Locker, key)
	txn := Client.Txn(ctx)
	txn.If(clientv3.Compare(clientv3.CreateRevision(lockey), "=", 0)).
		Then(clientv3.OpPut(lockey, val, clientv3.WithLease(res.ID))).
		Else(clientv3.OpGet(lockey))
	txnResp, err := txn.Commit()
	if err != nil {
		log.Println(err)
		return false
	}

	if !txnResp.Succeeded {
		return false
	} else {
		return true
	}
}
