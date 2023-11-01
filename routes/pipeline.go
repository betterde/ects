package routes

import (
	"github.com/betterde/ects/controllers/pipeline"
	"github.com/betterde/ects/services"
	"github.com/kataras/iris/v12/mvc"
)

func registerPipeline(application *mvc.Application) {
	application.Register(services.NewPipelineService())
	application.Handle(new(pipeline.Controller))
}
