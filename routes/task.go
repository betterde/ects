package routes

import (
	"github.com/betterde/ects/controllers/task"
	"github.com/betterde/ects/services"
	"github.com/kataras/iris/v12/mvc"
)

func registerTask(application *mvc.Application) {
	application.Register(services.NewTaskService())
	application.Handle(new(task.Controller))
}
