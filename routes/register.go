package routes

import (
	"github.com/betterde/ects/internal/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Register(app *iris.Application) {
	mvc.New(app.PartyFunc("/api", func(api iris.Party) {
		api.Logger().Info("register api")
		mvc.Configure(api.Party("/auth"), authentication)
		middleware.Init()
		//api.Use(middleware.JWTHandler)
		mvc.Configure(api.Party("/task"), registerTask)
		mvc.Configure(api.Party("/node"), registerNode)
		mvc.Configure(api.Party("/pipeline"), registerPipeline)
		mvc.Configure(api.Party("/dashboard"), registerDashboard)
		mvc.Configure(api.Party("/user"), registerUser)
		mvc.Configure(api.Party("/log"), registerLog)
		mvc.Configure(api.Party("/setting"), registerSetting)
		mvc.Configure(api.PartyFunc("/account", func(account iris.Party) {
			mvc.Configure(account.Party("/profile"), registerProfile)
		}))
	}))
}
