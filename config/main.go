package config

import (
	"io/ioutil"
	"log"
	"os"
)

type (
	Etcd struct {
		Killer    string   `json:"killer" yaml:"killer" validate:"required"`
		Locker    string   `json:"locker" yaml:"locker" validate:"required"`
		Service   string   `json:"service" yaml:"service" validate:"required"`
		Pipeline  string   `json:"pipeline" yaml:"pipeline" validate:"required"`
		Config    string   `json:"config" yaml:"config" validate:"required"`
		EndPoints []string `json:"endpoints" yaml:"endpoints" validate:"required"`
		Timeout   int64    `json:"timeout" yaml:"timeout" validate:"required"`
	}
	Database struct {
		Host string `json:"host" yaml:"host" validate:"required"`
		Port int    `json:"port" yaml:"port" validate:"required"`
		Name string `json:"name" yaml:"name" validate:"required"`
		User string `json:"user" yaml:"user" validate:"required"`
		Pass string `json:"pass" yaml:"pass" validate:"required"`
		Char string `json:"char" yaml:"char" validate:"required"`
	}
	User struct {
		Name    string `json:"name" validate:"required" yaml:"-"`
		Email   string `json:"email" validate:"required" yaml:"-"`
		Pass    string `json:"pass" validate:"required" yaml:"-"`
		Confirm string `json:"confirm" validate:"required" yaml:"-"`
	}
	Auth struct {
		Secret string `json:"secret" yaml:"secret" validate:"required"`
		TTL    int64  `json:"ttl" yaml:"ttl" validate:"required"`
	}
	Config struct {
		Database `json:"database"`
		Auth     `json:"auth"`
		Etcd     `json:"etcd"`
	}
)

var (
	Conf *Config
	Path string
)

func Init() *Config {
	return &Config{}
}

// 检查配置文件是否存在
func CheckConfigFile(path string) (bool, error) {
	_, err := os.Stat(path)
	exist := !os.IsNotExist(err)
	return exist, err
}

// 创建配置文件目录
func CreateConfigDir(dir string) {
	_, err := os.Stat(dir)

	if err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
}

// 检查配置文件目录是否有权限
func CheckConfigDirPermisson(dir string) bool {
	info, err := os.Stat(dir)
	if err != nil {
		log.Println(err)
	}
	mode := info.Mode()
	perm := mode.Perm()
	flag := perm & os.FileMode(493)
	if flag == 493 {
		return true
	}

	return false
}

// 写入配置文件
func WriteConfigToFile(file string, content []byte) bool {
	if err := ioutil.WriteFile(file, content, 0644); err != nil {
		log.Println(os.IsNotExist(err))
		log.Println(err)
	}
	return true
}
