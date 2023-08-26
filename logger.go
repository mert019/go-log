package golog

import (
	"fmt"
	"sync"

	"github.com/mert019/go-log/gologcore"
)

var loggerInstance gologcore.ILogger

type Logger struct {
	logDestinations []gologcore.LogDestinationConfiguration
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

	return loggerInstance, nil
}

func (l *Logger) Log(logModel *gologcore.LogModel) {

	var wg sync.WaitGroup

	errChan := make(chan error, len(l.logDestinations))

	for _, logDestination := range l.logDestinations {

		if logDestination.MinimumLogLevel > logModel.LogType {
			continue
		}

		wg.Add(1)

		log := logModel.MapToLog()

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
