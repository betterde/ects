package config

import "time"

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

var Conf *Config

type Auth struct {
	Secret string
	TTL time.Duration
}

func Init() *Config {
	conf := &Config{}
	return conf
}