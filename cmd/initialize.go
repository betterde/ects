package cmd

import (
	"context"
	"encoding/json"
	"github.com/betterde/ects/config"
	"github.com/betterde/ects/internal/discover"
	"github.com/betterde/ects/internal/system"
	"github.com/betterde/ects/internal/utils"
	"github.com/betterde/ects/models"
	"github.com/betterde/ects/routes"
	"github.com/betterde/ects/web"
	"github.com/coreos/etcd/clientv3"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"runtime"
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
			switch mode {
			case "web":
				startInitializeWeb()
				break
			case "json":
				startInitializeByConfigFile()
			case "yaml":
				startInitializeByConfigFile()
				break
			default:
				log.Println("Unsupport file extension")
				os.Exit(1)
			}
		},
	}

	mode string
	path string

	user = &models.User{
		Id:        uuid.NewV4().String(),
		Manager:   true,
		CreatedAt: utils.Time(time.Now()),
		UpdatedAt: utils.Time(time.Now()),
	}
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	config.Conf = config.Init()
	system.Info = &system.Information{
		Version: rootCmd.Version,
	}
	rootCmd.AddCommand(initializeCmd)
	initializeCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "web", "Set initialize mode with web ui or json, yaml config file")
	initializeCmd.PersistentFlags().StringVarP(&path, "path", "p", "", "Set config file path")
	initializeCmd.PersistentFlags().StringVarP(&user.Name, "name", "n", "", "Set admin name")
	initializeCmd.PersistentFlags().StringVarP(&user.Email, "email", "e", "", "Set admin email")
	initializeCmd.PersistentFlags().StringVarP(&user.Password, "pass", "P", "", "Set admin pass")
}

func startInitializeWeb() {
	app := iris.New()
	app.Logger().SetLevel("disable")
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
		sctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		if err := app.Shutdown(sctx); err != nil {
			log.Println(err)
		}
	})

	if err := app.Run(iris.Addr(":9701"), iris.WithOptimizations, iris.WithCharset("UTF-8"), iris.WithoutInterruptHandler); err != nil {
		log.Println(err)
	}
}

func startInitializeByConfigFile() {
	validatePath()
	if user.Name == "" || user.Email == "" || user.Password == "" {
		log.Println("Please enter admin user info")
		os.Exit(1)
	}
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if mode == "json" {
		if err := json.Unmarshal(buf, config.Conf); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	if mode == "yaml" {
		if err := yaml.Unmarshal(buf, config.Conf); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	discover.NewClient()

	buf, err = json.Marshal(config.Conf)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if res, err := discover.Client.Put(context.TODO(), config.Conf.Etcd.Config, string(buf), clientv3.WithPrevKV()); err != nil {
		log.Println(err)
		os.Exit(1)
	} else {
		if len(res.PrevKv.Value) > 0 {
			log.Printf("%s %s \n", "OLD CONFIG IS", string(res.PrevKv.Value))
		}
	}

	if err := utils.CreateDatabase(); err != nil {
		log.Println("Failed to create database", err)
		os.Exit(1)
	}

	if models.Engine == nil {
		// Create database engine
		models.Engine, err = models.Connection()
		if err != nil {
			log.Println("Failed to connect to database", err)
			os.Exit(1)
		}
	}

	if err := models.Migrate(); err != nil {
		log.Println("Failed to migrate the table", err)
		os.Exit(1)
	}

	pass, err := models.GeneratePassword(user.Password)
	user.Password = string(pass)

	if _, err := models.Engine.Insert(user); err != nil {
		log.Println("Failed to create system manager", err)
		os.Exit(1)
	}

	log.Println("INITIALIZE SUCCESSFUL")
}

func validatePath() {
	if path == "" {
		log.Println("Please enter your config file path or use --mode=web")
		os.Exit(1)
	}

	exist, err := config.CheckConfigFile(path)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if exist == false {
		log.Println("Config file does not exist")
		os.Exit(1)
	}
}
