package core

type ILogger interface {
	Log(*LogModel)
	Close()
}

type ILogDestination interface {
	Log(Log) error
	Close() error
}
