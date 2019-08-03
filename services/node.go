package services

import (
	"errors"
	"github.com/betterde/ects/models"
	"log"
)

type (
	NodeInterface interface {
		FindByID(id string) (*models.Node, error)
	}

	NodeService struct {
	}
)

func NewNodeService() NodeInterface {
	return &NodeService{}
}

// Get node by id
func (service *NodeService) FindByID(id string) (*models.Node, error) {
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
