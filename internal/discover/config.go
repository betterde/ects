package discover

import (
	"context"
	"encoding/json"
	"github.com/betterde/ects/config"
	"github.com/coreos/etcd/clientv3"
	"log"
	"os"
)

// Get config from etcd
func GetConf(key string) {
	if res, err := Client.Get(context.TODO(), key, clientv3.WithFirstKey()...); err != nil {
		log.Println(err)
		os.Exit(1)
	} else if res != nil {
		if err := json.Unmarshal(res.Kvs[0].Value, &config.Conf); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
}
