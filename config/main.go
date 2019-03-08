package config

import (
	"os"
	"time"
)

type Config struct {
	Service struct {
		Host string
		Port int
	}
	Database struct {
		Host string
		Port int
		Name string
		User string
		Pass string
		Char string
	}
	Auth struct{
		Secret string
		TTL time.Duration
	}
}

var (
	Conf *Config
	Path string
)

type Auth struct {
	Secret string
	TTL time.Duration
}

func Init() *Config {
	conf := &Config{}
	return conf
}

// 检查配置文件是否存在
func CheckConfigFile(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}