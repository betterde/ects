package config

import (
	"github.com/betterde/ects/internal/utils/system"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Service struct {
		Host string `json:"host" yaml:"host"`
		Port int    `json:"port" yaml:"port"`
	}
	Database struct {
		Host string `json:"host" yaml:"host"`
		Port int    `json:"port" yaml:"port"`
		Name string `json:"name" yaml:"name"`
		User string `json:"user" yaml:"user"`
		Pass string `json:"pass" yaml:"pass"`
		Char string `json:"char" yaml:"char"`
	}
	Auth struct {
		Secret string `json:"secret" yaml:"secret"`
		TTL    int    `json:"ttl" yaml:"ttl"`
	}
}

var (
	Conf *Config
	Path string
)

func Init() *Config {
	conf := &Config{}
	return conf
}

// 检查配置文件是否存在
func CheckConfigFile(path string) (bool, bool, error) {
	_, err := os.Stat(path)
	permission := os.IsPermission(err)
	exist := os.IsExist(err)
	return exist, permission, err
}

// 写入配置文件
func WriteConfigToFile(path string, content []byte) bool {
	if system.Info.Permission {
		if err := ioutil.WriteFile(path, content, 0755); err != nil {
			panic(err)
		}
	} else {
		log.Printf("Write %s Permission denied \n", path)
	}
	return true
}
