package services

import (
	"errors"
	"github.com/betterde/ects/models"
	"log"
)

type (
	WorkerInterface interface {
		FindByID(id string) (*models.Worker, error)
	}

	WorkerService struct {
	}
)

func NewWorkerService() WorkerInterface {
	return &WorkerService{}
}

// 根据ID获取节点信息
func (service *WorkerService) FindByID(id string) (*models.Worker, error) {
	var worker models.Worker
	result, err := models.Engine.Id(id).Get(&worker)
	if err != nil {
		log.Println(err)
	}

	if result {
		return &worker, nil
	}

	return &worker, errors.New("节点信息不存在")
}
