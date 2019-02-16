package config

type Config struct {
	Service struct{
		Host string
		Port int
	}
	Database struct{
		Host string
		Port int
		Name string
		User string
		Pass string
	}
}

func Init() *Config {
	conf := &Config{}
	return conf
}