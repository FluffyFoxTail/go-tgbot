package config

import (
	"sync"
	"tgbot/internal/logging"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug bool   `yaml:"is_debug"`
	Token   string `yaml:"token" evn:"TGTOKEN"`
	Timeout int    `yaml:"timeout"`
}

var instance *Config
var once sync.Once

func GetConfig(logger *logging.BotLogger) (*Config, error) {
	once.Do(func() {
		logger.Info("Read bot config")

		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Error(err.Error())
		}
	})
	return instance, nil
}
