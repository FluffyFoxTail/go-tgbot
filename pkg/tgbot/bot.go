package tgbot

import (
	"context"
	"fmt"
	"tgbot/internal/config"
	"tgbot/internal/logging"
	"tgbot/pkg/tgbot/handlers"
	"tgbot/pkg/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api      *tgbotapi.BotAPI
	logger   *logging.BotLogger
	config   *config.Config
	services map[string]handlers.Handler
}

func NewBot(config *config.Config, logger *logging.BotLogger, services map[string]handlers.Handler) (*Bot, error) {
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
	return &Bot{api: api, config: config, logger: logger, services: services}, nil
}

func (b *Bot) Start(ctx context.Context) {
	update := tgbotapi.NewUpdate(0)
	update.Timeout = b.config.Timeout
	updates := b.api.GetUpdatesChan(update)
	go b.listen(ctx, updates)
}

func (b *Bot) listen(ctx context.Context, updates tgbotapi.UpdatesChannel) {
	for {
		select {
		case <-ctx.Done():
			b.logger.Warn(fmt.Sprintf("Msg conn closed: %v", ctx.Err()))
			return
		case u := <-updates:
			if u.Message == nil {
				continue
			}

			if !u.Message.IsCommand() {
				b.reply(u.Message.Chat.ID, u.Message.MessageID, "just reply: "+u.Message.Text)
				continue
			}

			// TODO match a service and set of it`s commands
			if utils.Contains([]string{"hello", "help"}, u.Message.Command()) {
				answer, err := b.services["greeting"].HandleMsg(u.Message.Command(), []string{})
				if err != nil {
					b.logger.Error(err.Error())
				}
				go b.reply(u.Message.Chat.ID, u.Message.MessageID, answer)
			} else {
				go b.reply(u.Message.Chat.ID, 0, "I dont now this command")
			}
		}
	}
}

func (b *Bot) Stop() {
	b.api.StopReceivingUpdates()
}

func (b *Bot) reply(chatId int64, msgID int, text string) {
	b.logger.Info(fmt.Sprintf("Sending reply for: %d", chatId))
	msg := tgbotapi.NewMessage(chatId, fmt.Sprint(text))
	if msgID != 0 {
		msg.ReplyToMessageID = msgID
	}

	if _, err := b.api.Send(msg); err != nil {
		b.logger.Error(fmt.Sprintf("Unable to send msg: %v", err))
		return
	}
}
