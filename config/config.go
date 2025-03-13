package config

import (
	"errors"
	"fmt"
	"github.com/betterde/ects/internal/journal"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

type (
	Etcd struct {
		Killer    string   `json:"killer" yaml:"killer" mapstructure:"" validate:"required"`
		Locker    string   `json:"locker" yaml:"locker" mapstructure:"" validate:"required"`
		Service   string   `json:"service" yaml:"service" mapstructure:"" validate:"required"`
		Pipeline  string   `json:"pipeline" yaml:"pipeline" mapstructure:"" validate:"required"`
		Config    string   `json:"config" yaml:"config" mapstructure:"" validate:"required"`
		EndPoints []string `json:"endpoints" yaml:"endpoints" mapstructure:"" validate:"required"`
		Timeout   int64    `json:"timeout" yaml:"timeout" mapstructure:"" validate:"required"`
	}
	Database struct {
		Host string `json:"host" yaml:"host" mapstructure:"" validate:"required"`
		Port int    `json:"port" yaml:"port" mapstructure:"" validate:"required"`
		Name string `json:"name" yaml:"name" mapstructure:"" validate:"required"`
		User string `json:"user" yaml:"user" mapstructure:"" validate:"required"`
		Pass string `json:"pass" yaml:"pass" mapstructure:"" validate:"required"`
		Char string `json:"char" yaml:"char" mapstructure:"" validate:"required"`
	}
	User struct {
		Name    string `json:"name" yaml:"-" mapstructure:"" validate:"required"`
		Email   string `json:"email" yaml:"-" mapstructure:"" validate:"required"`
		Pass    string `json:"pass" yaml:"-" mapstructure:"" validate:"required"`
		Confirm string `json:"confirm" yaml:"-" mapstructure:"" validate:"required"`
	}
	Auth struct {
		Secret string `json:"secret" yaml:"secret" mapstructure:"" validate:"required"`
		TTL    int64  `json:"ttl" yaml:"ttl" mapstructure:"" validate:"required"`
	}
	Notification struct {
		Url        string `json:"url" yaml:"url" mapstructure:"" validate:"required"`
		Host       string `json:"host" yaml:"host" mapstructure:"" validate:"required"`
		Port       int    `json:"port" yaml:"port" mapstructure:"" validate:"numeric"`
		User       string `json:"user" yaml:"user" mapstructure:"" validate:"required"`
		Pass       string `json:"pass" yaml:"pass" mapstructure:"" validate:"required"`
		Name       string `json:"name" yaml:"name" mapstructure:"" validate:"required"`
		Protocol   string `json:"protocol" yaml:"protocol" mapstructure:"" validate:"required"`
		Encryption string `json:"encryption" yaml:"encryption" mapstructure:"" validate:"required"`
	}

	HTTP struct {
		Listen  string `yaml:"listen" mapstructure:"LISTEN"`
		TLSKey  string `yaml:"tlsKey" mapstructure:"TLSKEY"`
		TLSCert string `yaml:"tlsCert" mapstructure:"TLSCERT"`
	}

	GRPC struct {
		Listen  string `yaml:"listen" mapstructure:"LISTEN"`
		TLSKey  string `yaml:"tlsKey" mapstructure:"TLSKEY"`
		TLSCert string `yaml:"tlsCert" mapstructure:"TLSCERT"`
	}

	Logging struct {
		Level string `yaml:"level" mapstructure:"LEVEL"`
	}

	Config struct {
		Auth         Auth         `json:"auth" yaml:"auth"`
		Etcd         Etcd         `json:"etcd" yaml:"etcd"`
		HTTP         HTTP         `json:"http" yaml:"http" mapstructure:"HTTP"`
		GRPC         GRPC         `json:"grpc" yaml:"grpc" mapstructure:"GRPC"`
		Logging      Logging      `json:"logging" yaml:"logging" mapstructure:"LOGGING"`
		Database     Database     `json:"database" yaml:"database" mapstructure:"DATABASE"`
		Notification Notification `json:"notification" yaml:"notification" mapstructure:"NOTIFICATION"`
	}
)

var (
	keys = []string{
		"ENV",
		"HTTP.LISTEN",
		"HTTP.TLSKEY",
		"HTTP.TLSCERT",
		"GRPC.LISTEN",
		"GRPC.TLSKEY",
		"GRPC.TLSCERT",
		"LOGGING.LEVEL",
		"DATABASE.URL",
		"DATABASE.DATABASE",
		"DATABASE.USERNAME",
		"DATABASE.PASSWORD",
		"DATABASE.NAMESPACE",
	}
	Conf *Config
)

func Init() *Config {
	return &Config{}
}

// CheckConfigFile 检查配置文件是否存在
func CheckConfigFile(path string) (bool, error) {
	_, err := os.Stat(path)
	exist := !os.IsNotExist(err)
	return exist, err
}

// CreateConfigDir 创建配置文件目录
func CreateConfigDir(dir string) {
	_, err := os.Stat(dir)

	if err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
}

// CheckConfigDirPermisson 检查配置文件目录是否有权限
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

// WriteConfigToFile 写入配置文件
func WriteConfigToFile(file string, content []byte) bool {
	if err := os.WriteFile(file, content, 0644); err != nil {
		log.Println(os.IsNotExist(err))
		log.Println(err)
	}
	return true
}

func Parse(file string, envPrefix string) {
	if file != "" {
		viper.SetConfigFile(file)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName(".focusly")
		viper.AddConfigPath("/etc/focusly")
	}

	replacer := strings.NewReplacer(".", "_")

	// read in environment variables that match
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix(envPrefix)

	var notFoundError viper.ConfigFileNotFoundError

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil && errors.As(err, &notFoundError) {
		for _, key := range keys {
			if err := viper.BindEnv(key, fmt.Sprintf("%s_%s", envPrefix, replacer.Replace(key))); err != nil {
				journal.Logger.Error(err)
			}
		}
	}

	// read in environment variables that match
	viper.AutomaticEnv()

	err := viper.Unmarshal(&Conf)
	if err != nil {
		journal.Logger.Errorf("Unable to decode into config struct, %v", err)
		os.Exit(1)
	}
}
