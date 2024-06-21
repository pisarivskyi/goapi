package goapi

type Level int16

// verbose, debug, warning, error
const (
	VerboseLevel Level = 1
	DebugLevel   Level = 2
	InfoLevel    Level = 3
	WarningLevel Level = 4
	ErrorLevel   Level = 5
)

func (l Level) String() string {
	switch l {
	case VerboseLevel:
		return "[verbose]"
	case DebugLevel:
		return "[debug]"
	case InfoLevel:
		return "[info]"
	case WarningLevel:
		return "[warning]"
	case ErrorLevel:
		return "[error]"
	default:
		return ""
	}
}

func (l Level) Color() Color {
	switch l {
	case VerboseLevel:
		return CyanColor
	case DebugLevel:
		return BlueColor
	case InfoLevel:
		return GreenColor
	case WarningLevel:
		return YellowColor
	case ErrorLevel:
		return RedColor
	default:
		return CyanColor
	}
}
