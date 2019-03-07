package routes

import (
	"github.com/betterde/ects/web"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
)


func Register(application *iris.Application) {
	// SPA 路由刷新404 处理
	application.OnErrorCode(404, func(ctx iris.Context) {
		ctx.StatusCode(200)
		if err := ctx.View("index.html"); err != nil {
			log.Println(err)
		}
	})

	application.RegisterView(iris.HTML("./web/dist", ".html").Binary(web.Asset, web.AssetNames))

	// 页面路由
	application.Get("/", func(ctx iris.Context) {
		if err := ctx.View("index.html"); err != nil {
			log.Println(err)
		}
	})

	assetHandler := iris.StaticEmbeddedHandler("./web/dist",  web.Asset,  web.AssetNames, false)
	application.SPA(assetHandler).AddIndexName("index.html")

	// 注册鉴权路由
	mvc.Configure(application.Party("/auth"), authentication)

	// 注册用户管理路由
	mvc.Configure(application.Party("/organization/user"), users)
}