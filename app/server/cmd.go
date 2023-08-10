package server

import (
	"crac55/app/server/routers"

	"github.com/gorilla/mux"
)

// Router http入口
func Router() *mux.Router {
	r := routers.Router{}
	return r.SetRouter()
}
