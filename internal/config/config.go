package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	Listen struct {
		Port string `yaml:"port" env-default:"8080"`
	}
	DB struct {
		URI string `yaml:"uri" env-required:"true"`
	} `yaml:"db" env-required:"true"`
	IpApiURL  string `yaml:"ip_api_url" env-required:"true"`
	JWTSecret string `yaml:"jwt_secret" env-required:"true"`
}

var instance *Config
var once sync.Once

func GetConfig(path string) *Config {
	once.Do(func() {
		log.Println("read application config")
		instance = &Config{}
		if err := cleanenv.ReadConfig(path, instance); err != nil {
			log.Fatal(err)
		}
	})
	return instance
}
