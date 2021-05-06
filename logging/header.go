package logging

import "fmt"

type Level int

const (
	Fatal Level = iota // 0
	Error
	Warning
	Info
	Debug
)

func appendHeader(l Level, msg string) string {
	prepend := ""
	switch l {
	case Fatal:
		prepend = "Fatal"
	case Error:
		prepend = "Error"
	case Warning:
		prepend = "Warning"
	case Info:
		prepend = "Info"
	case Debug:
		prepend = "Debug"

	}
	return fmt.Sprintf("%s - %s", prepend, msg)

}
