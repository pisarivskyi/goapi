package goapi

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Logger struct {
	output          io.Writer
	allowedLevel    Level
	timeStampFormat string
}

func NewLogger() *Logger {
	return &Logger{
		output:          os.Stdout,
		allowedLevel:    DebugLevel,
		timeStampFormat: "02-Jan-2006 15:04:05.00",
	}
}

func (l *Logger) Verbose(msg string, a ...any) {
	l.writeLog(VerboseLevel, msg, a...)
}

func (l *Logger) Debug(msg string, a ...any) {
	l.writeLog(DebugLevel, msg, a...)
}

func (l *Logger) Info(msg string, a ...any) {
	l.writeLog(InfoLevel, msg, a...)
}

func (l *Logger) Warning(msg string, a ...any) {
	l.writeLog(WarningLevel, msg, a...)
}

func (l *Logger) Error(msg string, a ...any) {
	l.writeLog(ErrorLevel, msg, a...)
}

func (l *Logger) writeLog(lvl Level, msg string, a ...any) {
	if lvl >= l.allowedLevel {
		color := lvl.Color()
		timeStamp := time.Now().Format(l.timeStampFormat)
		logLevel := colorize(lvl.String(), color)

		sb := strings.Builder{}

		sb.WriteString(
			fmt.Sprintf("%s\t%s: %s\n", timeStamp, logLevel, colorize(msg, color)),
		)

		if len(a) > 0 {
			for _, item := range a {
				switch v := item.(type) {
				case error:
					sb.WriteString(colorize(fmt.Sprintf("%v\n", v.Error()), color))
				default:
					sb.WriteString(colorize(fmt.Sprintf("%v\n", v), color))
				}
			}
		}

		fmt.Fprint(l.output, sb.String())
	}
}
