# go-log
A versatile Go logger for multiple destinations.

## Installation
```console
go get github.com/mert019/go-log
```

## Usage
Create a log destination instance that implements the `gologcore.ILogDestination`.
You can use go-log's pre-implemented destionations, or you can implement your own log destination.
```go
rabbitmqLogger, err = logdestinations.NewRabbitMQLogger(
		logdestinations.RabbitMQLoggerConfiguration{
			Url:       "amqp://RabbitMQUser:Rabbit123456@localhost:5672/",
			QueueName: "logs",
		},
	)
```

Create a logger configuration. Add the log destinations with name and specify the log level.
```go
loggerConfiguration := gologcore.LoggerConfiguration{
		LogDestinationConfigurations: []gologcore.LogDestinationConfiguration{
			{
				Name:            "RabbitMQ",
				MinimumLogLevel: gologcore.Info,
				Destination:     rabbitmqLogger,
			},
			// Add other destinations
		},
	}
```

Initialize the global logger or create a new logger instance.
```go
// Initialize Global Logger
err := golog.InitializeLogger(loggerConfiguration)

if err != nil {
    panic(err)
}

logger, err := golog.GetLogger()
```
```go
// Create new logger instance
logger, err := golog.NewLogger(loggerConfiguration)
```

Send log by using Log method. You can use the pre-defined log model fields or  use custom parameters.
```go
customParameters := map[string]string{
    "Key1": "Value1",
    "Key2": "Value2",
}

logger.Log(
    &gologcore.LogModel{
        LogType:    gologcore.Info,
        Message:    "My first go-log message",
        Parameters: customParameters,
    },
)
```

Call Close method at the end of the program to dispose the logger.
```go
logger.Close()
```

## Log Levels
```console
- Debug
- Info
- Warning
- Error
- Critical
```

## Project Status
This project is currently in development and may not be production-ready. Please use it with caution and feel free to contribute or report issues.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.