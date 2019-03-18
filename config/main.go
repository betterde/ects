package config

import (
	"io/ioutil"
	"log"
	"os"
	"path"
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
	permission := true
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0666)

	if err != nil {
		if os.IsPermission(err) {
			permission = false
			log.Printf("Write %s Permission denied \n", path)
			//os.Exit(1)
		}
	}

	defer func() {
		if err := file.Close(); err != nil {
			// TODO
		}
	}()

	exist := !os.IsNotExist(err)
	return exist, permission, err
}

// 写入配置文件
func WriteConfigToFile(file string, content []byte) bool {
	dir := path.Dir(file)
	_, err := os.Stat(dir)
	if err != nil {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Panicln(err)
		}
	}
	if err := ioutil.WriteFile(file, content, 0666); err != nil {
		log.Println(os.IsNotExist(err))
		log.Println(err)
	}
	return true
}
