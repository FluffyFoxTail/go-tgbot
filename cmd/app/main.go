package main

import (
	"tgbot/internal/config"
	"tgbot/internal/logging"
	"tgbot/pkg/tgbot"
)

// TODO implement bot Start/Stop and listen methonds
func main() {
	logger := logging.NewLogger()

	cfg, err := config.GetConfig(logger)
	if err != nil {
		panic(err)
	}

	bot, err := tgbot.NewBot(cfg, logger)
	if err != nil {
		panic(err)
	}

	_ = bot
	// TODO
}
