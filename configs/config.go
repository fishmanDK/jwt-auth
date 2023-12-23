package configs

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HttpServer `yaml:"h_http-server"`
}

type HttpServer struct {
	Address      string        `yaml:"address" env-default:"localhost:8080"`
	ReadTimeout  time.Duration `yaml:"readTimeout"  env-default:"3s"`
	WriteTimeout time.Duration `yaml:"writeTimeout"  env-default:"3s"`
}

func NewConfig() *Config {
	config_path := "configs/config.yaml"

	if config_path == "" {
		log.Fatal("config_path is not set")
	}

	if _, err := os.Stat(config_path); err != nil {
		log.Fatalf("config file does not exist: %s", config_path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(config_path, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err.Error())
	}
	return &cfg
}
