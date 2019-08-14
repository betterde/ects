package cmd

import (
	"crypto/tls"
	"fmt"
	"github.com/betterde/ects/models"
	"github.com/gorhill/cronexpr"
	"github.com/spf13/cobra"
	"gopkg.in/gomail.v2"
	"log"
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
		d := gomail.NewDialer("smtp.qiye.aliyun.com", 465, "system@betterde.com", "System@2019")
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		m := gomail.NewMessage()
		m.SetHeader("From", fmt.Sprintf("%s<%s>", "ECTS", "system@betterde.com"))
		m.SetHeader("To", "george@betterde.com")
		m.SetHeader("Subject", "Notification")
		m.SetBody("text/html", "Hello Goerge")
		if err := d.DialAndSend(m); err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(singleCmd)
}
