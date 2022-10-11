package greeting

import (
	"tgbot/internal/config"
	"tgbot/pkg/tgbot/handlers"
)

type handler struct {
	Commands config.CommandsList
}

func NewHandler(commands config.CommandsList) handlers.Handler {
	return &handler{Commands: commands}
}

func (h *handler) HandleMsg(command string, args []string) (answer string, err error) {
	answer, err = h.Commands.GetAnswer(command)
	if err != nil {
		return "", err
	}
	return answer, nil
}
