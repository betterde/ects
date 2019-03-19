package config

import (
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
	User struct{
		Name    string `json:"name" validate:"required" yaml:"-"`
		Email   string `json:"email" validate:"required" yaml:"-"`
		Pass    string `json:"pass" validate:"required" yaml:"-"`
		Confirm string `json:"confirm" validate:"required" yaml:"-"`
	} `yaml:"-"`
	Auth struct {
		Secret string `json:"secret" yaml:"secret"`
		TTL    int64    `json:"ttl" yaml:"ttl"`
	}
}

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
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0666)

	defer func() {
		if err := file.Close(); err != nil {
			// TODO
		}
	}()

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
