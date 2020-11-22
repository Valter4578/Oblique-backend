package api

import (
	"github.com/gorilla/mux"
)

type API struct {
	IsDebugMode bool
	Router *mux.Router
}
