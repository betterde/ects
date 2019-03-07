package routes

import (
	"github.com/betterde/ects/internal/controllers/organization"
	"github.com/kataras/iris/mvc"
)

func teams(application *mvc.Application) {
	application.Handle(new(organization.TeamController))
}
