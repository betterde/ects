package routes

import (
	"github.com/betterde/ects/controllers/task"
	"github.com/kataras/iris/mvc"
)

func registerTask(application *mvc.Application) {
	application.Handle(new(task.Controller))
}
