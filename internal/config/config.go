package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	Server      `yaml:"http_server"`
}

type Server struct {
	Address     string        `yaml:"address" env-default:"localhost:8182"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./config/local.yaml"
	}

	//Проверка на существование файла
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file %s is not exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Can not read file: %s", configPath)
	}

	return &cfg
}
