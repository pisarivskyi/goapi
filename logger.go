package goapi

import (
	"fmt"
	"time"
)

type Level int16

const (
	InfoLevel    Level = 1
	NoticeLevel  Level = 2
	WarningLevel Level = 3
	ErrorLevel   Level = 4
	DebugLevel   Level = 5
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

type logger struct {
}

func (l Level) String() string {
	switch l {
	case InfoLevel:
		return "[info]"
	case NoticeLevel:
		return "[notice]"
	case ErrorLevel:
		return "[error]"
	default:
		return ""
	}
}

func (logger) Info(msg string) {
	printLog(msg, InfoLevel)
}

func (logger) Error(msg string) {
	printLog(msg, ErrorLevel)
}

func printLog(msg string, l Level) {
	fmt.Printf("%s %s: %s\n", time.Now().UTC().Format(time.DateTime), colorize(l.String(), l), msg)
}

func colorize(s string, l Level) string {
	switch l {
	case InfoLevel:
		return fmt.Sprintf(InfoColor, s)
	case NoticeLevel:
		return fmt.Sprintf(NoticeColor, s)
	case ErrorLevel:
		return fmt.Sprintf(ErrorColor, s)
	default:
		return s
	}
}
