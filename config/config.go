package config

import (
	"errors"
	"github.com/betterde/ects/internal/journal"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

type (
	Etcd struct {
		Config    string   `json:"config" yaml:"config" mapstructure:"config" validate:"required"`
		Killer    string   `json:"killer" yaml:"killer" mapstructure:"killer" validate:"required"`
		Locker    string   `json:"locker" yaml:"locker" mapstructure:"locker" validate:"required"`
		Service   string   `json:"service" yaml:"service" mapstructure:"service" validate:"required"`
		Pipeline  string   `json:"pipeline" yaml:"pipeline" mapstructure:"pipeline" validate:"required"`
		Timeout   uint32   `json:"timeout" yaml:"timeout" mapstructure:"timeout" validate:"required"`
		EndPoints []string `json:"endpoints" yaml:"endpoints" mapstructure:"endpoints" validate:"required"`
	}
	Database struct {
		Host string `json:"host" yaml:"host" mapstructure:"host" validate:"required"`
		Port uint16 `json:"port" yaml:"port" mapstructure:"port" validate:"required"`
		Name string `json:"name" yaml:"name" mapstructure:"name" validate:"required"`
		User string `json:"user" yaml:"user" mapstructure:"user" validate:"required"`
		Pass string `json:"pass" yaml:"pass" mapstructure:"pass" validate:"required"`
		Char string `json:"char" yaml:"char" mapstructure:"char" validate:"required"`
	}
	User struct {
		Name    string `json:"name" yaml:"-" mapstructure:"" validate:"required"`
		Email   string `json:"email" yaml:"-" mapstructure:"" validate:"required"`
		Pass    string `json:"pass" yaml:"-" mapstructure:"" validate:"required"`
		Confirm string `json:"confirm" yaml:"-" mapstructure:"" validate:"required"`
	}
	Auth struct {
		TTL    uint64 `json:"ttl" yaml:"ttl" mapstructure:"ttl" validate:"required"`
		Secret string `json:"secret" yaml:"secret" mapstructure:"secret" validate:"required"`
	}
	Notification struct {
		Url        string `json:"url" yaml:"url" mapstructure:"url" validate:"required"`
		Host       string `json:"host" yaml:"host" mapstructure:"host" validate:"required"`
		Port       uint16 `json:"port" yaml:"port" mapstructure:"port" validate:"numeric"`
		User       string `json:"user" yaml:"user" mapstructure:"user" validate:"required"`
		Pass       string `json:"pass" yaml:"pass" mapstructure:"pass" validate:"required"`
		Name       string `json:"name" yaml:"name" mapstructure:"name" validate:"required"`
		Protocol   string `json:"protocol" yaml:"protocol" mapstructure:"protocol" validate:"required"`
		Encryption string `json:"encryption" yaml:"encryption" mapstructure:"encryption" validate:"required"`
	}

	HTTP struct {
		Listen  string `json:"listen" yaml:"listen" mapstructure:"listen"`
		TLSKey  string `json:"tlsKey" yaml:"tlsKey" mapstructure:"tlskey"`
		TLSCert string `json:"tlsCert" yaml:"tlsCert" mapstructure:"tlscert"`
	}

	Logging struct {
		Level string `yaml:"level" mapstructure:"level"`
	}

	Config struct {
		Auth         Auth         `json:"auth" yaml:"auth" mapstructure:"auth"`
		Etcd         Etcd         `json:"etcd" yaml:"etcd" mapstructure:"etcd"`
		HTTP         HTTP         `json:"http" yaml:"http" mapstructure:"http"`
		Logging      Logging      `json:"logging" yaml:"logging" mapstructure:"logging"`
		Database     Database     `json:"database" yaml:"database" mapstructure:"database"`
		Notification Notification `json:"notification" yaml:"notification" mapstructure:"notification"`
	}
)

var Conf *Config

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
		viper.SetConfigType("yaml")
		viper.SetConfigName(".ects")
	}

	var notFoundError viper.ConfigFileNotFoundError

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil && errors.As(err, &notFoundError) {
		journal.Logger.Debugf("Config file not found, using defaults")
	}

	// read in environment variables that match
	viper.AutomaticEnv()

	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer("_", "."))

	err := viper.Unmarshal(&Conf)
	if err != nil {
		journal.Logger.Errorf("Unable to decode into config struct, %v", err)
		os.Exit(1)
	}
}
