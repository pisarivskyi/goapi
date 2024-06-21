package goapi

import "fmt"

type Color string

const (
	RedColor    Color = "\033[1;31m%s\033[0m"
	GreenColor  Color = "\033[1;32m%s\033[0m"
	YellowColor Color = "\033[1;33m%s\033[0m"
	BlueColor   Color = "\033[1;34m%s\033[0m"
	CyanColor   Color = "\033[1;36m%s\033[0m"
	DebugColor  Color = "\033[0;36m%s\033[0m"
)

func colorize(s string, c Color) string {
	return fmt.Sprintf(string(c), s)
}
