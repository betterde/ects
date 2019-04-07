package services

import (
	"github.com/betterde/ects/models"
	"github.com/go-xorm/builder"
	"log"
)

type TaskInterface interface {
	Tasks(pipline bool) []models.Task
	NormalTasks() ([]models.Task, error)
}

func NewTaskService() TaskInterface {
	return &TaskService{}
}

type TaskService struct {
}

// 获取任务列表
func (service *TaskService) Tasks(pipline bool) []models.Task {
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

// 获取需要执行的任务
func (service *TaskService) NormalTasks() ([]models.Task, error) {
	tasks := make([]models.Task, 0)
	err := models.Engine.Where(builder.Eq{"status": "normal"}).Find(&tasks)
	return tasks, err
}