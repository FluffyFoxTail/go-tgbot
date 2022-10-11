package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"tgbot/internal/config"
	"tgbot/internal/greeting"
	"tgbot/internal/logging"
	"tgbot/pkg/tgbot"
	"tgbot/pkg/tgbot/handlers"
)

// TODO REPLACE panic to beaty message
func main() {
	logger := logging.NewLogger()

	cfg, err := config.GetConfig(logger)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", cfg)
	greetingService := greeting.NewHandler(cfg.Commands)
	botServices := map[string]handlers.Handler{"greeting": greetingService}

	bot, err := tgbot.NewBot(cfg, logger, botServices)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger.Info("Start up bot")
	bot.Start(ctx)

	logger.Info("Bot on listen updates")
	defer bot.Stop()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	select {
	case signal := <-interrupt:
		logger.Warn(fmt.Sprintf("Signal was caught, app is stop: %v", signal))
		cancel()
	case <-ctx.Done():
		logger.Warn(fmt.Sprintf("Context was close, app is stop: %s", ctx.Err()))
		return
	}
}
