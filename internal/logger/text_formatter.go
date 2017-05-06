package logger

import (
	"fmt"
	"strings"

	"github.com/mgutz/ansi"
)

type TextFormatter struct {
	Colored bool
}

func colorForLogLevel(level LogLevel) string {
	switch level {
	case PanicLevel:
		return "red"
	case FatalLevel:
		return "red"
	case ErrorLevel:
		return "red"
	case WarnLevel:
		return "yellow"
	case InfoLevel:
		return "green"
	case DebugLevel:
		return ansi.DefaultFG
	}

	return ansi.DefaultFG
}

func appendNewLine(message string) string {
	if strings.HasSuffix(message, "\n") {
		return message
	}

	return fmt.Sprintf("%s\n", message)
}

func (f *TextFormatter) Format(level LogLevel, message string) ([]byte, error) {
	msg := appendNewLine(message)

	if f.Colored == false {
		return []byte(msg), nil
	}

	coloredMsg := ansi.Color(msg, colorForLogLevel(level))
	return []byte(coloredMsg), nil
}
