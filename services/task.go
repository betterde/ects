package services

import (
	"github.com/betterde/ects/models"
	"github.com/go-xorm/builder"
	"log"
)

type TaskService interface {
	Tasks(pipline bool) []models.Task
}

func NewTaskService() TaskService {
	return &taskService{}
}

type taskService struct {
}

func (service *taskService) Tasks(pipline bool) []models.Task {
	tasks := make([]models.Task, 0)
	session := models.Engine.NewSession()
	if pipline {
		session.Where(builder.Eq{"parent_id": ""})
	}

	err := session.Find(&tasks)
	session.Close()
	if err != nil {
		log.Panicln(err)
	}

	return []models.Task{}
}
