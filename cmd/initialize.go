package cmd

import (
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/system"
	"github.com/betterde/ects/routes"
	"github.com/betterde/ects/web"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/spf13/cobra"
	"log"
)

// installCmd represents the install command
var (
	initializeCmd = &cobra.Command{
		Use:     "init",
		Short:   "Run initialize elastic crontab system service",
		Long:    "Run initialize elastic crontab system service",
		Example: "ects init",
		Run: func(cmd *cobra.Command, args []string) {
			startInitialize()
		},
	}
)

func init() {
	config.Conf = config.Init()
	system.Info = &system.Information{
			Version: rootCmd.Version,
	}
	rootCmd.AddCommand(initializeCmd)
}

func startInitialize() {
	app := iris.New()
	app.OnErrorCode(404, func(ctx iris.Context) {
		ctx.Redirect("/", iris.StatusMovedPermanently)
	})

	app.RegisterView(iris.HTML("./web/dist", ".html").Binary(web.Asset, web.AssetNames))

	app.Get("/", func(ctx iris.Context) {
		if err := ctx.View("install.html"); err != nil {
			log.Println(err)
		}
	})

	mvc.Configure(app.Party("/api/initialize"), routes.Initialize)

	assetHandler := iris.StaticEmbeddedHandler("./web/dist", web.Asset, web.AssetNames, false)
	app.SPA(assetHandler).AddIndexName("install.html")
	if err := app.Run(iris.Addr(":9701"), iris.WithOptimizations, iris.WithCharset("UTF-8")); err != nil {
		log.Println(err)
	}
}
