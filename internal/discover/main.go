package discover

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/betterde/ects/models"
	"github.com/coreos/etcd/clientv3"
	"log"
	"sync"
	"time"
)

const (
	SERVICE_CONF = "/ects/config/"
	SERVICE_NODE = "/ects/nodes/"
	SERVICE_STATUS = "/ects/status/"
)

type (
	Service struct {
		node    *models.Node
		client  *clientv3.Client
		leaseID clientv3.LeaseID
		close   chan struct{}
		wg      sync.WaitGroup
	}
)

func NewService(node *models.Node, endpoints []string) (*Service, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 10 * time.Second,
	})

	if nil != err {
		return nil, err
	}

	service := &Service{
		client: client,
		close:  make(chan struct{}),
		node:   node,
	}

	return service, nil
}

// 注册服务
func (service *Service) Register(ttlSecond int64) error {
	res, err := service.client.Grant(context.TODO(), ttlSecond)
	if err != nil {
		return err
	}

	service.leaseID = res.ID

	val, err := json.Marshal(&service.node)
	if err != nil {
		log.Println(err)
	}

	var key string
	var cancelCtx context.Context

	key = fmt.Sprintf("%s%s", SERVICE_NODE, service.node.Id)

	cancelCtx, _ = context.WithCancel(context.TODO())

	// 创建带有节点详情的永久Key
	if _, err = service.client.Put(cancelCtx, key, string(val)); err != nil {
		return err
	}

	// 创建实时更新节点状态的Key，当节点出现故障则移除节点状态
	if _, err = service.client.Put(cancelCtx, fmt.Sprintf("%s%s", SERVICE_STATUS, service.node.Id), service.node.Id, clientv3.WithLease(service.leaseID)); err != nil {
		return err
	}


	log.Printf("Node is runing, register id is %s", service.node.Id)

	ch, err := service.client.KeepAlive(context.TODO(), service.leaseID)
	if nil != err {
		panic(err)
	}

	service.wg.Add(1)
	defer service.wg.Done()

	for {
		select {
		case <-service.close:
			return service.revoke()
		case <-service.client.Ctx().Done():
			return errors.New("server closed")
		case _, ok := <-ch:
			if !ok {
				return service.revoke()
			}
		}
	}
}

// 停止服务
func (service *Service) Stop() {
	close(service.close)
	service.wg.Wait()
	if err := service.client.Close(); err != nil {
		log.Println(err)
	}
}

func (service *Service) revoke() error {
	_, err := service.client.Revoke(context.TODO(), service.leaseID)
	if err != nil {
		log.Printf("[discovery] Service revoke key:[%s] error:[%s]", service.node.Id, err.Error())
	} else {
		log.Printf("[discovery] Service revoke successfully key:[%s]", service.node.Id)
	}

	return err
}
