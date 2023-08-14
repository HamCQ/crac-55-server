package routers

import (
	"crac55/app/server/handler"
	"crac55/app/server/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

// API 对外api
func (r *Router) API(router *mux.Router) {
	open := router.PathPrefix("/v1/55").Subrouter()
	open.Use(middleware.Middleware)
	{
		h := handler.NewCrac()
		open.HandleFunc("/search", h.Search).Methods(http.MethodGet)
		open.HandleFunc("/analyse/total/{year}", h.AnalyseTotal).Methods(http.MethodGet)
	}
}
