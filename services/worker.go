package services

import (
	"errors"
	"github.com/betterde/ects/models"
	"log"
)

type (
	WorkerInterface interface {
		FindByID(id string) (*models.Node, error)
	}

	WorkerService struct {
	}
)

func NewWorkerService() WorkerInterface {
	return &WorkerService{}
}

// 根据ID获取节点信息
func (service *WorkerService) FindByID(id string) (*models.Node, error) {
	var node models.Node
	result, err := models.Engine.Id(id).Get(&node)
	if err != nil {
		log.Println(err)
	}

	if result {
		return &node, nil
	}

	return &node, errors.New("节点信息不存在")
}
