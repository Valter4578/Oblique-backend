package logger

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

// LogError gets pointer to error and prints formated information about this one
func LogError(err *error) {
	color.Red("%v: %v", time.Now(), *err)
}

func LogInfo(inf string) {
	color.Red("%v: %v", time.Now(), inf)
}

// JSONMessage gets message and returns its formatted to json
func JSONMessage(msg string) string {
	return fmt.Sprintf("{\"message:\": \"%v\"}", msg)
}

// JSONError gets error and returns its formatted to json
func JSONError(err error) string {
	return fmt.Sprintf("{\"error:\": \"%v\"}", err.Error())
}
