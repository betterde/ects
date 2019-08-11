package cmd

import (
	"crypto/tls"
	"github.com/betterde/ects/models"
	"github.com/gorhill/cronexpr"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	"gopkg.in/gomail.v2"
	"log"
	"strings"
	"time"
)

var singleCmd = &cobra.Command{
	Use:   "single",
	Short: "Run a single node service",
	Long:  "Run a single node service on this server",
	Run: func(cmd *cobra.Command, args []string) {
		//discover.NewClient()
		//config.Conf = config.Init()
		//discover.GetConf("/ects/config")
		//pipeline.WatchPipelines("7df52971-4894-4f01-9171-7452c4ddceca")
		log.Println(cronexpr.MustParse("* * * * * *").Next(time.Now()).Format(models.DefaultTimeFormat))
		log.Println(strings.TrimPrefix("/var/local/laravel", "/var/local/"))
		log.Println(uuid.NewV4().String())
		d := gomail.NewDialer("", 27, "system@betterde.com", "system@2019")
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	},
}

func init() {
	rootCmd.AddCommand(singleCmd)
}
