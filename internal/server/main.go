package server

import (
	"ects/internal/api"
	"ects/web"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"log"
)

func Start() *iris.Application {
	server := iris.New()

	server.Use(iris.Gzip)
	server.Use(recover.New())
	server.Use(logger.New())

	// SPA 路由刷新404 处理
	server.OnErrorCode(404, func(ctx iris.Context) {
		ctx.StatusCode(200)
		if err := ctx.View("index.html"); err != nil {
			log.Println(err)
		}
	})

	server.RegisterView(iris.HTML("./web/dist", ".html").Binary(web.Asset, web.AssetNames))

	// 页面路由
	server.Get("/", func(ctx iris.Context) {
		if err := ctx.View("index.html"); err != nil {
			log.Println(err)
		}
	})

	assetHandler := iris.StaticEmbeddedHandler("./web/dist",  web.Asset,  web.AssetNames, false)
	server.SPA(assetHandler).AddIndexName("index.html")

	mvc.Configure(server.Party("/auth"), func(application *mvc.Application) {
		application.Handle(new(api.AuthenticationController))
	})
	return server
}
