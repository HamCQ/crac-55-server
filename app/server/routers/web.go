package routers

import (
	"crac55/app/server/handler"

	"github.com/gorilla/mux"
)

// Web 前端
func (r *Router) Web(router *mux.Router) {
	router.PathPrefix("/").Handler(handler.WebHandler{
		StaticPath: "web",
		IndexPath:  "index.html",
	})
}
