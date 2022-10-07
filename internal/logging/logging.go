package logging

import (
	"fmt"

	"github.com/FluffyFoxTail/extendlogger"
)

type BotLogger struct {
	*extendlogger.ExtendLogger
}

func NewLogger() *BotLogger {
	logger := extendlogger.NewExtendLogger()
	logger.SetLogLevel(extendlogger.LogLevelInfo)
	return &BotLogger{logger}
}

func (bl *BotLogger) Println(v ...interface{}) {
	fmt.Println(v...)
}

func (bl *BotLogger) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}
