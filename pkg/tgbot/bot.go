package tgbot

import (
	"fmt"
	"tgbot/internal/config"
	"tgbot/internal/logging"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api    *tgbotapi.BotAPI
	logger *logging.BotLogger
	config *config.Config
}

func NewBot(config *config.Config, logger *logging.BotLogger) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		return nil, fmt.Errorf("Unnable to connect TG: %w", err)
	}

	err = tgbotapi.SetLogger(logger)
	if err != nil {
		return nil, fmt.Errorf("Unnable to set logger: %w", err)
	}

	if config.IsDebug {
		logger.Info("Debug mod is one")
		api.Debug = true
	}

	logger.Info(fmt.Sprintf("Log in on: %s", api.Self.UserName))
	return &Bot{api: api, config: config, logger: logger}, nil
}
