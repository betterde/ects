package routes

import (
	"github.com/betterde/ects/internal/middleware"
	"github.com/betterde/ects/internal/utils/system"
	"github.com/betterde/ects/web"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
)

func Register(app *iris.Application) {
	// SPA 路由刷新404 处理
	app.OnErrorCode(404, func(ctx iris.Context) {
		ctx.StatusCode(200)
		if err := ctx.View("index.html"); err != nil {
			log.Println(err)
		}
	})

	app.RegisterView(iris.HTML("./web/dist", ".html").Binary(web.Asset, web.AssetNames))

	// 页面路由
	app.Get("/", func(ctx iris.Context) {
		if system.Installed == false {
			ctx.Redirect("/install")
		}
		if err := ctx.View("index.html"); err != nil {
			log.Println(err)
		}
	})

	// 页面路由
	app.Get("/install", func(ctx iris.Context) {
		if err := ctx.View("install.html"); err != nil {
			log.Println(err)
		}
	})

	assetHandler := iris.StaticEmbeddedHandler("./web/dist", web.Asset, web.AssetNames, false)
	app.SPA(assetHandler).AddIndexName("index.html")

	// 接口路由
	mvc.Configure(app.PartyFunc("/api", func(api iris.Party) {
		// 安装路由
		mvc.Configure(api.Party("/install"), installation)
		// 鉴权路由
		mvc.Configure(api.Party("/auth"), authentication)

		api.Use(middleware.JWTHandler.Serve)
		// 组织路由
		mvc.Configure(api.PartyFunc("/organization", func(org iris.Party) {
			// 用户路由
			mvc.Configure(org.Party("/user"), users)
			// 团队路由
			mvc.Configure(org.Party("/team"), teams)
		}))
	}))

	app.Get("/ping", myHandler)
}

func myHandler(ctx iris.Context) {
	user := ctx.Values().Get("jwt").(*jwt.Token)

	claims, _ := user.Claims.(jwt.MapClaims)

	ctx.Writef("%d", int64(claims["sub"].(float64)))
}
