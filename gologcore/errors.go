package gologcore

import "fmt"

type LoggerIsNotInitializedError struct{}

func (e *LoggerIsNotInitializedError) Error() string {
	return "logger is not initialized"
}

type LogDestinationConnectionError struct {
	Destination     string
	ConnectionError error
}

func (e *LogDestinationConnectionError) Error() string {
	return fmt.Sprintf("failed to connect '%s': %v", e.Destination, e.ConnectionError)
}

type GoLogValidationError struct {
	Message string
}

func (e *GoLogValidationError) Error() string {
	return e.Message
}
