package golog

import (
	"fmt"
	"sync"

	"github.com/mert019/go-log/gologcore"
	"github.com/mert019/go-log/logdestinations"
)

var loggerInstance gologcore.ILogger

type Logger struct {
	logDestinations []gologcore.LogDestinationConfiguration

	consoleLogger gologcore.ILogDestination
}

func InitializeLogger(configuration gologcore.LoggerConfiguration) error {

	newLogger, err := NewLogger(configuration)
	if err != nil {
		return err
	}

	loggerInstance = newLogger
	return nil
}

func GetLogger() (gologcore.ILogger, error) {
	if loggerInstance == nil {
		return nil, &gologcore.LoggerIsNotInitializedError{}
	}
	return loggerInstance, nil
}

func NewLogger(configuration gologcore.LoggerConfiguration) (gologcore.ILogger, error) {

	err := configuration.Validate()
	if err != nil {
		return nil, err
	}

	loggerInstance := &Logger{
		logDestinations: configuration.LogDestinationConfigurations,
	}

	if configuration.OnErrorWriteToConsole {
		consoleLogger, err := logdestinations.NewConsoleLogger()
		if err != nil {
			return nil, err
		}

		loggerInstance.consoleLogger = consoleLogger
	}

	return loggerInstance, nil
}

func (l *Logger) Log(logModel *gologcore.LogModel) {

	if len(l.logDestinations) == 0 {
		return
	}

	var wg sync.WaitGroup

	errChan := make(chan error, len(l.logDestinations))

	log := logModel.MapToLog()

	for _, logDestination := range l.logDestinations {

		if logDestination.MinimumLogLevel > logModel.LogType {
			continue
		}

		wg.Add(1)

		go func(d gologcore.LogDestinationConfiguration) {
			defer wg.Done()

			err := d.Destination.Log(log)
			if err != nil {
				errChan <- fmt.Errorf("error occured while sending log to '%s': %v", d.Name, err)
			}
		}(logDestination)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	var errors []error
	for err := range errChan {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		fmt.Printf("encountered %d error(s) while logging: %v\n", len(errors), errors)
	}

	if l.consoleLogger != nil && len(errors) == len(l.logDestinations) {
		l.consoleLogger.Log(log)
	}

}

func (l *Logger) Close() {

	var errors []error

	for _, logDestination := range l.logDestinations {
		err := logDestination.Destination.Close()
		if err != nil {
			errors = append(errors, fmt.Errorf("error occurred while closing log destination '%s': %d", logDestination.Name, err))
		}
	}

	if len(errors) > 0 {
		fmt.Printf("encountered %d error(s) while closing: %v\n", len(errors), errors)
	}
}
