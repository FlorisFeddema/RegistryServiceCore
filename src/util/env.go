package util

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Port     string `env:"PORT" env-default:"8080"`
	Dsn		 string `env:"DSN" env-default:""`
}

var cfg Config
var read = false

func getConfig(){
	err := cleanenv.ReadConfig(".env", &cfg)
	if err == nil{
		read = true
	}
}

func GetConfig() Config {
	if !read {
		getConfig()
	}
	return cfg
}
