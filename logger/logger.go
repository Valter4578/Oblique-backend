package logger

import (
	"fmt"
)

// JSONMessage gets message and returns its formatted to json
func JSONMessage(msg string, args ...interface{}) string {
	return fmt.Sprintf("{\"message:\": \"%v %v\"}", msg, args)
}

// JSONError gets error and returns its formatted to json
func JSONError(err error, args ...interface{}) string {
	return fmt.Sprintf("{\"error:\": \"%v %v\"}", err.Error(), args)
}
