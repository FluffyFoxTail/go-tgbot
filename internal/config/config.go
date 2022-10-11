package config

import (
	"errors"
	"fmt"
	"sync"
	"tgbot/internal/logging"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug  bool         `yml:"is_debug"`
	Token    string       `yml:"token" evn:"TGTOKEN"`
	Timeout  int          `yml:"timeout"`
	Commands CommandsList `yml:"commands"`
}

type CommandsList []Command

type Command struct {
	Key     string `yml:"key"`
	Message string `yml:"message"`
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

func (cl CommandsList) GetAnswer(key string) (string, error) {
	for _, command := range cl {
		if command.Key == key {
			return command.Message, nil
		}
	}
	return "", errors.New(fmt.Sprintf("Such command not found: %s", key))
}
