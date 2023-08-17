package routers

import (
	"github.com/gorilla/mux"
)

type Router struct{}

func (r *Router) SetRouter() *mux.Router {
	//总路由
	router := mux.NewRouter()
	{
		r.API(router)
	}
	{
		r.Web(router)
	}
	return router
}
