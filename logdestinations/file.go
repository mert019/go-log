package logdestinations

import (
	"fmt"
	"os"
	"sync"

	"github.com/mert019/go-log/gologcore"
)

type FileLoggerConfiguration struct {
	FileName string
}

func NewFileLogger(configuration FileLoggerConfiguration) (gologcore.ILogDestination, error) {

	return &FileLogger{
		configuration: configuration,
		fileName:      configuration.FileName,
	}, nil
}

type FileLogger struct {
	configuration FileLoggerConfiguration
	fileName      string
	mu            sync.Mutex
}

func (fileLogger *FileLogger) Log(logModel gologcore.Log) error {

	fileLogger.mu.Lock()
	defer fileLogger.mu.Unlock()

	message := gologcore.GetContent(logModel)

	file, err := os.OpenFile(fileLogger.fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, message)
	return err
}

func (fileLogger *FileLogger) Close() error {
	return nil
}
