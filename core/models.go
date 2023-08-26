package core

import "time"

type LogType int

const (
	Debug LogType = iota
	Info
	Warning
	Error
	Critical
)

var logTypeDescriptions = map[LogType]string{
	Critical: "critical",
	Debug:    "debug",
	Error:    "error",
	Info:     "info",
	Warning:  "warning",
}

func getLogTypeDescription(logType LogType) string {
	description, ok := logTypeDescriptions[logType]
	if !ok {
		return "unknown"
	}
	return description
}

type LogModel struct {
	Message      string
	LogType      LogType
	Error        error
	RequestBody  string
	ResponseBody string
	Parameters   map[string]string
}

type Log struct {
	Message      string
	LogType      LogType
	Error        string
	RequestBody  string
	ResponseBody string
	Parameters   map[string]string

	LogTypeDescription string
	LogDateTime        time.Time
}

func (lm *LogModel) MapToLog() Log {

	log := Log{
		Message:      lm.Message,
		LogType:      lm.LogType,
		RequestBody:  lm.RequestBody,
		ResponseBody: lm.ResponseBody,
		Parameters:   lm.Parameters,

		LogTypeDescription: getLogTypeDescription(lm.LogType),
		LogDateTime:        time.Now().UTC(),
	}

	if lm.Error != nil {
		log.Error = lm.Error.Error()
	}

	return log
}
