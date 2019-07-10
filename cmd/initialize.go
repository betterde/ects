package cmd

import (
	"context"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/system"
	"github.com/betterde/ects/routes"
	"github.com/betterde/ects/web"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/spf13/cobra"
	"log"
	"sync"
	"time"
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
		if err := ctx.View("initialize.html"); err != nil {
			log.Println(err)
		}
	})

	mvc.Configure(app.Party("/api/initialize"), routes.Initialize)

	assetHandler := iris.StaticEmbeddedHandler("./web/dist", web.Asset, web.AssetNames, false)
	app.SPA(assetHandler).AddIndexName("initialize.html")

	sg := new(sync.WaitGroup)
	defer sg.Wait()

	iris.RegisterOnInterrupt(func() {
		sg.Add(1)
		defer sg.Done()
		sctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
		defer cancel()
		if err := app.Shutdown(sctx); err != nil {
			log.Println(err)
		}
	})

	if err := app.Run(iris.Addr(":9701"), iris.WithOptimizations, iris.WithCharset("UTF-8"), iris.WithoutInterruptHandler); err != nil {
		log.Println(err)
	}
}
