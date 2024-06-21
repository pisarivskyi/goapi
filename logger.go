package goapi

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Logger struct {
	Output       io.Writer
	AllowedLevel Level
}

func NewLogger() *Logger {
	return &Logger{
		Output:       os.Stdout,
		AllowedLevel: DebugLevel,
	}
}

func (l *Logger) Verbose(msg string) {
	writeLog(l.Output, VerboseLevel, l.AllowedLevel, msg)
}

func (l *Logger) Debug(msg string) {
	writeLog(l.Output, DebugLevel, l.AllowedLevel, msg)
}

func (l *Logger) Info(msg string) {
	writeLog(l.Output, InfoLevel, l.AllowedLevel, msg)
}

func (l *Logger) Warning(msg string) {
	writeLog(l.Output, WarningLevel, l.AllowedLevel, msg)
}

func (l *Logger) Error(msg string) {
	writeLog(l.Output, ErrorLevel, l.AllowedLevel, msg)
}

func writeLog(w io.Writer, l Level, allowedLevel Level, msg string) {
	if l >= allowedLevel {
		fmt.Fprintf(w, "%s %s: %s\n", time.Now().UTC().Format(time.DateTime), colorize(l.String(), l.Color()), msg)
	}
}
