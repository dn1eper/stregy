package config

import (
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug  *bool  `yaml:"is_debug" env-required:"true"`
	LogLevel string `yaml:"log_level" env-default:"error"`
	Listen   struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
	PosgreSQL PosgreSQLConfig `yaml:"postgresql"`
}

type PosgreSQLConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yaml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Info(help)
			log.Fatal(err)
		}
	})
	return instance
}
