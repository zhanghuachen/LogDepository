package myLog

type Level int

const (
	DEBUG Level = iota
	TRACE
	INFO
	WARN
	ERROR
	FATAL
)

func getLevel(level Level)string{
	switch level {
	case 0:
		return "DEBUG"
	case 1:
		return "TRACE"
	case 2:
		return "INFO"
	case 3:
		return "WARN"
	case 4:
		return "ERROR"
	case 5:
		return "FATAL"
	default:
		return "DEBUG"
	}
}
