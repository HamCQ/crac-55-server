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
		open.HandleFunc("/analyse/total", h.AnalyseTotal).Methods(http.MethodGet)
		open.HandleFunc("/analyse/rank/top5", h.AnalyseRankTop5).Methods(http.MethodGet)
		open.HandleFunc("/analyse/rank/all", h.AnalyseRankAll).Methods(http.MethodGet)
		open.HandleFunc("/analyse/barchart/by09", h.AnalyseBY09BarChar).Methods(http.MethodGet)
		open.HandleFunc("/analyse/barchart/bncra", h.AnalyseBnCraBarChart).Methods(http.MethodGet)
		open.HandleFunc("/analyse/province", h.AnalyseByProvince).Methods(http.MethodGet)
	}
}
