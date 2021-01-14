package util

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Port     		string `env:"PORT" env-default:"8080"`
	Dsn		 		string `env:"DSN" env-default:""`
	DatabaseHost 	string `env:"DATABASE_HOST" env-default:"http://localhost"`
	DatabasePort	string `env:"DATABASE_PORT" env-default:"5984"`
	DatabaseUser	string `env:"DATABASE_USER" env-default:"E5wT5GjSJbaLUR9tRvnkYqW59XpAZbuB"`
	DatabasePassword	string `env:"DATABASE_PASSWORD" env-default:"FbxJYzPUM2tL4kkKYdQFmv5XdGUwDPKz"`
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
