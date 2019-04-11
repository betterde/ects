package routes

import (
	"github.com/betterde/ects/controllers/organization"
	"github.com/betterde/ects/services"
	"github.com/kataras/iris/mvc"
)

func registerTeam(application *mvc.Application) {
	application.Register(services.NewTeamService())
	application.Handle(new(organization.TeamController))
}
