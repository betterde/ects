package routes

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"

	"github.com/betterde/ects/internal/middleware"
	"github.com/betterde/ects/web"
)

func Register(app *iris.Application) {
	// SPA 404 handler
	app.OnErrorCode(404, func(ctx iris.Context) {
		ctx.StatusCode(200)
		if err := ctx.View("index.html"); err != nil {
			log.Println(err)
		}
	})

	mvc.Configure(app.PartyFunc("/api", func(api iris.Party) {
		mvc.Configure(api.Party("/auth"), authentication)
		api.Use(middleware.JWTHandler.Serve)
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

	app.Use(iris.Gzip)
	app.RegisterView(iris.HTML("./web/dist", ".html").Binary(web.Asset, web.AssetNames))

	app.Get("/", func(ctx iris.Context) {
		if err := ctx.View("index.html"); err != nil {
			log.Println(err)
		}
	})

	assetHandler := iris.StaticEmbeddedHandler("./web/dist", web.Asset, web.AssetNames, false)
	app.SPA(assetHandler).AddIndexName("index.html")
}
