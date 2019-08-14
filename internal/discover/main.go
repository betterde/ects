package discover

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/models"
	"github.com/coreos/etcd/clientv3"
	"log"
	"sync"
	"time"
)

type (
	Service struct {
		node    *models.Node
		leaseID clientv3.LeaseID
		close   chan struct{}
		wg      sync.WaitGroup
	}
)

var (
	err    error
	Client *clientv3.Client
)

// New ETCD V3 Client
func NewClient() {
	if Client, err = clientv3.New(clientv3.Config{
		Endpoints:   config.Conf.Etcd.EndPoints,
		DialTimeout: 10 * time.Second,
	}); err != nil {
		log.Println(err)
	}
}

func NewService(node *models.Node) (*Service, error) {
	if nil != err {
		return nil, err
	}

	service := &Service{
		close: make(chan struct{}),
		node:  node,
	}

	return service, nil
}

// Register service
func (service *Service) Register(ttlSecond int64) error {
	res, err := Client.Grant(context.TODO(), ttlSecond)
	if err != nil {
		return err
	}

	service.leaseID = res.ID

	val, err := json.Marshal(&struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Host        string `json:"host"`
		Port        int    `json:"port"`
		Mode        string `json:"mode"`
		Status      string `json:"status"`
		Version     string `json:"version"`
		Description string `json:"description"`
	}{
		service.node.Id,
		service.node.Name,
		service.node.Host,
		service.node.Port,
		service.node.Mode,
		service.node.Status,
		service.node.Version,
		service.node.Description,
	})
	if err != nil {
		log.Println(err)
	}

	key := fmt.Sprintf("%s/%s", config.Conf.Etcd.Service, service.node.Id)

	if _, err = Client.Put(context.TODO(), key, string(val), clientv3.WithLease(service.leaseID)); err != nil {
		return err
	}

	log.Printf("启动成功, ID为: %s", service.node.Id)

	ch, err := Client.KeepAlive(context.TODO(), service.leaseID)
	if nil != err {
		return err
	}

	service.wg.Add(1)
	defer service.wg.Done()

	for {
		select {
		case <-service.close:
			return service.revoke()
		case <-Client.Ctx().Done():
			return errors.New("server closed")
		case _, ok := <-ch:
			if !ok {
				return service.revoke()
			}
		}
	}
}

// Stop service
func (service *Service) Stop() {
	close(service.close)
	service.wg.Wait()
	if err := Client.Close(); err != nil {
		log.Println(err)
	}
}

// Revoke service leaseID
func (service *Service) revoke() error {
	_, err := Client.Revoke(context.TODO(), service.leaseID)
	if err != nil {
		log.Printf("[discovery] Service revoke %s error: %s", service.node.Id, err.Error())
	} else {
		log.Printf("[discovery] Service revoke successfully %s", service.node.Id)
	}

	return err
}
