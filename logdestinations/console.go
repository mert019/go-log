package logdestinations

import (
	"log"

	"github.com/mert019/go-log/gologcore"
)

func NewConsoleLogger() (gologcore.ILogDestination, error) {

	return &ConsoleLogger{}, nil
}

type ConsoleLogger struct{}

func (consoleLogger *ConsoleLogger) Log(logModel gologcore.Log) error {

	message := gologcore.GetContent(logModel)

	log.Println(message)

	return nil
}

func (consoleLogger *ConsoleLogger) Close() error {
	return nil
}
