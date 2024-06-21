package goapi

type Level int16

// verbose, debug, warning, error
const (
	VerboseLevel Level = 1 + iota
	DebugLevel
	InfoLevel
	WarningLevel
	ErrorLevel
)

func (l Level) String() string {
	switch l {
	case VerboseLevel:
		return "[VERBOSE]"
	case DebugLevel:
		return "[DEBUG]"
	case InfoLevel:
		return "[INFO]"
	case WarningLevel:
		return "[WARNING]"
	case ErrorLevel:
		return "[ERROR]"
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
