package configs

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Server struct {
	Port string `yaml:"port"`
}




func NewServer() Server {
	config_path := "configs/config.yaml"

	if config_path == "" {
		log.Fatal("config_path is not set")
	}

	if _, err := os.Stat(config_path); err != nil {
		log.Fatalf("config file does not exist: %s", config_path)
	}

	var cfg Server

	if err := cleanenv.ReadConfig(config_path, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err.Error())
	}
	return cfg
}


