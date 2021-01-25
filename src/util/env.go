package util

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Server struct {
		Http struct {
			Port     	int 	`yaml:"port" env:"SERVER_HTTP_PORT" env-default:"8080"`
			Mode     	string 	`yaml:"mode" env:"SERVER_HTTP_MODE" env-default:"production"`

		} `yaml:"http"`
	}  	`yaml:"server"`
	Database struct {
		Type 		string 	`yaml:"type" env:"DATABASE_TYPE" env-default:""`
		Host 		string 	`yaml:"host" env:"DATABASE_HOST" env-default:""`
		Port		int 	`yaml:"port" env:"DATABASE_PORT" env-default:"5984"`
		Username	string 	`yaml:"username" env:"DATABASE_USERNAME" env-default:""`
		Password	string 	`yaml:"password" env:"DATABASE_PASSWORD" env-default:""`
		Name	string 	`yaml:"name" env:"DATABASE_NAME" env-default:""`
	}	`yaml:"database"`
	Sentry struct {
		Dsn			string 	`yaml:"dsn" env:"SENTRY_DSN" env-default:""`
	}	`yaml:"sentry"`
}

var cfg Config
var read = false

func getConfig(){
	err := cleanenv.ReadConfig("config.yaml", &cfg)
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
