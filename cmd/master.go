// Copyright © 2019 George <george@betterde.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"ects/config"
	"ects/internal/server"
	"fmt"
	"github.com/kataras/iris"
	"github.com/spf13/cobra"
	"log"
)

var conf = config.Init()

// masterCmd represents the master command
var masterCmd = &cobra.Command{
	Use:   "master",
	Short: "Run a master node service",
	Long:  "Run a master node service on this server",
	Run: func(cmd *cobra.Command, args []string) {
		start(fmt.Sprintf("%s:%d", conf.Service.Host, conf.Service.Port))
	},
}

func init() {
	rootCmd.AddCommand(masterCmd)
	masterCmd.PersistentFlags().StringVar(&conf.Service.Host, "host", "0.0.0.0", "Set listen on IP")
	masterCmd.PersistentFlags().IntVar(&conf.Service.Port, "port", 7368, "Set listen on port")
	masterCmd.PersistentFlags().StringVar(&conf.Database.Host, "db_host", "localhost", "Set mysql service host")
	masterCmd.PersistentFlags().IntVar(&conf.Database.Port, "db_port", 3306, "Set mysql service host")
	masterCmd.PersistentFlags().StringVar(&conf.Database.Name, "db_name", "ects", "Set mysql service db name")
	masterCmd.PersistentFlags().StringVar(&conf.Database.User, "db_user", "root", "Set mysql service user")
	masterCmd.PersistentFlags().StringVar(&conf.Database.Host, "db_pass", "", "Set mysql service pass")
}

func start(addr string)  {
	app := server.Start()
	if err := app.Run(iris.Addr(addr), iris.WithoutInterruptHandler); err != nil {
		log.Println(err)
	}
}
