package routers

import (
	"crac55/app/server/handler"
	"crac55/app/server/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

// API 对外api
func (r *Router) API(router *mux.Router) {
	open := router.PathPrefix("/v1/crac55").Subrouter()
	open.Use(middleware.Middleware)
	{
		h := handler.NewCrac()
		open.HandleFunc("/analyse/total", h.AnalyseTotal).Methods(http.MethodGet)
	}
}
