package cmd

import (
	"github.com/betterde/ects/internal/system"
	"github.com/betterde/ects/web"
	"github.com/kataras/iris"
	"github.com/spf13/cobra"
	"log"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Run install elastic crontab system service",
	Long:  "Run install elastic crontab system service",
	Run: func(cmd *cobra.Command, args []string) {
		app := iris.New()
		app.OnErrorCode(404, func(ctx iris.Context) {
			ctx.Redirect("/", iris.StatusMovedPermanently)
		})

		app.RegisterView(iris.HTML("./web/dist", ".html").Binary(web.Asset, web.AssetNames))

		// 页面路由
		app.Get("/", func(ctx iris.Context) {
			if err := ctx.View("install.html"); err != nil {
				log.Println(err)
			}
		})

		app.Get("/api/install", func(ctx iris.Context) {
			system.Info = &system.Information{
				Version: rootCmd.Version,
			}

			if _, err := ctx.JSON(system.Info); err != nil {
				log.Println(err)
			}
		})

		app.Post("/api/install", func(ctx iris.Context) {

		})

		assetHandler := iris.StaticEmbeddedHandler("./web/dist", web.Asset, web.AssetNames, false)
		app.SPA(assetHandler).AddIndexName("install.html")
		if err := app.Run(iris.Addr(":9701"), iris.WithOptimizations, iris.WithCharset("UTF-8")); err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
