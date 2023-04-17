package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type Config struct {
	PostgreSQL struct {
		Username string `yaml:"username" env:"PSQL_USERNAME" env-required:"true"`
		Password string `yaml:"password" env:"PSQL_PASSWORD" env-required:"true"`
		Host     string `yaml:"host" env:"PSQL_HOST" env-required:"true"`
		Port     string `yaml:"port" env:"PSQL_PORT" env-required:"true"`
		Database string `yaml:"database" env:"PSQL_DATABASE" env-required:"true"`
	} `yaml:"postgresql"`
	Auth struct {
		HashSalt   string        `yaml:"hash_salt" env-required:"true"`
		SigningKey string        `yaml:"signing_key" env-required:"true"`
		TokenTtl   time.Duration `yaml:"token_ttl" env-required:"true"`
	} `yaml:"auth"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logrus.Info("read application config")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logrus.Info(help)
			logrus.Fatal(err)
		}
	})
	return instance
}
