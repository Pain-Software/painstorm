package config

import (
	"log"
	"log/slog"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App AppConfig `yaml:"app"`
}

type AppConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

var (
	instance *Config
	once     sync.Once
)

func Get() *Config {
	once.Do(func() {
		var cfg Config

		slog.Info("Reading configuration file.")
		data, err := os.ReadFile("./config.yaml")
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		slog.Info("Parsing configuration to struct")
		err = yaml.Unmarshal(data, &cfg)

		if err != nil {
			log.Fatalf("error: %v", err)
		}

		instance = &cfg
		slog.Info("Configuration has been successfully loaded.")
	})

	return instance
}
