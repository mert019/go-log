package gologcore

import (
	"fmt"
	"strings"
)

func GetContent(logModel Log) string {

	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("<datetime>=%v", logModel.LogDateTime))

	if len(strings.TrimSpace(logModel.LogTypeDescription)) != 0 {
		builder.WriteString(fmt.Sprintf(" <logtype>=%s", logModel.LogTypeDescription))
	}

	if len(strings.TrimSpace(logModel.Message)) != 0 {
		builder.WriteString(fmt.Sprintf(" <message>=%s", logModel.Message))
	}

	if len(strings.TrimSpace(logModel.Error)) != 0 {
		builder.WriteString(fmt.Sprintf(" <error>=%s", logModel.Error))
	}

	if len(strings.TrimSpace(logModel.RequestBody)) != 0 {
		builder.WriteString(fmt.Sprintf(" <requestbody>=%s", logModel.RequestBody))
	}

	if len(strings.TrimSpace(logModel.ResponseBody)) != 0 {
		builder.WriteString(fmt.Sprintf(" <responsebody>=%s", logModel.ResponseBody))
	}

	for key, value := range logModel.Parameters {
		builder.WriteString(fmt.Sprintf(" <%s>=%s", key, value))
	}

	return builder.String()
}
