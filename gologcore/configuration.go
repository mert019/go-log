package gologcore

import "strings"

type LoggerConfiguration struct {
	LogDestinationConfigurations []LogDestinationConfiguration

	OnErrorWriteToConsole bool
}

type LogDestinationConfiguration struct {
	Name            string
	MinimumLogLevel LogType
	Destination     ILogDestination
}

func (lc *LoggerConfiguration) Validate() error {

	names := map[string]struct{}{}

	for _, destinationConfiguration := range lc.LogDestinationConfigurations {

		if destinationConfiguration == (LogDestinationConfiguration{}) {
			return &GoLogValidationError{Message: "LogDestinationConfiguration cannot be empty"}
		}

		if len(strings.TrimSpace(destinationConfiguration.Name)) == 0 {
			return &GoLogValidationError{Message: "LogDestinationConfiguration.Name cannot be empty"}
		}

		if destinationConfiguration.Destination == nil {
			return &GoLogValidationError{Message: "LogDestinationConfiguration.Destination cannot be empty"}
		}

		_, exists := names[destinationConfiguration.Name]
		if exists {
			return &GoLogValidationError{Message: "LogDestinationConfiguration.Name must be unique"}
		} else {
			names[destinationConfiguration.Name] = struct{}{}
		}
	}

	return nil
}
